package DetailUnit

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
