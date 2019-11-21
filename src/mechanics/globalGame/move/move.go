package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
)

func CheckExit(user *player.Player, moveUnit *unit.Unit, path *[]*unit.PathUnit, i, countExit *int, followTarget bool) bool {
	// юнит или отряд умер
	if user.GetSquad() == nil || moveUnit == nil || !moveUnit.OnMap || moveUnit.HP <= 0 {
		return true
	}

	if moveUnit.ActualPath == nil || moveUnit.ActualPath != path {
		// если актуальный путь сменился то выполняем еще 1 итерацию из старого пути дабы дать время сгенерить новый путь
		if *countExit <= 0 {
			return true
		}
		*countExit--

		if moveUnit.LastPathCell == nil && len(*path)-1 >= *i+*countExit {
			moveUnit.LastPathCell = (*path)[*i+*countExit]
		} else {
			if len(*path)-1 < *i+*countExit {
				moveUnit.LastPathCell = (*path)[len(*path)-1]
			}
		}
	}

	// выходим если это атакующее движение при условиях:
	if followTarget {

		// если цель изменилась или юнит больше не преследует
		target := moveUnit.GetTarget()
		if target == nil || !target.Follow {
			return true
		}

		////или цель в зоне поражения
		//if CheckFireToTarget(moveUnit, mp, target) {
		//	return
		//}
	}

	// если клиент отключился то останавливаем его
	if globalGame.Clients.GetById(user.GetID()) == nil {
		return true
	}

	return false
}
