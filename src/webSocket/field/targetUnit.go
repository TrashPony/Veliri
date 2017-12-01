package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"strconv"
)

func TargetUnit(msg FieldMessage, ws *websocket.Conn) {
	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client := usersFieldWs[ws]
	activeGame := Games[client.GameID]

	if find {
		coordinates := game.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
		passed := false

		for _, target := range coordinates {
			if target.X == msg.TargetX && target.Y == msg.TargetY {
				target, ok := client.HostileUnits[msg.TargetX][msg.TargetY]
				if ok {
					game.SetTarget(*unit, strconv.Itoa(target.X)+":"+strconv.Itoa(target.Y), activeGame.stat.Id)
					unit.Target = &game.Coordinate{X: target.X, Y: target.Y}
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
