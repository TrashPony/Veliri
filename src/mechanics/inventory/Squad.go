package inventory
// todo создавать отряд есть у игрока его нет, и кидать в него стартовый эквип, тело мазершипа, тело 2х юнитов, 3 оружия, 2 типа патронов и эквип

/*func (squad *Squad) GetSquadMatherShip() {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id_mother_ship FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	gameMatherShip := &matherShip.MatherShip{}

	for rows.Next() {

		err := rows.Scan(&gameMatherShip.ID)
		if err != nil {
			log.Fatal(err)
		}

		gameMatherShip = GetTypeMatherShip(gameMatherShip.ID)
	}

	squad.MatherShip = gameMatherShip
}

func (squad *Squad) AddMatherShip(id int) () {

	gameMatherShip := GetTypeMatherShip(id)
	squad.MatherShip = gameMatherShip

	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO squad_mother_ship (id_squad, id_mother_ship) "+
		"VALUES ($1, $2)", squad.ID, gameMatherShip.ID)
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
	/*rows, err := dbConnect.GetDBConnect().Query("SELECT slot_in_mother_ship, id_chassis, id_weapon, id_tower, id_body, id_radar FROM squad_units WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make(map[int]*unit.Unit)

	for rows.Next() {

		var gameUnit unit.Unit

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

		if weaponID.Valid {
			weapon := get.GetWeapon(int(weaponID.Int64))
			gameUnit.SetWeapon(weapon)
		}

		if bodyID.Valid {
			body := get.GetBody(int(bodyID.Int64))
			gameUnit.SetBody(body)
		}

		units[matherSlot] = &gameUnit
	}

	squad.Units = units
}

func (squad *Squad) AddUnit(unit *unit.Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {

		squad.Units[slot] = unit

		ChassisID := sql.NullInt64{}
		WeaponID := sql.NullInt64{}
		TowerID := sql.NullInt64{}
		BodyID := sql.NullInt64{}
		RadarID := sql.NullInt64{}

		if unit.Weapon != nil {
			WeaponID = sql.NullInt64{int64(unit.Weapon.Id), true}
		}

		if unit.Body != nil {
			BodyID = sql.NullInt64{int64(unit.Body.Id), true}
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

func (squad *Squad) ReplaceUnit(unit *unit.Unit, slot int) {
	if squad.MatherShip.UnitSlots > slot {
		squad.DelUnit(slot)
		squad.AddUnit(unit, slot)
	}
}

func (squad *Squad) GetSquadEquip() {
	rows, err := dbConnect.GetDBConnect().Query("SELECT slot_in_mother_ship, id_equipping FROM squad_equipping WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equips = make(map[int]*equip.Equip)

	for rows.Next() {

		var matherSlot int

		var equipType equip.Equip

		err := rows.Scan(&matherSlot, &equipType.Id)
		if err != nil {
			log.Fatal(err)
		}

		equipType = GetTypeEquip(equipType.Id)

		equips[matherSlot] = &equipType
	}

	squad.Equip = equips
}

func (squad *Squad) AddEquip(equip *equip.Equip, slot int) {
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

func (squad *Squad) ReplaceEquip(equip *equip.Equip, slot int) {
	if squad.MatherShip.EquipmentSlots > slot {
		squad.DelEquip(slot)
		squad.AddEquip(equip, slot)
	}
}
*/