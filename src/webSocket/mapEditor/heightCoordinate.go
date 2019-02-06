package mapEditor

import (
	"../../mechanics/db/mapEditor"
	"../../mechanics/factories/maps"
	"github.com/gorilla/websocket"
)

func heightCoordinate(msg Message, ws *websocket.Conn, height int) {
	mapChange, _ := maps.Maps.GetByID(msg.ID)

	coordinateMap, ok := mapChange.GetCoordinate(msg.Q, msg.R)
	if ok {
		coordinateMap.Level += height
	}

	go mapEditor.ChangeHeightCoordinate(msg.ID, msg.Q, msg.R, height)
	selectMap(msg, ws)
}