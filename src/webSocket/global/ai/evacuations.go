package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
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
				x, y := globalGame.GetXYCenterHex(mapBase.Q, mapBase.R)
				transport.X = x
				transport.Y = y

				// TODO выедает производительность по неведомой причине, возможная причина расчеты в недостижимые точки
				// после выставления минимальной скорость 5 и стартовой скорости 5 производительность выросла в разы
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
		x := int(float64(radius) * math.Cos(radRotate))
		y := int(float64(radius) * math.Sin(radRotate))

		xBase, yBase := globalGame.GetXYCenterHex(transportBase.Q, transportBase.R)

		x += xBase // докидываем положение базы
		y += yBase // докидываем положение базы

		// формируем путь для движения
		_, path := globalGame.MoveTo(float64(transport.X), float64(transport.Y), 15, 5, 5,
			float64(x), float64(y), transport.Rotate, 10, mp, true, nil, false, false, nil)
		// запускаем транспорт
		FlyTransport(transport, transportBase, mp, path)
		time.Sleep(1 * time.Second)
	}
}

func FlyTransport(transport *base.Transport, transportBase *base.Base, mp *_map.Map, path []unit.PathUnit) {
	for _, pathUnit := range path {
		time.Sleep(200 * time.Millisecond)

		for transport.Job {
			// если транспорт начал свою работу то ждем пока он не освободится)
			time.Sleep(200 * time.Millisecond)
		}

		go wsGlobal.SendMessage(wsGlobal.Message{Event: "FreeMoveEvacuation", PathUnit: pathUnit,
			BaseID: transportBase.ID, TransportID: transport.ID, IDMap: mp.Id})

		transport.X = pathUnit.X
		transport.Y = pathUnit.Y
		transport.Rotate = pathUnit.Rotate
		transport.Speed = pathUnit.Speed
	}
}

func TransportMonitor(transportBase *base.Base, mp *_map.Map) {
	for {
		for _, coordinate := range mp.HandlersCoordinates {

			xHandle, yHandle := globalGame.GetXYCenterHex(coordinate.Q, coordinate.R)
			xBase, yBase := globalGame.GetXYCenterHex(transportBase.Q, transportBase.R)

			dist := int(globalGame.GetBetweenDist(xBase, yBase, xHandle, yHandle))
			if dist < transportBase.GravityRadius {
				if coordinate.Transport {
					wsGlobal.CheckTransportCoordinate(coordinate.Q, coordinate.R, 20, 60, mp.Id)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}
