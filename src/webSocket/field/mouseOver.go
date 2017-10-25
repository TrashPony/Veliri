package field

import (
	"websocket-master"
	"strconv"
	"../../game/objects"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].PermittedCoordinates
	units := usersFieldWs[ws].Units
	for i := 0; i < len(coordinates); i++{
		if strconv.Itoa(coordinates[i].X) == msg.X && strconv.Itoa(coordinates[i].Y) == msg.Y {
			for j := 0; j < len(units); j++ {
				if coordinates[i].X == units[j].X && coordinates[i].Y == units[j].Y {
					unit, errUnitParams := objects.GetXYUnits(msg.IdGame, msg.X, msg.Y)
					if errUnitParams == nil {
						resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: strconv.Itoa(unit.Hp),
							UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
							Init: strconv.Itoa(unit.Init), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
							AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
						fieldPipe <- resp
					}
				}
			}
		}
	}
}
