package globalGame

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"math"
)

type VisibleAnomaly struct {
	Signal      int `json:"signal"`
	Rotate      int `json:"rotate"`
	TypeAnomaly int `json:"type_anomaly"`
}

func GetVisibleAnomaly(user *player.Player, slot *detail.BodyEquipSlot) (visibleAnomalies []VisibleAnomaly, err error) {
	anomalies := maps.Maps.GetAllMapAnomaly(user.GetSquad().MapID)
	if slot == nil || slot.Equip == nil || slot.Equip.Applicable != "geo_scan" || slot.HP < 0 {
		return nil, errors.New("no anomaly")
	}

	visibleAnomalies = make([]VisibleAnomaly, 0)

	for _, anomaly := range anomalies {

		if anomaly == nil {
			continue
		}

		x, y := GetXYCenterHex(anomaly.GetQ(), anomaly.GetR())

		dist := GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, x, y)
		maxDist := (anomaly.GetPower() + slot.Equip.Radius) * 100

		if int(dist) < maxDist {
			signal := (int(dist) * 100) / maxDist // уровень сигнала от 0 до 100%

			//  math.Atan2 куда у - текущие у, куда х - текущие х, получаем угол
			needRad := math.Atan2(float64(y-user.GetSquad().MatherShip.Y), float64(x-user.GetSquad().MatherShip.X))
			// переводим в градусы
			needRotate := int(needRad * 180 / 3.14)

			diffRotate := user.GetSquad().MatherShip.Rotate - needRotate
			if diffRotate < 0 {
				diffRotate = 360 - diffRotate
			}

			if int(dist) < slot.Equip.Radius*100 {
				// если номалия ближе радиуса действия сканера, то мы определяем тип аномалии
				visibleAnomalies = append(visibleAnomalies, VisibleAnomaly{Signal: signal, Rotate: diffRotate, TypeAnomaly: anomaly.Type})
			} else {
				// иначе это неопределенная аноималия
				visibleAnomalies = append(visibleAnomalies, VisibleAnomaly{Signal: signal, Rotate: diffRotate, TypeAnomaly: 999})
			}
		}
	}

	if len(visibleAnomalies) > 0 {
		return visibleAnomalies, nil
	} else {
		return nil, errors.New("no anomaly")
	}
}
