package equip

import "../effect"

type Equip struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	Active        bool             `json:"active"`
	Specification string           `json:"specification"`
	Applicable    string           `json:"applicable"`
	Region        int              `json:"region"`
	Radius        int              `json:"radius"`
	TypeSlot      int              `json:"type_slot"`
	Reload        int              `json:"reload"`
	Power         int              `json:"power"`
	UsePower      int              `json:"use_power"`
	Effects       []*effect.Effect `json:"effects"`
	Used          bool             `json:"used"`
}
