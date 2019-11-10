package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/satori/go.uuid"
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

	if resp.Event == "FreeMoveEvacuation" {
		view, radar := client.GetSquad().CheckViewCoordinate(resp.PathUnit.X, resp.PathUnit.Y)
		if view {
			return resp
		}

		if radar {
			// получаем метку, подменяем обьект на метку затирая методанные транспорта, а путь оставляем прежним
			radarMark := client.GetSquad().GetVisibleObjectByID("transport" + strconv.Itoa(resp.TransportID))
			// если метки нет значит радар еще не нашел нехера и тупо игнорируем
			if radarMark != nil {
				resp.Event = "markMove"
				resp.RadarMark = radarMark
				resp.BaseID = 0
				resp.TransportID = 0
				return resp
			} else {
				return nil
			}
		}
	}

	return nil
}

func RadarWorker(user *player.Player, mp *_map.Map) {
	// TODO ошибка многопоточности, если клиент обновляет страницу больше 1го раза, создаются еще функции воркеры.
	// TODO надо гарантировать 1 воркер на 1 отряд
	// функция должна отслеживать что обьект вышел за пределы радара/обзора и сообщать об этом клиент
	// и наоборот небыл видим стал видим обьект вошел в зону радара/обзора
	//    -- для этого надо хранить предыдущие состояния (в прошлый раз видел, теперь нет - обьект вышел из поля зрения)
	// каждый обьект в зоне радара должен иметь метку например: objectType + id
	// каждой метке радара надо давать uuid что бы можно было ее двигать и удалять
	// так же метод получения UUID метки по objectType + id для фильтра исходящих сообщений в метода CheckView()
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
		// смотрим транспорты видим мы их или нет
		for _, gameBase := range bases.Bases.GetBasesByMap(mp.Id) {
			for _, transport := range gameBase.Transports {

				oldVisible := user.GetSquad().GetVisibleObjectByID("transport" + strconv.Itoa(transport.ID))
				view, radar := user.GetSquad().CheckViewCoordinate(transport.X, transport.Y)

				markEvent, objEvent, newMark := checkObjects(oldVisible, transport.ID, "fly", "transport", view, radar)
				sendRadarMessage(markEvent, objEvent, newMark, transport, transport.X, transport.Y)
			}
		}

		// все не обновленные обьекты считаются потеряными из виду, например телепорт смерть и тд
		user.GetSquad().RadarLock()
		for _, vObj := range user.GetSquad().VisibleObjects {
			if !vObj.Update {
				sendRadarMessage("removeRadarMark", "removeObj", vObj, nil, 0, 0)

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
