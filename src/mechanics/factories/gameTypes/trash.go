package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/trashItem"
)

type trashStore struct {
	trashItems map[int]trashItem.TrashItem
}

var TrashItems = newTrashStore()

func newTrashStore() *trashStore {
	return &trashStore{trashItems: get.TrashItems()}
}

func (t *trashStore) GetAllType() map[int]trashItem.TrashItem {
	return t.trashItems
}

func (t *trashStore) GetByID(id int) (*trashItem.TrashItem, bool) {
	trashItem, ok := t.trashItems[id]
	return &trashItem, ok
}
