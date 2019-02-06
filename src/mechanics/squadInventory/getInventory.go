package squadInventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/get"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func GetInventory(client *player.Player) {
	squads, err := get.UserSquads(client.GetID())
	if err != nil {
		println("error, get Squads")
		log.Fatal(err)
	}

	if len(squads) > 0 {
		client.SetSquads(squads)

		for _, activeSquad := range squads {
			if activeSquad.Active {
				client.SetSquad(activeSquad)
				return
			}
		}
	} else {
		// TODO проверять есть ли на базе где игрок МС корпуса
		//newSquad := squad.FirstSquad(client.GetID())
		//newSquad.Inventory = get.SquadInventory(newSquad.ID)
		//client.SetSquad(newSquad)
	}
}
