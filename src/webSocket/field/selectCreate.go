package field

import (
	"../../game/objects"
	"github.com/gorilla/websocket"
)

func SelectCoordinateCreate(ws *websocket.Conn) {
	var coordinates []*objects.Coordinate
	respawn := usersFieldWs[ws].Respawn

	for _, coordinate := range usersFieldWs[ws].CreateZone {
		_, ok := usersFieldWs[ws].Units[coordinate.X][coordinate.Y]
		if !ok {
			coordinates = append(coordinates, coordinate)
		}
	}

	for i := 0; i < len(coordinates); i++ {
		if !(coordinates[i].X == respawn.X && coordinates[i].Y == respawn.Y) {
			var createCoordinates = FieldResponse{Event: "SelectCoordinateCreate", UserName: usersFieldWs[ws].Login, X: coordinates[i].X, Y: coordinates[i].Y}
			fieldPipe <- createCoordinates
		}
	}
}
