package Squad

import "log"

func DeleteSquad(id int)  {
	// удаляем мазер шипы
	_, err := db.Exec("DELETE FROM squad_mother_ship WHERE id_squad=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	// юнитов
	_, err = db.Exec("DELETE FROM squad_units WHERE id_squad=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	// эквип
	_, err = db.Exec("DELETE FROM squad_equipping WHERE id_squad=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	// отряд
	_, err = db.Exec("DELETE FROM squads WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	//todo по хорошему это должна быть транзакиця
}

