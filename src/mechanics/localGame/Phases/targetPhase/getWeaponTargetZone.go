package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../Phases"
	"strconv"
)

func GetWeaponTargetCoordinate(gameUnit *unit.Unit, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {

	if gameUnit.GetWeaponSlot() != nil {
		targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

		unitCoordinate, find := activeGame.Map.GetCoordinate(gameUnit.GetQ(), gameUnit.GetR())

		if find {
			radiusCoordinates := coordinate.GetCoordinatesRadius(unitCoordinate, gameUnit.GetWeaponSlot().Weapon.Range)
			deadZone := coordinate.GetCoordinatesRadius(unitCoordinate, gameUnit.GetWeaponSlot().Weapon.MinAttackRange)

			zone := filter(gameUnit, radiusCoordinates, activeGame, gameUnit.GetWeaponSlot().Weapon.Artillery)

			for _, gameCoordinate := range zone {
				if !(gameCoordinate.Q == gameUnit.Q && gameCoordinate.R == gameUnit.R) {
					_, find := activeGame.GetUnit(gameCoordinate.Q, gameCoordinate.R)

					if gameUnit.GetWeaponSlot().Weapon.Type == "firearms" {
						Phases.AddCoordinate(targetCoordinate, gameCoordinate)
					} else {
						if find {
							Phases.AddCoordinate(targetCoordinate, gameCoordinate)
						}
					}
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
	return nil
}
