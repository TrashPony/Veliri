package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func updateThorium(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	if user != nil {

		// "squadInventory" потому что в глобальной игре нет больше инвентарей
		squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot, "squadInventory")

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		Move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})

		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MapID})
	}
}

func removeThorium(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	if user != nil {

		squadInventory.RemoveThorium(user, msg.ThoriumSlot, true)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		Move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MapID})
	}
}
