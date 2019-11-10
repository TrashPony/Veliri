package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	wsGlobal "github.com/TrashPony/Veliri/src/webSocket/global"
	"math"
	"math/rand"
	"time"
)

func EvacuationsLife() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		mapBases := bases.Bases.GetBasesByMap(mp.Id)
		for _, mapBase := range mapBases {

			// запускаем отслеживать транспортные товки
			go TransportMonitor(mapBase, mp)

			for _, transport := range mapBase.Transports {

				// задаем начальное положение эвакуаторов как у баз
				transport.X = mapBase.X
				transport.Y = mapBase.Y

				go LaunchTransport(transport, mapBase, mp)
			}
		}
	}
}

func LaunchTransport(transport *base.Transport, transportBase *base.Base, mp *_map.Map) {
	// рандомные полеты без дела в территории СВОЕЙ БАЗЫ
	// мониторить ячейки для эвакуации, если они в ПРЕДЕЛАХ БАЗЫ

	// находим рандомную точку окружности что бы туда следовать

	for {
		radRotate := float64(rand.Intn(360)) * math.Pi / 180 // берем рандомный угол

		radius := rand.Intn(transportBase.GravityRadius) // и рандомную дальность в радиусе базы
		x := int(float64(radius)*math.Cos(radRotate)) + transportBase.X
		y := int(float64(radius)*math.Sin(radRotate)) + transportBase.Y

		// формируем путь для движения
		// минимальная и текущая скорость должна быть 1 иначе будут мертвые зоны и дедлоки
		err, path := move.To(float64(transport.X), float64(transport.Y), 70, 10, 10,
			float64(x), float64(y), transport.Rotate, 100, 300)

		if err != nil {
			continue
		}

		// запускаем транспорт
		if FlyTransport(transport, transportBase, mp, path) {
			continue
		}
	}
}

func FlyTransport(transport *base.Transport, transportBase *base.Base, mp *_map.Map, path []*unit.PathUnit) bool {
	for _, pathUnit := range path {

		time.Sleep(time.Duration(pathUnit.Millisecond) * time.Millisecond)

		for transport.Job {
			// если транспорт начал свою работу то ждем пока он не освободится)
			time.Sleep(200 * time.Millisecond)
		}

		go wsGlobal.SendMessage(wsGlobal.Message{Event: "FreeMoveEvacuation", PathUnit: pathUnit,
			BaseID: transportBase.ID, TransportID: transport.ID, IDMap: mp.Id, NeedCheckView: true})

		transport.X = pathUnit.X
		transport.Y = pathUnit.Y
		transport.Rotate = pathUnit.Rotate
		transport.Speed = pathUnit.Speed
	}

	return true
}

func TransportMonitor(transportBase *base.Base, mp *_map.Map) {
	for {
		for _, coordinate := range mp.HandlersCoordinates {

			dist := int(game_math.GetBetweenDist(transportBase.X, transportBase.Y, coordinate.X, coordinate.Y))
			if dist < transportBase.GravityRadius {
				if coordinate.Transport {
					wsGlobal.CheckTransportCoordinate(coordinate.X, coordinate.Y, 20, 60, mp.Id)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}
