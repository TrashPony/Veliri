package game

type Coordinate struct {
	Type 	string
	Texture string
	X, Y, State int
	H, G, F     int
	Parent      *Coordinate
}