package field

import (
	"../../mechanics/game"
	"sync"
)

type GamesStore struct {
	mx sync.Mutex
	games map[int]*game.Game
}

var Games = NewGamesStore()

func NewGamesStore() *GamesStore {
	return &GamesStore{
		games: make(map[int]*game.Game),
	}
}

func (gamesStore *GamesStore) Get(id int) (*game.Game, bool) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	val, ok := gamesStore.games[id]
	return val, ok
}

func (gamesStore *GamesStore) Add(id int, game *game.Game) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	gamesStore.games[id] = game
}
