package gateio

import (
	"fmt"
	"testing"
	"time"
)

func GetGateWss() *GateWSAgent {
	config := &Config{}

	config.PublicEndpoint = "http://data.gateio.co/"
	config.PrivateEndpoint = "https://api.gateio.co/"
	config.WSEndpoint = "wss://ws.gate.io/v3"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = ENGLISH

	wss := &GateWSAgent{}
	err := wss.Start(config)
	if err != nil {
		fmt.Println("okex wss connect failed", err)
	}

	return wss
}

func TestConnection(t *testing.T) {
	t.Log(time.Now().Unix())

	connections := GetGateWss()

	var channel string
	var err error
	params := []interface{}{}

	channel = "server.ping"
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	channel = "server.time"
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	channel = "ticker.subscribe"
	params = []interface{}{"EOS_USDT", "BTC_USDT"}
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	channel = "trades.subscribe"
	params = []interface{}{"EOS_USDT", "BTC_USDT"}
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	channel = "depth.subscribe"
	params = []interface{}{
		[]interface{}{"ETH_USDT", 5, "0.0001"},
		[]interface{}{"BTC_USDT", 5, "0.0001"},
	}
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	channel = "kline.subscribe"
	params = []interface{}{"ETH_USDT", 1800}
	err = connections.Subscribe(channel, params, DefaultPrintData)
	if err != nil {
		t.Error("subscribe failed", err)
	}

	chain := make(chan bool)
	go func() {
		time.Sleep(time.Second * 60)
		chain <- true
	}()
	<-chain
}
