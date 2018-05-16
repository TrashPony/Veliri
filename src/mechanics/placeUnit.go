package mechanics

import (
	"./unit"
	"./db"
)

func PlaceUnit(gameUnit *unit.Unit, x,y int) error {
	gameUnit.SetX(x)
	gameUnit.SetY(y)
	gameUnit.SetOnMap(true)

	err := db.UpdateUnit(gameUnit)
	if err != nil {
		return err
	}
	return nil
}
