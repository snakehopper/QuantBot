package models

type AccountsData struct {
	Currency  string  `json:"currency"`
	Available float64 `json:"available,string"`
	Locked    float64 `json:"locked,string"`
	Timestamp int64   `json:"timestamp"`
}

type AccountsReturn struct {
	Data []AccountsData `json:"data"` // 用户数据
	errorResponse
}
