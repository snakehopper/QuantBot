package models

type SymbolsData struct {
	Symbol string `json:"symbol"`
	QuoteAssetName string `json:"quote_asset_name"`
	BaseAssetUnit string `json:"base_asset_unit"`
	BaseAssetName string `json:"base_asset_name"`
	BaseAsset string `json:"base_asset"`
	TickSize float64 `json:"tick_size,string"`
	QuoteAsset string `json:"quote_asset"`
	QuoteAssetUnit string `json:"quote_asset_unit"`
	Active bool `json:"active"`
	MinTrade float64 `json:"min_trade,string"`
	Status string `json:"status"`
}

type SymbolsReturn struct {
	Data    []SymbolsData `json:"data"`   // 交易及精度数据
	errorResponse
}
