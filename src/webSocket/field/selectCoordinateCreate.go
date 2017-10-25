package field

import (
	"websocket-master"
	"strconv"
	"../../game/objects"
)

func SelectCoordinateCreate(ws *websocket.Conn)  {
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	units := usersFieldWs[ws].Units
	unitsCoordinate := objects.GetUnitsCoordinate(units)
	responseCoordinate := subtraction(coordinates, unitsCoordinate)

	for i := 0; i < len(responseCoordinate); i++ {
		if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
			var createCoordinates= FieldResponse{Event: "SelectCoordinateCreate", UserName: usersFieldWs[ws].Login, X: strconv.Itoa(responseCoordinate[i].X), Y: strconv.Itoa(responseCoordinate[i].Y)}
			fieldPipe <- createCoordinates
		}
	}
}
