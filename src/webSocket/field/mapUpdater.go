package field

import (
	"../../game"
)

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
	for _, closeCoor := range closeCoordinates {
		closeCoordinate(login, closeCoor.X, closeCoor.Y)
	}

	for _, open := range openCoordinates {
		openCoordinate(login, open.X, open.Y)
	}
}
