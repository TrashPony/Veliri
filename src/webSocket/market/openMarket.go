package market

import (
	"../../mechanics/market"
	"github.com/gorilla/websocket"
)

func OpenMarket(msg Message, ws *websocket.Conn) {
	ws.WriteJSON(Message{Event: msg.Event, Orders: market.Orders.GetOrders()})
}
