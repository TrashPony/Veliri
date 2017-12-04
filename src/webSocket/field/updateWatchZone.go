package field

import (
	"../../game"
)

func updateMyUnit(client *game.Player)  {
	var unitsParameter InitUnit
	for _, xLine := range client.Units { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit(unit, client.GetLogin())
		}
	}
}

func updateMyStructure(client *game.Player)  {
	var structureParameter InitStructure
	for _, xLine := range client.Structure { // отправляем параметры своих структур
		for _, structure := range xLine {
			structureParameter.initStructure(structure, client.GetLogin())
		}
	}
}

func updateOpenCoordinate(client *game.Player, oldWatchZone map[int]map[int]*game.Coordinate) {

	for _, xLine := range client.Watch { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
			if !ok && newCoordinate.X >= 0 && newCoordinate.Y >= 0 {
				openCoordinate(client.GetLogin(), newCoordinate.X, newCoordinate.Y)
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.Watch[oldCoordinate.X][oldCoordinate.Y]
			_, findUnit := client.Units[oldCoordinate.X][oldCoordinate.Y]
			if !find && !findUnit{
				delete(client.Watch[oldCoordinate.X], oldCoordinate.Y)
				closeCoordinate(client.GetLogin(), oldCoordinate.X, oldCoordinate.Y)
			}
		}
	}
}

func updateHostileUnit(client *game.Player, oldWatchUnit map[int]map[int]*game.Unit) {
	for _, xLine := range client.HostileUnits { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
			if !ok {
				var unitsParameter InitUnit
				unitsParameter.initUnit(hostile, client.GetLogin())
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.HostileUnits[hostile.X][hostile.Y]
			if !find {
				closeCoordinate(client.GetLogin(), hostile.X, hostile.Y)
			}
		}
	}
}

func updateHostileStrcuture(client *game.Player, oldWatchHostileStructure map[int]map[int]*game.Structure)  {
	for _, xLine := range client.HostileStructure { // добавляем новые вражеские структуры которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchHostileStructure[hostile.X][hostile.Y]
			if !ok {
				var structureParams InitStructure
				structureParams.initStructure(hostile, client.GetLogin())
			}
		}
	}

	for _, xLine := range oldWatchHostileStructure {
		for _, hostile := range xLine {
			_, find := client.HostileStructure[hostile.X][hostile.Y]
			if !find {
				closeCoordinate(client.GetLogin(), hostile.X, hostile.Y)
			}
		}
	}
}
