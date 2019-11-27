package gateio

// GetDefaultConfig 获取测试用配置文件
func GetDefaultConfig() *Config {
	var config Config

	config.PublicEndpoint = "http://data.gateio.co/"
	config.PrivateEndpoint = "https://api.gateio.co/"
	config.WSEndpoint = "wss://ws.gate.io/v3"
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = ENGLISH

	// set your own ApiKey, SecretKey, Passphrase here
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""

	return &config
}

// NewTestClient 获取测试用 Client
func NewTestClient() *Client {
	// Set GateIO API's config
	client := NewClient(*GetDefaultConfig())

	return client
}
