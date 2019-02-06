package gameTypes

import (
	"../../db/get"
	"../../gameObjects/detail"
	"github.com/getlantern/deepcopy"
)

type bodyStore struct {
	bodies map[int]detail.Body
}

var Bodies = newBodyStore()

func newBodyStore() *bodyStore {
	return &bodyStore{bodies: get.BodiesType()}
}

func (b *bodyStore) GetByID(id int) (*detail.Body, bool) {
	var newBody detail.Body
	factoryBody, ok := b.bodies[id]

	err := deepcopy.Copy(&newBody, &factoryBody) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &newBody, ok
}

func (b *bodyStore) GetAllType() (map[int]detail.Body) {
	return b.bodies
}
