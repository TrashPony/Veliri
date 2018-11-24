package mapEditor

import (
	"../../mechanics/db/get"
	"github.com/gorilla/websocket"
)

func getAllCoordinate(msg Message, ws *websocket.Conn) {
	ws.WriteJSON(Response{Event: msg.Event, TypeCoordinates: get.AllTypeCoordinate()})
}
