package maps

import (
	"../../gameObjects/box"
	"../../gameObjects/map"
	"../../gameObjects/resource"
	"math/rand"
)

type Anomaly struct {
	q     int
	r     int
	MapID int `json:"map_id"`
	power int // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type  int

	/* награда */
	box      *box.Box
	resource *resource.Map
	text     string
}

func (a *Anomaly) GetQ() int {
	return a.q
}

func (a *Anomaly) GetR() int {
	return a.r
}

func (a *Anomaly) GetPower() int {
	return a.power
}

func anomalyGenerator(mp *_map.Map, m *MapStore) {
	i := 0

	for i < 5 {

		typeAnomaly := rand.Intn(4)

		// коробка с ресурсами 2+лвл
		if typeAnomaly == 0 {
			// todo
		}

		// коробка с чертежом
		if typeAnomaly == 1 {
			// todo
		}

		// руда
		if typeAnomaly == 2 {
			// todo
		}

		// текс
		if typeAnomaly == 3 {
			// todo
		}

		power := rand.Intn(6) // радиус

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)
		coordinatePlace, _ := mp.GetCoordinate(q, r)

		if coordinatePlace.Move {
			if m.anomaly[mp.Id] == nil {
				m.anomaly[mp.Id] = make([]*Anomaly, 0)
			}

			anomaly := &Anomaly{q: q, r: r, Type: typeAnomaly, power: power, MapID: mp.Id}

			m.anomaly[mp.Id] = append(m.anomaly[mp.Id], anomaly)
			i++
		}
	}
}

func (m *MapStore) GetAllMapAnomaly(mapID int) []*Anomaly {
	return Maps.anomaly[mapID]
}
