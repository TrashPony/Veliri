package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
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
	ID       string                  `json:"id"`
	Response bool                    `json:"response"`
	Accept   bool                    `json:"accept"`
	Credits  int                     `json:"credits"`
	Slots    map[int]*inventory.Slot `json:"slots"`
}

var diplomacyRequests = make(map[string]*diplomacyRequest)

// метод, когда игрок предложил мир другому игроку
func armisticePact(msg Message, client *player.Player) {

	activeGame, findGame := games.Games.Get(client.GetGameID())
	if findGame {

		for _, user := range activeGame.GetPlayers() {
			// отправляем игроку сообщение о намерение заключить мир
			if !user.Leave && user.GetLogin() == msg.ToUser {

				_, find := diplomacyRequests[user.GetLogin()+client.GetLogin()]

				if !activeGame.CheckPacts(client.GetID(), user.GetID()) && !find {

					if msg.Slots == nil && msg.Credits < user.GetCredits() || client.GetSquad().Inventory.ViewItemsBySlots(msg.Slots) && msg.Credits < user.GetCredits() {

						// мы удостоверились в том что все слоты и нужное количество присутсвует в инвентаре
						// в теории это удаление безопасно
						for number, slots := range client.GetSquad().Inventory.Slots {
							realSlot, _ := client.GetSquad().Inventory.Slots[number]
							realSlot.RemoveItemBySlot(slots.Quantity)
						}

						// отнимает кридиты у юзера
						client.SetCredits(client.GetCredits() - msg.Credits)

						request := diplomacyRequest{
							ID:      client.GetLogin() + user.GetLogin(),
							Credits: msg.Credits,
							Slots:   msg.Slots,
						}

						diplomacyRequests[client.GetLogin()+user.GetLogin()] = &request

						SendMessage(
							Message{
								Event:            "DiplomacyRequests",
								ToUser:           client.GetLogin(),
								DiplomacyRequest: &request,
							},
							user.GetID(),
							activeGame.Id,
						)

						go requestTimer(client.GetLogin()+user.GetLogin(), client, user, activeGame, &request)

					} else {
						SendMessage(ErrorMessage{Event: msg.Event, Error: "few items"}, client.GetID(), activeGame.Id)
					}

				} else {
					SendMessage(ErrorMessage{Event: msg.Event, Error: "pact already"}, client.GetID(), activeGame.Id)
				}
			}
		}
	}
}

func requestTimer(id string, client, toUser *player.Player, game *localGame.Game, request *diplomacyRequest) {

	rejectFunc := func() {
		delete(diplomacyRequests, id)

		client.SetCredits(client.GetCredits() + request.Credits)
		for _, slot := range request.Slots {
			client.GetSquad().Inventory.AddItemFromSlot(slot)
		}
	}

	for i := 15; i > 0; i -- {
		time.Sleep(1 * time.Second)

		if game.CheckPacts(client.GetID(), toUser.GetID()) {
			// проверка на то что союза небыло раньше
			SendMessage(ErrorMessage{Event: "ArmisticePact", Error: "pact already"}, client.GetID(), game.Id)
			rejectFunc()
			return
		}

		if request.Response {

			if request.Accept {

				toUser.SetCredits(toUser.GetCredits() + request.Credits)

				// TODO под мс client выкидываем протектор ящик и отдаем пароль user тот должен запомнится на фронте

				if game.CheckEndGame() {
					SendAllMessage(Message{Event: "EndGame"}, game)
				}

				game.Pacts = append(game.Pacts, &localGame.Pact{UserID1: toUser.GetID(), UserID2: client.GetID()})
				update.Game(game)

			} else {
				rejectFunc()
				SendMessage(Message{Event: "DiplomacyRequestsReject", ToUser: client.GetLogin()}, client.GetID(), game.Id)
			}

			return
		}
	}

	rejectFunc()
	SendMessage(Message{Event: "timeOutDiplomacyRequests", ToUser: toUser.GetLogin()}, client.GetID(), game.Id)
}

// метод когда игрок соглашается или нет с перемирием которое ему предложили
func acceptArmisticePact(msg Message, client *player.Player) {
	_, findGame := games.Games.Get(client.GetGameID())

	if findGame {

		request, find := diplomacyRequests[msg.ToUser+client.GetLogin()]
		if find && !request.Response {

			request.Response = true
			request.Accept = msg.Accept
		}
	}
}
