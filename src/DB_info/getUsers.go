package DB_info

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetUsers(query string)(User)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id, name, mail FROM users " + query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Mail)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user
}