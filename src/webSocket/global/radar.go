package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/getlantern/deepcopy"
	"github.com/satori/go.uuid"
	"reflect"
	"strconv"
	"time"
)

func CheckView(client *player.Player, resp *Message) *Message {
	// эта функция проверяет видит то или иное действие игрок
	// если игрок видит прямой видимостью то сообщение остается без изменений и отсылается
	// если игрок видит действие радаром, то надо создать новое сообщение с меткой а не обьектом
	// если игрок не видит совсем то нечего не отсылаем
	// для каждого типа сообщения своя трансформация

	// радар может отслеживать:
	// - передвижения эвакуаторов, юнитов, ящиков
	// - уничтожение эвакуаторов, юнитов, ящиков
	// - уничтожение руд

	// радар не ослеживает: действия и стрельбу

	// наверно это не супер оптимально но я не очень умный С:

	var msg Message // создаем копию сообщения что бы не испортить его для других пользователей
	err := deepcopy.Copy(&msg, &resp)
	if err != nil {
		println(err.Error())
	}

	if msg.Event == "FreeMoveEvacuation" || msg.Event == "startMoveEvacuation" || msg.Event == "MoveEvacuation" ||
		msg.Event == "placeEvacuation" || msg.Event == "ReturnEvacuation" || msg.Event == "stopEvacuation" {
		view, radar := client.GetSquad().CheckViewCoordinate(msg.PathUnit.X, msg.PathUnit.Y)
		if view {
			return &msg
		}

		if radar {
			// получаем метку, подменяем обьект на метку затирая методанные транспорта, а путь оставляем прежним
			radarMark := client.GetSquad().GetVisibleObjectByID("transport" + strconv.Itoa(msg.TransportID))
			// если метки нет значит радар еще не нашел нехера и тупо игнорируем
			if radarMark != nil {
				msg.Event = "markMove"
				msg.RadarMark = radarMark
				msg.BaseID = 0
				msg.TransportID = 0
				return &msg
			} else {
				return nil
			}
		}
	}

	if msg.Event == "MoveTo" {
		view, radar := client.GetSquad().CheckViewCoordinate(msg.PathUnit.X, msg.PathUnit.Y)
		if view {
			return &msg
		}

		if radar {
			// получаем метку, подменяем обьект на метку затирая методанные, а путь оставляем прежним
			radarMark := client.GetSquad().GetVisibleObjectByID("unit" + strconv.Itoa(msg.ShortUnit.ID))
			// если метки нет значит радар еще не нашел нехера и тупо игнорируем
			if radarMark != nil {
				msg.Event = "markMove"
				msg.RadarMark = radarMark
				msg.ShortUnit = nil
				return &msg
			}
		}
	}

	if msg.Event == "BoxTo" {
		view, radar := client.GetSquad().CheckViewCoordinate(msg.PathUnit.X, msg.PathUnit.Y)
		if view {
			return &msg
		}

		if radar {
			// получаем метку, подменяем обьект на метку затирая методанные, а путь оставляем прежним
			radarMark := client.GetSquad().GetVisibleObjectByID("box" + strconv.Itoa(msg.BoxID))
			// если метки нет значит радар еще не нашел нехера и тупо игнорируем
			if radarMark != nil {
				msg.Event = "markMove"
				msg.RadarMark = radarMark
				msg.BoxID = 0
				return &msg
			}
		}
	}

	if msg.Event == "FlyBullet" {
		view, _ := client.GetSquad().CheckViewCoordinate(msg.Bullet.X, msg.Bullet.Y)
		if view {
			return &msg
		}
		// TODO если пуля ушла в туман войны ее надо удалять, надо задействовать для этого RadarWorker,
		//  а ракеты можно и на радаре показывать
	}

	// ExplosionBullet - взрыв, визуал
	if msg.Event == "ExplosionBullet" {
		view, _ := client.GetSquad().CheckViewCoordinate(msg.Bullet.X, msg.Bullet.Y)
		if view {
			return &msg
		}
	}

	// MoveStop - месага чисто для визуала смысла не несет под радаром
	// PlaceUnit - отработает за счет радара
	// RemoveUnit - отработает за счет радара
	// startMining - видить действие могут ток те кто видит юнита визуально
	// stopMining - видить действие могут ток те кто видит юнита визуально
	// RotateGun - что бы видить поворот башни, надо видить юнита
	if msg.Event == "MoveStop" || msg.Event == "PlaceUnit" || msg.Event == "RemoveUnit" || msg.Event == "startMining" ||
		msg.Event == "stopMining" || msg.Event == "RotateGun" {
		view, _ := client.GetSquad().CheckViewCoordinate(msg.ShortUnit.X, msg.ShortUnit.Y)
		if view {
			return &msg
		}
	}

	// NewBox - отработает за счет радара
	if msg.Event == "NewBox" {
		view, _ := client.GetSquad().CheckViewCoordinate(msg.Box.X, msg.Box.Y)
		if view {
			return &msg
		}
	}

	// UpdateBox - видить содержимое могут только те кто видят ящик
	// openBox - отрыть ящик могут только те кто видят его
	// updateReservoir - видить обновление статы можно ток визуально
	// destroyReservoir - отработает радар в зоне радара
	// useDigger - видить действие могут ток те кто видит визуально, хотя тут не все так однозначно
	// FireWeapon - что бы видеть начало выстрела надо, видить визуально начало выстрела)
	if msg.Event == "UpdateBox" || msg.Event == "openBox" || msg.Event == "updateReservoir" || msg.Event == "destroyReservoir" ||
		msg.Event == "useDigger" || msg.Event == "FireWeapon" {
		view, _ := client.GetSquad().CheckViewCoordinate(msg.X, msg.Y)
		if view {
			return &msg
		}
	}

	return nil
}

func RadarWorker(user *player.Player, mp *_map.Map) {
	// функция должна отслеживать что обьект вышел за пределы радара/обзора и сообщать об этом клиент
	// и наоборот небыл видим стал видим обьект вошел в зону радара/обзора
	//    -- для этого надо хранить предыдущие состояния (в прошлый раз видел, теперь нет - обьект вышел из поля зрения)
	// каждый обьект в зоне радара должен иметь метку например: objectType + id
	// каждой метке радара надо давать uuid что бы можно было ее двигать и удалять
	// так же метод получения UUID метки по objectType + id для фильтра исходящих сообщений в метода CheckView()

	user.GetSquad().RadarWorkerWork = true
	defer func() {
		user.GetSquad().RadarWorkerWork = false
	}()

	// если мы заходим в метод значит произошла перезагрузка, радар инитим заного
	user.GetSquad().VisibleObjects = make(map[string]*squad.VisibleObjects)

	checkObjects := func(oldObj *squad.VisibleObjects, id int, typeMark, typeObject string, view, radar bool) (string, string, *squad.VisibleObjects) {

		defer func() {
			if oldObj != nil {
				oldObj.Update = true
			}
		}()

		if oldObj == nil && view {
			// мы не видили обьект совсем а теперь видим визуально
			oldObj = &squad.VisibleObjects{
				IDObject:   id,
				TypeObject: typeObject,
				UUID:       uuid.NewV1().String(),
				View:       view, Radar: radar, Type: typeMark}

			user.GetSquad().AddVisibleObject(oldObj)

			return "createRadarMark", "createObj", oldObj
		}

		if oldObj == nil && !view && radar {
			// мы не видили обьект совсем и видим теперь его на радаре

			oldObj = &squad.VisibleObjects{
				IDObject:   id,
				TypeObject: typeObject,
				UUID:       uuid.NewV1().String(),
				View:       view, Radar: radar, Type: typeMark}

			user.GetSquad().AddVisibleObject(oldObj)

			return "createRadarMark", "", oldObj
		}

		if oldObj != nil && !oldObj.View && oldObj.Radar && view {
			// мы видили обьект на радаре а теперь видим его визуально
			oldObj.View = true
			oldObj.Radar = true
			return "hideRadarMark", "createObj", oldObj
		}

		if oldObj != nil && oldObj.View && !view && radar {
			// мы видили обьект визуально а теперь видим только на радаре
			oldObj.View = false
			oldObj.Radar = true
			return "unhideRadarMark", "removeObj", oldObj
		}

		if oldObj != nil && oldObj.View && !view && !radar {
			// мы видили обьект визуально и он пропал
			user.GetSquad().RemoveVisibleObject(oldObj)
			return "removeRadarMark", "removeObj", oldObj
		}

		if oldObj != nil && !oldObj.View && oldObj.Radar && !view && !radar {
			// мы видили обьект на радаре и он пропал
			user.GetSquad().RemoveVisibleObject(oldObj)
			return "removeRadarMark", "", oldObj
		}

		return "", "", oldObj
		// во всем остальных случаях изменение состояния не произошло (но это не точно)
	}

	sendRadarMessage := func(markEvent, objEvent string, newMark *squad.VisibleObjects, object interface{}, x, y int) {
		if markEvent != "" || objEvent != "" {

			if objEvent != "createObj" {
				object = nil
			}

			go SendMessage(Message{Event: "radarWork",
				RadarMark: newMark, ActionMark: markEvent, ActionObject: objEvent, Object: object, X: x, Y: y,
				IDUserSend: user.GetID(), IDMap: mp.Id, Bot: user.Bot})
		}
	}

	for {
		// TODO сравнение и обновление обьектов

		if user.GetSquad().RadarWorkerExit {
			user.GetSquad().RadarWorkerExit = false
			return
		}

		// смотрим транспорты видим мы их или нет
		for _, gameBase := range bases.Bases.GetBasesByMap(mp.Id) {
			for _, transport := range gameBase.Transports {
				oldVisible := user.GetSquad().GetVisibleObjectByID("transport" + strconv.Itoa(transport.ID))
				view, radar := user.GetSquad().CheckViewCoordinate(transport.X, transport.Y)

				markEvent, objEvent, newMark := checkObjects(oldVisible, transport.ID, "fly", "transport", view, radar)
				go sendRadarMessage(markEvent, objEvent, newMark, transport, transport.X, transport.Y)
			}
		}

		// смотрим ящики, мои ящики.. видим мы их или нет
		for _, gameBox := range boxes.Boxes.GetAllBoxByMapID(mp.Id) {
			oldVisible := user.GetSquad().GetVisibleObjectByID("box" + strconv.Itoa(gameBox.ID))
			view, radar := user.GetSquad().CheckViewCoordinate(gameBox.X, gameBox.Y)

			markEvent, objEvent, newMark := checkObjects(oldVisible, gameBox.ID, "structure", "box", view, radar)
			go sendRadarMessage(markEvent, objEvent, newMark, gameBox, gameBox.X, gameBox.Y)
		}

		// смотрим на других юнитов которые не наши)
		for _, otherUnit := range globalGame.Clients.GetAllShortUnits(mp.Id) {
			oldVisible := user.GetSquad().GetVisibleObjectByID("unit" + strconv.Itoa(otherUnit.ID))
			view, radar := user.GetSquad().CheckViewCoordinate(otherUnit.X, otherUnit.Y)

			markEvent, objEvent, newMark := checkObjects(oldVisible, otherUnit.ID, "ground", "unit", view, radar)
			go sendRadarMessage(markEvent, objEvent, newMark, otherUnit, otherUnit.X, otherUnit.Y)
		}

		// смотрим на залежи ресурсов
		for _, x := range mp.Reservoir {
			for _, reservoir := range x {
				oldVisible := user.GetSquad().GetVisibleObjectByID("reservoir" + strconv.Itoa(reservoir.ID))
				view, radar := user.GetSquad().CheckViewCoordinate(reservoir.X, reservoir.Y)

				markEvent, objEvent, newMark := checkObjects(oldVisible, reservoir.ID, "resource", "reservoir", view, radar)
				go sendRadarMessage(markEvent, objEvent, newMark, reservoir, reservoir.X, reservoir.Y)
			}
		}

		// смотрим динамические обьекты, в отличие от прошлых обьектов эти можно увидеть только визуально
		// при этом пользователь их запоминает, тоесть он 1 раз увидил куст, ушел в другой конец карты,
		// куст будет виден через туман войны, НО если куст убьют то игрок будет всеравно его видеть
		// в тумане пока не откроет его зону снова.

		// проверяем видит ли юзер новые обьекты
		for _, x := range mp.GetCopyMapDynamicObjects() {
			for _, obj := range x {

				view, _ := user.GetSquad().CheckViewCoordinate(obj.X, obj.Y)
				memoryObj := user.GetMapDynamicObject(mp.Id, obj.X, obj.Y)

				if view && memoryObj == nil {
					user.AddDynamicObject(obj, mp.Id)
					go SendMessage(Message{Event: "radarWork", RadarMark: &squad.VisibleObjects{TypeObject: "dynamic_objects", IDObject: obj.ID},
						ActionMark: "", ActionObject: "createObj", Object: obj, X: obj.X, Y: obj.Y, IDUserSend: user.GetID(), IDMap: mp.Id, Bot: user.Bot})
				}

				if view && memoryObj != nil && !reflect.DeepEqual(obj, memoryObj) {
					user.RemoveDynamicObject(memoryObj, mp.Id)
					user.AddDynamicObject(obj, mp.Id)

					go SendMessage(Message{Event: "radarWork", RadarMark: &squad.VisibleObjects{TypeObject: "dynamic_objects", IDObject: obj.ID},
						ActionMark: "", ActionObject: "updateObj", Object: obj, X: obj.X, Y: obj.Y, IDUserSend: user.GetID(), IDMap: mp.Id, Bot: user.Bot})
				}
			}
		}

		// проверяем видит ли место где были старые обьекты но их уже нет
		for _, x := range user.GetMapDynamicObjects(mp.Id) {
			for _, memoryObj := range x {

				view, _ := user.GetSquad().CheckViewCoordinate(memoryObj.X, memoryObj.Y)
				obj := mp.GetDynamicObjects(memoryObj.X, memoryObj.Y)

				if view && obj == nil {
					user.RemoveDynamicObject(memoryObj, mp.Id)
					go SendMessage(Message{Event: "radarWork", RadarMark: &squad.VisibleObjects{TypeObject: "dynamic_objects", IDObject: memoryObj.ID},
						ActionMark: "", ActionObject: "removeObj", Object: memoryObj, X: memoryObj.X, Y: memoryObj.Y,
						IDUserSend: user.GetID(), IDMap: mp.Id, Bot: user.Bot})
				}
			}
		}

		// TODO пули

		// все не обновленные обьекты считаются потеряными из виду, например телепорт смерть и тд
		user.GetSquad().RadarLock()
		for _, vObj := range user.GetSquad().VisibleObjects {
			if !vObj.Update {
				go sendRadarMessage("removeRadarMark", "removeObj", vObj, nil, 0, 0)

				// что то пошло не так с мьютексами)
				user.GetSquad().RadarUnlock()
				user.GetSquad().RemoveVisibleObject(vObj)
				user.GetSquad().RadarLock()
			} else {
				vObj.Update = false
			}
		}
		user.GetSquad().RadarUnlock()

		time.Sleep(100 * time.Millisecond)
	}
}
