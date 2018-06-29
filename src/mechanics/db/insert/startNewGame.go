package insert

import (
	"../../../dbConnect"
	"../../lobby"
	"log"
)

func StartNewGame(game *lobby.Game) (int, bool) {

	id := 0

	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5) RETURNING id", // добавляем новую игру в БД
		game.Name, game.Map.Id, 0, "move", "").Scan(&id) // название игры, id карты, 0 - ход, Фаза движения, победитель

	if err != nil {
		println("add new game error")
		log.Fatal(err)
		return id, false
	}

	for _, user := range game.Users {
		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_squads (id_game, id_squad) VALUES ($1, $2)",
			id, user.GetSquad().ID)
			// todo обновление информации внутри сквада для мазр шипов, положение, на карте, снять все прошлые эффекты и тд
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

		/*for _, gameUnit := range user.Squad.Units {
			// todo обновление информации внутри сквада для юнитов, положение, на карте, снять все прошлые эффекты и тд
		}*/

		/*for _, equip := range user.Squad.Equip {
		    // todo обновление информации внутри сквада для эквипа, обнулить перезарядку
		}*/
	}

	err = AddCoordinateEffects(game.Map.Id, id)
	if err != nil {
		println("error db add coordinate effect new game")
		log.Fatal(err)
	}

	return id, true
}

func AddCoordinateEffects(mapID, gameID int) error {

	rows, err := dbConnect.GetDBConnect().Query("SELECT mc.x, mc.y, cte.id_effect "+
		"FROM map_constructor mc, coordinate_type ct, coordinate_type_effect cte "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id AND ct.id = cte.id_type; ", mapID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var x, y, effectID int

		err := rows.Scan(&x, &y, &effectID)
		if err != nil {
			println("start game get coordinate effects")
			log.Fatal(err)
			return err
		}

		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_zone_effects (id_game, id_effect, x, y, left_steps) VALUES ($1, $2, $3, $4, $5)",
			gameID, effectID, x, y, 999)

		if err != nil {
			println("start game add coordinate effects")
			log.Fatal(err)
			return err

		}
	}
	return nil
}
