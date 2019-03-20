package localGame

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func StartNewGame(gameName string, mapID int, users []*player.Player) (int, error) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("start new game error: " + err.Error())
	}

	id := 0

	err = tx.QueryRow("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5) RETURNING id", // добавляем новую игру в БД
		gameName, mapID, 1, "move", "").Scan(&id) // название игры, id карты, 0 - ход, Фаза движения, победитель

	if err != nil {
		println("add new game error")
		return id, err
	}

	for _, user := range users {
		_, err = tx.Exec("INSERT INTO action_game_squads (id_game, id_squad) VALUES ($1, $2)", id, user.GetSquad().ID)
		if err != nil {
			println("add matherShip error")
			return id, err
		}

		_, err = tx.Exec("INSERT INTO action_game_user (id_game, id_user, ready) VALUES ($1, $2, $3)", id, user.GetID(), false)
		if err != nil {
			println("add user game error")
			return id, err
		}

		for _, slotUnit := range user.GetSquad().MatherShip.Units {
			if slotUnit.Unit != nil {
				slotUnit.Unit.Q = 0
				slotUnit.Unit.R = 0
				slotUnit.Unit.OnMap = false
				slotUnit.Unit.Target = nil
				slotUnit.Unit.Move = false
				slotUnit.Unit.CalculateParams()
				slotUnit.Unit.ActionPoints = slotUnit.Unit.Speed
			}
		}

		// todo снять все прошлые эффекты
		// todo обнулить перезарядку у всего эквипа

		user.GetSquad().MatherShip.Target = nil
		user.GetSquad().MatherShip.Move = false
		user.GetSquad().MatherShip.CalculateParams()
		user.GetSquad().MatherShip.ActionPoints = user.GetSquad().MatherShip.Speed

		// говорим что отряд в бою
		user.GetSquad().InGame = true

		update.Squad(user.GetSquad(), true)
	}

	err = AddCoordinateEffects(tx, mapID, id)
	if err != nil {
		log.Fatal(err, "error db add coordinate effect new game")
	}

	tx.Commit()

	return id, nil
}

func AddCoordinateEffects(tx *sql.Tx, mapID, gameID int) error {

	rows, err := dbConnect.GetDBConnect().Query("SELECT " +
		"mc.q, " +
		"mc.r, " +
		"cte.id_effect " +
		""+
		"FROM map_constructor mc, coordinate_type ct, coordinate_type_effect cte "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id AND ct.id = cte.id_type; ", mapID)

	if err != nil {
		log.Fatal(err, "start game get coordinate effects")
	}

	for rows.Next() {
		var q, r, effectID int

		err := rows.Scan(&q, &r, &effectID)
		if err != nil {
			log.Fatal(err, "start game get coordinate effects")
			return err
		}

		_, err = tx.Exec("INSERT INTO action_game_zone_effects (id_game, id_effect, q, r, left_steps) VALUES ($1, $2, $3, $4, $5)",
			gameID, effectID, q, r, 999)

		if err != nil {
			log.Fatal(err, "start game add coordinate effects")
			return err

		}
	}
	return nil
}
