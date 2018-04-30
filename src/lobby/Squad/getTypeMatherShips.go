package Squad

import (
	"log"
)

func GetTypeMatherShips() []MatherShip {

	rows, err := db.Query("select * from mother_ship_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var matherShips = make([]MatherShip, 0)
	var matherShip MatherShip

	for rows.Next() {
		err := rows.Scan(&matherShip.Id, &matherShip.Type, &matherShip.HP, &matherShip.Armor, &matherShip.UnitSlots, &matherShip.UnitSlotSize, &matherShip.EquipmentSlots, &matherShip.RangeView)
		if err != nil {
			log.Fatal(err)
		}
		matherShips = append(matherShips, matherShip)
	}

	return matherShips
}

func GetTypeMatherShip(id int) *MatherShip {

	rows, err := db.Query("select * from mother_ship_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var matherShip MatherShip

	for rows.Next() {
		err := rows.Scan(&matherShip.Id, &matherShip.Type, &matherShip.HP, &matherShip.Armor, &matherShip.UnitSlots, &matherShip.UnitSlotSize, &matherShip.EquipmentSlots, &matherShip.RangeView)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &matherShip
}

type MatherShip struct {
	Id             int    `json:"id"`
	Type           string `json:"type"`
	HP             int    `json:"hp"`
	Armor          int    `json:"armor"`
	UnitSlots      int    `json:"unit_slots"`
	UnitSlotSize   int    `json:"unit_slot_size"`
	EquipmentSlots int    `json:"equipment_slots"`
	RangeView      int    `json:"range_view"`
}