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

	body.EquippingI = make(map[int]*detail.BodyEquipSlot) // todo чето как то пиздец
	body.EquippingII = make(map[int]*detail.BodyEquipSlot)
	body.EquippingIII = make(map[int]*detail.BodyEquipSlot)
	body.EquippingIV = make(map[int]*detail.BodyEquipSlot)
	body.EquippingV = make(map[int]*detail.BodyEquipSlot)

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
			if slotType == 1 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingI[slotNumber] = &equipSlot
			}
			if slotType == 2 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingII[slotNumber] = &equipSlot
			}
			if slotType == 3 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingIII[slotNumber] = &equipSlot
			}
			if slotType == 4 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingIV[slotNumber] = &equipSlot
			}
			if slotType == 5 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingV[slotNumber] = &equipSlot
			}
		}
	}
}

type Boder interface {
	GetBody() *detail.Body
	GetID() int
}

func BodyEquip(ship Boder, table string) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id_equipping, slot_in_body, type, type_slot, quantity "+
		" FROM " + table +
		" WHERE id_squad_unit = $1", ship.GetID())
	if err != nil {
		log.Fatal("get body equip" + err.Error())
	}
	defer rows.Close()

	var idEquip int
	var slot int
	var slotType int
	var equipType string
	var quantity int

	for rows.Next() {
		err := rows.Scan(&idEquip, &slot, &equipType, &slotType, &quantity)
		if err != nil {
			log.Fatal("scan body equip " + err.Error())
		}

		if slotType == 1 {
			ParseTypeSlot(ship.GetBody().EquippingI, idEquip, slot, equipType)
		}
		if slotType == 2 {
			ParseTypeSlot(ship.GetBody().EquippingII, idEquip, slot, equipType)
		}
		if slotType == 3 {
			ParseTypeSlot(ship.GetBody().EquippingIII, idEquip, slot, equipType)
		}
		if slotType == 4 {
			ParseTypeSlot(ship.GetBody().EquippingIV, idEquip, slot, equipType)
		}
		if slotType == 5 {
			ParseTypeSlot(ship.GetBody().EquippingV, idEquip, slot, equipType)
		}

		if equipType == "weapon" || equipType == "ammo" {
			for _, bodyWeaponSlot := range ship.GetBody().Weapons {
				if equipType == "weapon" {
					ship.GetBody().Weapons[bodyWeaponSlot.Number].Weapon = Weapon(idEquip)
				}
				if equipType == "ammo" {
					ship.GetBody().Weapons[bodyWeaponSlot.Number].Ammo = Ammo(idEquip)
					ship.GetBody().Weapons[bodyWeaponSlot.Number].AmmoQuantity = quantity
				}
			}
		}
	}
}

func ParseTypeSlot(equipping map[int]*detail.BodyEquipSlot, idEquip int, slot int, equipType string) {
	for i, bodyEquipSlot := range equipping {
		if bodyEquipSlot.Number == slot {
			if equipType == "equip" {
				equipping[i].Equip = TypeEquip(idEquip)
			}
		}
	}
}
