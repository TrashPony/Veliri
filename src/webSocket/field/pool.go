package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var sendMessagePipe = make(chan message)
var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	localGame.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента
	println("WS field Сессия: login: " + login + " id: " + strconv.Itoa(id))

	fieldReader(ws)
}

func fieldReader(ws *websocket.Conn) {
	for {

		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			localGame.Clients.DelClientByWS(ws)
			return
		}

		if msg.Event == "InitGame" {
			mutex.Lock() // это тут что бы при создание новой игры когда все пишут одновременно не создавались копии игр
			loadGame(msg, ws)
			mutex.Unlock()
		}

		client := localGame.Clients.GetByWs(ws)
		if client != nil && client.GetSquad().InGame && !client.ToLeave {

			if msg.Event == "Ready" {
				Ready(client)
			}

			if msg.Event == "SelectUnit" || msg.Event == "SelectStorageUnit" {
				SelectUnit(msg, client)
			}

			if msg.Event == "GetTargetZone" {
				GetTargetZone(msg, client)
			}

			if msg.Event == "GetPreviewPath" {
				GetPreviewPath(msg, client)
			}

			if msg.Event == "MoveUnit" || msg.Event == "PlaceUnit" {
				MoveUnit(msg, client)
			}

			if msg.Event == "SkipMoveUnit" {
				SkipMoveUnit(msg, client)
			}

			if msg.Event == "SetWeaponTarget" {
				SetTarget(msg, client)
			}

			if msg.Event == "Defend" {
				DefendTarget(msg, client)
			}

			if msg.Event == "SetTargetMapEquip" {
				SetTargetMapEquip(msg, client)
			}

			if msg.Event == "SetTargetUnitEquip" {
				SetTargetUnitEquip(msg, client)
			}

			if msg.Event == "SelectWeapon" {
				SelectWeapon(msg, client)
			}

			if msg.Event == "SelectEquip" {
				SelectEquip(msg, client)
			}

			if msg.Event == "InitLeave" {
				initFlee(msg, client)
			}

			if msg.Event == "FleeBattle" {
				fleeBattle(msg, client)
			}

			if msg.Event == "softFlee" {
				softFlee(client)
			}

			if msg.Event == "initReload" {
				initAmmoReload(msg, client)
			}

			if msg.Event == "Reload" {
				ammoReload(msg, client)
			}

			if msg.Event == "OpenDiplomacy" {
				openDiplomacy(client)
			}

			if msg.Event == "ArmisticePact" {
				armisticePact(msg, client)
			}

			if msg.Event == "AcceptArmisticePact" {
				acceptArmisticePact(msg, client)
			}

			if msg.Event == "initBuyOut" {
				initBuyOut(msg, client)
			}

			if msg.Event == "Mining" {
				// TODO добыча ресурсов в локальной игре
			}
		}
	}
}

func SendAllMessage(senderMessage interface{}, game *localGame.Game) {
	for _, user := range game.GetPlayers() {
		SendMessage(senderMessage, user.GetID(), game.Id)
	}
}

func SendMessage(senderMessage interface{}, userID, gameID int) {
	moves := message{message: senderMessage, userID: userID, gameID: gameID}
	sendMessagePipe <- moves
}

func Sender() {
	for {
		// Помни что пока отправляются данные, они могут изменится!
		resp := <-sendMessagePipe

		users, mx := localGame.Clients.GetAllConnects()

		for ws, client := range users {
			if client.GetID() == resp.userID && client.GetGameID() == resp.gameID {
				err := ws.WriteJSON(resp.message)
				if err != nil {
					mx.Unlock()
					localGame.Clients.DelClientByWS(ws)
				}
			}
		}
		mx.Unlock()
	}
}
