package dbPlayer

import (
	"../../../dbConnect"
	"log"
)

func AddCash(userID, appendCash int)  {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE users "+
		"SET credits = credits + $2  "+
		"WHERE id = $1",
		userID, appendCash)
	if err != nil {
		log.Fatal("update user " + err.Error())
	}
}
