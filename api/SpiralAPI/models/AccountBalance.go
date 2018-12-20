package models

// 子账户结构
type SubAccount struct {
	Currency string `json:"currency"` // 币种
	Balance  string `json:"balance"`  // 结余
	Type     string `json:"type"`     // 类型, trade: 交易余额, frozen: 冻结余额
}

type Balance struct {
	Currency  string  `json:"currency"`
	Available float64 `json:"available,string"`
	Locked    float64 `json:"locked,string"`
	Timestamp int64   `json:"timestamp"`
}

type BalanceReturn struct {
	Data []Balance `json:"data"` // 账户余额
	errorResponse
}
