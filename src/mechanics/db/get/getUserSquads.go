package get

import (
	"../../../dbConnect"
	"../../gameObjects/coordinate"
	"../../gameObjects/squad"
	"../../gameObjects/unit"
	"database/sql"
	"log"
	"strconv"
	"strings"
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

			userSquad.MatherShip.Units = make(map[int]*unit.Slot)

			for _, slot := range userSquad.MatherShip.Body.EquippingIV {
				unitSlot := unit.Slot{}
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

func SquadMatherShip(squadID int) (ship *unit.Unit) {

	rows, err := dbConnect.GetDBConnect().Query(
		"Select id, id_body, hp, q, r, rotate, action, target, queue_attack, use_equip, power, mother_ship, action_point "+
			"FROM squad_units "+
			"WHERE id_squad=$1 AND mother_ship=$2", squadID, true)
	if err != nil {
		log.Fatal("get ship squad " + err.Error())
	}
	defer rows.Close()

	ship = &unit.Unit{}

	var target string

	for rows.Next() {
		var idBody sql.NullInt64

		err = rows.Scan(&ship.ID, &idBody, &ship.HP, &ship.Q, &ship.R, &ship.Rotate, &ship.Action, &target,
			&ship.QueueAttack, &ship.UseEquip, &ship.Power, &ship.MS, &ship.ActionPoints)

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

func SquadUnits(squadID int, slot int) *unit.Unit {

	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT id, id_body, hp, q, r, rotate, action, target, queue_attack, on_map, use_equip, power, mother_ship, action_point "+
			"FROM squad_units "+
			"WHERE id_squad=$1 AND slot=$2 AND mother_ship=$3", squadID, slot, false)
	if err != nil {
		log.Fatal("get units squad " + err.Error())
	}
	defer rows.Close()

	var squadUnit unit.Unit
	var idBody int
	var target string

	for rows.Next() {
		err = rows.Scan(&squadUnit.ID, &idBody, &squadUnit.HP, &squadUnit.Q, &squadUnit.R, &squadUnit.Rotate,
			&squadUnit.Action, &target, &squadUnit.QueueAttack, &squadUnit.OnMap, &squadUnit.UseEquip, &squadUnit.Power,
			&squadUnit.MS, &squadUnit.ActionPoints)
		if err != nil {
			log.Fatal("get units squad " + err.Error())
		}
	}

	squadUnit.Body = Body(idBody)
	BodyEquip(&squadUnit)
	squadUnit.Target = ParseTarget(target)

	if squadUnit.ID != 0 {
		return &squadUnit
	} else {
		return nil
	}
}

func SquadInventory(squadID int) (inventory map[int]*squad.InventorySlot) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
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

		err := rows.Scan(&slot, &inventorySlot.Type, &inventorySlot.ItemID, &inventorySlot.Quantity, &inventorySlot.HP)
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
