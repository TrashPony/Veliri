package Squad

import (
	"log"
	"../DetailUnit"
	"errors"
)

type Squad struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	MatherShip *MatherShip        `json:"mather_ship"`
	Units      map[int]*Unit      `json:"units"`
	Equip      map[int]*Equipping `json:"equip"`
}

func (squad *Squad) GetSquadMatherShip() {

	rows, err := db.Query("Select id_mother_ship FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	matherShip := &MatherShip{}

	for rows.Next() {

		err := rows.Scan(&matherShip.Id)
		if err != nil {
			log.Fatal(err)
		}

		matherShip = GetTypeMatherShip(matherShip.Id)
	}

	squad.MatherShip = matherShip
}

func (squad *Squad) AddMatherShip(id int) () {

	matherShip := GetTypeMatherShip(id)
	squad.MatherShip = matherShip

	_, err := db.Exec("INSERT INTO squad_mother_ship (id_squad, id_mother_ship) "+
		"VALUES ($1, $2)", squad.ID, matherShip.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func (squad *Squad) DelMatherShip() () {

	squad.MatherShip = nil

	_, err := db.Exec("DELETE FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}

}

func (squad *Squad) ReplaceMatherShip(id int) () {
	squad.DelMatherShip()
	squad.AddMatherShip(id)
}

func (squad *Squad) GetSquadUnits() {
	rows, err := db.Query("Select slot_in_mother_ship, id_chassis, id_weapon, id_tower, id_body, id_radar FROM squad_units WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make(map[int]*Unit)

	for rows.Next() {

		var unit Unit

		var matherSlot int

		var chassis DetailUnit.Chassis
		var weapon DetailUnit.Weapon
		var tower DetailUnit.Tower
		var body DetailUnit.Body
		var radar DetailUnit.Radar

		err := rows.Scan(&matherSlot, &chassis.Id, &weapon.Id, &tower.Id, &body.Id, &radar.Id)
		if err != nil {
			log.Fatal(err)
		}

		chassis = DetailUnit.GetChass(chassis.Id)
		weapon = DetailUnit.GetWeapon(weapon.Id)
		tower = DetailUnit.GetTower(tower.Id)
		body = DetailUnit.GetBody(body.Id)
		radar = DetailUnit.GetRadar(radar.Id)

		unit.SetChassis(&chassis)
		unit.SetWeapon(&weapon)
		unit.SetTower(&tower)
		unit.SetBody(&body)
		unit.SetRadar(&radar)

		units[matherSlot] = &unit
	}

	squad.Units = units
}

func (squad *Squad) AddUnit(unit *Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {

		squad.Units[slot] = unit

		_, err := db.Exec("INSERT INTO squad_units (id_squad, slot_in_mother_ship, id_chassis, id_weapon, id_tower, id_body, id_radar) "+
			"VALUES ($1, $2, $3, $4, $5, $6, &7)", squad.ID, slot, unit.Chassis.Id, unit.Weapon.Id, unit.Tower.Id, unit.Body.Id, unit.Radar.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (squad *Squad) DelUnit(slot int) (error) {
	if squad.MatherShip.UnitSlots > slot {

		squad.Units[slot] = nil

		_, err := db.Exec("DELETE FROM squad_units WHERE id_squad=$1, slot_in_mother_ship=$2", squad.ID, slot)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	} else {
		return errors.New("wrong slot")
	}
}

func (squad *Squad) ReplaceUnit(unit *Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {
		squad.DelUnit(slot)
		squad.AddUnit(unit, slot)
	}
}

func (squad *Squad) GetSquadEquip() {
	rows, err := db.Query("Select slot_in_mother_ship, id_equipping FROM squad_equipping WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equips = make(map[int]*Equipping)

	for rows.Next() {

		var matherSlot int

		var equip Equipping

		err := rows.Scan(&matherSlot, &equip.Id)
		if err != nil {
			log.Fatal(err)
		}

		equip = GetTypeEquip(equip.Id)

		equips[matherSlot] = &equip
	}

	squad.Equip = equips
}

func (squad *Squad) AddEquip(equip *Equipping, slot int) {
	if squad.MatherShip.EquipmentSlots > slot {

		squad.Equip[slot] = equip

		_, err := db.Exec("INSERT INTO squad_equipping (id_squad, slot_in_mother_ship, id_equipping) "+
			"VALUES ($1, $2, $3)", squad.ID, slot, equip.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (squad *Squad) DelEquip(slot int) (error){
	if squad.MatherShip.EquipmentSlots > slot {

		squad.Equip[slot] = nil

		_, err := db.Exec("DELETE FROM squad_equipping WHERE id_squad=$1, slot_in_mother_ship=$2", squad.ID, slot)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	} else {
		return errors.New("wrong slot")
	}
}

func (squad *Squad) ReplaceEquip(equip *Equipping, slot int) {
	if squad.MatherShip.EquipmentSlots > slot {
		squad.DelEquip(slot)
		squad.AddEquip(equip, slot)
	}
}