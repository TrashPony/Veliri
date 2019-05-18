package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
)

func outBase(user *player.Player, msg Message) {

	lobbyPipe <- Message{Event: "StartOutBase", UserID: user.GetID()}
	err := lobby.OutBase(user)

	// todo запускать метод в отдельной горутине
	// todo флаг выхода с базы, т.к. пока освобождается респаун игрок может передумать

	if err != nil {
		lobbyPipe <- Message{Event: "Error", Error: err.Error(), UserID: user.GetID()}
	} else {
		lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID()}
	}
}
