package boxes

import (
	dbBox "github.com/TrashPony/Veliri/src/mechanics/db/box"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/box"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"math/rand"
	"sync"
	"time"
)

type store struct {
	mx    sync.Mutex
	boxes map[int]*boxInMap.Box
}

var Boxes = newBoxStore()

func newBoxStore() *store {
	return &store{
		boxes: dbBox.Boxes(),
	}
}

func (b *store) GetAllBoxByMapID(mapID int) []*boxInMap.Box {
	b.mx.Lock()
	defer b.mx.Unlock()

	mapBoxes := make([]*boxInMap.Box, 0)

	for _, mapBox := range b.boxes {
		if mapBox.MapID == mapID {
			mapBoxes = append(mapBoxes, mapBox)
		}
	}

	return mapBoxes
}

func (b *store) Get(id int) (*boxInMap.Box, *sync.Mutex) {
	b.mx.Lock()
	for _, mapBox := range b.boxes {
		if mapBox.ID == id {
			return mapBox, &b.mx
		}
	}
	return nil, &b.mx
}

func (b *store) GetByQR(q, r, mapID int) (*boxInMap.Box, *sync.Mutex) {
	b.mx.Lock()
	for _, mapBox := range b.boxes {
		if mapBox.MapID == mapID && mapBox.Q == q && mapBox.R == r {
			return mapBox, &b.mx
		}
	}
	return nil, &b.mx
}

func (b *store) DestroyBox(destroyBox *boxInMap.Box) {
	b.mx.Lock()
	defer b.mx.Unlock()
	dbBox.Destroy(destroyBox)
	delete(b.boxes, destroyBox.ID)
}

func (b *store) UpdateBox(updateBox *boxInMap.Box) {
	dbBox.Inventory(updateBox)
}

func (b *store) InsertNewBox(newBox *boxInMap.Box) *boxInMap.Box {
	b.mx.Lock()
	defer b.mx.Unlock()
	dbBox.Insert(newBox)
	b.boxes[newBox.ID] = newBox
	return newBox
}

func (b *store) GetAnomalyRandomBox(typeAnomaly int, boxType *box.Box) *boxInMap.Box {

	newBox := boxInMap.Box{Rotate: rand.Intn(360), TypeID: boxType.TypeID, DestroyTime: time.Now()}
	newBox.GetStorage().Slots = make(map[int]*inventory.Slot)

	if boxType.Protect {
		newBox.SetPassword(rand.Intn(9999))
	}

	// коробка с ресурсами 2+лвл в количестве 1-2 штук типов
	if typeAnomaly == 0 {
		// TODO
	}

	// коробка с чертежом
	if typeAnomaly == 1 {
		// TODO
	}

	return b.InsertNewBox(&newBox)
}
