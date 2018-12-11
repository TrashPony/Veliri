package base

import (
	"../../../dbConnect"
	"log"
)

func UserIntoBase(userID, baseID int) {
	// удаляем все записи пользователя (вообще их быть не должно)
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM base_users WHERE user_id=$1", userID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO base_users (user_id, base_id) VALUES ($1, $2)", userID, baseID)
	if err != nil {
		log.Fatal("add user to base" + err.Error())
	}
}
