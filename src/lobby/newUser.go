package lobby

import "log"

func CreateUser(login, email, password string)  {
	var err error

	_, err = db.Exec("INSERT INTO users (name, password, mail) VALUES ($1, $2, $3)", login, password, email)

	if err != nil {
		log.Fatal(err)
	}
}
