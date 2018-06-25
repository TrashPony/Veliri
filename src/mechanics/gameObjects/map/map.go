package _map

import (
	"../../localGame/map/coordinate"
)

type Map struct {
	Id            int
	Name          string
	XSize         int
	YSize         int
	DefaultTypeID int
	DefaultLevel  int
	Specification string
	OneLayerMap   map[int]map[int]*coordinate.Coordinate
	Respawns      int
}

func (mp *Map) GetCoordinate(x, y int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[x][y]
	return
}
