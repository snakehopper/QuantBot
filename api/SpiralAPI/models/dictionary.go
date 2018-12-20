package models

type Side string
type OrderType string
type OrderStatus string
type Currency string

const (
	BidSide Side = "bid"
	AskSide Side = "ask"

	LimitOrderType  OrderType = "limit"
	MarketOrderType OrderType = "market"

	Submitted       OrderStatus = "submitted"
	Accepted        OrderStatus = "accepted"
	Waiting         OrderStatus = "waiting"
	Rejected        OrderStatus = "rejected"
	PartialFilled   OrderStatus = "partial_filled"
	Filled          OrderStatus = "filled"
	CancelRequested OrderStatus = "cancel_requested"
	CancelRejected  OrderStatus = "cancel_rejected"
	Cancelled       OrderStatus = "cancelled"
	ModifyRequested OrderStatus = "modify_requested"
	ModifyRejected  OrderStatus = "modify_rejected"
	Modified        OrderStatus = "modified"
	Unknown         OrderStatus = "unknown"

	USDT Currency = "USDT"
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	LTC  Currency = "LTC"
	BCH  Currency = "BCH"
)
