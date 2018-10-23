package insert

import (
	"../../../dbConnect"
	"../../lobby"
	"../updateSquad"
	"log"
)

func StartNewGame(game *lobby.Game) (int, bool) {

	id := 0

	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5) RETURNING id", // добавляем новую игру в БД
		game.Name, game.Map.Id, 1, "move", "").Scan(&id) // название игры, id карты, 0 - ход, Фаза движения, победитель

	if err != nil {
		println("add new game error")
		log.Fatal(err)
		return id, false
	}

	for _, user := range game.Users {
		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_squads (id_game, id_squad) VALUES ($1, $2)",
			id, user.GetSquad().ID)
		if err != nil {
			println("add matherShip error")
			log.Fatal(err)
			return id, false
		}

		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_user (id_game, id_user, ready) VALUES ($1, $2, $3)",
			id, user.GetID(), "false")
		if err != nil {
			println("add user game error")
			log.Fatal(err)
			return id, false
		}

		for _, slotUnit := range user.GetSquad().MatherShip.Units {
			if slotUnit.Unit != nil {
				slotUnit.Unit.Q = 0
				slotUnit.Unit.R = 0
				slotUnit.Unit.OnMap = false
				slotUnit.Unit.Target = nil
				slotUnit.Unit.QueueAttack = 0
				slotUnit.Unit.CalculateParams()
				slotUnit.Unit.ActionPoints = slotUnit.Unit.Speed
			}
		}

		// todo обновление информации внутри сквада для мазр шипов, положение, на карте, снять все прошлые эффекты и тд
		// todo обновление информации внутри сквада для эквипа, обнулить перезарядку

		user.GetSquad().MatherShip.Q = user.GetRespawn().X
		user.GetSquad().MatherShip.R = user.GetRespawn().Y
		user.GetSquad().MatherShip.Target = nil
		user.GetSquad().MatherShip.QueueAttack = 0

		updateSquad.Squad(user.GetSquad())
	}

	err = AddCoordinateEffects(game.Map.Id, id)
	if err != nil {
		println("error db add coordinate effect new game")
		log.Fatal(err)
	}

	return id, true
}

func AddCoordinateEffects(mapID, gameID int) error {

	rows, err := dbConnect.GetDBConnect().Query("SELECT mc.q, mc.r, cte.id_effect "+
		"FROM map_constructor mc, coordinate_type ct, coordinate_type_effect cte "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id AND ct.id = cte.id_type; ", mapID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var q, r, effectID int

		err := rows.Scan(&q, &r, &effectID)
		if err != nil {
			println("start game get coordinate effects")
			log.Fatal(err)
			return err
		}

		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_zone_effects (id_game, id_effect, q, r, left_steps) VALUES ($1, $2, $3, $4, $5)",
			gameID, effectID, q, r, 999)

		if err != nil {
			println("start game add coordinate effects")
			log.Fatal(err)
			return err

		}
	}
	return nil
}
