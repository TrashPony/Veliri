package game

import (
	"log"
)

type UserStat struct {
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

	rows, err := db.Query("Select agu.id_game, users.name, agu.ready FROM action_game_user as agu, users WHERE agu.id_user=users.id AND agu.id_game=$1", idGame)
	if err != nil {
		println("gate game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*UserStat, 0)
	for rows.Next() {
		var user UserStat
		err := rows.Scan(&user.IdGame, &user.Name, &user.Ready)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}

	return users
}

func (user *UserStat) GetEquip()  {

}
