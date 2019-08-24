package bases

import (
	dbBase "github.com/TrashPony/Veliri/src/mechanics/db/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"math/rand"
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

func (b *store) GetCapital(fraction string) *base.Base {
	b.mx.Lock()
	defer b.mx.Unlock()

	for _, gameBase := range b.bases {
		if gameBase.Capital && gameBase.Fraction == fraction {
			return gameBase
		}
	}
	return nil
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

func (b *store) GetRandomBase() *base.Base {
	for {
		// TODO возможны проблемы))
		count := 0
		count2 := rand.Intn(len(b.bases))

		for id := range b.bases {
			if count == count2 {
				gameBase, _ := b.Get(id)
				return gameBase
			}
			count++
		}
	}
}
