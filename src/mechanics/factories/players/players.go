package players

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"sync"
)

type usersStore struct {
	mx    sync.Mutex
	users map[int]*player.Player
}

var Users = newUsersStore()

func newUsersStore() *usersStore {
	return &usersStore{
		users: make(map[int]*player.Player),
	}
}

func (usersStore *usersStore) Get(id int) (*player.Player, bool) {
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()
	val, ok := usersStore.users[id]
	return val, ok
}

func (usersStore *usersStore) Add(id int, login string) *player.Player {
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()

	newUser := dbPlayer.User(id, login)

	squadInventory.GetInventory(newUser)
	usersStore.users[id] = newUser

	return newUser
}

func (usersStore *usersStore) AddCash(userID, appendCash int) { // appendCash насколько увеличить баланс
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()

	user, find := usersStore.users[userID]
	if find { // если юзер уже в карте то обновляем его инфу
		user.SetCredits(user.GetCredits() + appendCash) // добавляем кредиты юзеру
	}

	dbPlayer.AddCash(userID, appendCash) // обновляем кеш в бд
}
