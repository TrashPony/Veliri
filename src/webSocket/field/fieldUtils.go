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

func subtraction(slice1 []*objects.Coordinate, slice2 []*objects.Coordinate) (ab []objects.Coordinate)  {
	mb := map[objects.Coordinate]bool{}
	for _, x := range slice2 {
		mb[*x] = true
	}
	for _, x := range slice1 {
		if _, ok := mb[*x]; !ok {
			ab = append(ab, *x)
		}
	}
	return ab
}

func PermissionCoordinates(client Clients, unit *objects.Unit, units map[string]*objects.Unit) (allCoordinate map[string]*objects.Coordinate, unitsCoordinate map[int]map[int]*objects.Unit, Err error) {
	allCoordinate = make(map[string]*objects.Coordinate)
	unitsCoordinate =  make(map[int]map[int]*objects.Unit)
	login := client.Login
	respawn := client.Respawn

	if login == unit.NameUser {
		PermissCoordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.WatchZone)
		for i := 0; i < len(PermissCoordinates); i++ {
			if !(PermissCoordinates[i].X == respawn.X && PermissCoordinates[i].Y == respawn.Y) {
				allCoordinate[strconv.Itoa(PermissCoordinates[i].X) + ":" + strconv.Itoa(PermissCoordinates[i].Y)] = PermissCoordinates[i]
			}
			x := strconv.Itoa(PermissCoordinates[i].X)
			y := strconv.Itoa(PermissCoordinates[i].Y)
			unitInMap, ok := units[x + ":"+ y]
			if ok {
				if unitsCoordinate[PermissCoordinates[i].X] != nil {
					unitsCoordinate[PermissCoordinates[i].X][PermissCoordinates[i].Y] = unitInMap
				} else {
					unitsCoordinate[PermissCoordinates[i].X] = make(map[int]*objects.Unit)
					unitsCoordinate[PermissCoordinates[i].X][PermissCoordinates[i].Y] = unitInMap
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, nil
}

func SendWatchCoordinate(ws *websocket.Conn, unit *objects.Unit){
	for _, coordinate := range unit.Watch {
		var emptyCoordinates = InitUnit{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: coordinate.X, Y: coordinate.Y}
		initUnit <- emptyCoordinates
	}

	for _, xLine := range unit.WatchUnit {
		for _, unit := range xLine{
			var unitsParametr = InitUnit{Event: "InitUnit", UserName: usersFieldWs[ws].Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
			HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y}
			initUnit <- unitsParametr // отправляем параметры каждого юнита отдельно
		}
	}
}

func ActionGameUser(players []objects.UserStat)  (activeUser []*Clients) {
	for _, clients := range usersFieldWs {
		add := false
		for _, userStat := range players {
			if clients.Login == userStat.Name && clients.GameStat.Id == userStat.IdGame {
				add = true
			}
		}
		if add {
			activeUser = append(activeUser, clients)
		}
	}
	return
}