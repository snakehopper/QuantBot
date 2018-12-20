package config

// API KEY
var (
	ACCESS_KEY string = ""
	SECRET_KEY string = ""
	ACCOUNT_ID string = ""
)

// API请求地址, 不要带最后的/
const (
	API_URL string = "https://api.spiral.exchange"
	WEBSOCKER_URL string = "wss://ws.spiral.exchange"

	WS_PING = ""
)

type WsEvent string
const (
	Ping WsEvent = "ping"
	Authenticate = "authenticate"

)
