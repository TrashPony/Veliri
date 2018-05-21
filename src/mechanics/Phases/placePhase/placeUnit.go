package placePhase

import (
	"../../unit"
	"../../db"
	"../../game"
	"../../player"
)

func PlaceUnit(gameUnit *unit.Unit, x,y int, actionGame *game.Game, client *player.Player) error {
	gameUnit.SetX(x)
	gameUnit.SetY(y)
	gameUnit.SetOnMap(true)

	actionGame.DelUnitStorage(gameUnit.Id)
	actionGame.SetUnit(gameUnit)

	client.DelUnitStorage(gameUnit.Id)
	client.AddUnit(gameUnit)

	err := db.UpdateUnit(gameUnit)
	if err != nil {
		return err
	}
	return nil
}
