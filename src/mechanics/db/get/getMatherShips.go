package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/matherShip"
)

func TypeMatherShips() []matherShip.MatherShip {

	rows, err := dbConnect.GetDBConnect().Query("select * from mother_ship_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var matherShips = make([]matherShip.MatherShip, 0)
	var gameMatherShip matherShip.MatherShip

	for rows.Next() {
		err := rows.Scan(&gameMatherShip.ID, &gameMatherShip.Type, &gameMatherShip.HP, &gameMatherShip.Armor,
			&gameMatherShip.UnitSlots, &gameMatherShip.UnitSlotSize, &gameMatherShip.EquipmentSlots, &gameMatherShip.RangeView)
		if err != nil {
			log.Fatal(err)
		}
		matherShips = append(matherShips, gameMatherShip)
	}

	return matherShips
}

func TypeMatherShip(id int) *matherShip.MatherShip {

	rows, err := dbConnect.GetDBConnect().Query("select * from mother_ship_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var gameMatherShip matherShip.MatherShip

	for rows.Next() {
		err := rows.Scan(&gameMatherShip.ID, &gameMatherShip.Type, &gameMatherShip.HP,
			&gameMatherShip.Armor, &gameMatherShip.UnitSlots, &gameMatherShip.UnitSlotSize,
			&gameMatherShip.EquipmentSlots, &gameMatherShip.RangeView)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &gameMatherShip
}
