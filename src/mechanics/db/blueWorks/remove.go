package blueWorks

import (
	"../../../dbConnect"
	"../../gameObjects/blueprints"

	"log"
)

func Remove(work *blueprints.BlueWork) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM created_blueprint WHERE id=$1", work.ID)
	if err != nil {
		log.Fatal("remove blueWork" + err.Error())
	}
}
