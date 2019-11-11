package squad

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"math"
	"strconv"
	"sync"
)

type Squad struct {
	ID                   int                        `json:"id"`
	Name                 string                     `json:"name"`
	Active               bool                       `json:"active"`
	MatherShip           *unit.Unit                 `json:"mather_ship"`
	InGame               bool                       `json:"in_game"`
	BaseID               int                        `json:"base_id"` /* если отряд не у игрока то он храниться на этой базе */
	SoftTransition       bool                       `json:"soft_transition"`
	VisibleObjects       map[string]*VisibleObjects `json:"-"` // key id_object+type_object
	updateVisibleObjects sync.RWMutex               `json:"-"`
	updateDB             sync.Mutex                 `json:"-"`

	/* необходимые флаги что бы обеспечить правильную перегрзку методов воркеров*/
	RecoveryPowerWork bool `json:"-"`
	RecoveryPowerExit bool `json:"-"`

	GunWorkerWork bool `json:"-"`
	GunWorkerExit bool `json:"-"`

	RadarWorkerWork bool `json:"-"`
	RadarWorkerExit bool `json:"-"`
}

type VisibleObjects struct {
	IDObject   int    `json:"id_object"`
	TypeObject string `json:"type_object"`
	UUID       string `json:"uuid"`
	View       bool   `json:"view"`  // в прямой видимости
	Radar      bool   `json:"radar"` // видим только радаром
	Type       string `json:"type"`  // fly(летающий), ground(наземный), structure(структура)
	Update     bool   `json:"update"`
}

func (s *Squad) GetVisibleObjectByID(id string) *VisibleObjects {
	s.updateVisibleObjects.RLock()
	defer s.updateVisibleObjects.RUnlock()

	object, ok := s.VisibleObjects[id]
	if ok {
		return object
	}

	return nil
}

func (s *Squad) AddVisibleObject(newObj *VisibleObjects) {
	s.updateVisibleObjects.Lock()
	defer s.updateVisibleObjects.Unlock()

	s.VisibleObjects[newObj.TypeObject+strconv.Itoa(newObj.IDObject)] = newObj
}

func (s *Squad) RemoveVisibleObject(removeObj *VisibleObjects) {
	s.updateVisibleObjects.Lock()
	defer s.updateVisibleObjects.Unlock()

	delete(s.VisibleObjects, removeObj.TypeObject+strconv.Itoa(removeObj.IDObject))
}

func (s *Squad) RadarLock() {
	s.updateVisibleObjects.Lock()
}

func (s *Squad) RadarUnlock() {
	s.updateVisibleObjects.Unlock()
}

func (s *Squad) UpdateLock() {
	s.updateDB.Lock()
}

func (s *Squad) UpdateUnlock() {
	s.updateDB.Unlock()
}

func (s *Squad) GetShortUnits() map[int]*unit.ShortUnitInfo {
	shortUnits := make(map[int]*unit.ShortUnitInfo)

	shortUnits[s.MatherShip.ID] = s.MatherShip.GetShortInfo()

	for _, unitSlot := range s.MatherShip.Units {
		if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap {
			shortUnits[unitSlot.Unit.ID] = unitSlot.Unit.GetShortInfo()
		}
	}

	return shortUnits
}

func (s *Squad) GetUnitByID(id int) *unit.Unit {
	if s.MatherShip.ID == id {
		return s.MatherShip
	}

	for _, unitSlot := range s.MatherShip.Units {
		if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.ID == id {
			return unitSlot.Unit
		}
	}

	return nil
}

func (s *Squad) GetFormationCoordinate(x, y int) (int, int) {
	//хз почему +90 но работает)
	alpha := float64(s.MatherShip.Rotate+90) * math.Pi / 180
	newX := float64(x)*math.Cos(alpha) - float64(y)*math.Sin(alpha) + float64(s.MatherShip.X)
	newY := float64(x)*math.Sin(alpha) + float64(y)*math.Cos(alpha) + float64(s.MatherShip.Y)

	return int(newX), int(newY)
}

func (s *Squad) CheckViewCoordinate(x, y int) (bool, bool) {

	radarView := false

	if s == nil || s.MatherShip == nil {
		return false, false
	}

	view, radar := s.MatherShip.CheckViewCoordinate(x, y)
	if view {
		return true, true
	}

	if radar {
		radarView = true
	}

	for _, unitSlot := range s.MatherShip.Units {
		if unitSlot != nil && unitSlot.Unit != nil {

			view, radar := unitSlot.Unit.CheckViewCoordinate(x, y)
			if view {
				return true, true
			}

			if radar {
				radarView = true
			}
		}
	}

	return false, radarView
}
