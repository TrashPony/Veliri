package field

import (
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/player"
)

func updateUnitHostileUser(client *player.Player, activeGame *localGame.Game, gameUnit *unit.Unit) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(gameUnit.Q, gameUnit.R)
			if watch {
				targetPipe <- Unit{Event: "UpdateUnit", UserName: user.GetLogin(), GameID: activeGame.Id, Unit: gameUnit}
			}
		}
	}
}
