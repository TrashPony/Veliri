package field

import (
	"log"
	"websocket-master"
	"../../game/objects"
	"strconv"
)

var fieldPipe = make(chan FieldResponse)
var usersFieldWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения

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
		}

		if msg.Event == "SelectCoordinateCreate" {
			SelectCoordinateCreate(ws)
		}

		if msg.Event == "CreateUnit" {
			CreateUnit(msg, ws)
		}

		if msg.Event == "MouseOver" {
			MouseOver(msg, ws)
		}

		if msg.Event == "Ready" {
			Ready(msg, ws)
		}
	}
}

func FieldReposeSender() {
	for {
		resp := <- fieldPipe // TODO : разделить пайп на множество под каждую фазу
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
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id int
	PermittedCoordinates []objects.Coordinate
	Units []objects.Unit
	Respawn objects.Respawn
	CreateZone []objects.Coordinate
	Players []objects.UserStat
}

