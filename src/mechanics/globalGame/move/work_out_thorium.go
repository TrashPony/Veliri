package move

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"

func WorkOutThorium(thoriumSlots map[int]*detail.ThoriumSlot, afterburner, highGravity bool) float64 {
	fullCount := 0

	for _, slot := range thoriumSlots {
		if slot.Count > 0 {
			fullCount++
		}
	}

	efficiency := float64((fullCount * 100) / len(thoriumSlots))

	for _, slot := range thoriumSlots {
		if slot.Count > 0 && efficiency > 0 {

			// формула выроботки топлива, если работает только 1 ячейка из 3х то ее эффективность в 66% больше
			thorium := 1 / float32(100/efficiency)

			if !highGravity && !afterburner { // если не форсах и не высокая гравитация, то не тратим топливо
				return efficiency
			}

			if highGravity && afterburner { // если активирован форсаж и высокая гравитация то топливо тратиться х15
				thorium = thorium * 15
			}

			if !highGravity && afterburner { // если активирован форсаж и низкая гравитация то топливо тратиться х5
				thorium = thorium * 5
			}

			if slot.Inversion {
				slot.WorkedOut--
				if slot.WorkedOut <= 0 {
					slot.Count--
					return efficiency
				}
			} else {
				slot.WorkedOut += thorium
				if slot.WorkedOut >= 100 {
					slot.Count--
					slot.WorkedOut = 0
				}
			}
		}
	}

	if afterburner {
		efficiency *= 2
	}

	return efficiency
}
