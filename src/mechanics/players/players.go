package players

import (
	"../db/dbPlayer"
	"../player"
	"../squadInventory"
	"sync"
)

type UsersStore struct {
	mx    sync.Mutex
	users map[int]*player.Player
}

var Users = NewUsersStore()

func NewUsersStore() *UsersStore {
	return &UsersStore{
		users: make(map[int]*player.Player),
	}
}

func (usersStore *UsersStore) Get(id int) (*player.Player, bool) {
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()
	val, ok := usersStore.users[id]
	return val, ok
}

func (usersStore *UsersStore) Add(id int, login string) *player.Player {
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()

	newUser := dbPlayer.User(id, login)

	squadInventory.GetInventory(newUser)
	usersStore.users[id] = newUser

	return newUser
}

func (usersStore *UsersStore) AddCash(userID, appendCash int) { // appendCash насколько увеличить баланс
	usersStore.mx.Lock()
	defer usersStore.mx.Unlock()

	user, find := usersStore.Get(userID)
	if find { // если юзер уже в карте то обновляем его инфу
		user.SetCredits(user.GetCredits() + appendCash) // добавляем кредиты юзеру
	}

	dbPlayer.AddCash(userID, appendCash) // обновляем кеш в бд
}