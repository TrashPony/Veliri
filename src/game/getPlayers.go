package game

import (
	"log"
)

func GetPlayer(game *Game) []*Player {

	rows, err := db.Query("Select users.name, agu.ready, users.id "+
		"FROM action_game_user as agu, users "+
		"WHERE agu.id_user=users.id AND agu.id_game=$1", game.GetStat().Id)
	if err != nil {
		println("get game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*Player, 0)

	for rows.Next() {
		var player Player
		err := rows.Scan(&player.login, &player.ready, &player.id)
		if err != nil {
			log.Fatal(err)
		}

		equip := GetEquip(player, game)
		units := GetNotGameUnits(player, game)

		player.SetGameID(game.GetStat().Id)
		player.SetEquip(equip)
		player.SetNotGameUnits(units)
		player.UpdateWatchZone(game)

		users = append(users, &player)
	}

	return users
}

func GetNotGameUnits(player Player, game *Game) []*Unit {
	units := make([]*Unit, 0)

	for _, unit := range game.GetNotGameUnits() {
		if unit.Owner == player.login {
			units = append(units, unit)
		}
	}

	return units
}

type Equip struct {
	Type string `json:"type"`
	Used bool   `json:"used"`
}

func GetEquip(player Player, game *Game) []*Equip {

	rows, err := db.Query("Select type.type, equip.used "+
		"FROM action_game_equipping as equip, equipping_type as type, users "+
		"WHERE users.id=$1 AND type.id=equip.id_type AND equip.id_user=$1 AND equip.id_game=$2", player.GetID(), game.GetStat().Id)
	if err != nil {
		println("get user equip stat")
		log.Fatal(err)
	}
	defer rows.Close()

	equips := make([]*Equip, 0)
	for rows.Next() {

		var equip Equip

		err := rows.Scan(&equip.Type, &equip.Used)
		if err != nil {
			log.Fatal(err)
		}

		equips = append(equips, &equip)
	}

	return equips
}
