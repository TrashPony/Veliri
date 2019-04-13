package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/attackPhase"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/hexLineDraw"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
)

type AttackMessage struct {
	Event        string                     `json:"event"`
	UserName     string                     `json:"user_name"`
	GameID       int                        `json:"game_id"`
	ResultBattle []attackPhase.ResultBattle `json:"result_battle"`
}

func attack(activeGame *localGame.Game) {

	resultBattle := attackPhase.AttackPhase(activeGame)

	if activeGame.CheckEndGame() {
		SendAllMessage(Message{Event: "EndGame"}, activeGame)
	}

	for _, gamePlayer := range activeGame.GetPlayers() {
		SendMessage(
			AttackMessage{
				Event:        "AttackPhase",
				UserName:     gamePlayer.GetLogin(),
				GameID:       gamePlayer.GetGameID(),
				ResultBattle: dataPreparation(resultBattle, gamePlayer, activeGame),
			},
			gamePlayer.GetID(),
			activeGame.Id,
		)

		for _, q := range gamePlayer.GetUnits() {
			for _, userUnit := range q {
				// обновляем всю стату юнитов у всех пользователей
				SendMessage(
					Move{
						Event:    "UpdateUnit",
						UserName: gamePlayer.GetLogin(),
						GameID:   activeGame.Id,
						Unit:     userUnit,
					},
					gamePlayer.GetID(),
					activeGame.Id,
				)
			}
		}
	}

	QueueSender(activeGame)
}

func dataPreparation(resultBattle []*attackPhase.ResultBattle, gamePlayer *player.Player, activeGame *localGame.Game) []attackPhase.ResultBattle {
	watchResultBattle := make([]attackPhase.ResultBattle, 0)

	for _, actionBattle := range resultBattle {
		var preloadActionBattle attackPhase.ResultBattle

		_, findAttackUnit := actionBattle.GetUserWatchCoordinate(gamePlayer.GetID(), actionBattle.AttackUnit.Q, actionBattle.AttackUnit.R)
		if findAttackUnit {
			// если игрок видит юнита добавляем все что относиться к атаке
			preloadActionBattle.AttackUnit = actionBattle.AttackUnit
			preloadActionBattle.RotateTower = actionBattle.RotateTower
		}

		_, findTargetCoordinate := actionBattle.GetUserWatchCoordinate(gamePlayer.GetID(), actionBattle.Target.Q, actionBattle.Target.R)
		if findTargetCoordinate {
			preloadActionBattle.Target = actionBattle.Target
		} else {
			preloadActionBattle.Target = coordinate.Coordinate{Type: "hide"}
		}

		// если игрок видит оружие/эквип или их воздействие то отдаем типы оружий и эквипа
		if findAttackUnit || findTargetCoordinate {
			if actionBattle.WeaponSlot.Weapon != nil {
				preloadActionBattle.WeaponSlot.Weapon = actionBattle.WeaponSlot.Weapon
			}
			if actionBattle.EquipSlot.Equip != nil {
				preloadActionBattle.EquipSlot.Equip = actionBattle.EquipSlot.Equip
			}
			if actionBattle.Reload {
				preloadActionBattle.Reload = true
			}
		}

		if !findAttackUnit && findTargetCoordinate {
			// если игрок не видит кто стреляет но видит цель то надо снаряду проложить траекторию полета от первой точки видимости
			unitCoordinate, find := activeGame.Map.GetCoordinate(actionBattle.AttackUnit.Q, actionBattle.AttackUnit.R)
			if find {
				flightPath := hexLineDraw.Draw(unitCoordinate, &actionBattle.Target, activeGame)

				for _, coordinatePath := range flightPath {
					_, findCoordinate := actionBattle.GetUserWatchCoordinate(gamePlayer.GetID(), coordinatePath.Q, coordinatePath.R)
					if findCoordinate {
						break
					}
					preloadActionBattle.StartWatchAttack = *coordinatePath
				}
			}
		}

		if findAttackUnit && !findTargetCoordinate {
			// если игрок не видит куда прилетает снаряд но видит кто пуляет то надо снаряду проложить траекторию полета до последней точки видимости
			unitCoordinate, find := activeGame.Map.GetCoordinate(actionBattle.AttackUnit.Q, actionBattle.AttackUnit.R)
			if find {
				flightPath := hexLineDraw.Draw(unitCoordinate, &actionBattle.Target, activeGame)

				var endNodePath coordinate.Coordinate

				for _, coordinatePath := range flightPath {
					_, findCoordinate := actionBattle.GetUserWatchCoordinate(gamePlayer.GetID(), coordinatePath.Q, coordinatePath.R)
					if !findCoordinate {
						endNodePath = *coordinatePath
					}
					if !findCoordinate {
						break
					}
				}
				preloadActionBattle.EndWatchAttack = endNodePath
			}
		}

		preloadActionBattle.TargetUnits = make([]attackPhase.TargetUnit, 0)

		for _, targetUnit := range actionBattle.TargetUnits {
			_, findTargetUnit := actionBattle.GetUserWatchCoordinate(gamePlayer.GetID(), targetUnit.Unit.Q, targetUnit.Unit.R)
			// если игрок видит юнита по которому стреляют добавляем его
			if findTargetUnit {
				preloadActionBattle.TargetUnits = append(preloadActionBattle.TargetUnits, targetUnit)
			}
		}

		preloadActionBattle.WatchNode = make(map[string]*watchZone.UpdaterWatchZone)

		for key, watch := range actionBattle.WatchNode {
			if strconv.Itoa(gamePlayer.GetID()) == key {
				preloadActionBattle.WatchNode[key] = watch
			}
		}

		watchResultBattle = append(watchResultBattle, preloadActionBattle)
	}

	return watchResultBattle
}
