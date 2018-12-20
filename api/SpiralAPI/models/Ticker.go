package models

type Ticker struct {
	ID     int64     `json:"id"`     // K线ID
	Amount float64   `json:"amount,string"` // 成交量
	Count  int64     `json:"count"`  // 成交笔数
	Open   float64   `json:"open,string"`   // 开盘价
	Close  float64   `json:"close,string"`  // 收盘价
	Low    float64   `json:"low,string"`    // 最低价
	High   float64   `json:"high,string"`   // 最高价
	Vol    float64   `json:"vol,string"`    // 成交额
	Bid    []float64 `json:"bid,string"`    // [买1价, 买1量]
	Ask    []float64 `json:"ask,string"`    // [卖1价, 卖1量]
}

type TickerReturn struct {
	Status  string `json:"status"` // 请求处理结果
	Ts      int64  `json:"ts"`     // 响应生成时间点
	Tick    Ticker `json:"tick"`   // K线聚合数据
	Ch      string `json:"ch"`     // 数据所属的Channel, 格式: market.$symbol.detail.merged
	errorResponse
}
