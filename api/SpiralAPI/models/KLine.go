package models

import (
	"encoding/json"
	"strconv"
)

type KLineData struct {
	OpenTs        int64
	Open          float64 // 开盘价
	High          float64 // 最高价
	Low           float64 // 最低价
	Close         float64 // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Vol           float64 // 成交额, 即SUM(每一笔成交价 * 该笔的成交数量)
	CloseTs       int64
	RESERVED      string
	NumberOfTrade int64
}

func (r *KLineData) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	if err := json.Unmarshal(bs, &arr); err != nil {
		return err
	}

	var err error

	r.OpenTs = int64(arr[0].(float64))
	if r.Open, err = strconv.ParseFloat(arr[1].(string), 64); err != nil {
		return err
	}
	if r.High, err = strconv.ParseFloat(arr[2].(string), 64); err != nil {
		return err
	}
	if r.Low, err = strconv.ParseFloat(arr[3].(string), 64); err != nil {
		return err
	}
	if r.Close, err = strconv.ParseFloat(arr[4].(string), 64); err != nil {
		return err
	}
	if r.Vol, err = strconv.ParseFloat(arr[5].(string), 64); err != nil {
		return err
	}
	r.CloseTs = int64(arr[6].(float64))
	r.RESERVED = arr[7].(string)
	r.NumberOfTrade = int64(arr[8].(float64))

	return nil
}

type KLineReturn struct {
	//Ts      int64       `json:"ts"`       // 响应生成时间点, 单位毫秒
	Data []KLineData `json:"data"` // KLine数据
	//Ch      string      `json:"ch"`       // 数据所属的Channel, 格式: market.$symbol.kline.$period
	errorResponse
}
