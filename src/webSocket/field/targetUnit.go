package field

func TargetUnit() {

	/*unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client := usersFieldWs[ws]
	activeGame := Games[client.GetGameID()]

	if find {
		coordinates := Mechanics.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
		passed := false

		for _, target := range coordinates {
			if target.X == msg.TargetX && target.Y == msg.TargetY {
				target, ok := client.GetHostileUnit(msg.TargetX, msg.TargetY)
				if ok {
					Mechanics.SetTarget(*unit, strconv.Itoa(target.X)+":"+strconv.Itoa(target.Y), activeGame.Id)
					unit.Target = &Mechanics.Coordinate{X: target.X, Y: target.Y}
					passed = true
					resp := Response{Event: msg.Event, UserName: client.GetLogin()}
					fieldPipe <- resp
					break
				}
			}
		}

		if !passed {
			resp := Response{Event: msg.Event, UserName: client.GetLogin(), Error: "not allow"}
			fieldPipe <- resp
		}
	} else {
		resp := Response{Event: msg.Event, UserName: client.GetLogin(), Error: "unit not found"}
		fieldPipe <- resp
	}*/
}
