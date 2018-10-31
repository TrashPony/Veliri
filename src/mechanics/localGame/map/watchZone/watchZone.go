package watchZone

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
)

type UpdaterWatchZone struct {
	CloseCoordinate []*coordinate.Coordinate `json:"close_coordinate"`
	OpenCoordinate  []*coordinate.Coordinate `json:"open_coordinate"`
	OpenUnit        []*unit.Unit             `json:"open_unit"`
}

// отправляем открытые ячейки, удаляем закрытые
func UpdateWatchZone(activeGame *localGame.Game, client *player.Player	) *UpdaterWatchZone {
	var updaterWatchZone UpdaterWatchZone

	oldWatchZone := client.GetWatchCoordinates()
	oldWatchHostileUnits := client.GetHostileUnits()

	client.SetUnits(nil)
	client.SetHostileUnits(nil)
	client.SetWatchCoordinates(nil)

	getAllWatchObject(activeGame, client)

	openCoordinate, closeCoordinate := updateOpenCoordinate(client, oldWatchZone)
	openUnit, closeUnit := updateHostileUnit(client, oldWatchHostileUnits)

	sendCloseCoordinate := parseCloseCoordinate(closeCoordinate, closeUnit, activeGame)

	updaterWatchZone.CloseCoordinate = sendCloseCoordinate
	updaterWatchZone.OpenCoordinate = openCoordinate
	updaterWatchZone.OpenUnit = openUnit

	return &updaterWatchZone
}
