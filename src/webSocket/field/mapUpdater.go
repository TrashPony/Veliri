package field

import (
	"../../game"
)

func UpdateWatchZone(client *game.Player, activeGame *game.Game, updaterWatchZone *game.UpdaterWatchZone)  {

	if updaterWatchZone == nil {
		updaterWatchZone = client.UpdateWatchZone(activeGame)
	}

	updateMyUnit(client)
	updateMyStructure(client)
	sendNewHostileUnit(updaterWatchZone.OpenUnit, client.GetLogin())
	sendNewHostileStructure(updaterWatchZone.OpenStructure, client.GetLogin())
	UpdateOpenCoordinate(updaterWatchZone.OpenCoordinate, updaterWatchZone.CloseCoordinate, client.GetLogin())
}

func updateMyUnit(client *game.Player)  {
	var unitsParameter InitUnit
	for _, xLine := range client.GetUnits() { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit("InitUnit", unit, client.GetLogin())
		}
	}
}

func updateMyStructure(client *game.Player)  {
	var structureParameter InitStructure
	for _, xLine := range client.GetStructures() { // отправляем параметры своих структур
		for _, structure := range xLine {
			structureParameter.initStructure("InitStructure", structure, client.GetLogin())
		}
	}
}

func sendNewHostileUnit(units []*game.Unit, login string )  {
	var UnitParams InitUnit
	for _, unit := range units {
		UnitParams.initUnit("InitUnit", unit, login)
	}
}

func sendNewHostileStructure(structures []*game.MatherShip, login string )  {
	var StructureParams InitStructure
	for _, structure := range structures {
		StructureParams.initStructure("InitStructure", structure, login)
	}
}

func UpdateOpenCoordinate(openCoordinates []*game.Coordinate, closeCoordinates []*game.Coordinate, login string)  {
	for _, closeCoor := range closeCoordinates {
		closeCoordinate(login, closeCoor.X, closeCoor.Y)
	}

	for _, open := range openCoordinates {
		openCoordinate(login, open.X, open.Y)
	}
}
