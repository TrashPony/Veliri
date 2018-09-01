package watchZone

import (
	"../../../player"
	"../../../gameObjects/coordinate"
	"strconv"
)

func updateOpenCoordinate(client *player.Player, oldWatchZone map[string]map[string]*coordinate.Coordinate) (openCoordinate []*coordinate.Coordinate, closeCoordinate []*coordinate.Coordinate) {
	for _, xLine := range client.GetWatchCoordinates() { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[strconv.Itoa(newCoordinate.Q)][strconv.Itoa(newCoordinate.R)]
			if !ok && newCoordinate.Q >= 0 && newCoordinate.R >= 0 {
				openCoordinate = append(openCoordinate, newCoordinate)
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.GetWatchCoordinate(oldCoordinate.Q, oldCoordinate.R)
			_, findUnit := client.GetUnit(oldCoordinate.Q, oldCoordinate.R)
			if !find && !findUnit {
				client.DelWatchCoordinate(oldCoordinate.Q, oldCoordinate.R)
				closeCoordinate = append(closeCoordinate, oldCoordinate)
			}
		}
	}
	return
}
