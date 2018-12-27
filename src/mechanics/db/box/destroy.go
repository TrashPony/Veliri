package box

import (
	"../../../dbConnect"
	"../../gameObjects/box"
	"log"
)

func Destroy(gameBox *box.Box) {
	// удаляем весь инвентарь
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM box_storage WHERE id_box=$1", gameBox.ID)
	if err != nil {
		log.Fatal("delete all items from box storage" + err.Error())
	}

	// удаляем сам ящик
	_, err = dbConnect.GetDBConnect().Exec("DELETE FROM box_in_map WHERE id=$1", gameBox.ID)
	if err != nil {
		log.Fatal("delete box in map" + err.Error())
	}
}
