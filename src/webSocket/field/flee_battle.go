package field

import (
	localGame2 "github.com/TrashPony/Veliri/src/mechanics/db/localGame"
	gameUpdate "github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/gorilla/websocket"
)

func fleeBattle(msg Message, ws *websocket.Conn) {
	client := localGame.Clients.GetByWs(ws)

	if client != nil {
		activeGame, findGame := games.Games.Get(client.GetGameID())

		if findGame && activeGame.Phase == "targeting" {

			// когда игрок ливает из боя все юниты которые на карте остаются на поле
			// боя как болванки и больше не принадлежат отряду и хранятся в отдельной таблице
			// оставлять игрока в игре со статусом leave true

			for _, unitSlot := range client.GetSquad().MatherShip.Units {
				if unitSlot.Unit != nil && unitSlot.Unit.OnMap {
					unitSlot.Unit.GameID = activeGame.Id
					localGame2.AddLeaveUnit(unitSlot.Unit, client.GetID(), activeGame.Id)
					unitSlot.Unit = nil
				}
			}

			client.GetSquad().InGame = false
			update.Squad(client.GetSquad(), true)

			client.Leave = true
			gameUpdate.Player(client)

			// заменяем игрока на фейка т.к. тот уже не игре
			activeGame.SetFakePlayer(client)

			activePlayersCount := 0
			for _, user := range activeGame.GetPlayers() {
				if !user.Leave {
					activePlayersCount++
				}
			}

			if activePlayersCount < 2 {
				// todo отправить последнему игроку в игре сообщение для выхода из игры
				// 1) мгноменно с потерей
				// 2) с ожиданием 15 сек но без потерь

				// TODO удалять игру
			}
		}
	}
}
