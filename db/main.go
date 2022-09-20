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

func Create(course float64) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
INSERT INTO priceBTC (price)
VALUES ($1)`
	_, err = db.Exec(sqlStatement, course)
	if err != nil {
		panic(err)
	}
}

func Get(limit string) []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	var prices []string
	rows, err := db.Query(fmt.Sprintf("select price from priceBTC ORDER BY id DESC LIMIT %s ", limit))
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
