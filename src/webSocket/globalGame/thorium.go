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

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}

		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID}
	}
}

func removeThorium(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	if user != nil {

		squadInventory.RemoveThorium(user, msg.ThoriumSlot, true)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID}
	}
}
