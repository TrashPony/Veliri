package db

import (
	"../game"
	"../player"
	"../unit"
	"../watchZone"
	"log"
)

func GetPlayer(game *game.Game) []*player.Player {

	rows, err := db.Query("Select users.name, agu.ready, users.id "+
		"FROM action_game_user as agu, users "+
		"WHERE agu.id_user=users.id AND agu.id_game=$1", game.Id)
	if err != nil {
		println("get game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*player.Player, 0)

	for rows.Next() {
		var client player.Player

		var login string
		var ready bool
		var id int
		
		err := rows.Scan(&login, &ready, &id)
		if err != nil {
			log.Fatal(err)
		}

		client.SetLogin(login)
		client.SetReady(ready)
		client.SetID(id)

		equip := GetEquip(client, game)
		units := GetNotGameUnits(client, game)

		client.SetGameID(game.Id)
		client.SetEquip(equip)
		client.SetUnitsStorage(units)

		watchZone.UpdateWatchZone(game, &client)

		users = append(users, &client)
	}

	return users
}

func GetNotGameUnits(player player.Player, game *game.Game) []*unit.Unit {
	units := make([]*unit.Unit, 0)

	for _, gameUnit := range game.GetUnitsStorage() {
		if gameUnit.Owner == player.GetLogin() {
			units = append(units, gameUnit)
		}
	}

	return units
}