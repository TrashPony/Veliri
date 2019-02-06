package market

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"github.com/TrashPony/Veliri/src/mechanics/market"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/TrashPony/Veliri/src/webSocket/storage"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersMarketWs = make(map[*websocket.Conn]*player.Player)

type Message struct {
	Event       string               `json:"event"`
	Orders      map[int]*order.Order `json:"orders"`
	Assortment  *market.Assortment   `json:"assortment"`
	OrderID     int                  `json:"order_id"`
	StorageSlot int                  `json:"storage_slot"`
	Price       int                  `json:"price"`
	Quantity    int                  `json:"quantity"`
	MinBuyOut   int                  `json:"min_buy_out"`
	Expires     int                  `json:"expires"`
	Error       string               `json:"error"`
	Credits     int                  `json:"credits"`
	ItemID      int                  `json:"item_id"`
	ItemType    string               `json:"item_type"`
	BaseName    string               `json:"base_name"`
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
			base, find := bases.Bases.Get(usersMarketWs[ws].InBaseID)

			var marketName string
			if !find {
				marketName = "Пустош"
			} else {
				marketName = base.Name
			}

			ws.WriteJSON(Message{Event: msg.Event, Orders: market.Orders.GetOrders(),
				Credits: usersMarketWs[ws].GetCredits(), Assortment: market.GetAssortment(),
				BaseName: marketName})
		}

		if msg.Event == "placeNewBuyOrder" {
			err := market.Orders.PlaceNewBuyOrder(msg.ItemID, msg.Price, msg.Quantity, msg.MinBuyOut, msg.Expires, msg.ItemType, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
			}
		}

		if msg.Event == "placeNewSellOrder" {
			err := market.Orders.PlaceNewSellOrder(msg.StorageSlot, msg.Price, msg.Quantity, msg.MinBuyOut, msg.Expires, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
			}
		}

		if msg.Event == "cancelOrder" {
			err := market.Orders.Cancel(msg.OrderID, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
				ws.WriteJSON(Message{Event: "getMyOrders", Orders: market.Orders.GetUserOrders(usersMarketWs[ws].GetID())})
			}
		}

		if msg.Event == "buy" { // покупака из существующего ордера
			err := market.Orders.Buy(msg.OrderID, msg.Quantity, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
			}
		}

		if msg.Event == "sell" { // продажа в существующий ордер
			err := market.Orders.Sell(msg.OrderID, msg.Quantity, usersMarketWs[ws])
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				storage.Updater(usersMarketWs[ws].GetID())
				OrderSender()
			}
		}

		if msg.Event == "getMyOrders" {
			ws.WriteJSON(Message{Event: msg.Event, Orders: market.Orders.GetUserOrders(usersMarketWs[ws].GetID())})
		}
	}
}

func OrderSender() {
	mutex.Lock()
	for ws := range usersMarketWs {
		err := ws.WriteJSON(Message{Event: "openMarket",
			Orders:  market.Orders.GetOrders(),
			Credits: usersMarketWs[ws].GetCredits()})
		if err != nil {
			log.Printf("error: %v", err)
			ws.Close()
			delete(usersMarketWs, ws)
		}
	}
	mutex.Unlock()
}
