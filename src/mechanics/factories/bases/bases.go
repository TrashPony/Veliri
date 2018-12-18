package bases

import (
	"sync"
	"../../gameObjects/base"
	"../../db/get"
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