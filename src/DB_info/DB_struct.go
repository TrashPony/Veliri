package DB_info


type User struct {
	id int
	name string
	password string
	mail string
}

type Games struct {
	nameGame      string
	nameMap       string
	nameCreator   string
	nameNewPlayer string
}

type ActiveGames struct {
	name      string
}

type Map struct {
	id	int
	name    string
	xSize   int
	ySize   int
	Type    string
}
