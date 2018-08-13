package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/squad"
	"../../gameObjects/unit"
	"../../gameObjects/matherShip"
	"../../gameObjects/coordinate"
	"strings"
	"strconv"
	"database/sql"
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

		if userSquad.MatherShip != nil && userSquad.MatherShip.Body != nil {

			userSquad.MatherShip.Units = make(map[int]*matherShip.UnitSlot)

			for _, slot := range userSquad.MatherShip.Body.EquippingIV {
				unitSlot := matherShip.UnitSlot{}
				unitSlot.Unit = SquadUnits(userSquad.ID, slot.Number)
				unitSlot.NumberSlot = slot.Number

				userSquad.MatherShip.Units[slot.Number] = &unitSlot
			}
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

	var target string

	for rows.Next() {
		var idBody sql.NullInt64

		err = rows.Scan(&ship.ID, &idBody, &ship.HP, &ship.X, &ship.Y, &ship.Rotate, &ship.Action, &target, &ship.QueueAttack)

		if err != nil {
			log.Fatal("scan get ship squad " + err.Error())
		}

		ship.Target = ParseTarget(target)

		if idBody.Valid {
			ship.Body = Body(int(idBody.Int64))
			BodyEquip(ship)
		} else {
			ship.Body = nil
		}
	}

	return
}

func SquadUnits(squadID int, slot int) (*unit.Unit) {

	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT id, id_body, hp, x, y, rotate, action, target, queue_attack "+
			"FROM squad_units "+
			"WHERE id_squad=$1 AND slot=$2", squadID, slot)
	if err != nil {
		log.Fatal("get units squad " + err.Error())
	}
	defer rows.Close()

	var squadUnit unit.Unit

	for rows.Next() {
		var idBody int

		err = rows.Scan(&squadUnit.ID, &idBody, &squadUnit.HP, &squadUnit.X, &squadUnit.Y, &squadUnit.Rotate, &squadUnit.Action, &squadUnit.Target, squadUnit.QueueAttack)
		if err != nil {
			log.Fatal(err)
		}

		squadUnit.Body = Body(idBody)

		BodyEquip(&squadUnit)
	}

	if squadUnit.ID != 0 {
		return &squadUnit
	} else {
		return nil
	}
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
		var slot int

		err := rows.Scan(&slot, &inventorySlot.Type, &inventorySlot.ItemID, &inventorySlot.Quantity)
		if err != nil {
			log.Fatal("get inventory squad " + err.Error())
		}

		if inventorySlot.Type == "weapon" {
			inventorySlot.Item = Weapon(inventorySlot.ItemID)
			inventory[slot] = &inventorySlot
		}

		if inventorySlot.Type == "ammo" {
			inventorySlot.Item = Ammo(inventorySlot.ItemID)
			inventory[slot] = &inventorySlot
		}

		if inventorySlot.Type == "equip" {
			inventorySlot.Item = TypeEquip(inventorySlot.ItemID)
			inventory[slot] = &inventorySlot
		}

		if inventorySlot.Type == "body" {
			inventorySlot.Item = Body(inventorySlot.ItemID)
			inventory[slot] = &inventorySlot
		}
	}

	return
}

func ParseTarget(targetKey string) *coordinate.Coordinate {
	targetCell := strings.Split(targetKey, ":")

	if len(targetCell) > 1 { // устанавливаем таргет если он есть
		x, ok := strconv.Atoi(targetCell[0])
		y, ok := strconv.Atoi(targetCell[1])
		if ok == nil {
			target := coordinate.Coordinate{X: x, Y: y}
			return &target
		} else {
			return nil
		}
	} else {
		return nil
	}
}
