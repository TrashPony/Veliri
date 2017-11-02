package field

import (
	"log"
	"websocket-master"
	"../../game/objects"
	"strconv"
	"sync"
)

var fieldPipe = make(chan FieldResponse)
var initUnit  = make(chan InitUnit)
var coordiante = make(chan Coordinate)
var usersFieldWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения

var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int)  {
	CheckDoubleLogin(login, &usersFieldWs)
	usersFieldWs[ws] = &Clients{Login:login, Id:id} // Регистрируем нового Клиента
	print("WS field Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	fieldReader(ws, usersFieldWs)
}

func fieldReader(ws *websocket.Conn, usersFieldWs map[*websocket.Conn]*Clients )  {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs , err)
			break
		}

		if msg.Event == "InitGame" {
			InitGame(msg, ws)
			continue
		}

		if msg.Event == "SelectCoordinateCreate" {
			SelectCoordinateCreate(ws)
			continue
		}

		if msg.Event == "CreateUnit" {
			CreateUnit(msg, ws) // TODO второй игрок не может сразу начать строить юнитов
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
		if msg.Event == "SelectUnit" {
			SelectUnit(msg, ws)
			continue
		}
		if msg.Event == "MoveUnit" {
			MoveUnit(msg, ws)
			continue
		}
		if msg.Event == "getPermittedCoordinates" {
			for _, unit := range usersFieldWs[ws].Units {
				SendWatchCoordinate(ws, unit)
			}
			continue
		}
	}
}

func FieldReposeSender() {
	for {
		resp := <- fieldPipe // TODO : разделить пайп на множество под каждую фазу
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
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
		resp := <- initUnit
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
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
		resp := <- coordiante
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
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

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id int
	Units map[string]*objects.Unit // KEY format X:Y
	HostileUnits map[string]*objects.Unit // KEY format X:Y
	Map objects.Map
	Respawn objects.Respawn
	CreateZone []*objects.Coordinate
	GameStat objects.Game
	Players []objects.UserStat
}

