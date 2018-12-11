package mapEditor

import (
	"../../mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func getMapList(msg Message, ws *websocket.Conn) {

	resp := Response{Event: "MapList", Maps: maps.Maps.GetAllMap()}

	ws.WriteJSON(resp)
}

func selectMap(msg Message, ws *websocket.Conn) {
	selectMap, _ := maps.Maps.GetByID(msg.ID)
	resp := Response{Event: "MapSelect", Map: *selectMap}

	ws.WriteJSON(resp)
}
