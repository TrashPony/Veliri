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
	game := Games[client.GameID]

	if !ok {
		delete(usersFieldWs, ws)
	} else {
		_, ok := coordinates[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
		if ok && !(msg.X == respawn.X && msg.Y == respawn.Y) {

			var unit objects.Unit
			unit, price, createError := mechanics.CreateUnit(msg.IdGame, strconv.Itoa(usersFieldWs[ws].Id), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {

				client.addUnit(&unit)
				game.addUnit(&unit)

				watchCoordinate, WatchUnit, _ := PermissionCoordinates(client, &unit, game.getUnits())

				for _, coordinate := range watchCoordinate {
					client.addCoordinate(coordinate) // TODO бага с тем что не затирается зона строительства если нехватает дальнссти видимости
					var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: coordinate.X, Y: coordinate.Y}
					fieldPipe <- emptyCoordinates
				}

				for _, xLine := range WatchUnit {    // TODO бага фазой хождения сразу после покупки юнитов что они становяться черным квадратом)
					for _, unit := range xLine {
						if client.Login != unit.NameUser {
							client.addHostileUnit(unit)
						}
					}
				}

				var unitsParameter InitUnit
				unitsParameter.initUnit(&unit, client.Login)

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