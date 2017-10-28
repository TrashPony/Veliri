package field

import (
	"websocket-master"
	"strconv"
	"../../game/mechanics"
	"../../game/objects"
)

func CreateUnit(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	var errorMsg bool = true
	for i := 0; i < len(coordinates); i++ {
		if strconv.Itoa(coordinates[i].X) == msg.X && strconv.Itoa(coordinates[i].Y) == msg.Y &&
			!(msg.X == strconv.Itoa(respawn.X) && msg.Y == strconv.Itoa(respawn.Y)){
			errorMsg = false
			unit, price, createError := mechanics.CreateUnit(msg.IdGame, strconv.Itoa(usersFieldWs[ws].Id), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {
				for _, userStat := range usersFieldWs[ws].Players {
					if userStat.Name == usersFieldWs[ws].Login {
						resp = FieldResponse{Event: msg.Event, UserName: userStat.Name, PlayerPrice: strconv.Itoa(price), X: strconv.Itoa(unit.X), Y: strconv.Itoa(unit.Y), TypeUnit: unit.NameType, UserOwned: usersFieldWs[ws].Login}
						fieldPipe <- resp
					}
				}
				var err error
				unit.Watch, unit.WatchUnit, err = sendPermissionCoordinates(msg.IdGame, ws, unit)
				if err != nil {
					break
				}
				for _, coordinate := range unit.Watch{
					var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: strconv.Itoa(coordinate.X), Y: strconv.Itoa(coordinate.Y)}
					fieldPipe <- emptyCoordinates
				}

				for _, unit := range unit.WatchUnit {
					var unitsParametr = FieldResponse{Event: "InitUnit", UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
						HP: strconv.Itoa(unit.Hp), UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), X: strconv.Itoa(unit.X), Y: strconv.Itoa(unit.Y)}
					fieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
				}

				usersFieldWs[ws].Units[objects.Coordinate{X: unit.X, Y:unit.Y}] = unit
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: createError.Error()}
				fieldPipe <- resp
			}
			break
		}
	}
	if errorMsg {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
		fieldPipe <- resp
	}
}