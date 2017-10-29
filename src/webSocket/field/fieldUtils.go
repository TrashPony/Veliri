package field

import (
	"websocket-master"
	"log"
	"../../game/mechanics"
	"../../game/objects"
	"strconv"
	"errors"
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

func subtraction(slice1 []*objects.Coordinate, slice2 []*objects.Coordinate) []objects.Coordinate  {
	mb := map[objects.Coordinate]bool{}
	for _, x := range slice2 {
		mb[*x] = true
	}
	ab := []objects.Coordinate{}
	for _, x := range slice1 {
		if _, ok := mb[*x]; !ok {
			ab = append(ab, *x)
		}
	}
	return ab
}

func sendPermissionCoordinates(idGame int, ws *websocket.Conn, unit *objects.Unit) ( map[string]*objects.Coordinate, map[string]*objects.Unit, error) {
	units := objects.GetAllUnits(idGame)
	unitsCoordinate := make(map[string]*objects.Unit)
	allCoordinate :=  make(map[string]*objects.Coordinate)
	login := usersFieldWs[ws].Login
	respawn := usersFieldWs[ws].Respawn

	if login == unit.NameUser {
		PermissCoordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.WatchZone)
		for i := 0; i < len(PermissCoordinates); i++ {
			if !(PermissCoordinates[i].X == respawn.X && PermissCoordinates[i].Y == respawn.Y) {
				allCoordinate[strconv.Itoa(PermissCoordinates[i].X) + ":" + strconv.Itoa(PermissCoordinates[i].Y)] = PermissCoordinates[i]
			}
			for j := 0; j < len(units); j++ {
				if (PermissCoordinates[i].X == units[j].X) && (PermissCoordinates[i].Y == units[j].Y) {
					unitsCoordinate[strconv.Itoa(units[j].X) + ":" + strconv.Itoa(units[j].Y)] = &units[j]
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, nil
}

func findUnit(msg FieldMessage, ws *websocket.Conn) (*objects.Unit, bool) {
	var findUnit objects.Unit
	for key, unit := range usersFieldWs[ws].Units {
		//hostileUnit := unit.WatchUnit
		if msg.X == unit.X && msg.Y == unit.Y {
			return  usersFieldWs[ws].Units[key], true
			break
		}
		/*for _, hostile := range hostileUnit {
			if msg.X == strconv.Itoa(hostile.X) && msg.Y == strconv.Itoa(hostile.Y) {
				return &hostile, true // TODO: если раскоментировать то юниты теряют все внутрение свойства :\
				break
			}
		}*/
	}
	return &findUnit, false
}

func SendWatchCoordinate(ws *websocket.Conn, unit *objects.Unit){
	for _, coordinate := range unit.Watch {
		var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: coordinate.X, Y: coordinate.Y}
		fieldPipe <- emptyCoordinates
	}

	for _, unit := range unit.WatchUnit {
		var unitsParametr = FieldResponse{Event: "InitUnit", UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
			HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), X: unit.X, Y: unit.Y}
		fieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
	}
}