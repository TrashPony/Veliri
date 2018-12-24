package boxes

import (
	dbBox "../../db/box"
	"../../gameObjects/box"
	"sync"
)

type store struct {
	mx    sync.Mutex
	boxes map[int]*box.Box
}

var Boxes = NewBoxStore()

func NewBoxStore() *store {
	return &store{
		boxes: dbBox.Boxes(),
	}
}

func (b *store) GetAllBoxByMapID(mapID int) []*box.Box {
	mapBoxes := make([]*box.Box, 0)

	for _, mapBox := range b.boxes {
		if mapBox.MapID == mapID {
			mapBoxes = append(mapBoxes, mapBox)
		}
	}

	return mapBoxes
}

func (b *store) Get(id int) *box.Box {
	for _, mapBox := range b.boxes {
		if mapBox.ID == id {
			return mapBox
		}
	}
	return nil
}

func (b *store) GetByQR(q, r, mapID int) *box.Box {
	for _, mapBox := range b.boxes {
		if mapBox.ID == mapID && mapBox.Q == q && mapBox.R == r{
			return mapBox
		}
	}
	return nil
}

func (b *store) DestroyBox(destroyBox *box.Box) {
	// todo удаление из бд, удаление из фабрики
}

func (b *store) UpdateBox(updateBox *box.Box) {
	dbBox.Inventory(updateBox)
}

func (b *store) InsertNewBox(newBox *box.Box) *box.Box {
	dbBox.Insert(newBox)
	b.boxes[newBox.ID] = newBox
	return newBox
}
