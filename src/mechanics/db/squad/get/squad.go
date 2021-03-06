package get

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"log"
)

func UserSquads(userID int) (squads []*squad.Squad, err error) {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, active, in_game, id_base FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*squad.Squad, 0)

	for rows.Next() {
		var userSquad squad.Squad

		err := rows.Scan(&userSquad.ID, &userSquad.Name, &userSquad.Active, &userSquad.InGame, &userSquad.BaseID)
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
			"x, "+
			"y, "+
			"rotate, "+
			"power, mother_ship, "+
			"action_point, "+
			"on_map, "+
			"defend, "+
			"move,"+
			"body_color_1,"+
			"body_color_2,"+
			"weapon_color_1,"+
			"weapon_color_2,"+
			"body_texture,"+
			"weapon_texture,"+
			"id_map "+
			""+
			"FROM squad_units "+
			"WHERE id_squad=$1 AND mother_ship=$2", squadID, true)
	if err != nil {
		log.Fatal("get ship squad " + err.Error())
	}
	defer rows.Close()

	ship = &unit.Unit{}

	for rows.Next() {
		var idBody sql.NullInt64

		err = rows.Scan(
			&ship.ID,
			&idBody,
			&ship.HP,
			&ship.X,
			&ship.Y,
			&ship.Rotate,
			&ship.Power,
			&ship.MS,
			&ship.ActionPoints,
			&ship.OnMap,
			&ship.Defend,
			&ship.Move,
			&ship.BodyColor1,
			&ship.BodyColor2,
			&ship.WeaponColor1,
			&ship.WeaponColor2,
			&ship.BodyTexture,
			&ship.WeaponTexture,
			&ship.MapID,
		)

		if err != nil {
			log.Fatal("scan get ship squad " + err.Error())
		}

		if idBody.Valid {
			ship.Body, _ = gameTypes.Bodies.GetByID(int(idBody.Int64))

			BodyEquip(ship)
			SquadThorium(ship, squadID)
			ship.Inventory = SquadInventory(ship.ID)

			ship.CalculateParams()
			ship.GunRotate = ship.Rotate
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
			"x, "+
			"y, "+
			"rotate, "+
			"on_map, "+
			"power, "+
			"mother_ship, "+
			"action_point, "+
			"defend, "+
			"move, "+
			"id_game, "+
			"body_color_1,"+
			"body_color_2,"+
			"weapon_color_1,"+
			"weapon_color_2,"+
			"body_texture,"+
			"weapon_texture,"+
			"id_map "+
			" "+
			"FROM squad_units "+
			"WHERE id_squad=$1 AND slot=$2 AND mother_ship=$3", squadID, slot, false)
	if err != nil {
		log.Fatal("get units squad " + err.Error())
	}
	defer rows.Close()

	var squadUnit unit.Unit
	var idBody int

	for rows.Next() {
		err = rows.Scan(
			&squadUnit.ID,
			&idBody,
			&squadUnit.HP,
			&squadUnit.X,
			&squadUnit.Y,
			&squadUnit.Rotate,
			&squadUnit.OnMap,
			&squadUnit.Power,
			&squadUnit.MS,
			&squadUnit.ActionPoints,
			&squadUnit.Defend,
			&squadUnit.Move,
			&squadUnit.GameID,
			&squadUnit.BodyColor1,
			&squadUnit.BodyColor2,
			&squadUnit.WeaponColor1,
			&squadUnit.WeaponColor2,
			&squadUnit.BodyTexture,
			&squadUnit.WeaponTexture,
			&squadUnit.MapID,
		)
		if err != nil {
			log.Fatal("get units squad " + err.Error())
		}
	}

	squadUnit.Body, _ = gameTypes.Bodies.GetByID(idBody)
	BodyEquip(&squadUnit)

	squadUnit.Inventory = SquadInventory(squadUnit.ID)
	squadUnit.GunRotate = squadUnit.Rotate

	squadUnit.CalculateParams()

	if squadUnit.ID != 0 {
		return &squadUnit
	} else {
		return nil
	}
}

func SquadInventory(unitID int) *inv.Inventory {
	var inventory inv.Inventory

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		"slot, "+
		"item_type, "+
		"item_id, "+
		"quantity, "+
		"hp,"+
		"place_user_id "+
		""+
		"FROM squad_units_inventory "+
		"WHERE id_unit = $1", unitID)
	if err != nil {
		log.Fatal("get inventory unit in squad " + err.Error())
	}
	defer rows.Close()

	inventory.Slots = make(map[int]*inv.Slot)
	inventory.SetSlotsSize(999)
	inventory.FillInventory(rows)

	return &inventory
}
