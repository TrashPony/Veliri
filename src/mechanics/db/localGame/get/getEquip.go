package get

import (
	"../../../player"
	"../../../localGame"
	"../../../gameObjects/equip"
	"../../../../dbConnect"
	"log"
)

func Equip(player player.Player, game *localGame.Game) []*equip.Equip {

	rows, err := dbConnect.GetDBConnect().Query("Select equip.id, type.type, equip.used, type.specification, "+
		"type.applicable, type.region " +
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

		err := rows.Scan(&userEquip.ID, &userEquip.Type, &userEquip.Used, &userEquip.Specification, &userEquip.Applicable, &userEquip.Region)
		if err != nil {
			log.Fatal(err)
		}

		EffectsEquip(&userEquip)

		equips = append(equips, &userEquip)
	}

	return equips
}
