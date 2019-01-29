package blueprints

type Blueprints struct {
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
}
