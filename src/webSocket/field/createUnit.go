package field

import (
	"websocket-master"
	"strconv"
	"../../game/mechanics"
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
				unitPermissCoordinate := sendPermissionCoordinates(msg.IdGame, ws, unit)
				for i := 0; i < len(unitPermissCoordinate); i++ {
					usersFieldWs[ws].PermittedCoordinates = append(usersFieldWs[ws].PermittedCoordinates, unitPermissCoordinate[i])
				}
				usersFieldWs[ws].Units = append(usersFieldWs[ws].Units, unit)
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