package watchZone

import (
	"../../../player"
	"../../../gameObjects/unit"
	"strconv"
)

func updateHostileUnit(client *player.Player, oldWatchUnit map[string]map[string]*unit.Unit) (openUnit []*unit.Unit, closeUnit []*unit.Unit) {
	for _, xLine := range client.GetHostileUnits() { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)]
			if !ok {
				openUnit = append(openUnit, hostile)
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.GetHostileUnit(hostile.X, hostile.Y)
			if !find {
				closeUnit = append(closeUnit, hostile)
			}
		}
	}

	return
}
