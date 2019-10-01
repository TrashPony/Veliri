package box

type Box struct {
	TypeID       int     `json:"type_id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	CapacitySize float32 `json:"capacity_size"`
	FoldSize     float32 `json:"fold_size"`
	Protect      bool    `json:"protect"`
	ProtectLvl   int     `json:"protect_lvl"`
	Underground  bool    `json:"underground"`
	Height       int     `json:"height"`
	Width        int     `json:"width"`
}
