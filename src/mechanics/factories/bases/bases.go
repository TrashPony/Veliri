package bases

import (
	dbBase "../../db/base"
	"../../gameObjects/base"
	"sync"
)

type store struct {
	mx    sync.Mutex
	bases map[int]*base.Base
}

var Bases = newBasesStore()

func newBasesStore() *store {
	return &store{
		bases: dbBase.Bases(),
	}
}

func UserIntoBase(userID, baseID int) {
	dbBase.UserIntoBase(userID, baseID)
}

func (b *store) Get(id int) (*base.Base, bool) {
	b.mx.Lock()
	defer b.mx.Unlock()
	val, ok := b.bases[id]
	return val, ok
}

func (b *store) GetBasesByMap(mapID int) map[int]*base.Base {
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
