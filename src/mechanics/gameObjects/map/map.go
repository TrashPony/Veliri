package _map

import (
	"../coordinate"
)

type Map struct {
	Id            int
	Name          string
	QSize         int
	RSize         int
	DefaultTypeID int
	DefaultLevel  int
	Specification string
	OneLayerMap   map[int]map[int]*coordinate.Coordinate
	Respawns      int
}

func (mp *Map) GetCoordinate(q, r int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[q][r]
	return
}
