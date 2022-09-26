package main

import (
	"RestAPICoinGecko/coingesco"
	"RestAPICoinGecko/db"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"

	"time"
)

func doEvery(d time.Duration, query *db.PriceQuery, f func(t time.Time, query *db.PriceQuery) error) {
	for x := range time.Tick(d) {
		err := f(x, query)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func GetAndSave(t time.Time, query *db.PriceQuery) error {
	priceBTC, err := coingesco.GetCoinGescoValue("bitcoin", "usd")
	if err != nil {
		return err
	}
	err = query.Update(priceBTC, "bitcoin")
	if err != nil {
		return err
	}
	priceETH, err := coingesco.GetCoinGescoValue("ethereum", "usd")
	if err != nil {
		return err
	}
	err = query.Update(priceETH, "ethereum")
	if err != nil {
		return err
	}
	return nil
}

//	func GetCourses(w http.ResponseWriter, r *http.Request) {
//		limit := r.URL.Query().Get("limit")
//		if limit == "" {
//			limit = "5"
//		}
//		name := r.URL.Query().Get("name")
//		if name == "" {
//			fmt.Fprintf(w, "What cryptocurrency do you need")
//
//		}
//		prices, _ := query.Get(limit, name)
//		fmt.Fprintf(w, strings.Join(prices, "\n"))
//
// }
func main() {
	priceQuery, err := db.NewPriceQuery()
	if err != nil {
		log.Fatal(err)
	}
	go doEvery(15*time.Second, priceQuery, GetAndSave)
	http.Handle("/course", GetCourses(priceQuery))
	http.ListenAndServe(":8080", nil)
}
func GetCourses(query *db.PriceQuery) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//ctx:=r.Context()
		//ctx = context.WithValue(ctx,"key",query)
		//priceQ:=r.Context().Value("key").(*db.PriceQuery)
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
