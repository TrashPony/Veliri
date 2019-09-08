package update

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"log"
)

func Squad(squad *squad.Squad, full bool) {

	squad.UpdateLock()

	tx, err := dbConnect.GetDBConnect().Begin()

	defer squad.UpdateUnlock()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("update squad tx error: " + err.Error())
	}

	_, err = tx.Exec("UPDATE squads SET active=$1, in_game=$2, id_base=$4, name=$5 WHERE id=$3",
		squad.Active, squad.InGame, squad.ID, squad.BaseID, squad.Name)

	if err != nil {
		log.Fatal("update squad" + err.Error())
	}

	if full {
		Units(squad, tx)
		InventorySquad(squad, tx)
	}

	MotherShip(squad, tx)
	SquadThorium(squad, tx)

	tx.Commit()
}

func SquadThorium(squad *squad.Squad, tx *sql.Tx) {
	_, err := tx.Exec("DELETE FROM squad_thorium_slots WHERE id_squad = $1", squad.ID)
	if err != nil {
		log.Fatal("delete thorium" + err.Error())
	}

	if squad.MatherShip.Body == nil {
		return
	}

	for _, slot := range squad.MatherShip.Body.ThoriumSlots {
		_, err := tx.Exec("INSERT INTO squad_thorium_slots (id_squad, slot, thorium) VALUES ($1, $2, $3)",
			squad.ID, slot.Number, slot.Count)
		if err != nil {
			log.Fatal("add new thorium slot" + err.Error())
		}
	}
}
