package field

import (
	"../../game/mechanics"
	"../../game/objects"
	"github.com/gorilla/websocket"
	"strconv"
)

func TargetUnit(msg FieldMessage, ws *websocket.Conn) {
	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client := usersFieldWs[ws]
	game := Games[client.GameID]

	if find {
		coordinates := objects.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
		passed := false

		for _, target := range coordinates {
			if target.X == msg.TargetX && target.Y == msg.TargetY {
				target, ok := client.HostileUnits[msg.TargetX][msg.TargetY]
				if ok {
					mechanics.SetTarget(*unit, strconv.Itoa(target.X)+":"+strconv.Itoa(target.Y), game.stat.Id)
					unit.Target = &objects.Coordinate{X: target.X, Y: target.Y}
					passed = true
					resp := FieldResponse{Event: msg.Event, UserName: client.Login}
					fieldPipe <- resp
					break
				}
			}
		}

		if !passed {
			resp := FieldResponse{Event: msg.Event, UserName: client.Login, Error: "not allow"}
			fieldPipe <- resp
		}
	} else {
		resp := FieldResponse{Event: msg.Event, UserName: client.Login, Error: "unit not found"}
		fieldPipe <- resp
	}
}
