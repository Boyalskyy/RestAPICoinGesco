package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwe1asd1"
	dbname   = "price_BTC"
)

func Create(course float64, name string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//	sqlStatement := `
	//INSERT INTO priceBTC (price,coinname)
	//VALUES ($1,$2)`
	//	_, err = db.Exec(sqlStatement, course, name)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	sqlStatement := `
UPDATE pricebtc
SET price = $2
WHERE coinname = $1;`
	_, err = db.Exec(sqlStatement, name, course)
	if err != nil {
		panic(err)
	}
}
func Get(limit string, name string) []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	var prices []string
	rows, err := db.Query(fmt.Sprintf("select price from priceBTC WHERE coinname='%s' ORDER BY id DESC LIMIT %s ", name, limit))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var price string
		err := rows.Scan(&price)
		if err != nil {
			log.Fatal(err)
		}
		prices = append(prices, price)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return prices
}
