package model

import (
	"database/sql"
	_ "github.com/lib/pq"

	"log"
)

const connStr string = "postgres://postgres@localhost:5432/slp_cd_db?sslmode=disable"

func GetDBConn() *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
