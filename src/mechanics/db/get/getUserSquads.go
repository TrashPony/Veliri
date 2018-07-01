package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/squad"
	"../../gameObjects/unit"
	"../../gameObjects/matherShip"
)

func UserSquads(userID int) (squads []*squad.Squad, err error) {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, active, in_game FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*squad.Squad, 0)

	for rows.Next() {
		var userSquad squad.Squad

		err := rows.Scan(&userSquad.ID, &userSquad.Name, &userSquad.Active, &userSquad.InGame)
		if err != nil {
			log.Fatal(err)
		}

		userSquad.MatherShip = SquadMatherShip(userSquad.ID)

		if userSquad.MatherShip != nil {
			userSquad.MatherShip.Units = SquadUnits(userSquad.ID)
		}

		userSquad.Inventory = SquadInventory(userSquad.ID)

		squads = append(squads, &userSquad)
	}

	return
}

func SquadMatherShip(squadID int) (ship *matherShip.MatherShip) {

	rows, err := dbConnect.GetDBConnect().Query(
		"Select id, id_body, hp, x, y, rotate, action, target, queue_attack "+
			"FROM squad_mother_ship "+
			"WHERE id_squad=$1", squadID)
	if err != nil {
		log.Fatal("get ship squad " + err.Error())
	}
	defer rows.Close()

	ship = &matherShip.MatherShip{}

	for rows.Next() {
		var idBody int

		err = rows.Scan(&ship.ID, &idBody, &ship.HP, &ship.X, &ship.Y, &ship.Rotate, &ship.Action, &ship.Target, ship.QueueAttack)
		if err != nil {
			log.Fatal("scan get ship squad " + err.Error())
		}

		ship.Body = Body(idBody)
		BodyEquip(ship)
	}

	return
}

func SquadUnits(squadID int) (units map[int]*unit.Unit) {

	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT id, id_body, hp, x, y, rotate, action, target, queue_attack, slot "+
			"FROM squad_units "+
			"WHERE id_squad=$1", squadID)
	if err != nil {
		log.Fatal("get units squad " + err.Error())
	}
	defer rows.Close()

	units = make(map[int]*unit.Unit)

	for rows.Next() {
		var squadUnit unit.Unit

		var idBody int
		var slot int

		err = rows.Scan(&squadUnit.ID, &idBody, &squadUnit.HP, &squadUnit.X, &squadUnit.Y, &squadUnit.Rotate, &squadUnit.Action, &squadUnit.Target, squadUnit.QueueAttack, slot)
		if err != nil {
			log.Fatal(err)
		}

		squadUnit.Body = Body(idBody)

		BodyEquip(&squadUnit)

		units[slot] = &squadUnit
	}

	return
}

func SquadInventory(squadID int) (inventory map[int]*squad.InventorySlot) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity "+
		"FROM squad_inventory "+
		"WHERE id_squad = $1", squadID)
	if err != nil {
		log.Fatal("get inventory squad " + err.Error())
	}
	defer rows.Close()

	inventory = make(map[int]*squad.InventorySlot)

	for rows.Next() {

		var inventorySlot = squad.InventorySlot{}
		var TypeItem string
		var idItem int
		var slot int

		err := rows.Scan(&slot, &TypeItem, &idItem, &inventorySlot.Quantity)
		if err != nil {
			log.Fatal("get inventory squad " + err.Error())
		}

		if TypeItem == "weapon" {
			inventorySlot.Item = Weapon(idItem)
			inventory[slot] = &inventorySlot
		}

		if TypeItem == "ammo" {
			inventorySlot.Item = Ammo(idItem)
			inventory[slot] = &inventorySlot
		}

		if TypeItem == "equip" {
			inventorySlot.Item = TypeEquip(idItem)
			inventory[slot] = &inventorySlot
		}

		if TypeItem == "body" {
			inventorySlot.Item = Body(idItem)
			inventory[slot] = &inventorySlot
		}
	}

	return
}
