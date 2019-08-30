package market

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/market"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
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
	MyOrders    map[int]*order.Order `json:"my_orders"`
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
	UserID      int                  `json:"user_id"`
	BaseName    string               `json:"base_name"`
	Count       int                  `json:"count"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()
	// todo concurrent map iteration and map write
	utils.CheckDoubleLogin(login, &usersMarketWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersMarketWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS market Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	mutex.Unlock()

	Reader(ws, newPlayer)
}

func Reader(ws *websocket.Conn, user *player.Player) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {
			mutex.Lock()
			ws.Close()
			delete(usersMarketWs, ws)
			mutex.Unlock()
			break
		}

		if msg.Event == "openMarket" || msg.Event == "getMyOrders" {
			OrderSender()
		}

		if msg.Event == "placeNewBuyOrder" {
			err := market.Orders.PlaceNewBuyOrder(msg.ItemID, msg.Price, msg.Quantity, msg.MinBuyOut, msg.Expires, msg.ItemType, user)
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				wsInventory.UpdateStorage(user.GetID())
				OrderSender()
			}
		}

		if msg.Event == "placeNewSellOrder" {
			err := market.Orders.PlaceNewSellOrder(msg.StorageSlot, msg.Price, msg.Quantity, msg.MinBuyOut, msg.Expires, user)
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				wsInventory.UpdateStorage(user.GetID())
				OrderSender()
			}
		}

		if msg.Event == "cancelOrder" {
			err := market.Orders.Cancel(msg.OrderID, user)
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				wsInventory.UpdateStorage(user.GetID())
				OrderSender()
			}
		}

		if msg.Event == "buy" { // покупака из существующего ордера
			err := market.Orders.Buy(msg.OrderID, msg.Quantity, user)
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				wsInventory.UpdateStorage(user.GetID())
				OrderSender()
			}
		}

		if msg.Event == "sell" { // продажа в существующий ордер
			err := market.Orders.Sell(msg.OrderID, msg.Quantity, user)
			if err != nil {
				ws.WriteJSON(Message{Event: msg.Event, Error: err.Error()})
			} else {
				wsInventory.UpdateStorage(user.GetID())
				OrderSender()
			}
		}

		if msg.Event == "getItemsInStorage" {
			// запрашивает все айтемы на складе которые можно продать (полностью починеные), на базе где размещен ордер
			count := 0

			find, marketOrder, mx := market.Orders.GetOrder(msg.OrderID)
			mx.Unlock()

			if find {
				storage, _ := storages.Storages.Get(user.GetID(), marketOrder.PlaceID)
				for _, slot := range storage.Slots {
					if slot.ItemID == msg.ItemID && slot.Type == msg.ItemType && slot.HP == slot.MaxHP {
						count++
					}
				}

				ws.WriteJSON(Message{Event: msg.Event, Count: count})
			} else {
				ws.WriteJSON(Message{Event: msg.Event, Error: "no find order"})
			}
		}
	}
}

func OrderSender() {
	mutex.Lock()
	for ws, user := range usersMarketWs {

		allOrders := market.Orders.GetOrders()
		myOrders := market.Orders.GetUserOrders(user.GetID())

		userBase, find := bases.Bases.Get(user.InBaseID)
		var marketName string
		var userMapID int

		if !find {
			marketName = "Пустош"
			userMapID = user.GetSquad().MatherShip.MapID
		} else {
			marketName = userBase.Name
			userMap, _ := maps.Maps.GetByID(userBase.MapID)
			userMapID = userMap.Id
		}

		// находить растояние в секторах от игрока до ордера
		var distSearch = func(orders *map[int]*order.Order) {
			for _, marketOrder := range *orders {

				if marketOrder.PlaceID == user.InBaseID {
					marketOrder.PathJump = -1
					continue
				}

				orderBase, _ := bases.Bases.Get(marketOrder.PlaceID)
				orderMap, _ := maps.Maps.GetByID(orderBase.MapID)

				pathMaps, _ := maps.Maps.FindGlobalPath(userMapID, orderMap.Id)
				marketOrder.PathJump = len(pathMaps) - 1
			}
		}

		distSearch(&allOrders)
		distSearch(&myOrders)

		err := ws.WriteJSON(
			Message{
				Event:      "openMarket",
				Orders:     allOrders,
				Credits:    user.GetCredits(),
				MyOrders:   myOrders,
				BaseName:   marketName,
				Assortment: market.GetAssortment(),
			},
		)

		if err != nil {
			log.Printf("error: %v", err)
			ws.Close()
			delete(usersMarketWs, ws)
		}
	}
	mutex.Unlock()
}
