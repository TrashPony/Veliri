package game

type Structure struct {
	Type      string `json:"type"`
	NameUser  string `json:"owner"`
	WatchZone int    `json:"watch_zone"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}

func (structure *Structure) getX() int {
	return structure.X
}

func (structure *Structure) getY() int {
	return structure.Y
}

func (structure *Structure) getWatchZone() int {
	return structure.WatchZone
}

func (structure *Structure) getNameUser() string {
	return structure.NameUser
}
