package equip

import (
	"../effect"
	"../resource"
)

type Equip struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Active          bool             `json:"active"`
	Specification   string           `json:"specification"`
	Applicable      string           `json:"applicable"`
	Region          int              `json:"region"`
	Radius          int              `json:"radius"`
	TypeSlot        int              `json:"type_slot"`
	Reload          int              `json:"reload"`
	Power           int              `json:"power"`
	UsePower        int              `json:"use_power"`
	Effects         []*effect.Effect `json:"effects"`
	MaxHP           int              `json:"max_hp"`
	StepsTime       int              `json:"steps_time"`
	Size            float32          `json:"size"`
	Initiative      int              `json:"initiative"`
	MiningChecker   bool             `json:"move_checker"`
	MiningReservoir *resource.Map    `json:"mining_reservoir"`
	miningExit      chan bool
}

func (e *Equip) CreateMining() {
	if e.miningExit != nil {
		close(e.miningExit)
	}
	e.miningExit = make(chan bool)
}

func (e *Equip) GetMining() chan bool {
	return e.miningExit
}
