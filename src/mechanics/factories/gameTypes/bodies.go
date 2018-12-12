package gameTypes

import (
	"../../db/get"
	"../../gameObjects/detail"
	"github.com/getlantern/deepcopy"
)

type BodyStore struct {
	bodies map[int]detail.Body
}

var Bodies = NewBodyStore()

func NewBodyStore() *BodyStore {
	return &BodyStore{bodies: get.BodiesType()}
}

func (b *BodyStore) GetByID(id int) (*detail.Body, bool) {
	var newBody detail.Body
	factoryBody, ok := b.bodies[id]

	err := deepcopy.Copy(&newBody, &factoryBody) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &newBody, ok
}

func (b *BodyStore) GetAllType() (map[int]detail.Body) {
	return b.bodies
}
