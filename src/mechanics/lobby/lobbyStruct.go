package lobby

type DontEndGames struct {
	Id     string
	Name   string
	Step   string
	Phase  string
	Winner string
	Ready  string
}

type ActiveGames struct {
	Id   int
	Name string
}
