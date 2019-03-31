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
		// TODO проверка во всех методах что игрок не ливнул
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

		if msg.Event == "Ready" {
			Ready(ws)
		}

		if msg.Event == "SelectUnit" || msg.Event == "SelectStorageUnit" {
			SelectUnit(msg, ws)
		}

		if msg.Event == "GetTargetZone" {
			GetTargetZone(msg, ws)
		}

		if msg.Event == "GetPreviewPath" {
			GetPreviewPath(msg, ws)
		}

		if msg.Event == "MoveUnit" || msg.Event == "PlaceUnit" {
			MoveUnit(msg, ws)
		}

		if msg.Event == "SkipMoveUnit" {
			SkipMoveUnit(msg, ws)
		}

		if msg.Event == "SetWeaponTarget" {
			SetTarget(msg, ws)
		}

		if msg.Event == "Defend" {
			DefendTarget(msg, ws)
		}

		if msg.Event == "SetTargetMapEquip" {
			SetTargetMapEquip(msg, ws)
		}

		if msg.Event == "SetTargetUnitEquip" {
			SetTargetUnitEquip(msg, ws)
		}

		if msg.Event == "SelectWeapon" {
			SelectWeapon(msg, ws)
		}

		if msg.Event == "SelectEquip" {
			SelectEquip(msg, ws)
		}

		if msg.Event == "FleeBattle" {
			fleeBattle(msg, ws)
		}

		if msg.Event == "Reload" {
			// TODO Перезарядка оружия
		}

		if msg.Event == "Diplomacy" {
			// TODO Дипломатия
		}

		if msg.Event == "Mining" {
			// TODO добыча ресурсов в локальной игре
		}
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
