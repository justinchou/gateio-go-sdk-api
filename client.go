package gateio

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client GateIO API 客户端
type Client struct {
	Config     Config
	HttpClient *http.Client
}

// NewClient Get a http client
func NewClient(config Config) *Client {
	var client Client

	client.Config = config
	timeout := config.TimeoutSecond

	if timeout <= 0 {
		timeout = 30
	}

	client.HttpClient = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	return &client
}

// GetSign 签名
func (client *Client) GetSign(params string) string {
	key := []byte(client.Config.SecretKey)
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(params))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// Request 发送 http 请求
func (client *Client) Request(method string, url string, param string) string {

	if client.Config.IsPrint {
		fmt.Println(method, url, param)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("request created failed", err)
		}
	}

	sign := client.GetSign(param)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", client.Config.ApiKey)
	req.Header.Set("sign", sign)

	resp, err := client.HttpClient.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if client.Config.IsPrint {
			fmt.Println("request read response failed", err)
		}
	}

	return string(body)
}
