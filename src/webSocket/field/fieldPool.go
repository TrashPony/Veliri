package field

import (
	"../../mechanics/player"
	"../../mechanics/players"
	"../utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var phasePipe = make(chan PhaseInfo)
var targetPipe = make(chan Unit)
var attackPipe = make(chan AttackMessage)
var move = make(chan Move)

var usersFieldWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения

var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersFieldWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersFieldWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS field Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	mutex.Unlock()

	fieldReader(ws, usersFieldWs)
}

func fieldReader(ws *websocket.Conn, usersFieldWs map[*websocket.Conn]*player.Player) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			utils.DelConn(ws, &usersFieldWs, err)
			break
		}

		if msg.Event == "InitGame" {
			mutex.Lock() // это тут что бы при создание новой игры когда все пишут одновременно не создавались копии игр
			loadGame(msg, ws)
			mutex.Unlock()
			continue
		}

		if msg.Event == "Ready" {
			Ready(ws)
			continue
		}

		if msg.Event == "SelectUnit" || msg.Event == "SelectStorageUnit" {
			SelectUnit(msg, ws)
			continue
		}

		if msg.Event == "GetTargetZone" {
			GetTargetZone(msg, ws)
			continue
		}

		if msg.Event == "GetPreviewPath" {
			GetPreviewPath(msg, ws)
			continue
		}

		if msg.Event == "MoveUnit" || msg.Event == "PlaceUnit" {
			MoveUnit(msg, ws)
			continue
		}

		if msg.Event == "SkipMoveUnit" {
			SkipMoveUnit(msg, ws)
			continue
		}

		if msg.Event == "SetWeaponTarget" {
			SetTarget(msg, ws)
			continue
		}

		if msg.Event == "Defend" {
			DefendTarget(msg, ws)
			continue
		}

		if msg.Event == "SetTargetMapEquip" {
			SetTargetMapEquip(msg, ws)
			continue
		}

		if msg.Event == "SetTargetUnitEquip" {
			SetTargetUnitEquip(msg, ws)
			continue
		}

		if msg.Event == "SelectWeapon" {
			SelectWeapon(msg, ws)
			continue
		}

		if msg.Event == "SelectEquip" {
			SelectEquip(msg, ws)
			continue
		}
	}
}

func PhaseSender() {
	for {
		resp := <-phasePipe
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName && client.GetGameID() == resp.GameID {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

func MoveSender() {
	for {
		resp := <-move
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName && client.GetGameID() == resp.GameID {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

func UnitSender() {
	for {
		resp := <-targetPipe
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName && client.GetGameID() == resp.GameID {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

func AttackSender() {
	for {
		resp := <-attackPipe
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName && client.GetGameID() == resp.GameID {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}
