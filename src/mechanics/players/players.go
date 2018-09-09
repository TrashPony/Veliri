package players

import (
	"../inventory"
	"../player"
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

	newUser := player.Player{}
	newUser.SetID(id)
	newUser.SetLogin(login)

	inventory.GetInventory(&newUser)
	//todo взятие ПОЛНОГО обьекта пользователя
	usersStore.users[id] = &newUser

	return &newUser
}
