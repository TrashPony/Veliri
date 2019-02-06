package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/box"
	"math/rand"
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

func (b *boxesStore) GetRandomBox() *box.Box {
	allType := make([]box.Box, 0)

	for _, typeRes := range b.boxes {
		allType = append(allType, typeRes)
	}

	randomIndex := rand.Intn(len(allType))
	newBox, _ := b.GetByID(allType[randomIndex].TypeID)

	return newBox
}
