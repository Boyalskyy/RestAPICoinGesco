package coingesco

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//type CoinGescoBTC struct {
//	BTC struct {
//		USD float64 `json:"usd"`
//	} `json:"bitcoin"`
//}
//type CoinGescoETH struct {
//	ETH struct {
//		USD float64 `json:"usd"`
//	} `json:"ethereum"`
//}

func GetCoinGescoValue(base string, quote string) (float64, error) {
	response, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", base, quote))

	if err != nil {
		return 0, err
	}
	result := make(map[string]map[string]float64)

	// read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	// close response body
	response.Body.Close()
	json.Unmarshal(body, &result)

	fmt.Println(string(body))
	return result[base][quote], nil
}

//func GetCoinGescoValueETH() (float64, string) {
//	responseETH, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	var coinGescoResponseETH CoinGescoETH
//
//	// read response body
//	body, err := ioutil.ReadAll(responseETH.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// close response body
//	responseETH.Body.Close()
//	json.Unmarshal(body, &coinGescoResponseETH)
//
//	fmt.Println(string(body))
//	ETH := "ethereum"
//	return coinGescoResponseETH.ETH.USD, ETH
//}
