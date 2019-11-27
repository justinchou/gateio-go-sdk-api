package gateio

// Config GateIO API Config
type Config struct {
	// Rest api endpoint url. eg: http://data.gateio.co/
	PublicEndpoint string

	// Rest api endpoint url. eg: https://api.gateio.co/
	PrivateEndpoint string

	// Rest websocket api endpoint url. eg:  wss://ws.gate.io/v3/
	WSEndpoint string

	// The user's api key provided by GateIO.
	ApiKey string
	// The user's secret key provided by GateIO. The secret key used to sign your request data.
	SecretKey string
	// The Passphrase will be provided by you to further secure your API access.
	Passphrase string
	// Http request timeout.
	TimeoutSecond int
	// Whether to print API information
	IsPrint bool
	// Internationalization @see file: constants.go
	I18n string
}
