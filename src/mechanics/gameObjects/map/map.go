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
	Emitters            []*Emitter               `json:"emitters"`
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

type Emitter struct {
	ID            int    `json:"id"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	MinScale      int    `json:"min_scale"`
	MaxScale      int    `json:"max_scale"`
	MinSpeed      int    `json:"min_speed"`
	MaxSpeed      int    `json:"max_speed"`
	TTL           int    `json:"ttl"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Color         string `json:"color"`
	Frequency     int    `json:"frequency"`
	MinAlpha      int    `json:"min_alpha"`
	MaxAlpha      int    `json:"max_alpha"`
	Animate       bool   `json:"animate"`
	AnimateSpeed  int    `json:"animate_speed"`
	NameParticle  string `json:"name_particle"`
	AlphaLoopTime int    `json:"alpha_loop_time"`
	Yoyo          bool   `json:"yoyo"`
}

func (mp *Map) GetCoordinate(q, r int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[q][r]
	return
}

func (mp *Map) GetResource(q, r int) *resource.Map {
	res, _ := mp.Reservoir[q][r]
	return res
}
