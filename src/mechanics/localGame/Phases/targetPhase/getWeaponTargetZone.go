package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../Phases"
	"../../../localGame"
)

func GetWeaponTargetCoordinate(gameUnit *unit.Unit, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {

	targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	weaponRange := 0

	for _, weaponSlot := range gameUnit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			weaponRange = weaponSlot.Weapon.Range
		}
	}

	RadiusCoordinates := coordinate.GetCoordinatesRadius(gameUnit.GetX(), gameUnit.GetY(), weaponRange)
	zone := filter(gameUnit, RadiusCoordinates, activeGame)

	for _, gameCoordinate := range zone {
		if !(gameCoordinate.X == gameUnit.X && gameCoordinate.Y == gameUnit.Y) {
			Phases.AddCoordinate(targetCoordinate, gameCoordinate)
		}
	}

	return targetCoordinate
}
