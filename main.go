package main

import (
	"RestAPICoinGecko/coingesco"
	"RestAPICoinGecko/db"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"

	"time"
)

func doEvery(d time.Duration, query *db.PriceQuery, f func(t time.Time, query *db.PriceQuery) error) {
	for x := range time.Tick(d) {
		err := f(x, query)
		if err != nil {
			log.Println(err)
		}
	}

}

func GetAndSave(t time.Time, query *db.PriceQuery) error {
	priceBTC, err := coingesco.GetCoinGescoValue("bitcoin", "usd")
	if err != nil {
		return errors.Wrap(err, "get btc value error")
	}
	err = query.Update(priceBTC, "bitcoin")
	if err != nil {
		return errors.Wrap(err, "update btc error")
	}
	priceETH, err := coingesco.GetCoinGescoValue("ethereum", "usd")
	if err != nil {
		return errors.Wrap(err, "get eth value error")
	}
	err = query.Update(priceETH, "ethereum")
	if err != nil {
		return errors.Wrap(err, "update eth error")
	}
	return nil
}

func main() {
	priceQuery, err := db.NewPriceQuery()
	if err != nil {
		log.Fatal(err)
	}
	go doEvery(5*time.Second, priceQuery, GetAndSave)
	http.Handle("/course", Recovery(GetCourses(priceQuery)))
	http.ListenAndServe(":8080", nil)
}
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
func GetCourses(query *db.PriceQuery) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "5"
		}
		name := r.URL.Query().Get("name")
		if name == "" || (name != "bitcoin" && name != "ethereum") {
			w.WriteHeader(400)
			fmt.Fprintf(w, "What cryptocurrency do you need")
			return
		}
		prices, err := query.Get(limit, name)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Fprintf(w, strings.Join(prices, "\n"))

	})
}
