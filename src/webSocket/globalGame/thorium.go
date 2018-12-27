package globalGame

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func updateThorium(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	if user != nil {

		squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}
	}
}

func removeThorium(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	if user != nil {

		squadInventory.RemoveThorium(user, msg.ThoriumSlot)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}
	}
}
