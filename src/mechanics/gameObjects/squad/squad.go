package squad

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"sync"
)

type Squad struct {
	ID              int                  `json:"id"`
	Name            string               `json:"name"`
	Active          bool                 `json:"active"`
	MatherShip      *unit.Unit           `json:"mather_ship"`
	Inventory       *inventory.Inventory `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame          bool                 `json:"in_game"`
	MapID           int                  `json:"map_id"`
	BaseID          int                  `json:"base_id"` /* если отряд не у игрока то он храниться на этой базе */
	Evacuation      bool                 `json:"evacuation"`
	InSky           bool                 `json:"in_sky"` /* отряд по той или иной причине летит Оо */
	MoveChecker     bool                 `json:"move_checker"`
	ForceEvacuation bool                 `json:"force_evacuation"`
	SoftTransition  bool                 `json:"soft_transition"`
	updateDB        sync.Mutex
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
