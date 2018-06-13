package equip

import "../effect"

type Equip struct {
	Id            int             `json:"id"`
	Type          string          `json:"type"`
	Specification string          `json:"specification"`
	Effects       []effect.Effect `json:"effects"`
	Used          bool            `json:"used"`
	Applicable    string          `json:"applicable"`
	Region        int             `json:"region"`
}
