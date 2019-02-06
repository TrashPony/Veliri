package watchZone

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
)

func updateHostileUnit(client *player.Player, oldWatchUnit map[string]map[string]*unit.Unit) (openUnit []*unit.Unit, closeUnit []*unit.Unit) {
	for _, xLine := range client.GetHostileUnits() { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[strconv.Itoa(hostile.Q)][strconv.Itoa(hostile.R)]
			if !ok {
				openUnit = append(openUnit, hostile)
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.GetHostileUnit(hostile.Q, hostile.R)
			if !find {
				closeUnit = append(closeUnit, hostile)
			}
		}
	}

	return
}
