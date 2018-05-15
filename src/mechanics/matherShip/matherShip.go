package matherShip

type MatherShip struct {
	Type      string `json:"type"`
	Owner     string `json:"owner"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	HP        int    `json:"hp"`
	Armor     int    `json:"armor"`
	RangeView int    `json:"range_view"`
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
