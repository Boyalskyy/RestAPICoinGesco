package main

import (
	"RestAPICoinGecko/coingesco"
	"RestAPICoinGecko/db"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func GetAndSave(t time.Time) {
	db.Create(coingesco.GetCoinGescoValueBTC())
	db.Create(coingesco.GetCoinGescoValueETH())
}
func GetCourses(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "5"
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "What cryptocurrency do you need")
		return
	}
	fmt.Fprintf(w, strings.Join(db.Get(limit, name), "\n"))
}
func main() {

	go doEvery(15*time.Second, GetAndSave)
	http.HandleFunc("/course", GetCourses)
	http.ListenAndServe(":8080", nil)

}
