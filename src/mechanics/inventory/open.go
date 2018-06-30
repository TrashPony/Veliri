package inventory

import (
	"../player"
	"../db/get"
	"log"
)

func Open(client *player.Player) {
	squads, err := get.UserSquads(client.GetID())
	if err != nil {
		println("error, get Squads")
		log.Fatal(err)
	}

	client.SetSquads(squads)

	for _, activeSquad := range squads {
		if activeSquad.Active {
			client.SetSquad(activeSquad)
			return
		}
	}
}
