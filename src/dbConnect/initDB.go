package dbConnect

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.100:5432/game?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func GetDBConnect() (*sql.DB)  {

	return DB
}
