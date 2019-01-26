package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/map"
	"../../mechanics/globalGame"
	"math"
	"math/rand"
	"time"
)

func EvacuationsLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		mapBases := bases.Bases.GetBasesByMap(mp.Id)
		for _, mapBase := range mapBases {
			for _, transport := range mapBase.Transports {

				// задаем начальное положение эвакуаторов как у баз
				x, y := globalGame.GetXYCenterHex(mapBase.Q, mapBase.R)
				transport.X = x
				transport.Y = y

				LaunchTransport(transport, mapBase, mp)
			}
		}
	}
}

func LaunchTransport(transport *base.Transport, transportBase *base.Base, mp *_map.Map) {
	// рандомные полеты без дела в территории СВОЕЙ БАЗЫ
	// мониторить ячейки для эвакуации, если они в ПРЕДЕЛАХ БАЗЫ

	// находим рандомную точку окружности что бы туда следовать

	radRotate := float64(rand.Intn(360)) * math.Pi / 180 // берем рандомный угол

	radius := rand.Intn(transportBase.GravityRadius) // и рандомную дальность в радиусе базы
	x := int(float64(radius) * math.Cos(radRotate))
	y := int(float64(radius) * math.Sin(radRotate))

	xBase, yBase := globalGame.GetXYCenterHex(transportBase.Q, transportBase.R)

	x += xBase // докидываем положение базы
	y += yBase // докидываем положение базы

	// формируем путь для движения
	_, path := globalGame.MoveTo(float64(transport.X), float64(transport.Y), 15, 15, 15,
		float64(x), float64(y), 0, mp, true, nil, false, false)

	// запускаем транспорт
	go FlyTransport(transport, transportBase, mp, path)
}

func FlyTransport(transport *base.Transport, transportBase *base.Base, mp *_map.Map, path []globalGame.PathUnit) {
	for _, pathUnit := range path {
		time.Sleep(400 * time.Millisecond)

		for transport.Job {
			// если транспорт начал свою работу то ждем пока он не освободится)
			time.Sleep(400 * time.Millisecond)
		}

		TransportMonitor(transport, transportBase, mp)

		globalPipe <- Message{Event: "FreeMoveEvacuation", PathUnit: pathUnit,
			BaseID: transportBase.ID, TransportID: transport.ID, idMap: mp.Id}

		transport.X = pathUnit.X
		transport.Y = pathUnit.Y
	}

	// как полетали создаем еще 1 рандомный путь для путеществия)
	go LaunchTransport(transport, transportBase, mp)
}

func TransportMonitor(transport *base.Transport, transportBase *base.Base, mp *_map.Map) {
	for _, coordinate := range mp.HandlersCoordinates {

		xHandle, yHandle := globalGame.GetXYCenterHex(coordinate.Q, coordinate.R)
		xBase, yBase := globalGame.GetXYCenterHex(transportBase.Q, transportBase.R)

		dist := int(globalGame.GetBetweenDist(xBase, yBase, xHandle, yHandle))
		if dist < transportBase.GravityRadius {
			if coordinate.Transport {
				CheckTransportCoordinate(coordinate.Q, coordinate.R, 10, 60, mp.Id)
			}
		}
	}
}
