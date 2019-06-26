package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/mapEditor"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func heightCoordinate(msg Message, ws *websocket.Conn, height int) {
	mapChange, _ := maps.Maps.GetByID(msg.ID)
	coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
	mapEditor.ChangeHeightCoordinate(coordinateMap, mapChange, height)
}
