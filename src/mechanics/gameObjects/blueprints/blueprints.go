package blueprints

type Blueprint struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ItemType        string `json:"item_type"`
	ItemId          int    `json:"item_id"`
	Icon            string `json:"icon"`
	CraftTime       int    `json:"craft_time"`
	Original        bool   `json:"original"`
	Copies          int    `json:"copies"`
	Count           int    `json:"count"`
	EnrichedThorium int    `json:"enriched_thorium"`
	Iron            int    `json:"iron"`
	Copper          int    `json:"copper"`
	Titanium        int    `json:"titanium"`
	Silicon         int    `json:"silicon"`
	Plastic         int    `json:"plastic"`
	Steel           int    `json:"steel"`
	Wire            int    `json:"wire"`
	Electronics     int    `json:"electronics"`
}

func (b *Blueprint) GetEnrichedThorium() int {
	return b.EnrichedThorium
}

func (b *Blueprint) GetIron() int {
	return b.Iron
}

func (b *Blueprint) GetCopper() int {
	return b.Copper
}

func (b *Blueprint) GetTitanium() int {
	return b.Titanium
}

func (b *Blueprint) GetSilicon() int {
	return b.Silicon
}

func (b *Blueprint) GetPlastic() int {
	return b.Plastic
}

func (b *Blueprint) GetSteel() int {
	return b.Steel
}

func (b *Blueprint) GetWire() int {
	return b.Wire
}

func (b *Blueprint) GetElectronics() int {
	return b.Electronics
}
