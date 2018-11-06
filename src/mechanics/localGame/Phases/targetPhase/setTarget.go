package targetPhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"errors"
	"math"
)

func SetTarget(gameUnit *unit.Unit, game *localGame.Game, targetQ, targetR int, client *player.Player) {

	target, _ := game.Map.GetCoordinate(targetQ, targetR)
	gameUnit.Target = target

	updateSquad.Squad(client.GetSquad())
}

func SetEquipTarget(useUnit *unit.Unit, useCoordinate *coordinate.Coordinate, useEquipSlot *detail.BodyEquipSlot, client *player.Player) error {
	if !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {
		// TODO проверка по энергии и отнимание энергии должна быть в фазе прицеливания, что бы сразу было понятно сколькко энергии еще осталось
		// TODO а при снимание цели возвращать энергию в тело
		useEquipSlot.Target = useCoordinate
		updateSquad.Squad(client.GetSquad())

		return nil
	} else {
		if useUnit.Power < useEquipSlot.Equip.UsePower {
			return errors.New("no power")
		}

		if useEquipSlot.Used {
			return errors.New("you not ready")
		}

		return errors.New("unknown error")
	}
}

func rotateUnit(unitCoordinate, target *coordinate.Coordinate) int {

	//http://zvold.blogspot.com/2010/01/bresenhams-line-drawing-algorithm-on_26.html

	var rotate float64

	if unitCoordinate.R%2 != 0 {
		if (target.R-unitCoordinate.R)%2 != 0 {
			rotate = math.Atan2(float64(target.R)-float64(unitCoordinate.R), (float64(target.Q)-0.5)-float64(unitCoordinate.Q))
		} else {
			rotate = math.Atan2(float64(target.R-unitCoordinate.R), float64(target.Q-unitCoordinate.Q))
		}
	} else {
		if (target.R-unitCoordinate.R)%2 != 0 {
			rotate = math.Atan2(float64(target.R)-float64(unitCoordinate.R), float64(target.Q)-(float64(unitCoordinate.Q)-0.5))
		} else {
			rotate = math.Atan2(float64(target.R-unitCoordinate.R), float64(target.Q-unitCoordinate.Q))
		}
	}

	rotate = rotate * 180 / math.Pi
	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
