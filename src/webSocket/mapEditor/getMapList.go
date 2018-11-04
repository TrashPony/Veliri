package mapEditor

import (
	"../../mechanics/db/get"
	"github.com/gorilla/websocket"
)

func getMapList(msg Message, ws *websocket.Conn) {

	resp := Response{Event: "MapList", Maps: get.InfoMapList()}

	ws.WriteJSON(resp)
}

func selectMap(msg Message, ws *websocket.Conn) {
	resp := Response{Event: "MapSelect", Map: get.Map(msg.ID)}

	ws.WriteJSON(resp)
}
