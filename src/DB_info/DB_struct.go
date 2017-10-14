package DB_info

type User struct {
	Id int
	Name string
	Password string
	Mail string
}

type Games struct {
	Name      string
	Map       string
	Creator   string
	NewPlayer string
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
	Id	int
	Name    string
	XSize   int
	YSize   int
	Type    string
}
