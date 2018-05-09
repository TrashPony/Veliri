package game

import "log"

type MatherShip struct {
	Type      string `json:"type"`
	Owner     string `json:"owner"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	HP        int    `json:"hp"`
	Armor     int    `json:"armor"`
	RangeView int    `json:"range_view"`
}

func (matherShip *MatherShip) getX() int {
	return matherShip.X
}

func (matherShip *MatherShip) getY() int {
	return matherShip.Y
}

func (matherShip *MatherShip) getWatchZone() int {
	return matherShip.RangeView
}

func (matherShip *MatherShip) getOwnerUser() string {
	return matherShip.Owner
}

func GetMatherShips(idGame int) (matherShips map[int]map[int]*MatherShip) {

	rows, err := db.Query("Select type.type, users.name, ship.x, ship.y, "+
		"type.hp, type.armor, type.range_view "+
		"FROM action_mother_ship as ship, mother_ship_type as type, users "+
		"WHERE users.id=ship.id_user AND type.id=ship.id_type AND id_game=$1", idGame)
	if err != nil {
		println("get game matherShip")
		log.Fatal(err)
	}

	matherShips = make(map[int]map[int]*MatherShip)

	for rows.Next() {
		var matherShip MatherShip
		err := rows.Scan(&matherShip.Type, &matherShip.Owner, &matherShip.X, &matherShip.Y, &matherShip.HP, &matherShip.Armor, &matherShip.RangeView)
		if err != nil {
			log.Fatal(err)
		}

		if matherShips[matherShip.X] != nil {
			matherShips[matherShip.X][matherShip.Y] = &matherShip
		} else {
			matherShips[matherShip.X] = make(map[int]*MatherShip)
			matherShips[matherShip.X][matherShip.Y] = &matherShip
		}
	}

	return
}
