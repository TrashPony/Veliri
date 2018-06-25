package matherShip

type MatherShip struct {
	ID             int    `json:"id"`
	Type           string `json:"type"`
	Owner          string `json:"owner"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
	HP             int    `json:"hp"`
	Armor          int    `json:"armor"`
	RangeView      int    `json:"range_view"`
	UnitSlots      int    `json:"unit_slots"`
	UnitSlotSize   int    `json:"unit_slot_size"`
	EquipmentSlots int    `json:"equipment_slots"`
}

func (matherShip *MatherShip) GetX() int {
	return matherShip.X
}

func (matherShip *MatherShip) GetY() int {
	return matherShip.Y
}

func (matherShip *MatherShip) GetWatchZone() int {
	return matherShip.RangeView
}

func (matherShip *MatherShip) GetOwnerUser() string {
	return matherShip.Owner
}
