package field

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn) {
	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	if !find {
		unit, find = usersFieldWs[ws].HostileUnits[msg.X][msg.Y]
	}

	if find {
		if unit.Target == nil {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: unit.Hp,
				UnitAction: strconv.FormatBool(unit.Action), Target: "", Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
				Init: strconv.Itoa(unit.Initiative), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
				AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
			initUnit <- resp
		} else {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: unit.Hp,
				UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target.X) + ":" + strconv.Itoa(unit.Target.Y), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
				Init: strconv.Itoa(unit.Initiative), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
				AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
			initUnit <- resp
		}
	} else {
		structure, find := usersFieldWs[ws].Structure[msg.X][msg.Y]
		if !find {
			structure, find = usersFieldWs[ws].HostileStructure[msg.X][msg.Y]
		}
		if find {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].Login, TypeUnit: structure.Type, UserOwned:structure.NameUser, RangeView: strconv.Itoa(structure.WatchZone)}
			initUnit <- resp
		}
	}
}
