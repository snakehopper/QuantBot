package api

import (
	"fmt"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/models"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/services"
	"math"
	"strings"
	"time"

	"github.com/miaolz123/conver"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/config"
	"github.com/phonegapX/QuantBot/constant"
	"github.com/phonegapX/QuantBot/model"
)

// Spiral the exchange struct of spiral.com
type Spiral struct {
	stockTypeMap     map[string]string
	tradeTypeMap     map[string]string
	recordsPeriodMap map[string]string
	minAmountMap     map[string]float64
	records          map[string][]Record
	logger           model.Logger
	option           Option

	limit     float64
	lastSleep int64
	lastTimes int64
}

// NewSpiral create an exchange struct of Spiral.com
func NewSpiral(opt Option) Exchange {
	config.ACCESS_KEY = opt.AccessKey
	config.SECRET_KEY = opt.SecretKey
	return &Spiral{
		stockTypeMap: map[string]string{
			"BTC/USDT": "BTCUSDT",
			"ETH/USDT": "ETHUSDT",
			"ETH/BTC":  "ETHBTC",
			"BCH/USDT": "BCHUSDT",
			"LTC/USDT": "LTCUSDT",
			"BCH/BTC":  "BCHBTC",
			"LTC/BTC":  "LTCBTC",
		},
		tradeTypeMap: map[string]string{
			"bid": constant.TradeTypeBuy,
			"ask": constant.TradeTypeSell,
		},
		recordsPeriodMap: map[string]string{
			"M":   "1",
			"M5":  "5",
			"M15": "15",
			"M30": "30",
			"H":   "60",
			"D":   "1440",
			"W":   "10080",
		},
		minAmountMap: map[string]float64{
			"BTC/USDT": 0.000001,
			"ETH/USDT": 0.00001,
			"ETH/BTC":  0.001,
			"BCH/USDT": 0.00001,
			"LTC/USDT": 0.00001,
			"BCH/BTC":  0.001,
			"LTC/BTC":  0.001,
		},
		records: make(map[string][]Record),
		logger:  model.Logger{TraderID: opt.TraderID, ExchangeType: opt.Type},
		option:  opt,

		limit:     10.0,
		lastSleep: time.Now().UnixNano(),
	}
}

// Log print something to console
func (e *Spiral) Log(msgs ...interface{}) {
	e.logger.Log(constant.INFO, "", 0.0, 0.0, msgs...)
}

// GetType get the type of this exchange
func (e *Spiral) GetType() string {
	return e.option.Type
}

// GetName get the name of this exchange
func (e *Spiral) GetName() string {
	return e.option.Name
}

// SetLimit set the limit calls amount per second of this exchange
func (e *Spiral) SetLimit(times interface{}) float64 {
	e.limit = conver.Float64Must(times)
	return e.limit
}

// AutoSleep auto sleep to achieve the limit calls amount per second of this exchange
func (e *Spiral) AutoSleep() {
	now := time.Now().UnixNano()
	interval := 1e+9/e.limit*conver.Float64Must(e.lastTimes) - conver.Float64Must(now-e.lastSleep)
	if interval > 0.0 {
		time.Sleep(time.Duration(conver.Int64Must(interval)))
	}
	e.lastTimes = 0
	e.lastSleep = now
}

// GetMinAmount get the min trade amonut of this exchange
func (e *Spiral) GetMinAmount(stock string) float64 {
	return e.minAmountMap[stock]
}

// GetAccount get the account detail of this exchange
func (e *Spiral) GetAccount() interface{} {
	accs, err := services.GetAccounts()
	if err != nil {
		e.logger.Log(constant.ERROR, "", 0.0, 0.0, "GetAccount() error, ", err)
		return false
	}

	e.logger.Log(constant.INFO, "", 0, 0, "GetAccount execute", accs)

	result := make(map[string]float64)
	for _, acc := range accs.Data {
		key := acc.Currency
		result[key] = acc.Available
		result["Frozen"+key] = acc.Locked
	}

	e.logger.Log(constant.INFO, "", 0, 0, "GetAccount response", result)

	return result
}

// Trade place an order
func (e *Spiral) Trade(tradeType string, stockType string, _price, _amount interface{}, msgs ...interface{}) interface{} {
	stockType = strings.ToUpper(stockType)
	tradeType = strings.ToUpper(tradeType)
	price := conver.Float64Must(_price)
	amount := conver.Float64Must(_amount)
	if amount < e.minAmountMap[stockType] {
		e.logger.Log(constant.ERROR, stockType, price, amount, "min trade amount is", e.minAmountMap[stockType])
	} else if mod := math.Mod(amount, e.minAmountMap[stockType]); mod < e.minAmountMap[stockType] {
		e.logger.Log(constant.INFO, stockType, price, amount, "adjust trade amount to match minimum trade amount", amount-mod)
		amount = amount - mod
	}
	if _, ok := e.stockTypeMap[stockType]; !ok {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "Trade() error, unrecognized stockType: ", stockType)
		return false
	}
	switch tradeType {
	case constant.TradeTypeBuy:
		return e.buy(stockType, price, amount, msgs...)
	case constant.TradeTypeSell:
		return e.sell(stockType, price, amount, msgs...)
	default:
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "Trade() error, unrecognized tradeType: ", tradeType)
		return false
	}
}

func (e *Spiral) buy(stockType string, price, amount float64, msgs ...interface{}) interface{} {
	var result models.PlaceReturn
	var err error
	if price == 0 {
		result, err = services.MarketBuy(amount, e.stockTypeMap[stockType])
	} else {
		result, err = services.LimitBuy(amount, price, e.stockTypeMap[stockType])
	}
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, price, amount, "Buy() error, ", err)
		return false
	}
	orderId := result.Order.Id
	if orderId <= 0 {
		e.logger.Log(constant.ERROR, stockType, price, amount, "Buy() errorCode, ", result.ErrorCode, result.Message)
		return false
	}
	e.logger.Log(constant.BUY, stockType, price, amount, msgs...)
	return fmt.Sprint(orderId)
}

func (e *Spiral) sell(stockType string, price, amount float64, msgs ...interface{}) interface{} {
	var result models.PlaceReturn
	var err error
	if price == 0 {
		result, err = services.MarketSell(amount, e.stockTypeMap[stockType])
	} else {
		result, err = services.LimitSell(amount, price, e.stockTypeMap[stockType])
	}
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, price, amount, "Sell() error, ", err)
		return false
	}
	orderId := result.Order.Id
	if orderId <= 0 {
		e.logger.Log(constant.ERROR, stockType, price, amount, "sell() errorCode, ", result.ErrorCode, result.Message)
		return false
	}
	e.logger.Log(constant.SELL, stockType, price, amount, msgs...)
	return fmt.Sprint(orderId)
}

// GetOrder get details of an order
func (e *Spiral) GetOrder(stockType, id string) interface{} {
	fmt.Println("GetOrder", stockType, id)
	res := e.GetOrders(stockType)
	switch res.(type) {
	case bool:
		return false
	default:
		ods := res.([]Order)
		fmt.Println("orders", ods)
		for _, od := range ods {
			if od.ID == id {
				return od
			}
		}
	}
	panic("unreachable")
}

// GetOrders get all unfilled orders
func (e *Spiral) GetOrders(stockType string) interface{} {
	stockType = strings.ToUpper(stockType)
	if _, ok := e.stockTypeMap[stockType]; !ok {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetOrders() error, unrecognized stockType: ", stockType)
		return false
	}
	result, err := services.GetOrders(e.stockTypeMap[stockType], nil)
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetOrders() error, ", err)
		return false
	}
	if result.ErrorCode != 0 {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetOrders() errorCode, ", result.ErrorCode, result.Message)
		return false
	}
	orders := []Order{}
	count := len(result.Orders)
	for i := 0; i < count; i++ {
		orders = append(orders, Order{
			ID:         fmt.Sprint(result.Orders[i].Id),
			Price:      result.Orders[i].FilledPrice,
			Amount:     result.Orders[i].Quantity,
			DealAmount: result.Orders[i].FilledQuantity,
			TradeType:  e.tradeTypeMap[string(result.Orders[i].Side)],
			StockType:  stockType,
		})
	}
	return orders
}

// GetTrades get all filled orders recently
func (e *Spiral) GetTrades(stockType string) interface{} {
	stockType = strings.ToUpper(stockType)
	if _, ok := e.stockTypeMap[stockType]; !ok {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetTrades() error, unrecognized stockType: ", stockType)
		return false
	}
	result, err := services.GetTrades(e.stockTypeMap[stockType])
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetTrades() error, ", err)
		return false
	}
	if result.ErrorCode != 0 {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, "GetTrades() errorCode, ", result.ErrorCode, result.Message)
		return false
	}
	orders := []Order{}
	count := len(result.Trades)
	for i := 0; i < count; i++ {
		r := result.Trades[i]
		orders = append(orders, Order{
			ID:         fmt.Sprint(r.ID),
			Price:      r.Price,
			Amount:     r.Quantity,
			DealAmount: r.Quantity,
			TradeType:  e.tradeTypeMap[string(r.Side)],
			StockType:  stockType,
		})
	}
	return orders
}

// CancelOrder cancel an order
func (e *Spiral) CancelOrder(order Order) bool {
	result, err := services.SubmitCancel(order.ID)
	if err != nil {
		e.logger.Log(constant.ERROR, "", 0.0, 0.0, "CancelOrder() error, ", err)
		return false
	}
	if result.ErrorCode != 0 {
		e.logger.Log(constant.ERROR, "", 0.0, 0.0, "CancelOrder() with errorCode, ", result.ErrorCode, result.Message)
		return false
	}
	e.logger.Log(constant.CANCEL, order.StockType, order.Price, order.Amount-order.DealAmount, order)
	return true
}

// getTicker get market ticker & depth
func (e *Spiral) getTicker(stockType string, sizes ...interface{}) (ticker Ticker, err error) {
	stockType = strings.ToUpper(stockType)
	if _, ok := e.stockTypeMap[stockType]; !ok {
		err = fmt.Errorf("GetTicker() error, unrecognized stockType: %+v", stockType)
		return
	}
	result, err := services.GetMarketDepth(e.stockTypeMap[stockType], "0")
	if err != nil {
		err = fmt.Errorf("GetTicker() error, %+v", err)
		return
	}
	if result.ErrorCode != 0 {
		err = fmt.Errorf("GetTicker() error, %v %v", result.ErrorCode, result.Message)
		return
	}
	count := len(result.Data)
	for i := 0; i < count; i++ {
		row := result.Data[i]
		switch row.Side {
		case models.BidSide:
			ticker.Bids = append(ticker.Bids, OrderBook{
				Price:  row.Price,
				Amount: row.Size,
			})
		case models.AskSide:
			ticker.Asks = append(ticker.Asks, OrderBook{
				Price:  row.Price,
				Amount: row.Size,
			})
		}
	}
	//rearrange bid array to make arr[0] is best deal price
	reverseSlice(ticker.Bids)

	if len(ticker.Bids) < 1 || len(ticker.Asks) < 1 {
		err = fmt.Errorf("GetTicker() error, can not get enough Bids or Asks")
		return
	}
	ticker.Buy = ticker.Bids[0].Price
	ticker.Sell = ticker.Asks[0].Price
	ticker.Mid = (ticker.Buy + ticker.Sell) / 2
	return
}

func reverseSlice(a []OrderBook) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// GetTicker get market ticker & depth
func (e *Spiral) GetTicker(stockType string, sizes ...interface{}) interface{} {
	ticker, err := e.getTicker(stockType, sizes...)
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, err)
		return false
	}
	return ticker
}

// GetRecords get candlestick data
func (e *Spiral) GetRecords(stockType, period string, sizes ...interface{}) interface{} {
	var count int64
	if len(sizes) == 1 {
		count = sizes[0].(int64)
	}
	res, err := e.getCandleStick(stockType, period, count)
	if err != nil {
		e.logger.Log(constant.ERROR, stockType, 0.0, 0.0, err)
		return false
	}
	return res
}

func (e *Spiral) getCandleStick(stockType, period string, sizes int64) (records []Record, err error) {
	stockType = strings.ToUpper(stockType)
	if _, ok := e.stockTypeMap[stockType]; !ok {
		err = fmt.Errorf("getCandleStick() error, unrecognized stockType: %+v", stockType)
		return
	}
	var nSize = int64(500)
	if sizes > 0 {
		nSize = sizes
	}
	result, err := services.GetKLine(e.stockTypeMap[stockType], e.recordsPeriodMap[period], nSize)
	if err != nil {
		err = fmt.Errorf("getCandleStick() error, %+v", err)
		return
	}
	if result.ErrorCode != 0 {
		err = fmt.Errorf("getCandleStick() error, %v %v", result.ErrorCode, result.Message)
		return
	}
	count := len(result.Data)
	for i := 0; i < count; i++ {
		row := result.Data[i]
		records = append(records, Record{
			Time:   row.OpenTs,
			Open:   row.Open,
			High:   row.High,
			Low:    row.Low,
			Close:  row.Close,
			Volume: row.Vol,
		})
	}
	return
}
