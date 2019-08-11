package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/getlantern/deepcopy"
)

func GetDetails(user *player.Player) {

	// расширеный метод BaseStatus(), дополнительно возвращает ресурсы со склада у игрока.
	// todo должен возвращать цены на продажу
	// todo продумать систему ценобразования на базах

	userBase, _ := bases.Bases.Get(user.InBaseID)
	userBase.GetSumEfficiency()
	userStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

	userBaseResource := make(map[int]*inv.Slot)
	for i, resourceSlot := range userBase.CurrentResources {

		userBase.GetRecyclePercent(resourceSlot.ItemID)

		var newSlot inv.Slot
		deepcopy.Copy(&newSlot, &resourceSlot)
		userBaseResource[i] = &newSlot
		userBaseResource[i].Quantity = userStorage.ViewQuantityItems(resourceSlot.ItemID, resourceSlot.Type)
	}

	lobbyPipe <- Message{
		Event:          "GetDetails",
		UserID:         user.GetID(),
		InventorySlots: userBaseResource,
		Base:           userBase,
	}
}
