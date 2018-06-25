package inventory

import (
	"log"
	"database/sql"
	"../detailUnit"
	"../dbConnect"
)

type Squad struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	MatherShip *MatherShip        `json:"mather_ship"`
	Units      map[int]*Unit      `json:"units"`
	Equip      map[int]*Equipping `json:"equip"`
}

func (squad *Squad) GetSquadMatherShip() {

	rows, err := dbConnect.GetDBConnect().Query("Select id_mother_ship FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
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

	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO squad_mother_ship (id_squad, id_mother_ship) "+
		"VALUES ($1, $2)", squad.ID, matherShip.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func (squad *Squad) DelMatherShip() () {

	squad.MatherShip = nil

	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}

}

func (squad *Squad) ReplaceMatherShip(id int) () {
	squad.DelMatherShip()
	squad.AddMatherShip(id)
}

func (squad *Squad) GetSquadUnits() {
	rows, err := dbConnect.GetDBConnect().Query("Select slot_in_mother_ship, id_chassis, id_weapon, id_tower, id_body, id_radar FROM squad_units WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make(map[int]*Unit)

	for rows.Next() {

		var unit Unit

		var matherSlot int

		chassisID := sql.NullInt64{}
		weaponID := sql.NullInt64{}
		towerID := sql.NullInt64{}
		bodyID := sql.NullInt64{}
		radarID := sql.NullInt64{}

		err := rows.Scan(&matherSlot, &chassisID, &weaponID, &towerID, &bodyID, &radarID)
		if err != nil {
			println("get unit")
			log.Fatal(err)
		}

		if chassisID.Valid {
			chassis := detailUnit.GetChass(int(chassisID.Int64))
			unit.SetChassis(chassis)
		}
		if weaponID.Valid {
			weapon := detailUnit.GetWeapon(int(weaponID.Int64))
			unit.SetWeapon(weapon)
		}
		if towerID.Valid {
			tower := detailUnit.GetTower(int(towerID.Int64))
			unit.SetTower(tower)
		}
		if bodyID.Valid {
			body := detailUnit.GetBody(int(bodyID.Int64))
			unit.SetBody(body)
		}
		if radarID.Valid {
			radar := detailUnit.GetRadar(int(radarID.Int64))
			unit.SetRadar(radar)
		}

		unit.CalculateParametersUnit()
		units[matherSlot] = &unit
	}

	squad.Units = units
}

func (squad *Squad) AddUnit(unit *Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {

		squad.Units[slot] = unit

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

		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO squad_units (id_squad, slot_in_mother_ship, id_chassis, id_weapon, id_tower, id_body, id_radar) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)", squad.ID, slot, ChassisID, WeaponID, TowerID, BodyID, RadarID)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (squad *Squad) DelUnit(slot int) (error) {

	squad.Units[slot] = nil

	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM squad_units WHERE id_squad=$1 AND slot_in_mother_ship=$2", squad.ID, slot)
	if err != nil {
		println("DelUnit")
		log.Fatal(err)
		return err
	}

	return nil
}

func (squad *Squad) ReplaceUnit(unit *Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {
		squad.DelUnit(slot)
		squad.AddUnit(unit, slot)
	}
}

func (squad *Squad) GetSquadEquip() {
	rows, err := dbConnect.GetDBConnect().Query("Select slot_in_mother_ship, id_equipping FROM squad_equipping WHERE id_squad=$1", squad.ID)
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

		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO squad_equipping (id_squad, slot_in_mother_ship, id_equipping) "+
			"VALUES ($1, $2, $3)", squad.ID, slot, equip.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (squad *Squad) DelEquip(slot int) (error) {
	squad.Equip[slot] = nil

	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM squad_equipping WHERE id_squad=$1 AND slot_in_mother_ship=$2", squad.ID, slot)
	if err != nil {
		println("DelEquip")
		log.Fatal(err)
		return err
	}

	return nil
}

func (squad *Squad) ReplaceEquip(equip *Equipping, slot int) {
	if squad.MatherShip.EquipmentSlots > slot {
		squad.DelEquip(slot)
		squad.AddEquip(equip, slot)
	}
}
