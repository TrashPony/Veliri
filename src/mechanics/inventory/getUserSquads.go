package inventory

import (
	"log"
	"../../dbConnect"
)

func GetUserSquads(userID int) (squads []*Squad, err error) {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*Squad, 0)

	for rows.Next() {
		var squad Squad

		err := rows.Scan(&squad.ID, &squad.Name)
		if err != nil {
			log.Fatal(err)
		}

		squad.GetSquadUnits()
		squad.GetSquadMatherShip()
		squad.GetSquadEquip()

		squads = append(squads, &squad)
	}

	return
}
