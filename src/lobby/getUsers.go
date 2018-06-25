package lobby

import (
	"log"
	"../dbConnect"
)

func GetUsers(query string) User {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name FROM users " + query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user
}
