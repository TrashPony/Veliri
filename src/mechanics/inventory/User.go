package inventory


type User struct {
	Id      int
	Name    string
	Squad   *Squad
	Squads  []*Squad
}