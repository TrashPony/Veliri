package globalGame

import (
	"github.com/gorilla/websocket"
	"../../mechanics/squadInventory"
)

func updateThorium(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	user, ok := usersGlobalWs[ws]
	if ok {

		squadInventory.SetThorium(user, msg.InventorySlot ,msg.ThoriumSlot)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg, stopMove, moveChecker) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}
	}
}

func removeThorium(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	user, ok := usersGlobalWs[ws]
	if ok {

		squadInventory.RemoveThorium(user, msg.ThoriumSlot)

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg, stopMove, moveChecker) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
		globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}
	}
}
