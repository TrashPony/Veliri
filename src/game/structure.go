package game

type Structure struct {
	Type string
	NameUser    string
	WatchZone   int
	X int
	Y int
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
