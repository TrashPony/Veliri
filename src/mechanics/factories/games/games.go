package games

import (
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"sync"
)

type store struct {
	mx    sync.Mutex
	games map[int]*localGame.Game
}

var Games = newGamesStore()

func newGamesStore() *store {
	return &store{
		games: make(map[int]*localGame.Game),
	}
}

func (gamesStore *store) Get(id int) (*localGame.Game, bool) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	val, ok := gamesStore.games[id]
	return val, ok
}

func (gamesStore *store) Add(id int, game *localGame.Game) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	gamesStore.games[id] = game
}
