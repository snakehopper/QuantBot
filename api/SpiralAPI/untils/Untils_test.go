package untils

import (
	"encoding/json"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/config"
	"testing"
)

func TestComputeHmac256(t *testing.T) {
	secret := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	txt := "GET/api/v1/instrument1518064236"
	want := "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00"

	actual := ComputeHmac256(txt, secret)

	if want != actual {
		t.Errorf("ComputeHmac256 return %v, want %v", actual, want)
	}
}

func TestSignature(t *testing.T) {
	config.SECRET_KEY = "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	txt := "GET/api/v1/instrument1518064236"
	want := "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00"

	actual := signature("GET", "/api/v1/instrument", "1518064236", nil)

	if want != actual {
		t.Errorf("%v signature return %v, want %v", txt, actual, want)
	}
}

func TestHttpPostRequest(t *testing.T) {
	config.SECRET_KEY = "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	var verb = "POST"
	var path = "/api/v1/order"
	var expires = "1518064238"
	//var data = `{"symbol":"BTCUSDT","price":219.0,"clOrdID":"mm_spiral/oemUeQ4CAJZgP3fjHsA","orderQty":98}`
	//var data2 = `{"symbol":"BTCUSDT","price":219,"clOrdID":"mm_spiral/oemUeQ4CAJZgP3fjHsA","orderQty":98}`

	data := struct {
		Symbol      string  `json:"symbol"`
		Price       float64 `json:"price"`
		CliendOrder string  `json:"clOrdID"`
		OrderQty    int64   `json:"orderQty"`
	}{
		"BTCUSDT", 219, "mm_spiral/oemUeQ4CAJZgP3fjHsA", 98,
	}
	bs,_:=json.Marshal(data)

	var want = "883a5e5fabf8a2c634aea1921afedd01dc0841bde128d4d72f3fbef3da50a3b3"
	//var want = "3613e2d7476cff0cf027422669561c62b5135b37b9150d2ab970de0aebfe2e90"
	actual := signature(verb, path, expires, string(bs))
	if want != actual {
		t.Errorf("signature return %v, want %v", actual, want)
	}

}
