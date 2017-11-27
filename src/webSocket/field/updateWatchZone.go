package field

import (
	"../../game/objects"
)

func updateMyUnit(client *Clients)  {
	var unitsParameter InitUnit
	for _, xLine := range client.Units { // отправляем параметры своих юнитов
		for _, unit := range xLine {
			unitsParameter.initUnit(unit, client.Login)
		}
	}
}

func updateMyStructure(client *Clients)  {
	var structureParameter InitStructure
	for _, xLine := range client.Structure { // отправляем параметры своих структур
		for _, structure := range xLine {
			structureParameter.initStructure(structure, client.Login)
		}
	}
}

func updateOpenCoordinate(client *Clients, oldWatchZone map[int]map[int]*objects.Coordinate) {
	for _, xLine := range client.Watch { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
			if !ok {
				client.addCoordinate(newCoordinate)
				openCoordinate(client.Login, newCoordinate.X, newCoordinate.Y)
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.Watch[oldCoordinate.X][oldCoordinate.Y]
			if !find {
				delete(client.Watch[oldCoordinate.X], oldCoordinate.Y)
				closeCoordinate(client.Login, oldCoordinate.X, oldCoordinate.Y)
			}
		}
	}
}

func updateHostileUnit(client *Clients, oldWatchUnit map[int]map[int]*objects.Unit) {
	for _, xLine := range client.HostileUnits { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
			if !ok {
				client.addHostileUnit(hostile)   // добавляем вражеского юнита в зону видимости
				var unitsParameter InitUnit
				unitsParameter.initUnit(hostile, client.Login)
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.HostileUnits[hostile.X][hostile.Y]
			if !find {
				closeCoordinate(client.Login, hostile.X, hostile.Y)
			}
		}
	}
}

func updateHostileStrcuture(client *Clients, oldWatchHostileStructure map[int]map[int]*objects.Structure)  {
	for _, xLine := range client.HostileStructure { // добавляем новые вражеские структуры которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchHostileStructure[hostile.X][hostile.Y]
			if !ok {
				client.addHostileStructure(hostile) // добавляем вражескую видимую структуру в зону видимости
				var structureParams InitStructure
				structureParams.initStructure(hostile, client.Login)
			}
		}
	}

	for _, xLine := range oldWatchHostileStructure {
		for _, hostile := range xLine {
			_, find := client.HostileUnits[hostile.X][hostile.Y]
			if !find {
				closeCoordinate(client.Login, hostile.X, hostile.Y)
			}
		}
	}
}
