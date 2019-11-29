package gateio

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

const WSSID = 20191129

// GateWSAgent GateIO Wss 客户端
type GateWSAgent struct {
	BaseURL string
	Config  *Config
	Conn    *websocket.Conn

	SubscribeResp chan *SubscribeResponse
	QueryResp     chan *QueryResponse
	ErrorCh       chan error

	SubscribeMap   map[string][]ReceivedDataCallback
	ActiveChannels map[string]bool

	processMutex sync.Mutex
}

// Start 开启 Wss 链接
func (gwa *GateWSAgent) Start(config *Config) (err error) {
	gwa.BaseURL = config.WSEndpoint
	conn, _, err := websocket.DefaultDialer.Dial(gwa.BaseURL, nil)
	if err != nil {
		log.Fatalf("dial:%+v", err)
		return err
	}

	gwa.Conn = conn
	gwa.Config = config

	if gwa.Config.IsPrint {
		log.Printf("Connected to %s", gwa.BaseURL)
	}

	gwa.SubscribeResp = make(chan *SubscribeResponse, 10)
	gwa.QueryResp = make(chan *QueryResponse, 10)
	gwa.ErrorCh = make(chan error)
	gwa.SubscribeMap = make(map[string][]ReceivedDataCallback)
	gwa.ActiveChannels = make(map[string]bool)

	go gwa.work()
	go gwa.receive()

	return nil
}

func (gwa *GateWSAgent) work() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	var err error

	for {
		select {
		case <-ticker.C:
			gwa.ping()
			gwa.time()
		case msg := <-gwa.SubscribeResp:
			err = gwa.handleSubscribe(msg)
			if err != nil {
				fmt.Println("handle subscribe failed", err)
			}
		case msg := <-gwa.QueryResp:
			err = gwa.handleQuery(msg)
			if err != nil {
				fmt.Println("handle query failed", err)
			}
		}
	}
}

// SendMessage 客户端向服务器发送数据结构体
type SendMessage struct {
	ID     int64         `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func (gwa *GateWSAgent) sendMsg(channel string, params []interface{}) (err error) {
	if gwa.Config.IsPrint {
		fmt.Println("send message", channel, params)
	}

	sendMsg := &SendMessage{WSSID, channel, params}
	sendStr, err := json.Marshal(sendMsg)
	if err != nil {
		if gwa.Config.IsPrint {
			fmt.Println("send json message", sendStr)
		}
		return err
	}

	if len(sendStr) == 0 {
		if gwa.Config.IsPrint {
			fmt.Println("send message empty return")
		}
		return nil
	}

	err = gwa.Conn.WriteMessage(websocket.TextMessage, []byte(sendStr))
	if err != nil {
		return err
	}

	return nil
}

func (gwa *GateWSAgent) ping() {
	gwa.sendMsg("server.ping", []interface{}{})
}

func (gwa *GateWSAgent) time() {
	gwa.sendMsg("server.time", []interface{}{})
}

func (gwa *GateWSAgent) handleSubscribe(msg *SubscribeResponse) (err error) {
	switch msg.Method {
	case "ticker.update":
		symbol := msg.Result[0].(string)
		ticker := &Ticker{}
		mapstructure.Decode(msg.Result[1], &ticker)

		if gwa.Config.IsPrint {
			str, _ := Struct2JsonString(ticker)
			fmt.Println("ticker", symbol, str)
		}
	case "trades.update":
		symbol := msg.Result[0].(string)
		trades := []*Trade{}
		switch t := msg.Result[1].(type) {
		case []interface{}:
			for _, v := range t {
				trade := &Trade{}
				mapstructure.Decode(v, &trade)
				trades = append(trades, trade)
			}
		}

		if gwa.Config.IsPrint {
			str, _ := Struct2JsonString(trades)
			fmt.Println("trades", symbol, str)
		}
	case "depth.update":
		symbol := msg.Result[2].(string)
		clean := msg.Result[0].(bool)
		asksbids := &_OrderBookItem{}
		mapstructure.Decode(msg.Result[1], asksbids)

		orderbook := &OrderBook{}

		orderbook.Asks = []Ask{}
		for _, ask := range asksbids.Asks {
			orderbook.Asks = append(
				[]Ask{
					Ask{ask[0], ask[1]},
				},
				orderbook.Asks...,
			)
		}

		orderbook.Bids = []Bid{}
		for _, bid := range asksbids.Bids {
			orderbook.Bids = append(
				orderbook.Bids,
				Bid{bid[0], bid[1]},
			)
		}

		if gwa.Config.IsPrint {
			str, _ := Struct2JsonString(orderbook)
			fmt.Println("depth", symbol, clean, str)
		}
	case "kline.update":

		var symbol string
		klines := []*KLine{}

		for _, v := range msg.Result {
			switch t := v.(type) {
			case []interface{}:
				kline := &KLine{
					t[0].(float64),
					t[1].(string),
					t[2].(string),
					t[3].(string),
					t[4].(string),
					t[5].(string),
					t[6].(string),
					t[7].(string),
				}
				if symbol == "" {
					symbol = t[7].(string)
				}
				klines = append(klines, kline)
			}
		}

		if gwa.Config.IsPrint {
			str, _ := Struct2JsonString(klines)
			fmt.Println("kline", symbol, str)
		}
	default:
		obj, err := Struct2JsonString(msg)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println("recv subscribe resp", obj)
	}
	return nil
}

func (gwa *GateWSAgent) handleQuery(msg *QueryResponse) (err error) {
	switch msg.Result.(type) {
	case string:
		fmt.Println("recv query resp", msg.Result)
	case float64:
		fmt.Println("recv query resp", int64(msg.Result.(float64)))
	default:
		obj, err := Struct2JsonString(msg)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println("recv query resp", obj)
	}

	return nil
}

// Subscribe 订阅
func (gwa *GateWSAgent) Subscribe(channel string, params []interface{}, cb ReceivedDataCallback) (err error) {
	gwa.processMutex.Lock()
	defer gwa.processMutex.Unlock()

	err = gwa.sendMsg(channel, params)
	if err != nil {
		return err
	}

	cbs := gwa.SubscribeMap[channel]
	if cbs == nil {
		cbs = []ReceivedDataCallback{}
		gwa.ActiveChannels[channel] = false
	}

	if cb != nil {
		cbs = append(cbs, cb)
		gwa.SubscribeMap[channel] = cbs
	}

	return nil
}

// GzipDecode 解压缩通过 Gzip 压缩过的数据
func (gwa *GateWSAgent) GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func (gwa *GateWSAgent) receive() {
	defer func() {
		a := recover()
		if a != nil {
			log.Printf("Receive End. Recover msg: %+v", a)
			debug.PrintStack()
		}
	}()

	for {
		messageType, message, err := gwa.Conn.ReadMessage()
		if err != nil {
			gwa.ErrorCh <- err
			break
		}

		txtMsg := message
		switch messageType {
		case websocket.TextMessage:
		case websocket.BinaryMessage:
			txtMsg, err = gwa.GzipDecode(message)
		}

		if gwa.Config.IsPrint {
			fmt.Println("recv msg type", messageType, "message", string(txtMsg), "err", err)
		}

		resp, err := loadResponse(txtMsg)
		if err != nil {
			gwa.ErrorCh <- err
			break
		}

		switch resp.(type) {
		case *SubscribeResponse:
			gwa.SubscribeResp <- resp.(*SubscribeResponse)
		case *QueryResponse:
			gwa.QueryResp <- resp.(*QueryResponse)
		}
	}
}
