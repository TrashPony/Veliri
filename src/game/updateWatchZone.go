package game

import "strconv"

func updateOpenCoordinate(client *Player, oldWatchZone map[string]map[string]*Coordinate) (openCoordinate []*Coordinate, closeCoordinate []*Coordinate){
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

func updateHostileUnit(client *Player, oldWatchUnit map[string]map[string]*Unit) (openUnit []*Unit, closeUnit []*Unit) {
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

func updateHostileStrcuture(client *Player, oldWatchHostileMatherShip map[string]map[string]*MatherShip) (openMatherShip []*MatherShip, closeMatherShip []*MatherShip) {
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

func parseCloseCoordinate(closeCoordinate []*Coordinate, closeUnit []*Unit, closeMatherShip []*MatherShip, game *Game) ([]*Coordinate)  {

	for _, unit := range closeUnit {
		//coordinate, find := game.GetMap().GetCoordinate(unit.X, unit.Y)
		//if find { TODO полностью инициализировать карту.....
		coordinate := Coordinate{X: unit.X, Y:unit.Y}
		closeCoordinate = append(closeCoordinate, &coordinate)
		//}
	}

	for _, matherShip := range closeMatherShip { // TODO я не понимаю что тут происходит >_<
		coordinate, find := game.GetMap().GetCoordinate(matherShip.X, matherShip.Y)
		if find {
			closeCoordinate = append(closeCoordinate, coordinate)
		}
	}

	return closeCoordinate
}
