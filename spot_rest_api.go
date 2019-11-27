package gateio

// GetPairs all support pairs
func (client *Client) GetPairs() string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/pairs"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// MarketInfo Market Info
func (client *Client) MarketInfo() string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/marketinfo"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// MarketList Market Details
func (client *Client) MarketList() string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/marketlist"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// Tickers tickers
func (client *Client) Tickers() string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/tickers"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// Ticker ticker
func (client *Client) Ticker(ticker string) string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/ticker" + "/" + ticker
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// OrderBooks Depth
func (client *Client) OrderBooks() string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/orderBooks"
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// OrderBook Depth of pair
func (client *Client) OrderBook(params string) string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/orderBook/" + params
	param := ""
	ret := client.Request(method, url, param)
	return ret
}

// TradeHistory Trade History
func (client *Client) TradeHistory(params string) string {
	method := "GET"
	url := client.Config.PublicEndpoint + "api2/1/tradeHistory/" + params
	param := ""
	ret := client.Request(method, url, param)
	return ret
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
