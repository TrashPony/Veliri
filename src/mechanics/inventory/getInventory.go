package inventory

import (
	"log"
	"../player"
	"../db/get"
	"../db/insert"
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
		newSquad := insert.FirstSquad(client.GetID())
		newSquad.Inventory = get.SquadInventory(newSquad.ID)
		client.SetSquad(newSquad)
	}
}
