package db

import (
	"../player"
	"../game"
	"../equip"
	"../effect"
	"log"
)

func GetEquip(player player.Player, game *game.Game) []*equip.Equip {

	rows, err := db.Query("Select equip.id, type.type, equip.used, type.specification "+
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

		err := rows.Scan(&userEquip.Id, &userEquip.Type, &userEquip.Used, &userEquip.Specification)
		if err != nil {
			log.Fatal(err)
		}

		GetEffectsEquip(&userEquip)

		equips = append(equips, &userEquip)
	}

	return equips
}

func GetEffectsEquip(equip *equip.Equip)  {

	equip.Effects = make([]effect.Effect, 0)

	rows, err := db.Query(" SELECT et.id, et.name, et.type, et.steps_time, et.parameter, et.quantity, et.percentages"+
		" FROM equip_effects, effects_type et WHERE equip_effects.id_equip=$1 AND equip_effects.id_effect=et.id;", equip.Id)
	if err != nil {
		println("get user equip effects")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var equipEffect effect.Effect

		err := rows.Scan(&equipEffect.TypeID, &equipEffect.Name, &equipEffect.Type, &equipEffect.StepsTime, &equipEffect.Parameter, &equipEffect.Quantity, &equipEffect.Percentages)
		if err != nil {
			log.Fatal(err)
		}

		equip.Effects = append(equip.Effects, equipEffect)
	}
}
