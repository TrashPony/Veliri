package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func changeSquad(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squad_inventory.ChangeSquad(user, msg.SquadID)

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}

func renameSquad(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	err := squad_inventory.RenameSquad(user, msg.Name)

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
