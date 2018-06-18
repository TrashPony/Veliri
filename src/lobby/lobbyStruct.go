package lobby

type DontEndGames struct {
	Id     string
	Name   string
	Map    Map
	Step   string
	Phase  string
	Winner string
	Ready  string
}

type ActiveGames struct {
	Id   int
	Name string
}

type Map struct {
	Id            int
	Name          string
	XSize         int
	YSize         int
	TypeID        int
	Respawns      int
	Specification string
}
