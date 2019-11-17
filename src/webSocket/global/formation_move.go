package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

func FormationInit(user *player.Player, unitsID []int) bool {
	for _, id := range unitsID {

		if user.GetSquad().MatherShip.ID == id && !user.GetSquad().MatherShip.Formation {
			user.GetSquad().MatherShip.Formation = true
			go FormationMove(user)
			return true
		} else {
			if user.GetSquad().MatherShip.Formation {
				return true
			}
		}

	}

	return false
}

func FormationMove(user *player.Player) {
	for {
		for _, unitSlot := range user.GetSquad().MatherShip.Units {

			if unitSlot.Unit != nil && unitSlot.Unit.OnMap && unitSlot.Unit.Formation {

				x, y := user.GetSquad().GetFormationCoordinate(unitSlot.Unit.FormationPos.X, unitSlot.Unit.FormationPos.Y)

				msg := Message{}
				msg.ToX, msg.ToY = float64(x), float64(y)
				msg.UnitsID = []int{unitSlot.Unit.ID}

				Move(user, msg, true)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}
