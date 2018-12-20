package models

type PlaceRequestParams struct {
	ClientOrderId string    `json:"clt_ord_id,omitempty"`
	Symbol        string    `json:"symbol"`
	Type          OrderType `json:"type"`
	Price         float64   `json:"price,string"`
	Quantity      float64   `json:"quantity,string"`
	Side          Side      `json:"side"`
}

type PlaceData struct {
	Id            int64       `json:"id"`
	ClientOrderId string      `json:"clt_ord_id"`
	UserId        int64       `json:"user_id"`
	Symbol        string      `json:"symbol"`
	Side          Side        `json:"side"`
	Price         float64     `json:"price,string"`
	Quantity      float64     `json:"quantity,string"`
	Type          OrderType   `json:"type"`
	Status        OrderStatus `json:"status"`
	CreateTime    int64       `json:"create_time"`
	UpdateTime    int64       `json:"update_time"`
}

type PlaceReturn struct {
	Order PlaceData `json:"order"`
	errorResponse
}
