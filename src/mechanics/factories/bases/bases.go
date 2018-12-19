package bases

import (
	"../../db/get"
	"../../gameObjects/base"
	"sync"
)

type Store struct {
	mx    sync.Mutex
	bases map[int]*base.Base
}

var Bases = NewBasesStore()

func NewBasesStore() *Store {
	return &Store{
		bases: get.Bases(),
	}
}

func (b *Store) Get(id int) (*base.Base, bool) {
	b.mx.Lock()
	defer b.mx.Unlock()
	val, ok := b.bases[id]
	return val, ok
}

func (b *Store) GetBasesByMap(mapID int) map[int]*base.Base {
	b.mx.Lock()
	defer b.mx.Unlock()

	bases := make(map[int]*base.Base)

	for _, mapBase := range b.bases {
		if mapBase.MapID == mapID {
			bases[mapBase.ID] = mapBase
		}
	}

	return bases
}
