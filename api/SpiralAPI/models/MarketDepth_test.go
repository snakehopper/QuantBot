package models

import (
	"encoding/json"
	"testing"
)

func TestMarketDepth_UnmarshalJSON(t *testing.T) {
	var res MarketDepthReturn

	if err:=json.Unmarshal([]byte(testMarketDepthResponse), &res);err!=nil{
		t.Fatal(err)
	}

	if res.Symbol != "BTCUSDT" {
		t.Error("unmarshal error")
	}
}

const testMarketDepthResponse = `
{"symbol":"BTCUSDT","last_update_id":59590231,
"data":[["3722.98","0.153068","bid"],["3725.30","0.278047","bid"],
["3727.92","0.505367","bid"],["3730.25","0.205977","bid"],
["3744.61","0.373208","bid"],["3751.36","0.061996","bid"],
["3751.62","0.083582","bid"],["3752.07","0.112671","bid"],
["3761.89","0.045852","bid"],["3771.45","0.033924","bid"],
["3781.94","0.033928","ask"],["3782.14","0.264589","ask"],
["3782.58","0.061984","ask"],["3782.70","0.083784","ask"],
["3783.94","0.045838","ask"],["3795.84","0.206226","ask"],
["3800.15","0.278449","ask"],["3802.22","0.112671","ask"],
["3806.91","0.152115","ask"],["3826.01","0.505317","ask"],
["3832.26","0.373208","ask"]]}`
