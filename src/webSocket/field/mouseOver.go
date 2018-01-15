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
		structure, find := usersFieldWs[ws].GetStructure(msg.X, msg.Y)
		if !find {
			structure, find = usersFieldWs[ws].GetHostileStructure(msg.X, msg.Y)
		}
		if find {
			var resp InitStructure
			resp.initStructure(msg.Event, structure, client.GetLogin())
		}
	}
}
