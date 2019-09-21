package _map

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"math/rand"
)

type Map struct {
	Id                  int `json:"id"`
	Name                string
	QSize               int
	RSize               int
	DefaultTypeID       int
	DefaultLevel        int
	Specification       string
	OneLayerMap         map[int]map[int]*coordinate.Coordinate
	Reservoir           map[int]map[int]*resource.Map `json:"reservoir"`
	Global              bool                          `json:"global"`
	InGame              bool                          `json:"in_game"`
	HandlersCoordinates []*coordinate.Coordinate      `json:"handlers_coordinates"`
	Beams               []*Beam                       `json:"beams"`
	Emitters            []*Emitter                    `json:"emitters"`
	GeoData             []*ObstaclePoint              `json:"geo_data"`

	// разделяем карту на зоны (100х100рх) при загрузке сервера,
	// однако квадраты будут размером 200х200 и будут накладыватся друг на друга (это нужно на тот случай если юнит стоит на границе)
	// добавляем в зону все поинты которые пересекают данных квадрат
	// когда надо найти колизию с юнитом делем его полизию на 100 и отбрасываем дровь так мы получим зону
	// например положение юнита 55:77 = зона 0:0, 159:234 = 1:2, 1654:2340 = 16:23
	// и смотрим только те точки которые находятся в данной зоне
	GeoZone [][][]*ObstaclePoint `json:"geo_zone"`

	Anomalies []*Anomalies `json:"anomalies"`

	// тут хранятся просчитанные шаблоны координат, что бы не проверять координаты при каждом поиске пути
	GeoDataMaps map[int]map[int]map[int]coordinate.Coordinate `json:"-"`

	// показывает позицию на карте мира, пока используется ради меню карты на фронте
	XGlobal int `json:"x_global"`
	YGlobal int `json:"y_global"`

	Fraction       string `json:"fraction"`
	PossibleBattle bool   `json:"possible_battle"`
}

type ShortInfoMap struct {
	Id                  int `json:"id"`
	Name                string
	QSize               int
	RSize               int
	Specification       string
	Global              bool                     `json:"global"`
	HandlersCoordinates []*coordinate.Coordinate `json:"handlers_coordinates"`
	XGlobal             int                      `json:"x_global"`
	YGlobal             int                      `json:"y_global"`
	Fraction            string                   `json:"fraction"`
	PossibleBattle      bool                     `json:"possible_battle"`
}

type Anomalies struct {
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Radius int    `json:"radius"`
	Type   string `json:"type"`
	Power  int    `json:"power"`
}

func (mp *Map) GetShortInfoMap() *ShortInfoMap {
	return &ShortInfoMap{
		Id:                  mp.Id,
		Name:                mp.Name,
		QSize:               mp.QSize,
		RSize:               mp.RSize,
		Specification:       mp.Specification,
		Global:              mp.Global,
		HandlersCoordinates: mp.HandlersCoordinates,
		XGlobal:             mp.XGlobal,
		YGlobal:             mp.YGlobal,
		Fraction:            mp.Fraction,
		PossibleBattle:      mp.PossibleBattle,
	}
}

func (mp *Map) GetRandomEntryBase() *coordinate.Coordinate {
	for {
		// TODO возможны проблемы))
		count := 0
		entryCount := 0

		for _, entry := range mp.HandlersCoordinates {
			if entry.Handler == "base" {
				entryCount++
			}
		}

		if entryCount == 0 {
			return nil
		}

		count2 := rand.Intn(len(mp.HandlersCoordinates))
		for _, entry := range mp.HandlersCoordinates {
			if count == count2 && entry.Handler == "base" {
				return entry
			}
			count++
		}
	}
}

func (mp *Map) GetEntryBase(baseID int) *coordinate.Coordinate {
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "base" && entry.ToBaseID == baseID {
			return entry
		}
	}
	return nil
}

func (mp *Map) GetRandomEntrySector() *coordinate.Coordinate {
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" {
			return entry
		}
	}
	return nil
}

// TODO надо переписать на интерфейсы
func (mp *Map) GetEntryTySector(sectorID int) *coordinate.Coordinate {
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" && entry.ToMapID == sectorID {
			return entry
		}
	}
	return nil
}

func (mp *ShortInfoMap) GetEntryTySector(sectorID int) *coordinate.Coordinate {
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" && entry.ToMapID == sectorID {
			return entry
		}
	}
	return nil
}

func (mp *Map) GetAllEntrySectors() []*coordinate.Coordinate {
	entrySectors := make([]*coordinate.Coordinate, 0)
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" {
			entrySectors = append(entrySectors, entry)
		}
	}

	return entrySectors
}

func (mp *ShortInfoMap) GetAllEntrySectors() []*coordinate.Coordinate {
	entrySectors := make([]*coordinate.Coordinate, 0)
	for _, entry := range mp.HandlersCoordinates {
		if entry.Handler == "sector" {
			entrySectors = append(entrySectors, entry)
		}
	}

	return entrySectors
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

func (mp *Map) SetXYSize(hexWidth, hexHeight, Scale int) (int, int) {
	return (mp.QSize * hexWidth) / Scale, int(float64(mp.RSize)*float64(hexHeight)*0.75) / Scale
}

// TODO GetMaxPriorityTexture, GetMaxPriorityObject близнецы
func (mp *Map) GetMaxPriorityTexture() int {
	max := 0

	for _, xLine := range mp.OneLayerMap {
		for _, coordinate := range xLine {
			if max < coordinate.TexturePriority {
				max = coordinate.TexturePriority
			}
		}
	}

	return max
}

func (mp *Map) GetMaxPriorityObject() int {
	max := 0

	for _, xLine := range mp.OneLayerMap {
		for _, coordinate := range xLine {
			if max < coordinate.ObjectPriority {
				max = coordinate.ObjectPriority
			}
		}
	}

	return max
}

func (mp *Map) AddResourceInMap(reservoir *resource.Map) {
	if mp.Reservoir[reservoir.Q] != nil {
		mp.Reservoir[reservoir.Q][reservoir.R] = reservoir
	} else {
		mp.Reservoir[reservoir.Q] = make(map[int]*resource.Map)
		mp.Reservoir[reservoir.Q][reservoir.R] = reservoir
	}
}
