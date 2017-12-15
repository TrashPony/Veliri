package lobby

import (
	"log"
)

func GetUsers(query string) User {

	rows, err := db.Query("Select id, name, mail, password FROM users " + query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Mail, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user
}
