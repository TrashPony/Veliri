package targetPhase

import (
	"../../../db/squad/update"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"errors"
)

func SetTarget(gameUnit *unit.Unit, game *localGame.Game, targetQ, targetR int, client *player.Player) {

	target, _ := game.Map.GetCoordinate(targetQ, targetR)
	gameUnit.Target = target

	update.Squad(client.GetSquad(), true)
}

func SetEquipTarget(useUnit *unit.Unit, useCoordinate *coordinate.Coordinate, useEquipSlot *detail.BodyEquipSlot, client *player.Player) error {
	if !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {
		// TODO проверка по энергии и отнимание энергии должна быть в фазе прицеливания, что бы сразу было понятно сколькко энергии еще осталось
		// TODO а при снимание цели возвращать энергию в тело
		useEquipSlot.Target = useCoordinate
		update.Squad(client.GetSquad(), true)

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
