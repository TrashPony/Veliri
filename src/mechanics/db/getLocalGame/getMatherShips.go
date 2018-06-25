package getLocalGame

import (
	"../../gameObjects/matherShip"
	"../../../dbConnect"
	"log"
)

func MatherShips(idGame int) (matherShips map[int]map[int]*matherShip.MatherShip) {

	rows, err := dbConnect.GetDBConnect().Query("Select type.type, users.name, ship.x, ship.y, "+
		"type.hp, type.armor, type.range_view "+
		"FROM action_mother_ship as ship, mother_ship_type as type, users "+
		"WHERE users.id=ship.id_user AND type.id=ship.id_type AND id_game=$1", idGame)
	if err != nil {
		println("get game matherShip")
		log.Fatal(err)
	}
	defer rows.Close()

	matherShips = make(map[int]map[int]*matherShip.MatherShip)

	for rows.Next() {
		var ship matherShip.MatherShip
		err := rows.Scan(&ship.Type, &ship.Owner, &ship.X, &ship.Y, &ship.HP, &ship.Armor, &ship.RangeView)
		if err != nil {
			log.Fatal(err)
		}

		if matherShips[ship.X] != nil {
			matherShips[ship.X][ship.Y] = &ship
		} else {
			matherShips[ship.X] = make(map[int]*matherShip.MatherShip)
			matherShips[ship.X][ship.Y] = &ship
		}
	}

	return
}
