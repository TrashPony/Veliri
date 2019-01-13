package storages

import (
	"../../db/storage"
	inv "../../gameObjects/inventory"
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
			p.storages[userId][baseId] = storage.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.Get(userId, baseId)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.Get(userId, baseId)
	}
}

func (p *Pool) AddItem(userId, baseId int, item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32, maxHP int) bool {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {

			ok := baseStorage.AddItem(item, itemType, itemID, quantity, hp, itemSize, maxHP)
			if ok {
				storage.Inventory(baseStorage, userId, baseId)
			}

			p.mx.Unlock()
			return ok
		} else {
			p.storages[userId][baseId] = storage.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize, maxHP)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize, maxHP)
	}
}

func (p *Pool) AddSlot(userId, baseId int, slot *inv.Slot) bool {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {

			ok := baseStorage.AddItemFromSlot(slot)
			if ok {
				storage.Inventory(baseStorage, userId, baseId)
			}

			p.mx.Unlock()
			return ok
		} else {
			p.storages[userId][baseId] = storage.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.AddSlot(userId, baseId, slot)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.AddSlot(userId, baseId, slot)
	}
}

func (p *Pool) RemoveItem(userId, baseId, numberSlot, quantityRemove int) (bool, int) {

	p.mx.Lock()
	defer p.mx.Unlock()

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {
			slot, ok := baseStorage.Slots[numberSlot]
			if ok {
				countRemove := slot.RemoveItemBySlot(quantityRemove)
				storage.Inventory(baseStorage, userId, baseId)
				return true, countRemove
			} else {
				return false, 0
			}
		} else {
			return false, 0
		}
	} else {
		return false, 0
	}
	return false, 0
}
