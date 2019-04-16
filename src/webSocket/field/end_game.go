package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"time"
)

func EndGame(game *localGame.Game) {
	SendAllMessage(Message{Event: "EndGame"}, game)

	// через 5 минут принудительно всех отключит от игры
	time.Sleep(1 * time.Minute)
	game.End = true

	for _, client := range game.GetPlayers() {
		if !client.Leave {
			leave(client, game, true)
		}
	}

	update.Game(game)
}
