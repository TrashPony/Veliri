package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/attackPhase"
	"../../mechanics/localGame/map/hexLineDraw"
	"../../mechanics/localGame/map/watchZone"
	"../../mechanics/player"
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

	for _, gamePlayer := range activeGame.GetPlayers() {
		attack := AttackMessage{Event: "AttackPhase", UserName: gamePlayer.GetLogin(), GameID: gamePlayer.GetGameID(),
			ResultBattle: dataPreparation(resultBattle, gamePlayer, activeGame)}
		attackPipe <- attack

		for _, q := range gamePlayer.GetUnits() {
			for _, userUnit := range q {
				// обновляем всю стату юнитов у всех пользователей
				moves := Move{Event: "UpdateUnit", UserName: gamePlayer.GetLogin(), GameID: activeGame.Id, Unit: userUnit}
				move <- moves
			}
		}
	}

	QueueSender(activeGame)
}

func dataPreparation(resultBattle []*attackPhase.ResultBattle, gamePlayer *player.Player, activeGame *localGame.Game) []attackPhase.ResultBattle {
	watchResultBattle := make([]attackPhase.ResultBattle, 0)
	// TODO на момент отдачи данных пользователи могут имеють другой обзор
	for _, actionBattle := range resultBattle {
		var preloadActionBattle attackPhase.ResultBattle

		_, findAttackUnit := gamePlayer.GetWatchCoordinate(actionBattle.AttackUnit.Q, actionBattle.AttackUnit.R)
		if findAttackUnit {
			// если игрок видит юнита добавляем все что относиться к атаке
			preloadActionBattle.AttackUnit = actionBattle.AttackUnit
			preloadActionBattle.RotateTower = actionBattle.RotateTower
		}

		_, findTargetCoordinate := gamePlayer.GetWatchCoordinate(actionBattle.Target.Q, actionBattle.Target.R)
		if findTargetCoordinate {
			preloadActionBattle.Target = actionBattle.Target
		}

		// если игрок видит оружие/эквип или их воздействие то отдаем типы оружий и эквипа
		if findAttackUnit || findTargetCoordinate {
			if actionBattle.WeaponSlot.Weapon != nil {
				preloadActionBattle.WeaponSlot.Weapon = actionBattle.WeaponSlot.Weapon
			}
			if actionBattle.EquipSlot.Equip != nil {
				preloadActionBattle.EquipSlot.Equip = actionBattle.EquipSlot.Equip
			}
		}

		if !findAttackUnit && findTargetCoordinate {
			// если игрок не видит кто стреляет но видит цель то надо снаряду проложить траекторию полета от первой точки видимости
			unitCoordinate, find := activeGame.Map.GetCoordinate(actionBattle.AttackUnit.Q, actionBattle.AttackUnit.R)
			if find {
				flightPath := hexLineDraw.Draw(unitCoordinate, &actionBattle.Target, activeGame)

				for _, coordinatePath := range flightPath {
					firstNodePath, findCoordinate := gamePlayer.GetWatchCoordinate(coordinatePath.Q, coordinatePath.R)
					if findCoordinate {
						preloadActionBattle.StartWatchAttack = *firstNodePath
						break
					}
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
					firstNodePath, findCoordinate := gamePlayer.GetWatchCoordinate(coordinatePath.Q, coordinatePath.R)
					if findCoordinate {
						endNodePath = *firstNodePath
					} else {
						break
					}
				}
				preloadActionBattle.EndWatchAttack = endNodePath
			}
		}

		preloadActionBattle.TargetUnits = make([]attackPhase.TargetUnit, 0)

		for _, targetUnit := range actionBattle.TargetUnits {
			_, findTargetUnit := gamePlayer.GetWatchCoordinate(targetUnit.Unit.Q, targetUnit.Unit.R)
			// если игрок видит юнита добавляем его
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
