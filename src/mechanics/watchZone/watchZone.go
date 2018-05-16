package watchZone

import (
	"../game"
	"../player"
	"../matherShip"
	"../unit"
	"../coordinate"
)

type UpdaterWatchZone struct {
	CloseCoordinate []*coordinate.Coordinate `json:"close_coordinate"`
	OpenCoordinate  []*coordinate.Coordinate `json:"open_coordinate"`
	OpenUnit        []*unit.Unit             `json:"open_unit"`
	OpenMatherShip  []*matherShip.MatherShip `json:"open_mather_ship"`
}

// отправляем открытые ячейки, удаляем закрытые
func UpdateWatchZone(activeGame *game.Game, client *player.Player) (*UpdaterWatchZone) {
	var updaterWatchZone UpdaterWatchZone

	oldWatchZone := client.GetWatchCoordinates()
	oldWatchHostileUnits := client.GetHostileUnits()
	oldWatchHostileMatherShips := client.GetHostileMatherShips()

	client.SetUnits(nil)
	client.SetMatherShip(nil)
	client.SetHostileUnits(nil)
	client.SetHostileMatherShips(nil)
	client.SetWatchCoordinates(nil)

	getAllWatchObject(activeGame, client)

	openCoordinate, closeCoordinate := updateOpenCoordinate(client, oldWatchZone)
	openUnit, closeUnit := updateHostileUnit(client, oldWatchHostileUnits)
	openMatherShip, closeMatherShip := updateHostileMatherShip(client, oldWatchHostileMatherShips)

	sendCloseCoordinate := parseCloseCoordinate(closeCoordinate, closeUnit, closeMatherShip, activeGame)

	updaterWatchZone.CloseCoordinate = sendCloseCoordinate
	updaterWatchZone.OpenCoordinate = openCoordinate
	updaterWatchZone.OpenUnit = openUnit
	updaterWatchZone.OpenMatherShip = openMatherShip

	return &updaterWatchZone
}
