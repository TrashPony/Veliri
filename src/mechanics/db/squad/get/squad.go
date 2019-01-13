package get

import (
	"../../../../dbConnect"
	"../../../factories/gameTypes"
	"../../../gameObjects/coordinate"
	inv "../../../gameObjects/inventory"
	"../../../gameObjects/squad"
	"../../../gameObjects/unit"
	"database/sql"
	"log"
	"strconv"
	"strings"
)

func UserSquads(userID int) (squads []*squad.Squad, err error) {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, active, in_game, q, r, id_map, id_base FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*squad.Squad, 0)

	for rows.Next() {
		var userSquad squad.Squad

		err := rows.Scan(&userSquad.ID, &userSquad.Name, &userSquad.Active, &userSquad.InGame, &userSquad.Q,
			&userSquad.R, &userSquad.MapID, &userSquad.BaseID)
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
		"SELECT "+
			"id, "+
			"id_body, "+
			"hp, "+
			"q, "+
			"r, "+
			"rotate, "+
			"target, "+
			"power, mother_ship, "+
			"action_point, "+
			"on_map, "+
			"defend, "+
			"move "+
			""+
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

		err = rows.Scan(
			&ship.ID,
			&idBody,
			&ship.HP,
			&ship.Q,
			&ship.R,
			&ship.Rotate,
			&target,
			&ship.Power,
			&ship.MS,
			&ship.ActionPoints,
			&ship.OnMap,
			&ship.Defend,
			&ship.Move,
		)

		if err != nil {
			log.Fatal("scan get ship squad " + err.Error())
		}

		ship.Target = ParseTarget(target)

		if idBody.Valid {
			ship.Body, _ = gameTypes.Bodies.GetByID(int(idBody.Int64))
			BodyEquip(ship)

			SquadThorium(ship, squadID)

			ship.CalculateParams()

		} else {
			ship.Body = nil
		}
	}

	return
}

func SquadThorium(ship *unit.Unit, squadID int) {
	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		"slot, "+
		"thorium "+
		""+
		"FROM squad_thorium_slots "+
		"WHERE id_squad = $1", squadID)
	if err != nil {
		log.Fatal("get thorium squad" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var slotNumber, thorium int

		err = rows.Scan(&slotNumber, &thorium)
		if err != nil {
			log.Fatal("get thorium squad " + err.Error())
		}

		for _, slot := range ship.Body.ThoriumSlots {
			if slot.Number == slotNumber {
				ship.Body.ThoriumSlots[slotNumber].Count = thorium
			}
		}
	}
}

func SquadUnits(squadID int, slot int) *unit.Unit {

	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT "+
			"id, "+
			"id_body, "+
			"hp, "+
			"q, "+
			"r, "+
			"rotate, "+
			"target, "+
			"on_map, "+
			"power, "+
			"mother_ship, "+
			"action_point, "+
			"defend,"+
			"move "+
			" "+
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
		err = rows.Scan(
			&squadUnit.ID,
			&idBody,
			&squadUnit.HP,
			&squadUnit.Q,
			&squadUnit.R,
			&squadUnit.Rotate,
			&target,
			&squadUnit.OnMap,
			&squadUnit.Power,
			&squadUnit.MS,
			&squadUnit.ActionPoints,
			&squadUnit.Defend,
			&squadUnit.Move,
		)
		if err != nil {
			log.Fatal("get units squad " + err.Error())
		}
	}

	squadUnit.Body, _ = gameTypes.Bodies.GetByID(idBody)
	BodyEquip(&squadUnit)
	squadUnit.Target = ParseTarget(target)

	squadUnit.CalculateParams()

	if squadUnit.ID != 0 {
		return &squadUnit
	} else {
		return nil
	}
}

func SquadInventory(squadID int) (inventory inv.Inventory) {

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		"slot, "+
		"item_type, "+
		"item_id, "+
		"quantity, "+
		"hp "+
		""+
		"FROM squad_inventory "+
		"WHERE id_squad = $1", squadID)
	if err != nil {
		log.Fatal("get inventory squad " + err.Error())
	}
	defer rows.Close()

	inventory.Slots = make(map[int]*inv.Slot)
	inventory.SetSlotsSize(40)
	inventory.FillInventory(rows)

	return
}

func ParseTarget(targetKey string) *coordinate.Coordinate {
	targetCell := strings.Split(targetKey, ":")

	if len(targetCell) > 1 { // устанавливаем таргет если он есть
		q, ok := strconv.Atoi(targetCell[0])
		r, ok := strconv.Atoi(targetCell[1])
		if ok == nil {
			target := coordinate.Coordinate{Q: q, R: r}
			return &target
		} else {
			return nil
		}
	} else {
		return nil
	}
}
