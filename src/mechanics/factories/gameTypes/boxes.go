package gameTypes

import (
	"../../db/get"
	"../../gameObjects/box"
)

type boxesStore struct {
	boxes map[int]box.Box
}

var Boxes = newBoxStore()

func newBoxStore() *boxesStore {
	return &boxesStore{boxes: get.Boxes()}
}

func (b *boxesStore) GetByID(id int) (*box.Box, bool) {
	var newBox box.Box
	newBox, ok := b.boxes[id]
	return &newBox, ok
}

func (b *boxesStore) GetAllType() map[int]box.Box {
	return b.boxes
}
