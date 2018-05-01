package lobby

type User struct {
	Id       int
	Name     string
	Password string
	Mail     string
}

type LobbyGames struct {
	Name     string
	Map      Map
	Creator  string
	Respawns []*Respawn
	Users    map[string]bool
}

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
	Type          string
	Respawns      int
	Specification string
}
