package globalGame

import "../gameObjects/detail"

func WorkOutThorium(thoriumSlots map[int]*detail.ThoriumSlot, afterburner bool) float64 {
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

			if afterburner { // если активирован форсаж то топливо тратиться х5
				thorium = thorium * 15
			}

			slot.WorkedOut += thorium

			if slot.WorkedOut >= 100 {
				slot.Count--
				slot.WorkedOut = 0
			}
		}
	}

	return efficiency
}


