package resource

import "time"

type Map struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	TypeID      int       `json:"type_id"`
	ResourceID  int       `json:"resource_id"`
	Resource    *Resource `json:"resource"`
	Count       int       `json:"count"`
	MapID       int       `json:"map_id"`
	Q           int       `json:"q"`
	R           int       `json:"r"`
	Rotate      int       `json:"rotate"`
	DestroyTime time.Time `json:"destroy_time"`
	MaxCount    int       `json:"max_count"`
	MinCount    int       `json:"min_count"`
}

type Resource struct {
	TypeID int     `json:"type_id"`
	Name   string  `json:"name"`
	Size   float32 `json:"size"`

	// описывает что выходит из этих ресурсов при переработке
	EnrichedThorium int `json:"enriched_thorium"`
	Iron            int `json:"iron"`
	Copper          int `json:"copper"`
	Titanium        int `json:"titanium"`
	Silicon         int `json:"silicon"`
	Plastic         int `json:"plastic"`
}

type RecycledResource struct {
	TypeID int     `json:"type_id"`
	Name   string  `json:"name"`
	Size   float32 `json:"size"`
}

type CraftDetail struct {
	TypeID int     `json:"id"`
	Name   string  `json:"name"`
	Size   float32 `json:"size"`

	/* количественное описание требуемых ресурсов для создания 1 штуки */
	/* примитивы */
	EnrichedThorium int `json:"enriched_thorium"`
	Iron            int `json:"iron"`
	Copper          int `json:"copper"`
	Titanium        int `json:"titanium"`
	Silicon         int `json:"silicon"`
	Plastic         int `json:"plastic"`

	/* детали тоесть фактически сылки на самих себя */
	Steel int `json:"steel"`
	Wire  int `json:"wire"`
}
