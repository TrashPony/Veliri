package field

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
	"../../game"
)

var fieldPipe = make(chan FieldResponse)
var initUnit = make(chan InitUnit)
var initStructure = make(chan InitStructure)
var coordinate = make(chan sendCoordinate)

var usersFieldWs = make(map[*websocket.Conn]*game.Player) // тут будут храниться наши подключения
var Games = make(map[int]*game.Game)

var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersFieldWs)
	newPlayer := game.Player{}
	newPlayer.SetLogin(login)
	newPlayer.SetID(id)

	usersFieldWs[ws] = &newPlayer // Регистрируем нового Клиента
	print("WS field Сессия: ")                        // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	fieldReader(ws, usersFieldWs)
}

func fieldReader(ws *websocket.Conn, usersFieldWs map[*websocket.Conn]*game.Player) {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs, err)
			break
		}

		if msg.Event == "InitGame" {
			toGame(msg, ws)
			continue
		}

		if msg.Event == "CreateUnit" {
			CreateUnit(msg, ws)
			continue
		}

		if msg.Event == "MouseOver" {
			MouseOver(msg, ws)
			continue
		}

		if msg.Event == "Ready" {
			Ready(msg, ws)
			continue
		}
		if msg.Event == "SelectUnit" || msg.Event == "SelectCoordinateCreate" {
			SelectUnit(msg, ws)
			continue
		}
		if msg.Event == "MoveUnit" {
			MoveUnit(msg, ws)
			continue
		}

		if msg.Event == "TargetUnit" {
			TargetUnit(msg, ws)
			continue
		}
	}
}

func FieldReposeSender() {
	for {
		resp := <-fieldPipe
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName {
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

func InitUnitSender() {
	for {
		resp := <-initUnit
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName {
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

func InitStructureSender()  {
	for {
		resp := <- initStructure
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName {
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

func CoordinateSender() {
	for {
		resp := <-coordinate
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.GetLogin() == resp.UserName {
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
