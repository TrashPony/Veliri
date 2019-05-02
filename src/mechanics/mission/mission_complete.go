package mission

import (
	"errors"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func Complete(client *player.Player, uuid string) (error, *dialog.Page) {
	endMission, ok := client.Missions[uuid]
	defer dbPlayer.UpdateUser(client)

	if ok {
		if endMission.Type == "delivery" {
			return deliveryComplete(client, endMission)
		}
	}

	return errors.New("wrong mission"), nil
}
