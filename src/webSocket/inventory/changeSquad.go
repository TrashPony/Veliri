package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func changeSquad(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.ChangeSquad(user, msg.SquadID)

	UpdateSquad(user, err, ws, msg)
}

func renameSquad(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	err := squadInventory.RenameSquad(user, msg.Name)

	UpdateSquad(user, err, ws, msg)
}
