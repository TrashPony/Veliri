package dbConnect

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"time"
)

var DB *sql.DB

func init() {

	//Генератор случайных чисел обычно нужно рандомизировать перед использованием, иначе, он, действительно,
	// будет выдавать одну и ту же последовательность.
	rand.Seed(time.Now().UnixNano())

	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.100:5432/game?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func GetDBConnect() *sql.DB {
	return DB
}
