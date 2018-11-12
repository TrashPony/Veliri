package market

import (
	"../../mechanics/player"
	"../../mechanics/players"
	"../utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"log"
)

var mutex = &sync.Mutex{}

var orderPipe = make(chan Order)
var usersMarketWs = make(map[*websocket.Conn]*player.Player)

type Order struct {
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersMarketWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersMarketWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS market Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	mutex.Unlock()

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		//var msg Message

		//err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		println()
	}
}

func OrderSender() {
	for {
		resp := <-orderPipe
		mutex.Lock()
		for ws := range usersMarketWs {
			err := ws.WriteJSON(resp)
			if err != nil {
				log.Printf("error: %v", err)
				ws.Close()
				delete(usersMarketWs, ws)
			}
		}
		mutex.Unlock()
	}
}
