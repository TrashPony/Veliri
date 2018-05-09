package field

import (
	"github.com/gorilla/websocket"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]
	unit, find := client.GetUnit(msg.X, msg.Y)
	if !find {
		unit, find = client.GetHostileUnit(msg.X, msg.Y)
	}

	if find && ok {
		var resp InitUnit
		resp.initUnit(msg.Event, unit, client.GetLogin())
	} else {
		matherShip := usersFieldWs[ws].GetMatherShip()
		if matherShip == nil {
			matherShip, find = usersFieldWs[ws].GetHostileMatherShip(msg.X, msg.Y)
			if find {
				var resp InitStructure
				resp.initMatherShip(msg.Event, matherShip, client.GetLogin())
			}
		} else {
			var resp InitStructure
			resp.initMatherShip(msg.Event, matherShip, client.GetLogin())
		}
	}
}
