package db

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "qwe1asd1"
	dbname   = "postgres"
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
		return nil, errors.Wrap(err, "database open error")
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
		return errors.Wrap(err, "database update error")
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
			return nil, errors.Wrap(err, "error getting value from database")
		}
		prices = append(prices, price)

	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error getting value from database")
	}
	return prices, nil
}
