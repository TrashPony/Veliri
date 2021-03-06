package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func SetThorium(user *player.Player, inventorySlot, numThoriumSlot int, source string) error {

	thoriumSlot, _ := user.GetSquad().MatherShip.Body.ThoriumSlots[numThoriumSlot]

	slot := getSlotBySource(user, inventorySlot, source)

	// торий это ресурс с ид 1 и типом "recycle"
	if thoriumSlot != nil && slot != nil && slot.ItemID == 1 && slot.Item != nil && slot.Type == "recycle" {

		needThorium := thoriumSlot.MaxCount - thoriumSlot.Count

		if needThorium <= slot.Quantity {
			thoriumSlot.Count += needThorium
			RemoveSlotBySource(user, inventorySlot, source, needThorium)
		} else {
			thoriumSlot.Count += slot.Quantity
			RemoveSlotBySource(user, inventorySlot, source, slot.Quantity)
		}

		go update.Squad(user.GetSquad(), true)
		return nil
	} else {
		return errors.New("no find slot")
	}
}

func RemoveThorium(user *player.Player, numThoriumSlot int) error {

	thoriumSlot, _ := user.GetSquad().MatherShip.Body.ThoriumSlots[numThoriumSlot]

	// торий это ресурс с ид 1 и типом "recycle"
	if thoriumSlot != nil && thoriumSlot.Count > 0 {
		item, _ := gameTypes.Resource.GetRecycledByID(1)

		user.GetSquad().MatherShip.Inventory.AddItem(item, "recycle", 1, thoriumSlot.Count, 1,
			item.Size/float32(thoriumSlot.Count), 1, false, user.GetID())

		thoriumSlot.Count = 0
		return nil
	} else {
		return errors.New("no thorium")
	}
}
