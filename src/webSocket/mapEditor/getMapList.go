package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func getMapList(msg Message, ws *websocket.Conn) {

	resp := Response{Event: "MapList", Maps: maps.Maps.GetAllMap()}

	ws.WriteJSON(resp)
}

func selectMap(msg Message, ws *websocket.Conn) {
	selectMap := get.MapByID(msg.ID)

	resp := Response{Event: "MapSelect", Map: *selectMap, Bases: bases.Bases.GetBasesByMap(selectMap.Id)}

	ws.WriteJSON(resp)
}
