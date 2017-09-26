package auth


import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetUsers()([]User)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users = make([]User, 0)
	var user User

	for rows.Next() {
		err := rows.Scan(&user.id, &user.name, &user.password, &user.mail)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}


type User struct {
	id int
	name string
	password string
	mail string
}
