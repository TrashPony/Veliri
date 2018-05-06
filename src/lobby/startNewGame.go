package lobby

import (
	"log"
	"database/sql"
)

func StartNewGame(game *LobbyGames) (int, bool) {
	id := 0

	err := db.QueryRow("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5) RETURNING id", // добавляем новую игру в БД
		game.Name, game.Map.Id, 0, "Init", "").Scan(&id) // название игры, id карты, 0 - ход, Фаза Инициализации (растановка войск), победитель

	if err != nil {
		println("add new game error")
		log.Fatal(err)
		return id, false
	}

	for _, user := range game.Users {
		_, err = db.Exec("INSERT INTO action_mother_ship (id_game, id_type, id_user, x, y) VALUES ($1, $2, $3, $4, $5)",
			id, user.Squad.MatherShip.Id, user.Id, user.Respawn.X, user.Respawn.Y)
		if err != nil {
			println("add matherShip error")
			log.Fatal(err)
			return id, false
		}

		_, err = db.Exec("INSERT INTO action_game_user (id_game, id_user, ready) VALUES ($1, $2, $3)",
			id, user.Id, "false")
		if err != nil {
			println("add user game error")
			log.Fatal(err)
			return id, false
		}

		for _, unit := range user.Squad.Units {

			ChassisID := sql.NullInt64{}
			WeaponID := sql.NullInt64{}
			TowerID := sql.NullInt64{}
			BodyID := sql.NullInt64{}
			RadarID := sql.NullInt64{}

			if unit.Chassis != nil {
				ChassisID = sql.NullInt64{int64(unit.Chassis.Id), true}
			}

			if unit.Weapon != nil {
				WeaponID = sql.NullInt64{int64(unit.Weapon.Id), true}
			}

			if unit.Tower != nil {
				TowerID = sql.NullInt64{int64(unit.Tower.Id), true}
			}

			if unit.Body != nil {
				BodyID = sql.NullInt64{int64(unit.Body.Id), true}
			}

			if unit.Radar != nil {
				RadarID = sql.NullInt64{int64(unit.Radar.Id), true}
			}

			_, err = db.Exec("INSERT INTO action_game_unit (id_user, id_game, "+ // вносим методанные юнита
				"id_chassis, id_weapons, id_tower, id_body, id_radar, "+ // части тела
				"Weight, Speed, Initiative, Damage, RangeAttack, MinAttackRange, AreaAttack, "+ // характиристики
				"TypeAttack, HP, Armor, EvasionCritical, VulKinetics, VulThermal, VulEM, VulExplosive, "+
				"RangeView, Accuracy, WallHack)"+ // TODO надо узнать как можно это сделать проще и лучше)
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)",
				user.Id, id, ChassisID, WeaponID, TowerID, BodyID, RadarID,
				unit.Weight, unit.Speed, unit.Initiative, unit.Damage, unit.RangeAttack,
				unit.MinAttackRange, unit.AreaAttack, unit.TypeAttack, unit.HP, unit.Armor,
				unit.EvasionCritical, unit.VulKinetics, unit.VulThermal, unit.VulEM, unit.VulExplosive,
				unit.RangeView, unit.Accuracy, unit.WallHack)
			if err != nil {
				println("add unit game error")
				log.Fatal(err)
				return id, false
			}
		}

		for _, equip := range user.Squad.Equip {
			_, err = db.Exec("INSERT INTO action_game_equipping (id_game, id_user, id_type, used) VALUES ($1, $2, $3, $4)",
				id, user.Id, equip.Id, false)
			if err != nil {
				println("add equip error")
				log.Fatal(err)
				return id, false
			}
		}
	}

	return id, true
}
