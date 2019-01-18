package mapEditor

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func getMapList(msg Message, ws *websocket.Conn) {

	resp := Response{Event: "MapList", Maps: maps.Maps.GetAllMap()}

	ws.WriteJSON(resp)
}

func selectMap(msg Message, ws *websocket.Conn) {
	selectMap, _ := maps.Maps.GetByID(msg.ID)

	resp := Response{Event: "MapSelect", Map: *selectMap, Bases: bases.Bases.GetBasesByMap(selectMap.Id)}

	ws.WriteJSON(resp)
}
