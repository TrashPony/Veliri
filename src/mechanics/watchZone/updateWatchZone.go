package watchZone

import (
	"strconv"
	"../game"
	"../player"
	"../matherShip"
	"../unit"
	"../coordinate"
)


func GetAllWatchObject(activeGame *game.Game, client *player.Player) {

	for _, xLine := range activeGame.GetUnits() {
		for _, gameUnit := range xLine {

			watchCoordinate, watchUnit, watchMatherShip, err := Watch(gameUnit, client.GetLogin(), activeGame) //PermissionCoordinates(client, unit, units)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.AddUnit(gameUnit)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileMatherShip(hostile)
						}
					}
				}

				for _, gameCoordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[gameCoordinate.X][gameCoordinate.Y]
					if !ok {
						client.AddCoordinate(gameCoordinate)
					}
				}
			}
		}
	}

	for _, xLine := range activeGame.GetMatherShips() {
		for _, gameMatherShip := range xLine {

			watchCoordinate, watchUnit, watchMatherShip, err := Watch(gameMatherShip, client.GetLogin(), activeGame)

			if err != nil { // если структура не моя то пропускаем дальнейшее действие
				continue
			} else {
				client.AddMatherShips(gameMatherShip)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileMatherShip(hostile)
						}
					}
				}

				for _, gameCoordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[gameCoordinate.X][gameCoordinate.Y]
					if ok {
						client.AddCoordinate(gameCoordinate)
					}
				}
			}
		}
	}
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

	GetAllWatchObject(activeGame, client)

	openCoordinate, closeCoordinate := updateOpenCoordinate(client, oldWatchZone)
	openUnit, closeUnit := updateHostileUnit(client, oldWatchHostileUnits)
	openMatherShip, closeMatherShip := updateHostileStrcuture(client, oldWatchHostileMatherShips)

	sendCloseCoordinate := parseCloseCoordinate(closeCoordinate, closeUnit, closeMatherShip, activeGame)

	updaterWatchZone.CloseCoordinate = sendCloseCoordinate
	updaterWatchZone.OpenCoordinate = openCoordinate
	updaterWatchZone.OpenUnit = openUnit
	updaterWatchZone.OpenMatherShip = openMatherShip

	return &updaterWatchZone
}

func updateOpenCoordinate(client *player.Player, oldWatchZone map[string]map[string]*coordinate.Coordinate) (openCoordinate []*coordinate.Coordinate, closeCoordinate []*coordinate.Coordinate){
	for _, xLine := range client.GetWatchCoordinates() { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[strconv.Itoa(newCoordinate.X)][strconv.Itoa(newCoordinate.Y)]
			if !ok && newCoordinate.X >= 0 && newCoordinate.Y >= 0 {
				openCoordinate = append(openCoordinate, newCoordinate)
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.GetWatchCoordinate(oldCoordinate.X, oldCoordinate.Y)
			_, findUnit := client.GetUnit(oldCoordinate.X, oldCoordinate.Y)
			if !find && !findUnit{
				client.DelWatchCoordinate(oldCoordinate.X, oldCoordinate.Y)
				closeCoordinate = append(closeCoordinate, oldCoordinate)
			}
		}
	}
	return
}

func updateHostileUnit(client *player.Player, oldWatchUnit map[string]map[string]*unit.Unit) (openUnit []*unit.Unit, closeUnit []*unit.Unit) {
	for _, xLine := range client.GetHostileUnits() { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)]
			if !ok {
				openUnit = append(openUnit, hostile)
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.GetHostileUnit(hostile.X, hostile.Y)
			if !find {
				closeUnit = append(closeUnit, hostile)
			}
		}
	}

	return
}

func updateHostileStrcuture(client *player.Player, oldWatchHostileMatherShip map[string]map[string]*matherShip.MatherShip) (openMatherShip []*matherShip.MatherShip, closeMatherShip []*matherShip.MatherShip) {
	for _, xLine := range client.GetHostileMatherShips() { // добавляем новые вражеские структуры которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchHostileMatherShip[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)]
			if !ok {
				openMatherShip = append(openMatherShip, hostile)
			}
		}
	}

	for _, xLine := range oldWatchHostileMatherShip {
		for _, hostile := range xLine {
			_, find := client.GetHostileMatherShip(hostile.X, hostile.Y)
			if !find {
				closeMatherShip = append(closeMatherShip, hostile)
			}
		}
	}
	return
}

func parseCloseCoordinate(closeCoordinate []*coordinate.Coordinate, closeUnit []*unit.Unit, closeMatherShip []*matherShip.MatherShip, game *game.Game) ([]*coordinate.Coordinate)  {

	for _, unit := range closeUnit {
		coordinate, find := game.GetMap().GetCoordinate(unit.X, unit.Y)
		if find {
			closeCoordinate = append(closeCoordinate, coordinate)
		}
	}

	for _, matherShip := range closeMatherShip { // TODO я не понимаю что тут происходит >_<
		coordinate, find := game.GetMap().GetCoordinate(matherShip.X, matherShip.Y)
		if find {
			closeCoordinate = append(closeCoordinate, coordinate)
		}
	}

	return closeCoordinate
}

type UpdaterWatchZone struct {
	CloseCoordinate []*coordinate.Coordinate `json:"close_coordinate"`
	OpenCoordinate  []*coordinate.Coordinate `json:"open_coordinate"`
	OpenUnit        []*unit.Unit       `json:"open_unit"`
	OpenMatherShip  []*matherShip.MatherShip `json:"open_mather_ship"`
}