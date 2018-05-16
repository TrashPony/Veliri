package watchZone

import (
	"../player"
	"../matherShip"
	"strconv"
)

func updateHostileMatherShip(client *player.Player, oldWatchHostileMatherShip map[string]map[string]*matherShip.MatherShip) (openMatherShip []*matherShip.MatherShip, closeMatherShip []*matherShip.MatherShip) {
	for _, xLine := range client.GetHostileMatherShips() { // добавляем новые вражеские структуры которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchHostileMatherShip[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)]
			if !ok {
				openMatherShip = append(openMatherShip, hostile)
			}
		}
	}

	for _, xLine := range oldWatchHostileMatherShip {
		for _, hostile := range xLine {
			_, find := client.GetHostileMatherShip(hostile.X, hostile.Y)
			if !find {
				closeMatherShip = append(closeMatherShip, hostile)
			}
		}
	}
	return
}
