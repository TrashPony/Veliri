package games

import (
	"../../localGame"
	"sync"
)

type Store struct {
	mx    sync.Mutex
	games map[int]*localGame.Game
}

var Games = NewGamesStore()

func NewGamesStore() *Store {
	return &Store{
		games: make(map[int]*localGame.Game),
	}
}

func (gamesStore *Store) Get(id int) (*localGame.Game, bool) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	val, ok := gamesStore.games[id]
	return val, ok
}

func (gamesStore *Store) Add(id int, game *localGame.Game) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	gamesStore.games[id] = game
}
