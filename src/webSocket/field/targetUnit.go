package field

import (
	"websocket-master"
	"strconv"
	"../../game/mechanics"

)

func TargetUnit(msg FieldMessage, ws *websocket.Conn)  {
	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client := usersFieldWs[ws]

	if find {
		coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
		passed := false
		for _, target := range coordinates {
			if target.X == msg.TargetX && target.Y == msg.TargetY {
				target, ok := client.HostileUnits[msg.TargetX][msg.TargetY]
				if ok {
					go mechanics.SetTarget(*unit, strconv.Itoa(target.X) + ":" + strconv.Itoa(target.Y), client.GameStat.Id)
					unit.Target = strconv.Itoa(target.X) + ":" + strconv.Itoa(target.Y)
					passed = true
					resp := FieldResponse{Event: msg.Event, UserName: client.Login}
					fieldPipe <- resp
					break
				}
			}
		}

		if passed {
			for _, gameUser := range client.Players {
				for _, user := range usersFieldWs {
					if user.Login != client.Login {
						if gameUser.Name == user.Login && client.GameStat.Id == user.GameStat.Id {
							hostileUnit, ok := user.HostileUnits[msg.X][msg.Y]
							if ok {
								hostileUnit.Target = strconv.Itoa(msg.TargetX) + ":" + strconv.Itoa(msg.TargetY)
							}
						}
					}
				}
			}
		} else {
			resp := FieldResponse{Event: msg.Event, UserName: client.Login, Error: "not allow"}
			fieldPipe <- resp
		}
	} else {
		resp := FieldResponse{Event: msg.Event, UserName: client.Login, Error: "unit not found"}
		fieldPipe <- resp
	}
}