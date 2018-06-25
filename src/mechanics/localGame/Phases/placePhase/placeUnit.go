package placePhase

import (
	"../../../unit"
	"../../../db"
	"../../../player"
	"../../../localGame"
)

func PlaceUnit(gameUnit *unit.Unit, x,y int, actionGame *localGame.Game, client *player.Player) error {

	gameUnit.SetX(x) //значит тут мы присваиваем юниту го координаты куда его поставили
	gameUnit.SetY(y) //значит тут мы присваиваем юниту го координаты куда его поставили
	gameUnit.SetOnMap(true) // устанавливаем ему параметр который говорит что он на игровом поле

	actionGame.DelUnitStorage(gameUnit.ID) // юдаяем его из трюма в обьекте игры
	actionGame.SetUnit(gameUnit)           // добавляем его как активного юнита в обьект игры

	client.DelUnitStorage(gameUnit.ID)     // юдаяем его из трюма в обьекте игры
	client.AddUnit(gameUnit)			   // добавляем его как активного юнита в обьект игры

	err := db.UpdateUnit(gameUnit)		   // обновляем его параметры в БД игры
	if err != nil {						   // если при добавление не случилось ишибки то отправляем nil что значит нет ошибок, юнит обновлен и стоит на карте
		return err
	}
	return nil

	// TODO что тебе надо сделать что бы изначально юнит получал не свои координаты куда его поставили а координаты MatherShip, client.GetMatherShip().X client.GetMatherShip().Y
	// TODO это будет точка откуда он идет в точку куда его поставили x,y
	// TODO надо алгоритмом поиска пути (править его не надо) найти для юнита путь от точки  client.GetMatherShip().X client.GetMatherShip().Y до точки x,y
	// TODO return ить от сюда его путь масивом обьектов "TruePatchNode"
}
