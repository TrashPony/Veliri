package field

import "github.com/gorilla/websocket"

func placeUnit(msg Message, ws *websocket.Conn)  {
	client, ok := usersFieldWs[ws]

	if !ok {
		delete(usersFieldWs, ws)
		return
	}

	unit, find := client.GetUnitStorage(msg.UnitID)
	if find {
		client.PlaceUnit(unit, msg.X, msg.Y)
	}
}
