package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../Phases"
	"strconv"
)

func GetWeaponTargetCoordinate(gameUnit *unit.Unit, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {

	targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	weaponRange := 0
	minWeaponRange := 0
	artillery := false

	for _, weaponSlot := range gameUnit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			weaponRange = weaponSlot.Weapon.Range
			minWeaponRange = weaponSlot.Weapon.MinAttackRange
			artillery = weaponSlot.Weapon.Artillery
		}
	}

	unitCoordinate, find := activeGame.Map.GetCoordinate(gameUnit.GetQ(), gameUnit.GetR())

	if find {
		radiusCoordinates := coordinate.GetCoordinatesRadius(unitCoordinate, weaponRange)
		deadZone := coordinate.GetCoordinatesRadius(unitCoordinate, minWeaponRange)

		zone := filter(gameUnit, radiusCoordinates, activeGame, artillery)

		for _, gameCoordinate := range zone {
			if !(gameCoordinate.Q == gameUnit.Q && gameCoordinate.R == gameUnit.R) {
				Phases.AddCoordinate(targetCoordinate, gameCoordinate)
			}
		}

		for _, deadCoordinate := range deadZone { // удаляем координаты которые находяться в мертвой зоне атаки
			_, find := targetCoordinate[strconv.Itoa(deadCoordinate.Q)][strconv.Itoa(deadCoordinate.R)]
			if find {
				delete(targetCoordinate[strconv.Itoa(deadCoordinate.Q)], strconv.Itoa(deadCoordinate.R))
			}
		}
	}

	return targetCoordinate
}
