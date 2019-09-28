package squad

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"math"
	"sync"
)

type Squad struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Active         bool       `json:"active"`
	MatherShip     *unit.Unit `json:"mather_ship"`
	InGame         bool       `json:"in_game"`
	BaseID         int        `json:"base_id"` /* если отряд не у игрока то он храниться на этой базе */
	SoftTransition bool       `json:"soft_transition"`
	updateDB       sync.Mutex
}

// TODO concurrent map read and map write Inventory

func (s *Squad) UpdateLock() {
	s.updateDB.Lock()
}

func (s *Squad) UpdateUnlock() {
	s.updateDB.Unlock()
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
