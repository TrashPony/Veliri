package lobby

import (
	"log"
	"./DetailUnit"
)

type Squad struct {
	ID         int
	Name       string
	MatherShip MatherShip
	Units      []*Unit
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
	rows, err := db.Query("Select id_chassis, id_weapon, id_tower, id_body, id_radar FROM squad_units WHERE id_squad=$1", squad.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make([]*Unit, 0)

	for rows.Next() {

		var unit Unit

		var chassis DetailUnit.Chassis
		var weapon DetailUnit.Weapon
		var tower DetailUnit.Tower
		var body DetailUnit.Body
		var radar DetailUnit.Radar

		err := rows.Scan(&chassis.Id, &weapon.Id, &tower.Id, &body.Id, &radar.Id)
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

		units = append(units, &unit)
	}

	squad.Units = units
}

func (squad *Squad)GetSquadMatherShip()  {

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

func (squad *Squad) AddUnit(unit *Unit) {
	if squad.MatherShip.UnitSlots > len(squad.Units) {
		squad.Units = append(squad.Units, unit)
		// TODO добавить его в базу
	}
}

func (squad *Squad) DelUnit(unit *Unit) {
	//TODO удалить из масива и базы
}
