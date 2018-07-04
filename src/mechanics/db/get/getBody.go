package get

import (
	"../../../dbConnect"
	"../../gameObjects/detail"
	"log"
)

func Body(id int) (body *detail.Body) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, "+
		"vulnerability_to_kinetics, vulnerability_to_thermo, vulnerability_to_em, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power "+
		"wall_hack "+
		"FROM body_type "+
		"WHERE id=$1", id)
	if err != nil {
		log.Fatal("get body: " + err.Error())
	}
	defer rows.Close()

	body = &detail.Body{}

	for rows.Next() {
		err = rows.Scan(&body.ID, &body.Name, &body.MotherShip, &body.Speed, &body.Initiative, &body.MaxHP, &body.Armor, &body.EvasionCritical,
			&body.VulToKinetics, &body.VulToThermo, &body.VulToEM, &body.VulToExplosion, &body.RangeView, &body.Accuracy, &body.MaxPower, &body.RecoveryPower)
		if err != nil {
			log.Fatal("get body: " + err.Error())
		}
	}

	BodySlots(body)

	return body
}

func BodySlots(body *detail.Body) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT type_slot, number_slot, weapon, weapon_type "+
		"FROM body_slots "+
		"WHERE id_body = $1", body.ID)
	if err != nil {
		log.Fatal("get body slot " + err.Error())
	}
	defer rows.Close()

	body.Equipping = make(map[int]*detail.BodyEquipSlot)
	body.Weapons = make(map[int]*detail.BodyWeaponSlot)

	for rows.Next() {
		var slotType int
		var slotNumber int
		var slotWeapon bool
		var slotWeaponType string

		err := rows.Scan(&slotType, &slotNumber, &slotWeapon, &slotWeaponType)
		if err != nil {
			log.Fatal("get body slot " + err.Error())
		}

		if slotWeapon {
			weaponSlot := detail.BodyWeaponSlot{Type: slotType, Number: slotNumber, WeaponType: slotWeaponType}
			body.Weapons[slotNumber] = &weaponSlot
		} else {
			equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
			body.Equipping[slotNumber] = &equipSlot
		}
	}
}

type Boder interface {
	GetBody() *detail.Body
	GetID() int
}

func BodyEquip(ship Boder) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id_equipping, slot_in_body, type"+
		"FROM squad_mother_ship_equipping "+
		"WHERE id_squad_mother_ship = $1", ship.GetID())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var idEquip int
	var slot int
	var equipType string

	for rows.Next() {
		err := rows.Scan(&idEquip, &slot, &equipType)
		if err != nil {
			log.Fatal("get body equip " + err.Error())
		}

		for i, bodyEquipSlot := range ship.GetBody().Equipping {
			if bodyEquipSlot.Number == slot {
				if equipType == "equip" {
					ship.GetBody().Equipping[i].Equip = TypeEquip(idEquip)
				}
			}
		}

		for _, bodyWeaponSlot := range ship.GetBody().Weapons {
			if bodyWeaponSlot.Weapon != nil {
				if equipType == "weapon" {
					ship.GetBody().Weapons[bodyWeaponSlot.Number].Weapon = Weapon(idEquip)
				}
				if equipType == "ammo" {
					ship.GetBody().Weapons[bodyWeaponSlot.Number].Ammo = Ammo(idEquip)
				}
			}
		}
	}
}
