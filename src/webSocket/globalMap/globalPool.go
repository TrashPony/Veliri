package globalMap

import (
	"../../lobby"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
	"time"
	"math/rand"
)

var mutex = &sync.Mutex{}

var usersGlobalWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var globalPipe = make(chan GlobalResponse)

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersGlobalWs)
	usersGlobalWs[ws] = &Clients{Login: login, Id: id, X: rand.Intn(1900), Y: rand.Intn(900), Rotate: rand.Intn(180)} // Регистрируем нового Клиента
	print("WS global Сессия: ")                        // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	GlobalReader(ws)
}

func GlobalReader(ws *websocket.Conn) {
	for {
		var msg GlobalMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersGlobalWs, err)
			break
		}

		if msg.Event == "Join" {
			for _, user := range usersGlobalWs {
				if user.Login != usersGlobalWs[ws].Login {
					var resp = GlobalResponse{Event: "Users", UserName: user.Login, X: user.X, Y: user.Y, Rotate: user.Rotate}
					ws.WriteJSON(resp)
				} else {
					var resp = GlobalResponse{Event: msg.Event, UserName: user.Login,  X: user.X, Y: user.Y, Rotate: user.Rotate}
					ws.WriteJSON(resp)
				}
			}
			var resp = GlobalResponse{Event: "Users", UserName: usersGlobalWs[ws].Login, X: usersGlobalWs[ws].X, Y: usersGlobalWs[ws].Y, Rotate: usersGlobalWs[ws].Rotate}
			globalPipe <- resp
		}

		if msg.Event == "PlayerMove" {
			if msg.UserName == usersGlobalWs[ws].Login {

				var resp = GlobalResponse{Event: msg.Event, UserName: usersGlobalWs[ws].Login, X: msg.X, Y: msg.Y, Rotate: msg.Rotate}
				globalPipe <- resp
			}
		}

		if msg.Event == "Fire" {
			var resp = GlobalResponse{Event: msg.Event, UserName: usersGlobalWs[ws].Login}
			globalPipe <- resp
		}
	}
}

func GlobalReposeSender() {
	for {
		resp := <-globalPipe
		mutex.Lock()
		for ws, client := range usersGlobalWs {
			err := ws.WriteJSON(resp)
			if err != nil {
				log.Printf("error: %v", err)
				lobby.DelLobbyGame(client.Login)
				ws.Close()
				delete(usersGlobalWs, ws)
			}

		}
		mutex.Unlock()
	}
}

func TimerSteep() {
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

type Clients struct {
	Login string
	Id    int
	X     int
	Y     int
	Rotate int
}
