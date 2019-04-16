package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func outBase(user *player.Player, msg Message) {

	err := lobby.OutBase(user)

	// todo запускать метод в отдельной горутине
	// todo флаг выхода с базы, т.к. пока освобождается респаун игрок может передумать

	if err != nil {
		lobbyPipe <- Message{Event: "Error", Error: err.Error(), UserID: user.GetID()}
	} else {
		lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID()}
	}
}
