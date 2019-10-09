package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
)

// функция смотри колизию коробки со всеми обьектами мира, если колизия с коробкой то возвращается коробка с которой колизия
func CheckBoxCollision(box *boxInMap.Box, mp *_map.Map, excludeUnitID int) (bool, *boxInMap.Box) {

	rect := GetCenterRect(float64(box.X), float64(box.Y), float64(box.Height), float64(box.Width))
	rect.Rotate(box.Rotate)

	// проверяем колизии со статичной картой
	possibleMove, _ := searchStaticMapCollisionByRect(box.X, box.Y, mp, false, rect, 0, 150)
	if !possibleMove {
		return possibleMove, nil
	}

	possibleMove, _ = checkMapReservoir(mp, rect)
	if !possibleMove {
		return possibleMove, nil
	}

	units := globalGame.Clients.GetAllShortUnits(mp.Id, true)
	free := checkCollisionsUnits(rect, units, mp.Id, excludeUnitID)
	if !free {
		return false, nil
	}

	collisionBox := checkCollisionsBoxes(mp.Id, rect, true)
	if collisionBox != nil && collisionBox.ID != box.ID {
		return false, collisionBox
	} else {
		return true, nil
	}
}
