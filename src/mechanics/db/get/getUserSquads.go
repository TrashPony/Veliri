package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/squad"
)

func UserSquads(userID int) (squads []*squad.Squad, err error) {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*squad.Squad, 0)

	for rows.Next() {
		var userSquad squad.Squad

		err := rows.Scan(&userSquad.ID, &userSquad.Name)
		if err != nil {
			log.Fatal(err)
		}

		//userSquad.GetSquadUnits()
		//userSquad.GetSquadMatherShip()
		//userSquad.GetSquadEquip()

		squads = append(squads, &userSquad)
	}

	return
}
