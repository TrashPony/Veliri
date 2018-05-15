package dbo

import (
	"../player"
	"../game"
	"../equip"
	"log"
)

func GetEquip(player player.Player, game *game.Game) []*equip.Equip {

	rows, err := db.Query("Select type.type, equip.used "+
		"FROM action_game_equipping as equip, equipping_type as type, users "+
		"WHERE users.id=$1 AND type.id=equip.id_type AND equip.id_user=$1 AND equip.id_game=$2", player.GetID(), game.Id)
	if err != nil {
		println("get user equip stat")
		log.Fatal(err)
	}
	defer rows.Close()

	equips := make([]*equip.Equip, 0)
	for rows.Next() {

		var userEquip equip.Equip

		err := rows.Scan(&userEquip.Type, &userEquip.Used)
		if err != nil {
			log.Fatal(err)
		}

		equips = append(equips, &userEquip)
	}

	return equips
}
