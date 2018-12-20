package models

import (
	"encoding/json"
	"strconv"
)

type MarketDepth struct {
	Price float64
	Size  float64
	Side  Side
}

func (r *MarketDepth) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	if err := json.Unmarshal(bs, &arr); err != nil {
		return err
	}

	var err error
	if r.Price, err = strconv.ParseFloat(arr[0].(string), 64); err != nil {
		return err
	}
	if r.Size, err = strconv.ParseFloat(arr[1].(string), 64); err != nil {
		return err
	}
	r.Side = Side(arr[2].(string))

	return nil
}

type MarketDepthReturn struct {
	//Status  string      `json:"status"` // 请求状态, ok或者error
	//Ts      int64       `json:"ts"`     // 响应生成时间点, 单位: 毫秒
	Symbol       string        `json:"symbol"`
	LastUpdateId int64         `json:"last_update_id"`
	Data         []MarketDepth `json:"data"` // Depth数据
	//Ch      string      `json:"ch"`     //  数据所属的Channel, 格式: market.$symbol.depth.$type
	errorResponse
}
