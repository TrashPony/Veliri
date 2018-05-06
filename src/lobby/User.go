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

func (user User) SetReady() bool {
	if user.Squad != nil {
		if CheckUnit(user.Squad) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func CheckUnit(squad *Squad.Squad) bool {
	if len(squad.Units) >= 1 {
		return true
	}
	return false
}
