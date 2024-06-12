package coingesco

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func GetCoinGescoValue(base string, quote string) (float64, error) {
	response, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", base, quote))
	if err != nil {
		return 0, errors.Wrap(err, "API request error")
	}
	result := make(map[string]map[string]float64)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, errors.Wrap(err, "API read error")
	}
	response.Body.Close()
	json.Unmarshal(body, &result)

	fmt.Println(string(body))
	return result[base][quote], nil
}
