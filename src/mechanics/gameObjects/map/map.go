package _map

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math/rand"
	"strconv"
)

type Map struct {
	Id                  int `json:"id"`
	Name                string
	XSize               int
	YSize               int
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

	// разделяем карту на зоны (DiscreteSize х DiscreteSize) при загрузке сервера,
	// добавляем в зону все поинты которые пересекают данных квадрат и ближайшие к нему
	// когда надо найти колизию с юнитом делем его полизию на 100 и отбрасываем дровь так мы получим зону
	// например положение юнита 55/DiscreteSize:77/DiscreteSize = зона 0:0, 257/DiscreteSize:400/DiscreteSize = 1:1, 1654/DiscreteSize:2340/DiscreteSize = 6:9
	// и смотрим только те точки которые находятся в данной зоне
	GeoZones [][]*Zone `json:"-"`

	Anomalies []*Anomalies `json:"anomalies"`

	// показывает позицию на карте мира, пока используется ради меню карты на фронте
	XGlobal int `json:"x_global"`
	YGlobal int `json:"y_global"`

	Fraction       string `json:"fraction"`
	PossibleBattle bool   `json:"possible_battle"`
}

type Zone struct {
	Size      int                      `json:"size"`
	DiscreteX int                      `json:"discrete_x"`
	DiscreteY int                      `json:"discrete_y"`
	Obstacle  []*ObstaclePoint         `json:"obstacle"`
	Regions   []*Region                `json:"regions"`
	Cells     []*coordinate.Coordinate `json:"cells"` // все координаты в зоне
}

func (z *Zone) GetNeighboursZone(mp *Map) []*Zone {

	neighboursZones := make([]*Zone, 0)
	checkRegion := func(zone *Zone) {
		if zone != nil {
			neighboursZones = append(neighboursZones, zone)
		}
	}
	//строго лево
	if z.DiscreteX-1 >= 0 {
		checkRegion(mp.GeoZones[z.DiscreteX-1][z.DiscreteY])
	}
	//строго право
	if len(mp.GeoZones[z.DiscreteX+1]) > 0 {
		checkRegion(mp.GeoZones[z.DiscreteX+1][z.DiscreteY])
	}
	//верх центр
	if z.DiscreteY-1 >= 0 {
		checkRegion(mp.GeoZones[z.DiscreteX][z.DiscreteY-1])
	}
	//низ центр
	checkRegion(mp.GeoZones[z.DiscreteX][z.DiscreteY+1])

	////верх лево
	//if z.DiscreteY-1 >= 0 && z.DiscreteX-1 > 0 {
	//	checkRegion(mp.GeoZones[z.DiscreteX-1][z.DiscreteY-1])
	//}
	//
	////верх право
	//if z.DiscreteY-1 >= 0 && len(mp.GeoZones[z.DiscreteX+1]) > 0 {
	//	checkRegion(mp.GeoZones[z.DiscreteX+1][z.DiscreteY-1])
	//}
	////низ лево
	//if z.DiscreteX-1 >= 0 {
	//	checkRegion(mp.GeoZones[z.DiscreteX-1][z.DiscreteY+1])
	//}
	////низ право
	//if len(mp.GeoZones[z.DiscreteX+1]) > 0 {
	//	checkRegion(mp.GeoZones[z.DiscreteX+1][z.DiscreteY+1])
	//}

	return neighboursZones
}

func (z *Zone) GetRegionsByXY(x, y int) *Region {

	for _, region := range z.Regions {
		if region != nil && region.Index != 0 {

			_, find := region.Cells[x/game_math.CellSize][y/game_math.CellSize]
			if find {
				return region
			}
		}
	}

	return nil
}

func (z *Zone) GetKey() string {
	return strconv.Itoa(z.DiscreteX) + "" + strconv.Itoa(z.DiscreteY)
}

type Region struct {
	Index       int                                    `json:"index"`
	Cells       map[int]map[int]*coordinate.Coordinate `json:"cells"`        // координаты принадлежащие району
	GlobalLinks map[string]*Link                       `json:"global_links"` // уникальные зоны
	Links       []*Link                                `json:"links"`        // зоны в которые можно пройти из этого региона по каждой клетке
	Zone        *Zone                                  `json:"zone"`         // что бы каждый регион знал своего родителя
}

func (r *Region) GetKey() string {
	return r.Zone.GetKey() + strconv.Itoa(r.Index)
}

type Link struct {
	Zone   *Zone   `json:"zone"`
	Region *Region `json:"region"`
	FromX  int     `json:"from_x"`
	FromY  int     `json:"from_y"`
	ToX    int     `json:"to_x"`
	ToY    int     `json:"to_y"`
}

func (l *Link) GetGlobalKey() string {
	return strconv.Itoa(l.Zone.DiscreteX) + "" + strconv.Itoa(l.Zone.DiscreteY) + "" + strconv.Itoa(l.Region.Index)
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
		QSize:               mp.XSize,
		RSize:               mp.YSize,
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

func (mp *Map) GetCoordinate(x, y int) (coordinate *coordinate.Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[x][y]
	return
}

func (mp *Map) GetResource(q, r int) *resource.Map {
	res, _ := mp.Reservoir[q][r]
	return res
}

func (mp *Map) SetXYSize(Scale int) (int, int) {
	return mp.XSize / Scale, mp.YSize / Scale
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
	if mp.Reservoir[reservoir.X] != nil {
		mp.Reservoir[reservoir.X][reservoir.Y] = reservoir
	} else {
		mp.Reservoir[reservoir.X] = make(map[int]*resource.Map)
		mp.Reservoir[reservoir.X][reservoir.Y] = reservoir
	}
}
