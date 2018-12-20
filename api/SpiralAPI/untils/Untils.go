package untils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/phonegapX/QuantBot/api/SpiralAPI/config"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"net/url"

	"time"

	//"golang.org/x/net/proxy"
)

var g struct {
	cli *gentleman.Client
}

func init() {
	g.cli = gentleman.New().URL(config.API_URL)
}

func signature(verb, path, expired string, data interface{}) string {
	if data == nil {
		data = ""
	}
	txt := fmt.Sprintf("%v%v%v%v", verb, path, expired, data)
	fmt.Println(txt)
	return ComputeHmac256(txt, config.SECRET_KEY)
}

func HttpGetRequest(path string, mapParams map[string]string) string {
	ul, err := url.Parse(path)
	if err != nil {
		return err.Error()
	}

	var val = ul.Query()
	for k, v := range mapParams {
		val.Set(k, v)
	}
	ul.RawQuery = val.Encode()

	g.cli.Use(headers.Set("api-key", config.ACCESS_KEY))

	expired := fmt.Sprint(time.Now().Add(5 * time.Second).Unix())
	g.cli.Use(headers.Set("api-expires", expired))

	sign := signature("GET", ul.String(), expired, nil)
	g.cli.Use(headers.Set("api-signature", sign))
	//fmt.Println("path:", ul.String())
	//fmt.Println("signature", sign)

	resp, err := g.cli.Path(path).Get().SetQueryParams(mapParams).Send()
	if err != nil {
		//fmt.Println(err)
		return err.Error()
	}
	//fmt.Println(resp.RawRequest.URL.String())
	return resp.String()
}

func HttpPostRequest(path string, data interface{}) string {
	g.cli.Use(headers.Set("api-key", config.ACCESS_KEY))

	expired := fmt.Sprint(time.Now().Add(5 * time.Second).Unix())
	g.cli.Use(headers.Set("api-expires", expired))

	bs, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	sign := signature("POST", path, expired, string(bs))
	g.cli.Use(headers.Set("api-signature", sign))

	resp, err := g.cli.Path(path).Post().JSON(bs).Send()
	if err != nil {
		return err.Error()
	}

	return resp.String()
}

func HttpDeleteRequest(path string) string {
	g.cli.Use(headers.Set("api-key", config.ACCESS_KEY))

	expired := fmt.Sprint(time.Now().Add(5 * time.Second).Unix())
	g.cli.Use(headers.Set("api-expires", expired))

	sign := signature("DELETE", path, expired, nil)
	g.cli.Use(headers.Set("api-signature", sign))

	resp, err := g.cli.Path(path).Delete().Send()
	if err != nil {
		return err.Error()
	}
	return resp.String()
}

// HMAC SHA256加密
// strMessage: 需要加密的信息
// strSecret: 密钥
// return: HEX 编码的密文
func ComputeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))

	return hex.EncodeToString(h.Sum(nil))
}
