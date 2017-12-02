package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"strconv"
)

func CreateUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	client, ok := usersFieldWs[ws]
	activeGame := Games[client.GameID]

	if !ok {
		delete(usersFieldWs, ws)
	} else {
		_, ok := coordinates[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
		if ok && !(msg.X == respawn.X && msg.Y == respawn.Y) {
			var unit game.Unit
			unit, price, createError := game.CreateUnit(msg.IdGame, strconv.Itoa(usersFieldWs[ws].Id), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {
				activeGame.SetUnit(&unit)
				client.updateWatchZone(activeGame)
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, PlayerPrice: price, X: unit.X, Y: unit.Y}
				fieldPipe <- resp

				var unitsParameter InitUnit
				unitsParameter.initUnit(&unit, client.Login)
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
