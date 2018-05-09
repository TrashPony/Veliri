package field

import (
	"../../game"
)

func UpdateWatchZone(client *game.Player, activeGame *game.Game, updaterWatchZone *game.UpdaterWatchZone) {

	if updaterWatchZone == nil {
		updaterWatchZone = client.UpdateWatchZone(activeGame)
	}

	updateMyUnit(client)
	updateMyMatherShip(client)
	sendNewHostileUnit(updaterWatchZone.OpenUnit, client.GetLogin())
	sendNewHostileMatherShip(updaterWatchZone.OpenStructure, client.GetLogin())
	UpdateOpenCoordinate(updaterWatchZone.OpenCoordinate, updaterWatchZone.CloseCoordinate, client.GetLogin())
}

func updateMyUnit(client *game.Player) {
	var unitsParameter InitUnit
	for _, xLine := range client.GetUnits() { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit("InitUnit", unit, client.GetLogin())
		}
	}
}

func updateMyMatherShip(client *game.Player) {
	var matherShipParameter InitStructure
	matherShipParameter.initMatherShip("InitStructure", client.GetMatherShip(), client.GetLogin())
}

func sendNewHostileUnit(units []*game.Unit, login string) {
	var UnitParams InitUnit
	for _, unit := range units {
		UnitParams.initUnit("InitUnit", unit, login)
	}
}

func sendNewHostileMatherShip(structures []*game.MatherShip, login string) {
	var matherShipParameter InitStructure
	for _, structure := range structures {
		matherShipParameter.initMatherShip("InitStructure", structure, login)
	}
}

func UpdateOpenCoordinate(openCoordinates []*game.Coordinate, closeCoordinates []*game.Coordinate, login string) {
	for _, closeCoor := range closeCoordinates {
		closeCoordinate(login, closeCoor.X, closeCoor.Y)
	}

	for _, open := range openCoordinates {
		openCoordinate(login, open.X, open.Y)
	}
}
