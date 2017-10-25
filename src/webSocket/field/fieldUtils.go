package field

import (
	"websocket-master"
	"log"
	"../../game/mechanics"
	"../../game/objects"
	"strconv"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients)  {
	for ws, client  := range *usersWs {
		if client.Login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func subtraction(slice1 []objects.Coordinate, slice2 []objects.Coordinate) []objects.Coordinate  {
	mb := map[objects.Coordinate]bool{}
	for _, x := range slice2 {
		mb[x] = true
	}
	ab := []objects.Coordinate{}
	for _, x := range slice1 {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func sendPermissionCoordinates(idGame string, ws *websocket.Conn, unit objects.Unit) ([]objects.Coordinate) {
	units := objects.GetAllUnits(idGame)
	var allCoordinate []objects.Coordinate
	login := usersFieldWs[ws].Login
	respawn := usersFieldWs[ws].Respawn

	if login == unit.NameUser {
		PermissCoordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.WatchZone)
		for i := 0; i < len(PermissCoordinates); i++ {
			allCoordinate = append(allCoordinate, PermissCoordinates[i])
			if !(PermissCoordinates[i].X == respawn.X && PermissCoordinates[i].Y == respawn.Y) {
				var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: login, X: strconv.Itoa(PermissCoordinates[i].X), Y: strconv.Itoa(PermissCoordinates[i].Y)}
				fieldPipe <- emptyCoordinates
			}
			for j := 0; j < len(units); j++ {
				if (PermissCoordinates[i].X == units[j].X) && (PermissCoordinates[i].Y == units[j].Y) {
					var unitsParametr = FieldResponse{Event: "InitUnit", UserName: login, TypeUnit: units[j].NameType, UserOwned: units[j].NameUser,
						HP: strconv.Itoa(units[j].Hp), UnitAction: strconv.FormatBool(units[j].Action), Target: strconv.Itoa(units[j].Target), X: strconv.Itoa(units[j].X), Y: strconv.Itoa(units[j].Y)}
					fieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
				}
			}
		}
	}
	return allCoordinate
}