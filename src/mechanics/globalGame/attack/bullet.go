package attack

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"

func GetZAndSpeedBullet(bullet *unit.Bullet, startDist, currentDist, startSpeed float64) float64 {

	K := 0.02 // коэффицент влияние на скорость и высоту для артилерии

	percentPath := 100 - (int((currentDist * 100) / startDist))

	// у самоводящизся ракет не меняется скорость и высота полета
	if bullet.Ammo.Type != "missile" {

		// обычные пули летят по прямой и падают под силой притяжения (от 1 до 0)
		if !bullet.Weapon.Artillery {

			bullet.Z = 1 - ((1.0 / startDist) * (startDist - currentDist))
		} else {
			// артилерийские пули летят сначало вверх до 50% и падают после 50%, так же меняется скорость полета пули
			if bullet.Weapon.Artillery {
				if percentPath < 50 {
					// уменьшение скорости
					bullet.Z = 1 + (float64(percentPath) * K)
					return startSpeed - (startSpeed * float64(percentPath) / 100)
				} else {
					// увеличение скорости
					minSpeed := startSpeed * float64(50) / 100
					maxZ := 1 + (50 * K)
					bullet.Z = maxZ - (float64(percentPath-50) * K)
					return minSpeed + ((startSpeed * float64(percentPath) / 100) - minSpeed)
				}
			}
		}
	}

	return startSpeed
}
