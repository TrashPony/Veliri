package lobby

import (

)

func StartNewGame(game *LobbyGames) (int, bool) {

	id := 0
	// TODO переделать на транзакции, передалть запросы под новую логику
	/*err := dbConnect.GetDBConnect().QueryRow("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5) RETURNING id", // добавляем новую игру в БД
		game.Name, game.Map.Id, 0, "Init", "").Scan(&id) // название игры, id карты, 0 - ход, Фаза Инициализации (растановка войск), победитель

	if err != nil {
		println("add new game error")
		log.Fatal(err)
		return id, false
	}

	for _, user := range game.Users {
		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_mother_ship (id_game, id_type, id_user, x, y) VALUES ($1, $2, $3, $4, $5)",
			id, user.Squad.MatherShip.Id, user.Id, user.Respawn.X, user.Respawn.Y)
		if err != nil {
			println("add matherShip error")
			log.Fatal(err)
			return id, false
		}

		_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_user (id_game, id_user, ready) VALUES ($1, $2, $3)",
			id, user.Id, "false")
		if err != nil {
			println("add user game error")
			log.Fatal(err)
			return id, false
		}

		for _, gameUnit := range user.Squad.Units {

			ChassisID := sql.NullInt64{}
			WeaponID := sql.NullInt64{}
			TowerID := sql.NullInt64{}
			BodyID := sql.NullInt64{}
			RadarID := sql.NullInt64{}


			if gameUnit.Weapon != nil {
				WeaponID = sql.NullInt64{int64(gameUnit.Weapon.Id), true}
			}

			if gameUnit.Body != nil {
				BodyID = sql.NullInt64{int64(gameUnit.Body.Id), true}
			}


			_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_unit (id_user, id_game, "+ // вносим методанные юнита
				"id_chassis, id_weapons, id_tower, id_body, id_radar, "+ // части тела
				"weight, speed, initiative, damage, range_attack, min_attack_range, area_attack, "+ // характиристики
				"type_attack, hp, armor, evasion_critical, vul_kinetics, vul_thermal, vul_em, vul_explosive, "+
				"range_view, accuracy, wall_hack, action, target, queue_attack, rotate, x, y, on_map, max_hp)"+ // TODO надо узнать как можно это сделать проще и лучше)
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33)",
				user.Id, id, ChassisID, WeaponID, TowerID, BodyID, RadarID,
				unit.Weight, unit.Speed, unit.Initiative, unit.Damage, unit.RangeAttack,
				unit.MinAttackRange, unit.AreaAttack, unit.TypeAttack, unit.HP, unit.Armor,
				unit.EvasionCritical, unit.VulKinetics, unit.VulThermal, unit.VulEM, unit.VulExplosive,
				unit.RangeView, unit.Accuracy, unit.WallHack, false, "", 0, 0, 0, 0, false, unit.HP)
			if err != nil {
				println("add unit game error")
				log.Fatal(err)
				return id, false
			}
		}

		for _, equip := range user.Squad.Equip {
			_, err = dbConnect.GetDBConnect().Exec("INSERT INTO action_game_equipping (id_game, id_user, id_type, used) VALUES ($1, $2, $3, $4)",
				id, user.Id, equip.Id, false)
			if err != nil {
				println("add equip error")
				log.Fatal(err)
				return id, false
			}
		}
	}

	err = AddCoordinateEffects(game.Map.Id, id)
	if err != nil {
		println("error db add coordinate effect new game")
		log.Fatal(err)
	}*/

	return id, true
}

func AddCoordinateEffects(mapID, gameID int) error {

	/*rows, err := dbConnect.GetDBConnect().Query("SELECT mc.x, mc.y, cte.id_effect "+
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
*/
	return nil
}
