package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwe1asd1"
	dbname   = "price_BTC"
)

type PriceQuery struct {
	db *sql.DB
}

func NewPriceQuery() (*PriceQuery, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &PriceQuery{db: db}, nil
}
func (p *PriceQuery) Update(course float64, name string) error {

	sqlStatement := `
UPDATE pricebtc
SET price = $2
WHERE coinname = $1;`
	_, err := p.db.Exec(sqlStatement, name, course)
	if err != nil {
		return err
	}
	return nil
}
func (p *PriceQuery) Get(limit string, name string) ([]string, error) {
	var prices []string
	rows, err := p.db.Query(fmt.Sprintf("select price from priceBTC WHERE coinname='%s' ORDER BY id DESC LIMIT %s ", name, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var price string
		err := rows.Scan(&price)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return prices, nil
}

//func Create(course float64, name string) error {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//	//	sqlStatement := `
//	//INSERT INTO priceBTC (price,coinname)
//	//VALUES ($1,$2)`
//	//	_, err = db.Exec(sqlStatement, course, name)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//}
//	sqlStatement := `
//UPDATE pricebtc
//SET price = $2
//WHERE coinname = $1;`
//	_, err = db.Exec(sqlStatement, name, course)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func Get(limit string, name string) []string {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//	var prices []string
//	rows, err := db.Query(fmt.Sprintf("select price from priceBTC WHERE coinname='%s' ORDER BY id DESC LIMIT %s ", name, limit))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var price string
//		err := rows.Scan(&price)
//		if err != nil {
//			log.Fatal(err)
//		}
//		prices = append(prices, price)
//
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//	return prices
//}
