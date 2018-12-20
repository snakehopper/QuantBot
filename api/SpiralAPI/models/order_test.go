package models

import (
	"encoding/json"
	"testing"
)

func TestOrderData_Unmarshal(t *testing.T) {
	var res OrderData
	if err := json.Unmarshal([]byte(testOrderDataResponse), &res); err != nil {
		t.Fatal(err)
	}

	if res.Id != 99453 || res.Type != "limit" || res.Status != "partial_filled" {
		t.Error("unmarshal error")
	}
}


func TestOrderReturn_Unmarshal(t *testing.T) {
	var res OrdersReturn
	if err := json.Unmarshal([]byte(testOrderReturnResponse), &res); err != nil {
		t.Fatal(err)
	}

	if len(res.Orders) != 2 {
		t.Error("unmarshal error")
	}
}

const testOrderDataResponse = `{"id":99453,"price":"0.12333","type":"limit","status":"partial_filled"}`
const testOrderReturnResponse = `
{"orders":[
{"id":40013496,"clt_ord_id":"H2N31Q55XT1B7D95","user_id":450,"symbol":"BTCUSDT","side":"ask","price":"3700.00","filled_price":"0.00","quantity":"0.131124","filled_quantity":"0.000000","type":"limit","status":"cancelled","create_time":1545054951981185,"update_time":1545054992007036},
{"id":40013524,"clt_ord_id":"NAHGIFM6QC2NYL3V","user_id":450,"symbol":"BTCUSDT","side":"ask","price":"3555.00","filled_price":"3555.00","quantity":"0.131124","filled_quantity":"0.131124","type":"limit","status":"filled","create_time":1545055003967381,"update_time":1545072590758712}
]}
`