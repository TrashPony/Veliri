package field

import (
	"websocket-master"
	"strconv"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn) {
	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	if !find {
		unit, find = usersFieldWs[ws].HostileUnits[msg.X][msg.Y]
	}

	if find {
		resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: unit.Hp,
			UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
			Init: strconv.Itoa(unit.Init), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
			AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
		initUnit <- resp
	}
}
