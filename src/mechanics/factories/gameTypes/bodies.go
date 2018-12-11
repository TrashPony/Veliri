package gameTypes

import (
	"../../db/get"
	"../../gameObjects/detail"
)

type BodyStore struct {
	bodies map[int]detail.Body
}

var Bodies = NewBodyStore()

func NewBodyStore() *BodyStore {
	return &BodyStore{bodies: get.BodiesType()}
}

func (b *BodyStore) GetByID(id int) (*detail.Body, bool) {
	// TODO копировать слоты снаряжения и оружия, т.к. это сылочные типы данных
	var newBody detail.Body
	newBody, ok := b.bodies[id]
	return &newBody, ok
}
