package _map

import (
	"../coordinate"
	"../resource"
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
	Reservoir     map[int]map[int]*resource.Map   `json:"reservoir"`
	Respawns      int
	Global        bool                            `json:"global"`
	InGame        bool                            `json:"in_game"`
}

func (mp *Map) GetCoordinate(q, r int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[q][r]
	return
}

func (mp *Map) GetResource(q, r int) *resource.Map {
	res, _ := mp.Reservoir[q][r]
	return res
}
