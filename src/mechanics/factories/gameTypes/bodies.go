package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/getlantern/deepcopy"
	"math/rand"
)

type bodyStore struct {
	bodies map[int]detail.Body
}

var Bodies = newBodyStore()

func newBodyStore() *bodyStore {
	return &bodyStore{bodies: get.BodiesType()}
}

func (b *bodyStore) GetByID(id int) (*detail.Body, bool) {
	var newBody detail.Body
	factoryBody, ok := b.bodies[id]

	err := deepcopy.Copy(&newBody, &factoryBody) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &newBody, ok
}

func (b *bodyStore) GetRandom() *detail.Body {
	for {
		// TODO возможны проблемы))
		count := 0
		count2 := rand.Intn(len(b.bodies))
		for id := range b.bodies {
			if count == count2 {
				body, _ := b.GetByID(id)
				if body.MotherShip {
					return body
				}
			}
			count++
		}
	}
}

func (b *bodyStore) GetAllType() map[int]detail.Body {
	return b.bodies
}
