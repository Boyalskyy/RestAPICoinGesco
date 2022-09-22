package coingesco

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinGescoBTC struct {
	BTC struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}
type CoinGescoETH struct {
	ETH struct {
		USD float64 `json:"usd"`
	} `json:"ethereum"`
}

func GetCoinGescoValueBTC() (float64, string) {
	responseBTC, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")

	if err != nil {
		fmt.Println(err)
	}
	var coinGescoResponseBTC CoinGescoBTC

	// read response body
	body, err := ioutil.ReadAll(responseBTC.Body)
	if err != nil {
		fmt.Println(err)
	}
	// close response body
	responseBTC.Body.Close()
	json.Unmarshal(body, &coinGescoResponseBTC)

	fmt.Println(string(body))
	BTC := "bitcoin"
	return coinGescoResponseBTC.BTC.USD, BTC
}

func GetCoinGescoValueETH() (float64, string) {
	responseETH, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")

	if err != nil {
		fmt.Println(err)
	}
	var coinGescoResponseETH CoinGescoETH

	// read response body
	body, err := ioutil.ReadAll(responseETH.Body)
	if err != nil {
		fmt.Println(err)
	}
	// close response body
	responseETH.Body.Close()
	json.Unmarshal(body, &coinGescoResponseETH)

	fmt.Println(string(body))
	ETH := "ethereum"
	return coinGescoResponseETH.ETH.USD, ETH
}
