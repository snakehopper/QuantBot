package models

type OrderData struct {
	Id             int64       `json:"id"`
	ClientOrderId  string      `json:"clt_ord_id"`
	Symbol         string      `json:"symbol"`
	Side           Side        `json:"side"`
	Price          float64     `json:"price,string"`
	FilledPrice    float64     `json:"filled_price,string"`
	Quantity       float64     `json:"quantity,string"`
	FilledQuantity float64     `json:"filled_quantity,string"`
	Type           OrderType   `json:"type"`
	Status         OrderStatus `json:"status"`
	CreateTime     int64       `json:"create_time"`
	UpdateTime     int64       `json:"update_time"`
}

type OrdersReturn struct {
	Orders []OrderData `json:"orders"`
	errorResponse
}
