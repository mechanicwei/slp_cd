package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"log"
)

var dbUrl = viper.GetString("pg_url")

func GetDBConn() *sqlx.DB {
	db, err := sqlx.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
