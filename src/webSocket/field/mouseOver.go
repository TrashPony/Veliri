package field

import (
	"websocket-master"
	"strconv"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	unit, find := findUnit(msg, ws)
	if find {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: strconv.Itoa(unit.Hp),
			UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
			Init: strconv.Itoa(unit.Init), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
			AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
		fieldPipe <- resp
	}
}
