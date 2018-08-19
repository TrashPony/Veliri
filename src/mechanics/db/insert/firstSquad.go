package insert

import (
	"../../gameObjects/squad"
	"../../../dbConnect"

	"log"
)

func FirstSquad(userID int) (userSquad *squad.Squad)  {

	err, userSquad := AddNewSquad("first", userID)
	if err != nil {
		log.Fatal("create first squad" + err.Error())
	}

	/* 2 танка */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 1, "body", 1, 2)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* 1 мазершип */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 2, "body", 2, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* 3 оружия */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 3, "weapon", 1, 2)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 4, "weapon", 2, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* ammo */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 5, "ammo", 1, 50)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 6, "ammo", 2, 50)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* equip */

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 7, "equip", 1, 2)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 8, "equip", 2, 2)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) " +
		"VALUES ($1, $2, $3, $4, $5)",
		userSquad.ID, 9, "equip", 3, 2)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	return
}
