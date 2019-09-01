package squad

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"log"
)

func AddNewSquad(name string, userID int) (err error, newSquad *squad.Squad) {
	id := 0
	err = dbConnect.GetDBConnect().QueryRow("INSERT INTO squads (name, active, id_user, in_game, id_base) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		name, true, userID, false, 1).Scan(&id)

	if err != nil {
		log.Fatal("add new squad " + err.Error())
		return err, nil
	}

	newSquad = &squad.Squad{ID: id, Active: true, Name: name, InGame: false}

	return nil, newSquad
}

func FirstSquad(userID, baseID int) {
	///* 3 разных танка */
	//_, err := dbConnect.GetDBConnect().Exec("INSERT INTO base_storage (base_id, user_id, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6, $7)",
	//	baseID, userID, 1, "body", 3, 1, 15)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 2, "body", 4, 1, 25)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 3, "body", 1, 1, 40)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///* 1 мазершип */
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 4, "body", 2, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///* weapon */
	//
	///*firearms*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 5, "weapon", 1, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 6, "weapon", 2, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///*missile*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 7, "weapon", 3, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 8, "weapon", 4, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///*laser*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 9, "weapon", 5, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 10, "weapon", 6, 1, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///* ammo */
	//
	///*firearms*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 11, "ammo", 1, 20, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 12, "ammo", 2, 10, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///*missile*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 13, "ammo", 3, 20, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 14, "ammo", 4, 10, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	///*laser*/
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 15, "ammo", 5, 10, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 16, "ammo", 6, 20, 1)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	///* equip */
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 17, "equip", 1, 2, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 18, "equip", 2, 2, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 19, "equip", 3, 2, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
	//	"VALUES ($1, $2, $3, $4, $5, $6)",
	//	userSquad.ID, 20, "equip", 5, 2, 100)
	//if err != nil {
	//	log.Fatal("filling first squad" + err.Error())
	//}
	//return
}
