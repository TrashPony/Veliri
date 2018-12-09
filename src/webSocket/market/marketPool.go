package market

import (
	"../../mechanics/db/get"
	"../../mechanics/gameObjects/order"
	"../../mechanics/market"
	"../../mechanics/player"
	"../../mechanics/players"
	"../storage"
	"../utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersMarketWs = make(map[*websocket.Conn]*player.Player)

type Order struct {
}

type Message struct {
	Event       string         `json:"event"`
	Orders      []*order.Order `json:"orders"`
	StorageSlot int            `json:"storage_slot"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
	MinBuyOut   int            `json:"min_buy_out"`
	Expires     int            `json:"expires"`
	Error       string         `json:"error"`
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
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {
			println(err.Error())
			utils.DelConn(ws, &usersMarketWs, err)
			break
		}

		if msg.Event == "openMarket" {
			OpenMarket(msg, ws)
		}

		if msg.Event == "placeNewBuyOrder" {
			// todo открытие нового ордера на покупку, оповестить других участников рынка
		}

		if msg.Event == "placeNewSellOrder" {
			err := market.PlaceNewSellOrder(msg.StorageSlot, msg.Price, msg.Quantity, msg.MinBuyOut, msg.Expires, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
			}
		}

		if msg.Event == "cancelBuyOrder" {
			// todo отмена ордера на продажу, оповестить других участников рынка
		}

		if msg.Event == "cancelSellOrder" {
			// todo отмена ордера на продажу, оповестить других участников рынка
		}

		if msg.Event == "buy" {
			// todo покупка в открытый оредар или частичный выкуп, оповестить других участников рынка
		}

		if msg.Event == "sell" {
			// todo продажа в открытый оредар или частичный выкуп, оповестить других участников рынка
		}
	}
}

func OrderSender() {
	mutex.Lock()
	for ws := range usersMarketWs {
		err := ws.WriteJSON(Message{Event: "openMarket", Orders: get.OpenOrders()})
		if err != nil {
			log.Printf("error: %v", err)
			ws.Close()
			delete(usersMarketWs, ws)
		}
	}
	mutex.Unlock()
}
