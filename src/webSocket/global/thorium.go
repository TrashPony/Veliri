package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
)

func updateThorium(user *player.Player, msg Message) {

	// "squadInventory" потому что в глобальной игре нет больше инвентарей
	squad_inventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot, "squadInventory")

	msg.ToX = user.GetSquad().MatherShip.ToX
	msg.ToY = user.GetSquad().MatherShip.ToY

	Move(user, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
	go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})

	go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(),
		ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MatherShip.MapID})
}

func removeThorium(user *player.Player, msg Message) {

	squad_inventory.RemoveThorium(user, msg.ThoriumSlot)

	msg.ToX = user.GetSquad().MatherShip.ToX
	msg.ToY = user.GetSquad().MatherShip.ToY

	Move(user, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
	go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
	go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(),
		ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MatherShip.MapID})
}
