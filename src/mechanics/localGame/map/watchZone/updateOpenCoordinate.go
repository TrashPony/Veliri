package watchZone

import (
	"../../../player"
	"../coordinate"
	"strconv"
)

func updateOpenCoordinate(client *player.Player, oldWatchZone map[string]map[string]*coordinate.Coordinate) (openCoordinate []*coordinate.Coordinate, closeCoordinate []*coordinate.Coordinate) {
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
			if !find && !findUnit {
				client.DelWatchCoordinate(oldCoordinate.X, oldCoordinate.Y)
				closeCoordinate = append(closeCoordinate, oldCoordinate)
			}
		}
	}
	return
}
