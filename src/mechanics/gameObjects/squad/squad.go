package squad

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
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
