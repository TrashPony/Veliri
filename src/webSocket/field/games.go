package field

import (
	"../../mechanics/localGame"
	"sync"
)

type GamesStore struct {
	mx sync.Mutex
	games map[int]*localGame.Game
}

var Games = NewGamesStore()

func NewGamesStore() *GamesStore {
	return &GamesStore{
		games: make(map[int]*localGame.Game),
	}
}

func (gamesStore *GamesStore) Get(id int) (*localGame.Game, bool) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	val, ok := gamesStore.games[id]
	return val, ok
}

func (gamesStore *GamesStore) Add(id int, game *localGame.Game) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	gamesStore.games[id] = game
}
