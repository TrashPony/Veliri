package map_editor

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/maps/mapEditor"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func heightCoordinate(msg Message, ws *websocket.Conn, height int) {
	mapChange, _ := maps.Maps.GetByID(msg.ID)
	coordinateMap, _ := mapChange.GetCoordinate(msg.X, msg.Y)
	mapEditor.ChangeHeightCoordinate(coordinateMap, mapChange, height)
}
