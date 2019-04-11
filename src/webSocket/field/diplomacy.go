package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"time"
)

// если между всеми участниками боя достигнут мир или остались только те игроки которые в мире то бой прекращается.

// отдает текущее состояние дипломатии, игроков, пакты которые уже заключены
func openDiplomacy(client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())
	if findGame {
		SendMessage(
			Message{
				Event:          "OpenDiplomacy",
				DiplomacyUsers: activeGame.Pacts,
			},
			client.GetID(),
			activeGame.Id,
		)
	}
}

type diplomacyRequest struct {
	ID       string `json:"id"`
	Response bool   `json:"response"`
}

var diplomacyRequests = make(map[string]diplomacyRequest)

// метод, когда игрок предложил мир другому игроку
func armisticePact(msg Message, client *player.Player) {

	activeGame, findGame := games.Games.Get(client.GetGameID())
	if findGame {

		for _, user := range activeGame.GetPlayers() {
			// отправляем игроку сообщение о намерение заключить мир
			if !user.Leave && user.GetLogin() == msg.ToUser {

				_, find := diplomacyRequests[user.GetLogin()+client.GetLogin()]

				if !activeGame.CheckPacts(client.GetID(), user.GetID()) && !find {

					SendMessage(
						Message{
							Event:  "DiplomacyRequests",
							ToUser: client.GetLogin(),
						},
						user.GetID(),
						activeGame.Id,
					)

					diplomacyRequests[client.GetLogin()+user.GetLogin()] = diplomacyRequest{ID: client.GetLogin() + user.GetLogin()}
					go requestTimer(client.GetLogin()+user.GetLogin(), client, user, activeGame.Id)

				} else {
					SendMessage(ErrorMessage{Event: msg.Event, Error: "pact already"}, client.GetID(), activeGame.Id)
				}

				return
			}
		}
	}
}

func requestTimer(id string, client, toUser *player.Player, gameID int) {
	for i := 15; i > 0; i -- {
		time.Sleep(1 * time.Second)
	}

	delete(diplomacyRequests, id)

	SendMessage(
		Message{
			Event:  "timeOutDiplomacyRequests",
			ToUser: toUser.GetLogin(),
		},
		client.GetID(),
		gameID,
	)
}

// метод когда игрок соглашается или нет с перемирием которое ему предложили
func acceptArmisticePact(msg Message, client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findGame {
		for _, user := range activeGame.GetPlayers() {
			if !user.Leave && user.GetLogin() == msg.ToUser {

				request, find := diplomacyRequests[user.GetLogin()+client.GetLogin()]

				if find && !request.Response {
					request.Response = true
					if msg.Accept {
						if !activeGame.CheckPacts(client.GetID(), user.GetID()) {
							activeGame.Pacts = append(activeGame.Pacts, &localGame.Pact{UserID1: user.GetID(), UserID2: client.GetID()})
							update.Game(activeGame)
						} else {
							SendMessage(ErrorMessage{Event: msg.Event, Error: "pact already"}, client.GetID(), activeGame.Id)
						}
					} else {
						SendMessage(
							Message{
								Event:  "DiplomacyRequestsReject",
								ToUser: client.GetLogin(),
							},
							user.GetID(),
							activeGame.Id,
						)
					}
				}
			}
		}
	}
}

// выкупи игрока за ресурсы, принудительный мир
func buyOut(msg Message, client *player.Player) {

}

// метод когда игрок соглашается или нет с выкупом
func acceptBuyOut(msg Message, client *player.Player) {

}
