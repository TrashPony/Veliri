package squad

import (
	"../../gameObjects/inventory"
	"../../gameObjects/unit"
	"sync"
)

type Squad struct {
	ID              int                 `json:"id"`
	Name            string              `json:"name"`
	Active          bool                `json:"active"`
	MatherShip      *unit.Unit          `json:"mather_ship"`
	Inventory       inventory.Inventory `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame          bool                `json:"in_game"`
	Q               int                 `json:"q"`
	R               int                 `json:"r"`
	GlobalX         int                 `json:"global_x"` /* вычасляема координата на пиксельной сетке */
	GlobalY         int                 `json:"global_y"` /* вычасляема координата на пиксельной сетке */
	ToX             float64             `json:"to_x"`     /* куда отряд двигается */
	ToY             float64             `json:"to_y"`     /* куда отряд двигается */
	MapID           int                 `json:"map_id"`
	BaseID          int                 `json:"base_id"`
	CurrentSpeed    float64             `json:"current_speed"`
	HighGravity     bool                `json:"high_gravity"`
	Afterburner     bool                `json:"afterburner"`
	Evacuation      bool                `json:"evacuation"`
	InSky           bool                `json:"in_sky"` /* отряд по той или иной причине летит Оо */
	MoveChecker     bool                `json:"move_checker"`
	ForceEvacuation bool                `json:"force_evacuation"`
	SoftTransition  bool                `json:"soft_transition"`
	stopMove        chan bool
	updateDB        sync.Mutex
}

func (s *Squad) UpdateLock() {
	s.updateDB.Lock()
}

func (s *Squad) UpdateUnlock() {
	s.updateDB.Unlock()
}

func (s *Squad) CreateMove() {

	if s.stopMove != nil {
		close(s.stopMove)
	}

	s.stopMove = make(chan bool)
}

func (s *Squad) GetMove() chan bool {
	return s.stopMove
}
