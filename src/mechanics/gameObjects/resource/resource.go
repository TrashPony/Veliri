package resource

type Resource struct {
	TypeID int     `json:"type_id"`
	Name   string  `json:"name"`
	Size   float32 `json:"size"`

	EnrichedThorium int `json:"enriched_thorium"`
}

type RecycledResource struct {
	TypeID int     `json:"type_id"`
	Name   string  `json:"name"`
	Size   float32 `json:"size"`
}
