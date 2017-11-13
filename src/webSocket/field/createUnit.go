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
				watchCoordinate, WatchUnit, _ := PermissionCoordinates(client, unit, units)
				client.addUnit(unit)

				for _, coordinate := range watchCoordinate {
					client.addCoordinate(coordinate)
					var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: coordinate.X, Y: coordinate.Y}
					fieldPipe <- emptyCoordinates
				}

				for _, xLine := range WatchUnit {
					for _, unit := range xLine {
						if client.Login == unit.NameUser {
							client.addUnit(unit)
						} else {
							client.addHostileUnit(unit)
						}
						var unitsParameters = InitUnit{Event: "InitUnit", UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
							HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y}
						initUnit <- unitsParameters

					}
				}

				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, PlayerPrice: price, X: unit.X, Y: unit.Y}
				fieldPipe <- resp

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