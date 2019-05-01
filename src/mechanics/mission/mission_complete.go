package mission

import (
	"errors"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func Complete(client *player.Player, uuid string) (error, *dialog.Page) {
	endMission, ok := client.Missions[uuid]
	if ok {
		if endMission.Type == "delivery" {
			return deliveryComplete(client, endMission)
		}
	}

	dbPlayer.UpdateUser(client)
	return errors.New("wrong mission"), nil
}
