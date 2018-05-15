package field

import (
	"../../mechanics/unit"
	"../../mechanics/player"
	"../../mechanics/game"
	"../../mechanics/watchZone"
	"../../mechanics/matherShip"
	"../../mechanics/coordinate"
)

func UpdateWatchZone(client *player.Player, activeGame *game.Game, updaterWatchZone *watchZone.UpdaterWatchZone) {

	if updaterWatchZone == nil {
		updaterWatchZone = watchZone.UpdateWatchZone(activeGame, client)
	}

	updateMyUnit(client)
	updateMyMatherShip(client)
	sendNewHostileUnit(updaterWatchZone.OpenUnit, client.GetLogin())
	sendNewHostileMatherShip(updaterWatchZone.OpenMatherShip, client.GetLogin())
	UpdateOpenCoordinate(updaterWatchZone.OpenCoordinate, updaterWatchZone.CloseCoordinate, client.GetLogin())
}

func updateMyUnit(client *player.Player) {
	var unitsParameter InitUnit
	for _, xLine := range client.GetUnits() { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit("InitUnit", unit, client.GetLogin())
		}
	}
}

func updateMyMatherShip(client *player.Player) {
	var matherShipParameter InitStructure
	matherShipParameter.initMatherShip("InitStructure", client.GetMatherShip(), client.GetLogin())
}

func sendNewHostileUnit(units []*unit.Unit, login string) {
	var UnitParams InitUnit
	for _, unit := range units {
		UnitParams.initUnit("InitUnit", unit, login)
	}
}

func sendNewHostileMatherShip(structures []*matherShip.MatherShip, login string) {
	var matherShipParameter InitStructure
	for _, structure := range structures {
		matherShipParameter.initMatherShip("InitStructure", structure, login)
	}
}

func UpdateOpenCoordinate(openCoordinates []*coordinate.Coordinate, closeCoordinates []*coordinate.Coordinate, login string) {
	for _, closeCoor := range closeCoordinates {
		closeCoordinate(login, closeCoor.X, closeCoor.Y)
	}

	for _, open := range openCoordinates {
		openCoordinate(login, open.X, open.Y)
	}
}
