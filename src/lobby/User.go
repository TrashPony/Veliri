package lobby

import (
	"./Squad"
)

type User struct {
	Id      int
	Name    string
	Ready   bool
	Squad   *Squad.Squad
	Squads  []*Squad.Squad
	Respawn *Respawn
	Game    string
}

func (user User) SetGame(game LobbyGames)  {
	println(game.Name)
	user.Game = game.Name
}
