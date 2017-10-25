package lobby

type User struct {
	Id int
	Name string
	Password string
	Mail string
}

type LobbyGames struct {
	Name      string
	Map       string
	Creator   string
	Respawns  map[Respawn]string
	Users     map[string]bool
}

type DontEndGames struct {
	Id		  string
	Name      string
	IdMap       string
	Step	  string
	Phase     string
	Winner    string
	Ready	  string
}

type ActiveGames struct {
	Id	      int
	Name      string
}

type Map struct {
	Id	     int
	Name     string
	XSize    int
	YSize    int
	Type     string
	Respawns int
}

type Respawn struct {
	Id	  int
	Name  string
	IdMap int
	X     int
	Y     int
}