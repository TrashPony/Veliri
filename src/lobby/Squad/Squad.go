package Squad

import (
	"log"
	"../DetailUnit"
)

type Squad struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	MatherShip MatherShip    `json:"mather_ship"`
	Units      map[int]*Unit `json:"units"`
}

func AddNewSquad(name string, userID int) (err error) {
	_, err = db.Exec("INSERT INTO squads (name, id_user) VALUES ($1, $2)", name, userID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetUserSquads(userID int) (squads []*Squad, err error) {

	rows, err := db.Query("Select id, name FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	squads = make([]*Squad, 0)

	for rows.Next() {
		var squad Squad

		err := rows.Scan(&squad.ID, &squad.Name)
		if err != nil {
			log.Fatal(err)
		}

		squad.GetSquadUnits()

		squads = append(squads, &squad)
	}

	return
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

func (squad *Squad) GetSquadMatherShip() {

	rows, err := db.Query("Select id_mother_ship FROM squad_mother_ship WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var matherShip MatherShip

	for rows.Next() {

		err := rows.Scan(&matherShip.Id)
		if err != nil {
			log.Fatal(err)
		}
		// TODO формировать полный обьект мазершипа включая его настройки и модули
		matherShip = GetMatherShip(matherShip.Id)
	}

	squad.MatherShip = matherShip
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

func (squad *Squad) DelUnit(slot int) {
	if squad.MatherShip.UnitSlots > slot {

		squad.Units[slot] = nil

		_, err := db.Exec("DELETE FROM squad_units WHERE id_squad=$1, slot_in_mother_ship=$2", squad.ID, slot)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (squad *Squad) ReplaceUnit(unit *Unit, slot int)  {
	if squad.MatherShip.UnitSlots > slot {
		squad.DelUnit(slot)
		squad.AddUnit(unit, slot)
	}
}
