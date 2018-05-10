package game

import (
	"log"
)

type UserStat struct {
	ID     int      `json:"id"`
	IdGame int      `json:"id_game"`
	Name   string   `json:"name"`
	Ready  bool     `json:"ready"`
	Equip  []*Equip `json:"equip"`
}

type Equip struct {
	Type string `json:"type"`
	Used bool   `json:"used"`
}

func GetUserStat(idGame int) []*UserStat {

	rows, err := db.Query("Select agu.id_game, users.name, agu.ready, users.id "+
		"FROM action_game_user as agu, users "+
		"WHERE agu.id_user=users.id AND agu.id_game=$1", idGame)
	if err != nil {
		println("get game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*UserStat, 0)
	for rows.Next() {
		var user UserStat
		err := rows.Scan(&user.IdGame, &user.Name, &user.Ready, &user.ID)
		if err != nil {
			log.Fatal(err)
		}

		user.GetEquip()
		users = append(users, &user)
	}

	return users
}

func (user *UserStat) GetEquip() {
	rows, err := db.Query("Select type.type, equip.used "+
		"FROM action_game_equipping as equip, equipping_type as type, users "+
		"WHERE type.id=equip.id_type AND equip.id_user=$1 AND equip.id_game=$2", user.ID, user.IdGame)
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

		user.GetEquip()
		equips = append(equips, &equip)
	}

	user.Equip = equips
}
