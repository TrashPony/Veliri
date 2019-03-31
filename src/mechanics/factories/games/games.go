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

func (gamesStore *store) GetPlayerID(playerID int) (*localGame.Game, bool) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	// игрок может быть одновременно только в 1 битве, поэтому это безопасно
	for _, game := range gamesStore.games {
		for _, user := range game.GetPlayers() {
			if user.GetID() == playerID && !user.Leave {
				return game, true
			}
		}
	}
	return nil, false
}

func (gamesStore *store) Add(id int, game *localGame.Game) {
	gamesStore.mx.Lock()
	defer gamesStore.mx.Unlock()
	gamesStore.games[id] = game
}
