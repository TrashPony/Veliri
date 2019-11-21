package move

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"

func StopMove(moveUnit *unit.Unit, resetSpeed bool) {
	if moveUnit != nil {
		moveUnit.ActualPath = nil // останавливаем прошлое движение

		if resetSpeed {
			moveUnit.CurrentSpeed = 0
		}
	}
}
