package get

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"log"
)

func WeaponsType() map[int]detail.Weapon {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" min_attack_range," +
		" range_attack," +
		" accuracy," +
		" ammo_capacity," +
		" artillery," +
		" power," +
		" max_hp," +
		" type," +
		" standard_size," +
		" size, " +
		" equip_damage," +
		" equip_critical_damage," +
		" x_attach," +
		" y_attach," +
		" rotate_speed," +
		" count_fire_bullet," +
		" bullet_speed," +
		" reload_time, " +
		" delay_following_fire," +
		" fire_positions " +
		"FROM weapon_type")
	if err != nil {
		log.Fatal("get all type weapon: " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]detail.Weapon)

	for rows.Next() {
		var weapon detail.Weapon
		var positions []byte

		err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy,
			&weapon.AmmoCapacity, &weapon.Artillery, &weapon.Power, &weapon.MaxHP, &weapon.Type, &weapon.StandardSize,
			&weapon.Size, &weapon.EquipDamage, &weapon.EquipCriticalDamage, &weapon.XAttach,
			&weapon.YAttach, &weapon.RotateSpeed, &weapon.CountFireBullet, &weapon.BulletSpeed, &weapon.ReloadTime,
			&weapon.DelayFollowingFire, &positions)
		if err != nil {
			log.Fatal("get all type scan weapon: " + err.Error())
		}

		err = json.Unmarshal(positions, &weapon.FirePositions)
		if err != nil {
			log.Fatal("get weapon type fire positions: ", err)
		}

		allType[weapon.ID] = weapon
	}

	return allType
}
