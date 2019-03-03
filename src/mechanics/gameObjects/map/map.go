package _map

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
)

type Map struct {
	Id                  int
	Name                string
	QSize               int
	RSize               int
	DefaultTypeID       int
	DefaultLevel        int
	Specification       string
	OneLayerMap         map[int]map[int]*coordinate.Coordinate
	Reservoir           map[int]map[int]*resource.Map `json:"reservoir"`
	Respawns            int
	Global              bool                     `json:"global"`
	InGame              bool                     `json:"in_game"`
	HandlersCoordinates []*coordinate.Coordinate `json:"handlers_coordinates"`
	GeoData             []*ObstaclePoint         `json:"geo_data"`
	Beams               []*Beam                  `json:"beams"`
}

type ObstaclePoint struct {
	ID     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
}

type Beam struct {
	ID     int    `json:"id"`
	XStart int    `json:"x_start"`
	YStart int    `json:"y_start"`
	XEnd   int    `json:"x_end"`
	YEnd   int    `json:"y_end"`
	Color  string `json:"color"`
}

func (mp *Map) GetCoordinate(q, r int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[q][r]
	return
}

func (mp *Map) GetResource(q, r int) *resource.Map {
	res, _ := mp.Reservoir[q][r]
	return res
}
