package game

func updateOpenCoordinate(client *Player, oldWatchZone map[int]map[int]*Coordinate) (openCoordinate []*Coordinate, closeCoordinate []*Coordinate){
	for _, xLine := range client.GetWatchCoordinates() { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
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

func updateHostileUnit(client *Player, oldWatchUnit map[int]map[int]*Unit) (openUnit []*Unit, closeUnit []*Unit) {
	for _, xLine := range client.GetHostileUnits() { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
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

func updateHostileStrcuture(client *Player, oldWatchHostileStructure map[int]map[int]*Structure) (openStructure []*Structure, closeStructure []*Structure) {
	for _, xLine := range client.GetHostileStructures() { // добавляем новые вражеские структуры которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchHostileStructure[hostile.X][hostile.Y]
			if !ok {
				openStructure = append(openStructure, hostile)
			}
		}
	}

	for _, xLine := range oldWatchHostileStructure {
		for _, hostile := range xLine {
			_, find := client.GetHostileStructure(hostile.X, hostile.Y)
			if !find {
				closeStructure = append(closeStructure, hostile)
			}
		}
	}
	return
}

func parseCloseCoordinate(closeCoordinate []*Coordinate, closeUnit []*Unit, closeStructure []*Structure, game *Game) ([]*Coordinate)  {

	for _, unit := range closeUnit {
		//coordinate, find := game.GetMap().GetCoordinate(unit.X, unit.Y)
		//if find { TODO полностью инициализировать карту
		coordinate := Coordinate{X: unit.X, Y:unit.Y}
		closeCoordinate = append(closeCoordinate, &coordinate)
		//}
	}

	for _, structure := range closeStructure { // TODO я не понимаю что тут происходит >_<
		coordinate, find := game.GetMap().GetCoordinate(structure.X, structure.Y)
		if find {
			closeCoordinate = append(closeCoordinate, coordinate)
		}
	}

	return closeCoordinate
}
