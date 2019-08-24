package storages

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/storage"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"sync"
)

type pool struct {
	mx       sync.Mutex
	storages map[int]map[int]*inv.Inventory // [user_ID, base_ID, Inventory]
}

var Storages = newStoragePoll()

func newStoragePoll() *pool {
	return &pool{
		storages: make(map[int]map[int]*inv.Inventory),
	}
}

func (p *pool) Get(userId, baseId int) (*inv.Inventory, bool) {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	if baseId == 0 {
		p.mx.Unlock()
		return nil, false
	}

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

func (p *pool) AddItem(userId, baseId int, item interface{}, itemType string, itemID, quantity, hp int, itemSize float32, maxHP int, newSlot bool) bool {
	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {

			ok := baseStorage.AddItem(item, itemType, itemID, quantity, hp, itemSize, maxHP, newSlot, userId)
			if ok {
				storage.Inventory(baseStorage, userId, baseId)
			}

			p.mx.Unlock()
			return ok
		} else {
			p.storages[userId][baseId] = storage.UserStorage(userId, baseId)
			p.mx.Unlock()
			return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize, maxHP, newSlot)
		}
	} else {
		p.storages[userId] = make(map[int]*inv.Inventory)
		p.mx.Unlock()
		return p.AddItem(userId, baseId, item, itemType, itemID, quantity, hp, itemSize, maxHP, newSlot)
	}
}

func (p *pool) AddSlot(userId, baseId int, slot *inv.Slot) bool {

	p.mx.Lock()
	// sync.Mutex не рекурсивен, поэтому возможно это не безопасно, и закрывается не через defer :\

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {

			ok := baseStorage.AddItemFromSlot(slot, userId)
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

func (p *pool) AddItemBySlot(userId, baseId, numberSlot, quantity int) bool {

	p.mx.Lock()
	defer p.mx.Unlock()

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {
			slot, ok := baseStorage.Slots[numberSlot]
			if ok {
				slot.AddItemBySlot(quantity, userId)
				storage.Inventory(baseStorage, userId, baseId)
				return true
			}
		}
	}

	return false
}

func (p *pool) RemoveItemBySlot(userId, baseId, numberSlot, quantityRemove int) (bool, int) {

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
			}
		}
	}

	return false, 0
}

func (p *pool) RemoveItem(userId, baseId, itemID int, itemType string, quantityRemove int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	userStorages, userOk := p.storages[userId]
	if userOk {
		baseStorage, baseOk := userStorages[baseId]
		if baseOk {
			err := baseStorage.RemoveItem(itemID, itemType, quantityRemove)
			if err == nil {
				storage.Inventory(baseStorage, userId, baseId)
				return true
			}
		}
	}

	return false
}
