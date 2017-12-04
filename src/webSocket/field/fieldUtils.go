package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"log"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*game.Player) {
	for ws, client := range *usersWs {
		if client.GetLogin() == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*game.Player, err error) {
	log.Printf("error: %v", err)
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func subtraction(slice1 []*game.Coordinate, slice2 []*game.Coordinate) (ab []game.Coordinate) {
	mb := map[game.Coordinate]bool{}
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

func ActionGameUser(players []*game.UserStat) (activeUser []*game.Player) {
	for _, clients := range usersFieldWs {
		add := false
		for _, userStat := range players {
			if clients.GetLogin() == userStat.Name && clients.GetGameID() == userStat.IdGame {
				add = true
			}
		}
		if add {
			activeUser = append(activeUser, clients)
		}
	}
	return
}

func UpdateWatchZone(client *game.Player, activeGame *game.Game)  {
	closeCoordinate, openCoordinate, openUnit, openStructure := client.UpdateWatchZone(activeGame)

	updateMyUnit(client)
	updateMyStructure(client)

	sendNewHostileUnit(openUnit, client.GetLogin())
	sendNewHostileStructure(openStructure, client.GetLogin())
	UpdateOpenCoordinate(openCoordinate, closeCoordinate, client.GetLogin())
}

func updateMyUnit(client *game.Player)  {
	var unitsParameter InitUnit
	for _, xLine := range client.GetUnits() { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit(unit, client.GetLogin())
		}
	}
}

func updateMyStructure(client *game.Player)  {
	var structureParameter InitStructure
	for _, xLine := range client.GetStructures() { // отправляем параметры своих структур
		for _, structure := range xLine {
			structureParameter.initStructure(structure, client.GetLogin())
		}
	}
}

func sendNewHostileUnit(units []*game.Unit, login string )  {
	var UnitParams InitUnit
	for _, unit := range units {
		UnitParams.initUnit(unit, login)
	}
}

func sendNewHostileStructure(structures []*game.Structure, login string )  {
	var StructureParams InitStructure
	for _, structure := range structures {
		StructureParams.initStructure(structure, login)
	}
}

func UpdateOpenCoordinate(openCoordinates []*game.Coordinate, closeCoordinates []*game.Coordinate, login string)  {
	for _, open := range openCoordinates {
		openCoordinate(login, open.X, open.Y)
	}

	for _, closeCoor := range closeCoordinates {
		closeCoordinate(login, closeCoor.X, closeCoor.Y)
	}
}