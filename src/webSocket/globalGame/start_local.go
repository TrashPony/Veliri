package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"strconv"
)

func startLocalGame(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	squadID, err := strconv.Atoi(msg.ToSquadID)
	if err != nil {
		// TODO на ботов тоже можно нападать но не сейчас)
		go SendMessage(Message{Event: "Error", Error: "it's bot", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
		return
	}

	toUser := globalGame.Clients.GetBySquadId(squadID)

	if user != nil && toUser != nil &&
		user.GetSquad() != nil && toUser.GetSquad() != nil &&
		user.GetSquad().MapID == toUser.GetSquad().MapID && user.GetID() != toUser.GetID() {

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, toUser.GetSquad().GlobalX, toUser.GetSquad().GlobalY)

		// проверяем что атакующий пользователь "видит" того на кого нападает, *4 костыль не продумайн гей дизайн
		if int(dist) < user.GetSquad().MatherShip.RangeView*globalGame.HexagonHeight*4 {

			gamePlayers := make([]*player.Player, 0)
			gamePlayers = append(gamePlayers, user)
			gamePlayers = append(gamePlayers, toUser)

			// TODO посмотреть кто находится в радиусе боя и предложить им участие в бою и добавить их в бой

			// костыль
			mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
			startCoordinate := globalGame.GetQRfromXY(user.GetSquad().GlobalX, user.GetSquad().GlobalY, mp)
			user.GetSquad().MatherShip.Q = startCoordinate.Q
			user.GetSquad().MatherShip.R = startCoordinate.R

			startCoordinate = globalGame.GetQRfromXY(toUser.GetSquad().GlobalX, toUser.GetSquad().GlobalY, mp)
			toUser.GetSquad().MatherShip.Q = startCoordinate.Q
			toUser.GetSquad().MatherShip.R = startCoordinate.R

			_, err := localGame.StartNewGame("", user.GetSquad().MapID, gamePlayers)
			if err == nil {

				gameShortPlayers := make([]*player.ShortUserInfo, 0)
				for _, gamePlayer := range gamePlayers {
					gameShortPlayers = append(gameShortPlayers, gamePlayer.GetShortUserInfo(true))
				}

				SendMessage(Message{Event: "LocalGame", IDMap: user.GetSquad().MapID, Bot: user.Bot, OtherUsers: gameShortPlayers})

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
