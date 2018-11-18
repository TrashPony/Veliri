package storage

import (
	"../db/get"
	"../db/updateStorage"
	inv "../gameObjects/inventory"
	"sync"
)

type Pool struct {
	mx       sync.Mutex
	storages map[int]map[int]*inv.Inventory // [user_ID, base_ID, Inventory]
}

var Storages = NewStoragePoll()

func NewStoragePoll() *Pool {
	return &Pool{
		storages: make(map[int]map[int]*inv.Inventory),
	}
}

func (p *Pool) Get(userId, baseId int) (*inv.Inventory, bool) {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {
			p.mx.Unlock()
			return baseStorage, true
		} else {
			p.storages[userId][baseId] = get.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.Get(userId, baseId)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.Get(userId, baseId)
	}
}

func (p *Pool) AddItem(userId, baseId int, item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32) bool {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {

			ok := baseStorage.AddItem(item, itemType, itemID, quantity, hp, itemSize)
			if ok {
				updateStorage.Inventory(baseStorage, userId, baseId)
			}

			p.mx.Unlock()
			return ok
		} else {
			p.storages[userId][baseId] = get.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize)
	}
}

func RemoveItem(userId, baseId, numberSlot, quantityRemove int) {
	// TODO добавлять/удалять итемы методом инвентаря что бы избежать проблем со слотами
	// TODO и обновлять данные в бд отдельным методом
}
