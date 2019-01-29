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

func (r *Resource) GetEnrichedThorium() int {
	return r.EnrichedThorium
}

func (r *Resource) GetIron() int {
	return r.Iron
}

func (r *Resource) GetCopper() int {
	return r.Copper
}

func (r *Resource) GetTitanium() int {
	return r.Titanium
}

func (r *Resource) GetSilicon() int {
	return r.Silicon
}

func (r *Resource) GetPlastic() int {
	return r.Plastic
}

func (r *Resource) GetSteel() int {
	return 0
}

func (r *Resource) GetWire() int {
	return 0
}

func (r *Resource) GetElectronics() int {
	return 0
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
}
