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
	b.mx.Lock()
	defer b.mx.Unlock()

	mapBoxes := make([]*box.Box, 0)

	for _, mapBox := range b.boxes {
		if mapBox.MapID == mapID {
			mapBoxes = append(mapBoxes, mapBox)
		}
	}

	return mapBoxes
}

func (b *store) Get(id int) (*box.Box, *sync.Mutex) {
	b.mx.Lock()
	for _, mapBox := range b.boxes {
		if mapBox.ID == id {
			return mapBox, &b.mx
		}
	}
	return nil, &b.mx
}

func (b *store) GetByQR(q, r, mapID int) (*box.Box, *sync.Mutex) {
	b.mx.Lock()
	for _, mapBox := range b.boxes {
		if mapBox.ID == mapID && mapBox.Q == q && mapBox.R == r {
			return mapBox, &b.mx
		}
	}
	return nil, &b.mx
}

func (b *store) DestroyBox(destroyBox *box.Box) {
	b.mx.Lock()
	defer b.mx.Unlock()
	dbBox.Destroy(destroyBox)
	delete(b.boxes, destroyBox.ID)
}

func (b *store) UpdateBox(updateBox *box.Box) {
	dbBox.Inventory(updateBox)
}

func (b *store) InsertNewBox(newBox *box.Box) *box.Box {
	b.mx.Lock()
	defer b.mx.Unlock()
	dbBox.Insert(newBox)
	b.boxes[newBox.ID] = newBox
	return newBox
}
