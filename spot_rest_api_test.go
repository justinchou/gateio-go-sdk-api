package gateio

import "testing"

func TestGetPairs(t *testing.T) {
	c := NewTestClient()
	ac, _ := c.GetPairs()

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestMarketInfo(t *testing.T) {
	c := NewTestClient()
	ac, err := c.MarketInfo()

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestMarketList(t *testing.T) {
	c := NewTestClient()
	ac, err := c.MarketList()

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestTickers(t *testing.T) {
	c := NewTestClient()
	ac, err := c.Tickers()

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.Ticker("btc_usdt")

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestOrderBooks(t *testing.T) {
	c := NewTestClient()
	ac, err := c.OrderBooks()

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestOrderBook(t *testing.T) {
	c := NewTestClient()
	ac, err := c.OrderBook("btc_usdt")

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}

func TestTradeHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.TradeHistory("btc_usdt")

	if err != nil {
		t.Error("invalid json parse")
	}

	jstr, _ := Struct2JsonString(ac)
	t.Log(jstr)
}
