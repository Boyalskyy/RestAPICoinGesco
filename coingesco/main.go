package coingesco

import (
	"RestAPICoinGecko/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinGesco struct {
	BTC struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func GetCoinGescoValue() float64 {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")

	if err != nil {
		fmt.Println(err)
	}
	var coinGescoResponse CoinGesco

	// read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// close response body
	response.Body.Close()
	json.Unmarshal(body, &coinGescoResponse)
	db.Create(coinGescoResponse.BTC.USD)
	fmt.Println(string(body))
	return coinGescoResponse.BTC.USD
}
