package field

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func MouseOver(msg FieldMessage, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]
	unit, find := client.GetUnit(msg.X, msg.Y)
	if !find {
		unit, find = client.GetHostileUnit(msg.X,msg.Y)
	}

	if find && ok{
		if unit.Target == nil {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: unit.Hp,
				UnitAction: strconv.FormatBool(unit.Action), Target: "", Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
				Init: strconv.Itoa(unit.Initiative), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
				AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
			initUnit <- resp
		} else {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: unit.Hp,
				UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target.X) + ":" + strconv.Itoa(unit.Target.Y), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
				Init: strconv.Itoa(unit.Initiative), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
				AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
			initUnit <- resp
		}
	} else {
		structure, find := usersFieldWs[ws].GetStructure(msg.X, msg.Y)
		if !find {
			structure, find = usersFieldWs[ws].GetHostileStructure(msg.X, msg.Y)
		}
		if find {
			resp := InitUnit{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), TypeUnit: structure.Type, UserOwned:structure.NameUser, RangeView: strconv.Itoa(structure.WatchZone)}
			initUnit <- resp
		}
	}
}
