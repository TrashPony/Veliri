package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
)

func startLocalGame(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	toUser := globalGame.Clients.GetById(msg.ToUserID)

	if user != nil && toUser != nil && user.GetSquad() != nil && toUser.GetSquad() != nil &&
		user.GetSquad().MapID == toUser.GetSquad().MapID && user.GetID() != toUser.GetID() {

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, toUser.GetSquad().GlobalX, toUser.GetSquad().GlobalY)

		// проверяем что атакующий пользователь "видит" того на кого нападает
		if int(dist) < user.GetSquad().MatherShip.RangeView*globalGame.HexagonHeight {

			gamePlayers := make([]*player.Player, 0)
			gamePlayers = append(gamePlayers, user)
			gamePlayers = append(gamePlayers, toUser)

			// TODO посмотреть кто находится в радиусе боя и предложить им участие в бою и добавить их в бой

			_, err := localGame.StartNewGame("", user.GetSquad().MapID, gamePlayers)
			if err == nil {

				gameShortPlayers := make([]*player.ShortUserInfo, 0)
				for _, gamePlayer := range gamePlayers {
					gameShortPlayers = append(gameShortPlayers, gamePlayer.GetShortUserInfo())
				}

				go SendMessage(Message{Event: "LocalGame", IDMap: user.GetSquad().MapID, Bot: user.Bot, OtherUsers: gameShortPlayers})

			} else {
				go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
			}
		} else {
			go SendMessage(Message{Event: "Error", Error: "Wrong range", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
		}
	} else {
		go SendMessage(Message{Event: "Error", Error: "Wrong target", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
	}
}
