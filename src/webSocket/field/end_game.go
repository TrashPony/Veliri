package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"time"
)

func EndGame(game *localGame.Game) {

	if game.StartEnd {
		return
	}
	game.StartEnd = true

	for i := 3600; i > 0; i-- {

		SendAllMessage(Message{Event: "EndGame", Seconds: i}, game)

		end := false

		for _, client := range game.GetPlayers() {
			if !client.Leave {
				end = true
			}
		}

		if end {
			game.End = true
			update.Game(game)
		}

		time.Sleep(1 * time.Second)
	}

	game.End = true
	for _, client := range game.GetPlayers() {
		if !client.Leave {
			leave(client, game, true)
		}
	}

	update.Game(game)
}
