package mission

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

// если миссия доставка, в миссии все предыдущие акшены выполнены или он всего 1
// то забераем итем который игрок должен был доставить и выдаем награду
func deliveryComplete(client *player.Player, endMission *mission.Mission) (error, *dialog.Page) {
	if len(endMission.Actions) == 1 && endMission.Actions[0].BaseID == client.InBaseID {
		if client.GetSquad().Inventory.RemoveItemsByOtherInventory(endMission.Actions[0].NeedItems) {
			client.SetCredits(client.GetCredits() + endMission.RewardCr)
			// TODO выдать награду итемами

			delete(client.Missions, endMission.UUID)
			client.NotifyQueue[endMission.UUID] = &player.Notify{Name: "mission", Event: "complete", UUID: endMission.UUID}

			return nil, endMission.Actions[0].Dialog.GetPageByType("success")
		} else {
			return errors.New("few items"), endMission.Actions[0].Dialog.GetPageByType("failure")
		}
	}

	return errors.New("wrong mission"), nil
}
