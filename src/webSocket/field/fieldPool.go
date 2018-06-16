package field

import (
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"../../mechanics/player"
	"log"
)

var watchPipe = make(chan Watch)
var phasePipe = make(chan PhaseInfo)
var targetPipe = make(chan Unit)
var equipPipe = make(chan SendUseEquip)
var move = make(chan Move)

var usersFieldWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения

var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersFieldWs)
	newPlayer := player.Player{}
	newPlayer.SetLogin(login)
	newPlayer.SetID(id)
	usersFieldWs[ws] = &newPlayer // Регистрируем нового Клиента
	print("WS field Сессия: ")    // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	fieldReader(ws, usersFieldWs)
}

func fieldReader(ws *websocket.Conn, usersFieldWs map[*websocket.Conn]*player.Player) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs, err)
			break
		}

		if msg.Event == "InitGame" {
			mutex.Lock() // это тут что бы при создание новой игры когда все пишут одновременно не создавались копии игр
			loadGame(msg, ws)
			mutex.Unlock()
			continue
		}

		if msg.Event == "SelectStorageUnit" {
			selectStorageUnit(msg, ws)
			continue
		}

		if msg.Event == "PlaceUnit" {
			placeUnit(msg, ws)
			continue
		}

		if msg.Event == "Ready" {
			Ready(ws)
			continue
		}

		if msg.Event == "SelectUnit" {
			SelectUnit(msg, ws)
			continue
		}

		if msg.Event == "GetTargetZone" {
			GetTargetZone(msg, ws)
		}

		if msg.Event == "MoveUnit" {
			MoveUnit(msg, ws)
			continue
		}

		if msg.Event == "SkipMoveUnit" {
			SkipMoveUnit(msg, ws)
			continue
		}

		if msg.Event == "SetTarget"{
			SetTarget(msg, ws)
			continue
		}

		if msg.Event == "Defend" {
			DefendTarget(msg, ws)
			continue
		}

		if msg.Event == "UseEquip" {
			UseEquip(msg, ws)
			continue
		}
	}
}

func WatchSender() {
	for {
		resp := <-watchPipe
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
		resp := <- targetPipe
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

func EquipSender() {
	for {
		resp := <- equipPipe
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
