package gateio

import (
	"encoding/json"
	"fmt"
)

// CoinPairs 交易对列表
type CoinPairs []string

// GetPairs 获取所有交易对列表
func (client *Client) GetPairs() (coinpairs *CoinPairs, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/pairs"
	param := ""
	ret := client.Request(method, url, param)

	coinpairs = &CoinPairs{}
	err = json.Unmarshal([]byte(ret), coinpairs)
	if err != nil {
		return nil, err
	}

	if client.Config.IsPrint {
		fmt.Println("coinpairs", coinpairs)
	}

	return coinpairs, nil
}

type _MarketInfoResp struct {
	Result string                   `json:"result"`
	Pairs  []map[string]*MarketInfo `json:"pairs"`
}

// MarketInfo 交易对基本信息
type MarketInfo struct {
	DecimalPlaces int     `json:"decimal_places"`
	MinAmount     float64 `json:"min_amount"`
	MinAmountA    float64 `json:"min_amount_a"`
	MinAmountB    float64 `json:"min_amount_b"`
	Fee           float64 `json:"fee"`
	TradeDisabled int     `json:"trade_disabled"`
	Symbol        string  `json:"symbol"`
}

// MarketInfo Market Info
func (client *Client) MarketInfo() (marketInfos *[]*MarketInfo, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/marketinfo"
	param := ""
	ret := client.Request(method, url, param)

	marketInfoResp := &_MarketInfoResp{}
	err = json.Unmarshal([]byte(ret), marketInfoResp)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	marketInfoList := []*MarketInfo{}
	for _, info := range marketInfoResp.Pairs {
		for k, v := range info {
			v.Symbol = k
			marketInfoList = append(marketInfoList, v)
		}
	}

	return &marketInfoList, nil
}

type _MarketListResp struct {
	Result string        `json:"result"`
	Data   *[]MarketItem `json:"data"`
}

// MarketItem 交易对信息
type MarketItem struct {
	No          int         `json:"no"`
	Symbol      string      `json:"symbol"`
	Name        string      `json:"name"`
	NameEn      string      `json:"name_en"`
	NameCn      string      `json:"name_cn"`
	Pair        string      `json:"pair"`
	Rate        string      `json:"rate"`
	VolA        string      `json:"vol_a"`
	VolB        string      `json:"vol_b"`
	CurrA       string      `json:"curr_a"`
	CurrB       string      `json:"curr_b"`
	CurrSuffix  string      `json:"curr_suffix"`
	RatePercent string      `json:"rate_percent"`
	Trend       string      `json:"trend"`
	Supply      int         `json:"supply"`
	Marketcap   interface{} `json:"marketcap"`
	Lq          string      `json:"lq"`
}

// MarketList Market Details
func (client *Client) MarketList() (marketList *[]MarketItem, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/marketlist"
	param := ""
	ret := client.Request(method, url, param)

	marketListResp := &_MarketListResp{}
	err = json.Unmarshal([]byte(ret), marketListResp)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	marketList = marketListResp.Data

	return marketList, nil
}

// TickerInfo Ticker 数据
type TickerInfo struct {
	Result        string `json:"result"`
	Last          string `json:"last"`
	LowestAsk     string `json:"lowestAsk"`
	HighestBid    string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
	High24Hr      string `json:"high24hr"`
	Low24Hr       string `json:"low24hr"`
	Elapsed       string `json:"elapsed"`
}

// Tickers tickers
func (client *Client) Tickers() (tickerInfos *map[string]*TickerInfo, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/tickers"
	param := ""
	ret := client.Request(method, url, param)

	tickerInfos = &map[string]*TickerInfo{}
	err = json.Unmarshal([]byte(ret), tickerInfos)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	return tickerInfos, nil
}

// Ticker ticker
func (client *Client) Ticker(symbol string) (tickerInfo *TickerInfo, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/ticker" + "/" + symbol
	param := ""
	ret := client.Request(method, url, param)
	fmt.Println(ret)

	tickerInfo = &TickerInfo{}
	err = json.Unmarshal([]byte(ret), tickerInfo)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	return tickerInfo, nil
}

type _OrderBookItem struct {
	Result string      `json:"result"`
	Asks   [][2]string `json:"asks"`
	Bids   [][2]string `json:"bids"`
}

// Ask 卖
type Ask struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

// Bid 买
type Bid struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

// OrderBook 挂单列表
type OrderBook struct {
	Asks    []Ask  `json:"asks"`
	Bids    []Bid  `json:"bids"`
	Elapsed string `json:"elapsed"`
}

// OrderBooks Depth
func (client *Client) OrderBooks() (retItems map[string]*OrderBook, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/orderBooks"
	param := ""
	ret := client.Request(method, url, param)

	orderbooksResp := map[string]*_OrderBookItem{}
	err = json.Unmarshal([]byte(ret), &orderbooksResp)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	orderbooks := map[string]*OrderBook{}
	for symbol, orderbook := range orderbooksResp {
		orderbooks[symbol] = &OrderBook{}

		orderbooks[symbol].Asks = []Ask{}
		for _, ask := range orderbook.Asks {
			orderbooks[symbol].Asks = append(
				[]Ask{
					Ask{ask[0], ask[1]},
				},
				orderbooks[symbol].Asks...,
			)
		}

		orderbooks[symbol].Bids = []Bid{}
		for _, bid := range orderbook.Bids {
			orderbooks[symbol].Bids = append(
				orderbooks[symbol].Bids,
				Bid{bid[0], bid[1]},
			)
		}
	}

	return orderbooks, nil
}

type _OrderbookResp struct {
	Asks    [][2]string `json:"asks"`
	Bids    [][2]string `json:"bids"`
	Elapsed string      `json:"elapsed"`
}

// OrderBook Depth of pair
func (client *Client) OrderBook(symbol string) (orderbooks *OrderBook, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/orderBook/" + symbol
	param := ""
	ret := client.Request(method, url, param)

	orderbooksResp := &_OrderbookResp{}
	err = json.Unmarshal([]byte(ret), orderbooksResp)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	orderbooks = &OrderBook{}

	orderbooks.Asks = []Ask{}
	for _, ask := range orderbooksResp.Asks {
		orderbooks.Asks = append(
			[]Ask{
				Ask{ask[0], ask[1]},
			},
			orderbooks.Asks...,
		)
	}

	orderbooks.Bids = []Bid{}
	for _, bid := range orderbooksResp.Bids {
		orderbooks.Bids = append(
			orderbooks.Bids,
			Bid{bid[0], bid[1]},
		)
	}

	orderbooks.Elapsed = orderbooksResp.Elapsed

	return orderbooks, nil
}

// TradeHistory 成交历史记录
type TradeHistory struct {
	TradeID   string `json:"tradeID"`
	Total     string `json:"total"`
	Date      string `json:"date"`
	Rate      string `json:"rate"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

type _TradeHistoryResp struct {
	Data    []TradeHistory `json:"data"`
	Result  string         `json:"result"`
	Elapsed string         `json:"elapsed"`
}

// TradeHistory Trade History
func (client *Client) TradeHistory(symbol string) (history *[]TradeHistory, err error) {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/tradeHistory/" + symbol
	param := ""
	ret := client.Request(method, url, param)

	resp := &_TradeHistoryResp{}
	err = json.Unmarshal([]byte(ret), resp)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("parse json err", err)
		}
		return nil, err
	}

	history = &resp.Data

	return history, nil
}

// Balances Get account fund balances
func (client *Client) Balances() string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "api2/1/private/balances"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// DepositAddress get deposit address
func (client *Client) DepositAddress(currency string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/depositAddress"
	param := "currency=" + currency
	ret := client.Request(method, url, param)
	return ret
}

// DepositsWithdrawals get deposit withdrawal history
func (client *Client) DepositsWithdrawals(start string, end string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/depositsWithdrawals"
	param := "start=" + start + "&end=" + end
	ret := client.Request(method, url, param)
	return ret
}

// Buy Place order buy
func (client *Client) Buy(currencyPair string, rate string, amount string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/buy"
	param := "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	ret := client.Request(method, url, param)
	return ret
}

// Sell Place order sell
func (client *Client) Sell(currencyPair string, rate string, amount string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/sell"
	param := "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	ret := client.Request(method, url, param)
	return ret
}

// CancelOrder Cancel order
func (client *Client) CancelOrder(orderNumber string, currencyPair string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/cancelOrder"
	param := "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	ret := client.Request(method, url, param)
	return ret
}

// CancelAllOrders Cancel all orders
func (client *Client) CancelAllOrders(types string, currencyPair string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/cancelAllOrders"
	param := "type=" + types + "&currencyPair=" + currencyPair
	ret := client.Request(method, url, param)
	return ret
}

// GetOrder Get order status
func (client *Client) GetOrder(orderNumber string, currencyPair string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/getOrder"
	param := "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	ret := client.Request(method, url, param)
	return ret
}

// OpenOrders Get my open order list
func (client *Client) OpenOrders() string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/openOrders"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// MyTradeHistory 获取我的24小时内成交记录
func (client *Client) MyTradeHistory(currencyPair string, orderNumber string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/tradeHistory"
	param := "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	ret := client.Request(method, url, param)
	return ret
}

// Withdraw 提现API
func (client *Client) Withdraw(currency string, amount string, address string) string {
	method := "POST"
	url := client.Config.PrivateEndpoint + "/api2/1/private/withdraw"
	param := "currency=" + currency + "&amount=" + amount + "&address=" + address
	ret := client.Request(method, url, param)
	return ret
}
