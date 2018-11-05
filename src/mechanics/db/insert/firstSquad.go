package insert

import (
	"../../../dbConnect"
	"../../gameObjects/squad"

	"log"
)

func FirstSquad(userID int) (userSquad *squad.Squad) {

	err, userSquad := AddNewSquad("first", userID)
	if err != nil {
		log.Fatal("create first squad" + err.Error())
	}

	/* 3 разных танка */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 1, "body", 3, 1, 15)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 2, "body", 4, 1, 25)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 3, "body", 1, 1, 40)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* 1 мазершип */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 4, "body", 2, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* weapon */

	/*firearms*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 5, "weapon", 1, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 6, "weapon", 2, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/*missile*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 7, "weapon", 3, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 8, "weapon", 4, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/*laser*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 9, "weapon", 5, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 10, "weapon", 6, 1, 100)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* ammo */

	/*firearms*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 11, "ammo", 1, 20, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 12, "ammo", 2, 10, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/*missile*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 13, "ammo", 3, 20, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 14, "ammo", 4, 10, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}
	/*laser*/
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 15, "ammo", 5, 10, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		userSquad.ID, 16, "ammo", 6, 20, 1)
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	/* equip */
	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp, target) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userSquad.ID, 17, "equip", 1, 2, 100, "")
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp, target) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userSquad.ID, 18, "equip", 2, 2, 100, "")
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp, target) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userSquad.ID, 19, "equip", 3, 2, 100, "")
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp, target) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userSquad.ID, 20, "equip", 5, 2, 100, "")
	if err != nil {
		log.Fatal("filling first squad" + err.Error())
	}
	return
}
