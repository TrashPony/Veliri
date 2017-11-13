package field

import (
	"websocket-master"
	"strconv"
	"../../game/mechanics"
	"../../game/objects"
)

func CreateUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	client, ok := usersFieldWs[ws]
	if !ok {
		delete(usersFieldWs, ws)
	} else {
		_, ok := coordinates[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
		if ok && !(msg.X == respawn.X && msg.Y == respawn.Y) {
			unit, price, createError := mechanics.CreateUnit(msg.IdGame, strconv.Itoa(usersFieldWs[ws].Id), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {

				units := objects.GetAllUnits(msg.IdGame)
				unit.Watch, unit.WatchUnit, _ = PermissionCoordinates(*client, unit, units)

				for _, coordinate := range unit.Watch {
					var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: coordinate.X, Y: coordinate.Y}
					fieldPipe <- emptyCoordinates
				}

				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, PlayerPrice: price, X: unit.X, Y: unit.Y}
				fieldPipe <- resp

				/*for _, userStat := range usersFieldWs[ws].Players {
					if userStat.Name == usersFieldWs[ws].Login {
						var unitsParametr = InitUnit{Event: "InitUnit", UserName: userStat.Name, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
							HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), X: unit.X, Y: unit.Y}
						initUnit <- unitsParametr
					}
				}*/

				for _, xLine := range unit.WatchUnit {
					for _, unit := range xLine {
						var unitsParametr = InitUnit{Event: "InitUnit", UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
							HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y}
						initUnit <- unitsParametr
					}
				}

				usersFieldWs[ws].addUnit(unit)

			} else {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: createError.Error()}
				fieldPipe <- resp
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
			fieldPipe <- resp
		}
	}
}