package market

import (
	"../../mechanics/db/get"
	"github.com/gorilla/websocket"
)

func OpenMarket(msg Message, ws *websocket.Conn) {
	ws.WriteJSON(Message{Event: msg.Event, Orders: get.OpenOrders()})
}
