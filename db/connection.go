package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetDbConnection() *sqlx.DB {
	if db != nil {
		return db
	}
	d, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	db = d
	return db
}
