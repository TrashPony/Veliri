package field

type Message struct {
	Event      string `json:"event"`
	IdGame     int    `json:"id_game"`
	UnitID     int    `json:"unit_id"`
	EquipID    int    `json:"equip_id"`
	IdTarget   string `json:"id_target"`
	TypeUnit   string `json:"type_unit"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	ToX        int    `json:"to_x"`
	ToY        int    `json:"to_y"`
	TargetX    int    `json:"target_x"`
	TargetY    int    `json:"target_y"`
	EquipType  int    `json:"equip_type"`
	NumberSlot int    `json:"number_slot"`
}

type ErrorMessage struct {
	Event string `json:"event"`
	Error string `json:"error"`
}
