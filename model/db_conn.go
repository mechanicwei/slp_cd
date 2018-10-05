package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

const connStr string = "postgres://postgres@localhost:5432/slp_cd_db?sslmode=disable"

func GetDBConn() *sqlx.DB {
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
