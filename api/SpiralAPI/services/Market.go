package services

import (
	"encoding/json"
	"fmt"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/models"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/untils"
)

// 批量操作的API下个版本再封装

//------------------------------------------------------------------------------------------
// 交易API

// 获取K线数据
// strSymbol: 交易对, btcusdt, bccbtc......
// strPeriod: K线类型, 1, 5, 15......
// nSize: 获取数量, [1-500]
// return: KLineReturn 对象
func GetKLine(strSymbol, strPeriod string, nSize int64) (r models.KLineReturn, err error) {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["limit"] = fmt.Sprint(nSize)

	url := "/api/v1/klines"

	jsonKLineReturn := untils.HttpGetRequest(url, mapParams)
	if err = json.Unmarshal([]byte(jsonKLineReturn), &r); err != nil {
		fmt.Println(jsonKLineReturn, err)
	}

	return
}

// 获取聚合行情
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TickReturn对象
func GetTicker(strSymbol string) (r models.TickerReturn, err error) {
	panic("use GetMarketDepth instead")

	return
}

// 获取交易深度信息
// strSymbol: 交易对, btcusdt, bccbtc......
// strType: Depth类型, step0、step1......stpe5 (合并深度0-5, 0时不合并)
// return: MarketDepthReturn对象
func GetMarketDepth(strSymbol, strType string) (r models.MarketDepthReturn, err error) {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["limit"] = strType

	url := "/api/v1/orderbook"

	jsonMarketDepthReturn := untils.HttpGetRequest(url, mapParams)
	if err = json.Unmarshal([]byte(jsonMarketDepthReturn), &r); err != nil {
		fmt.Println(jsonMarketDepthReturn, err)
	}

	return
}

// 获取交易细节信息
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TradeDetailReturn对象
func GetTrades(strSymbol string) (r models.TradesReturn, err error) {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	//mapParams["start"] = strSymbol
	//mapParams["reverse"] = strSymbol
	//mapParams["start_time"] = strSymbol
	//mapParams["end_time"] = strSymbol
	mapParams["count"] = "1000"

	url := "/api/v1/trades"

	jsonTradeDetailReturn := untils.HttpGetRequest(url, mapParams)
	if err = json.Unmarshal([]byte(jsonTradeDetailReturn), &r); err != nil {
		fmt.Println(jsonTradeDetailReturn, err)
	}

	return
}

// 获取Market Detail 24小时成交量数据
// strSymbol: 交易对, btcusdt, bccbtc......
// return: MarketDetailReturn对象
func GetMarketDetail(strSymbol string) (r models.MarketDetailReturn, err error) {
	panic("implement this")

	return
}

//------------------------------------------------------------------------------------------
// 公共API

// 查询系统支持的所有交易及精度
// return: SymbolsReturn对象
func GetSymbols() (r models.SymbolsReturn, err error) {
	url := "/api/v1/products"

	jsonCurrencysReturn := untils.HttpGetRequest(url, nil)
	if err = json.Unmarshal([]byte(jsonCurrencysReturn), &r); err != nil {
		fmt.Println(jsonCurrencysReturn, err)
	}

	return
}

// 查询系统支持的所有币种
// return: CurrencysReturn对象
func GetCurrencys() (r models.CurrencysReturn, err error) {
	url := "/api/v1/currencies"

	jsonCurrencysReturn := untils.HttpGetRequest(url, nil)
	if err = json.Unmarshal([]byte(jsonCurrencysReturn), &r); err != nil {
		fmt.Println(jsonCurrencysReturn, err)
	}

	return
}

// 查询系统当前时间戳
// return: TimestampReturn对象
func GetTimestamp() (r models.TimestampReturn, err error) {
	panic("implement this")
}

//------------------------------------------------------------------------------------------
// 用户资产API

// 查询当前用户的所有账户, 根据包含的私钥查询
// return: AccountsReturn对象
func GetAccounts() (r models.BalanceReturn, err error) {
	return GetAccountBalance("")
}

// 根据账户ID查询账户余额
// nAccountID: 账户ID, 不知道的话可以通过GetAccounts()获取, 可以只现货账户, C2C账户, 期货账户
// return: BalanceReturn对象
func GetAccountBalance(curr string) (r models.BalanceReturn, err error) {
	mapParams := make(map[string]string)
	mapParams["currency"] = curr

	url := "/api/v1/wallet/balances"

	jsonBanlanceReturn := untils.HttpGetRequest(url, mapParams)
	if err = json.Unmarshal([]byte(jsonBanlanceReturn), &r); err != nil {
		fmt.Println(jsonBanlanceReturn, err)
	}

	return
}

//------------------------------------------------------------------------------------------
// 交易API

// 下单
// params: 下单信息
// return: PlaceReturn对象
func Place(params models.PlaceRequestParams) (r models.PlaceReturn, err error) {
	strRequest := "/api/v1/order"

	jsonPlaceReturn := untils.HttpPostRequest(strRequest, nil, &params)
	if err = json.Unmarshal([]byte(jsonPlaceReturn), &r); err != nil {
		fmt.Println(jsonPlaceReturn, err)
	}

	return
}

func MarketBuy(amount float64, symbol string) (models.PlaceReturn, error) {
	p := models.PlaceRequestParams{
		ClientOrderId: "",
		Symbol:        symbol,
		Type:          models.MarketOrderType,
		Quantity:      amount,
		Side:          models.BidSide,
	}
	return Place(p)
}

func MarketSell(amount float64, symbol string) (models.PlaceReturn, error) {
	p := models.PlaceRequestParams{
		ClientOrderId: "",
		Symbol:        symbol,
		Type:          models.MarketOrderType,
		Quantity:      amount,
		Side:          models.AskSide,
	}
	return Place(p)
}

func LimitBuy(amount, price float64, symbol string) (models.PlaceReturn, error) {
	p := models.PlaceRequestParams{
		ClientOrderId: "",
		Symbol:        symbol,
		Type:          models.LimitOrderType,
		Price:         price,
		Quantity:      amount,
		Side:          models.BidSide,
	}
	return Place(p)
}

func LimitSell(amount, price float64, symbol string) (models.PlaceReturn, error) {
	p := models.PlaceRequestParams{
		ClientOrderId: "",
		Symbol:        symbol,
		Type:          models.LimitOrderType,
		Price:         price,
		Quantity:      amount,
		Side:          models.AskSide,
	}
	return Place(p)
}

// 申请撤销一个订单请求
// strOrderID: 订单ID
// return: PlaceReturn对象
func SubmitCancel(strOrderID string) (r models.PlaceReturn, err error) {
	params := map[string]string{
		"order_id": strOrderID,
	}
	jsonPlaceReturn := untils.HttpDeleteRequest("/api/v1/order", params)
	err = json.Unmarshal([]byte(jsonPlaceReturn), &r)
	if err != nil {
		fmt.Println(jsonPlaceReturn, err)
	}
	return
}

func SubmitCancelAll() (r models.PlaceReturn, err error) {
	jsonPlaceReturn := untils.HttpDeleteRequest("/api/v1/order/all", nil)
	if err = json.Unmarshal([]byte(jsonPlaceReturn), &r); err != nil {
		fmt.Println(jsonPlaceReturn, err)
	}

	return
}

func GetMyTrades(strSymbol string) (r models.TradesReturn, err error) {
	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	//mapParams["start"] = strSymbol
	mapParams["count"] = "100"
	//mapParams["reverse"] = "100"
	//mapParams["start_time"] = "100"
	//mapParams["end_time"] = "100"

	strRequest := "/api/v1/myTrades"

	jsonOrdersReturn := untils.HttpGetRequest(strRequest, mapParams)
	if err = json.Unmarshal([]byte(jsonOrdersReturn), &r); err != nil {
		fmt.Println(jsonOrdersReturn, err)
	}

	return
}

func GetOrders(strSymbol string, side *models.Side) (r models.OrdersReturn, err error) {
	mapParams := make(map[string]string)
	//mapParams["clt_ord_id"] = "clt_ord_id"
	mapParams["symbol"] = strSymbol
	mapParams["count"] = "1000" //1~1000
	if side != nil {
		mapParams["side"] = string(*side)
	}

	//if openOnly != nil {
	//	mapParams["filter"] = fmt.Sprintf("open:%v", openOnly)
	//}
	//mapParams["reverse"] = "100"
	//mapParams["start_time"] = "100"
	//mapParams["end_time"] = "100"

	strRequest := "/api/v1/order"

	jsonOrdersReturn := untils.HttpGetRequest(strRequest, mapParams)

	if err = json.Unmarshal([]byte(jsonOrdersReturn), &r); err != nil {
		fmt.Println(jsonOrdersReturn, err)
	}

	return
}
