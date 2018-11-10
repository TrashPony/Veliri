package attackPhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../localGame/Phases/movePhase"
	"../../../localGame/map/watchZone"
	"strconv"
)

func AttackPhase(game *localGame.Game) (resultBattle []*ResultBattle) {

	// формирует очередь боя
	sortItems := createQueueAttack(game.GetUnits())

	// отыгрываем бой
	resultBattle = attack(sortItems, game)

	// востаналиываем энерги, даем актив поинты и снимаем флаги использованого снаряжения, снимаем цели юнитов и эквипа
	recovery(game)

	// находим кто будет ходить первым
	movePhase.QueueMove(game)

	for _, player := range game.GetPlayers() {
		updateSquad.Squad(player.GetSquad()) // вносим все изменениея в базу данных
	}

	return
}

type ResultBattle struct {
	AttackUnit       unit.Unit                                               `json:"attack_unit"`
	RotateTower      int                                                     `json:"rotate_tower"`       // на сколько надо повернуть орудие
	TargetUnits      []TargetUnit                                            `json:"targets_units"`      // юниты на которых воздействует действие
	WeaponSlot       detail.BodyWeaponSlot                                   `json:"weapon_slot"`        // Чем воздействуем (если оружием то EquipSlot == nil)
	EquipSlot        detail.BodyEquipSlot                                    `json:"equip_slot"`         // Чем воздействуем (если снарягой то WeaponSlot == nil)
	StartWatchAttack coordinate.Coordinate                                   `json:"start_watch_attack"` // Первая ячейка откуда игрок вижит летящий снаряд, расчитывается перед отправкой
	EndWatchAttack   coordinate.Coordinate                                   `json:"end_watch_attack"`   // последняя ячейка где игрок вижит летящий снаряд, расчитывается перед отправкой
	Target           coordinate.Coordinate                                   `json:"target"`             // куда летит снаряд, действие
	watchPlayer      map[string]map[string]map[string]*coordinate.Coordinate // видимые координаты пользователем на данный шаг, нужно при отправке данных
	WatchNode        map[string]*watchZone.UpdaterWatchZone                  `json:"watch_node"` // расчет видимости на каждый экшен для каждого пользователя [user_ID]watch
	Error            string                                                  `json:"error"`
}

func (result *ResultBattle) GetUserWatchCoordinate(id, q, r int) (*coordinate.Coordinate, bool) {
	gameCoordinate, find := result.watchPlayer[strconv.Itoa(id)][strconv.Itoa(q)][strconv.Itoa(r)]
	return gameCoordinate, find
}

type TargetUnit struct {
	Unit          unit.Unit `json:"unit"`           // Юнит на который воздействует
	Damage        int       `json:"damage"`         // если юнит получает урон то сколько
	Heal          int       `json:"heal"`           // если юнит получает хил то сколько
	Power         int       `json:"Power"`          // отнимание или прибавление энергии
	BreakingEquip bool      `json:"breaking_equip"` // если сломался хотя бы 1 эквип говорить об этом клиенту
}

func attack(sortItems []*QueueAttack, game *localGame.Game) (resultBattle []*ResultBattle) {
	resultBattle = make([]*ResultBattle, 0)

	for _, item := range sortItems {

		var resultAction *ResultBattle

		if item.ActionUnit.HP > 0 {
			if item.WeaponSlot != nil {
				// firearms может пулять куда угодно
				if item.WeaponSlot.Weapon.Type == "firearms" {
					targetCoordinate, ok := game.Map.GetCoordinate(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
					if ok {
						resultAction = InitAttack(item.ActionUnit, targetCoordinate, game)
					}
				}
				// laser и missile только в юнитов
				if item.WeaponSlot.Weapon.Type == "laser" || item.WeaponSlot.Weapon.Type == "missile" {
					targetUnit, ok := game.GetUnit(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
					if ok && targetUnit.HP > 0 {
						targetCoordinate, _ := game.Map.GetCoordinate(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
						resultAction = InitAttack(item.ActionUnit, targetCoordinate, game)
					}
				}
			} else {
				if item.EquipSlot.HP > 0 {
					if item.EquipSlot.Equip.Applicable == "my_units" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok && targetUnit.Owner == item.ActionUnit.Owner {
							resultAction = ToUnit(item.ActionUnit, targetUnit, item.EquipSlot)
						}
					}

					if item.EquipSlot.Equip.Applicable == "hostile_units" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok && targetUnit.Owner != item.ActionUnit.Owner {
							resultAction = ToUnit(item.ActionUnit, targetUnit, item.EquipSlot)
						}
					}

					if item.EquipSlot.Equip.Applicable == "all" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok {
							resultAction = ToUnit(item.ActionUnit, targetUnit, item.EquipSlot)
						}
					}

					if item.EquipSlot.Equip.Applicable == "myself" {
						resultAction = ToUnit(item.ActionUnit, item.ActionUnit, item.EquipSlot)
					}

					if item.EquipSlot.Equip.Applicable == "myself_move" {
						resultAction = MoveEquip(item.ActionUnit, game, item.EquipSlot)
					}

					if item.EquipSlot.Equip.Applicable == "map" {
						targetCoordinate, ok := game.Map.GetCoordinate(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok {
							resultAction = ToMap(item.ActionUnit, targetCoordinate, game, item.EquipSlot)
						}
					}
				} else {
					resultAction = &ResultBattle{Error: "equip breaking"}
				}
			}
		} else {
			resultAction = &ResultBattle{Error: "unit is dead"}
		}

		if resultAction == nil {
			continue
		}

		resultAction.watchPlayer = make(map[string]map[string]map[string]*coordinate.Coordinate)
		resultAction.WatchNode = make(map[string]*watchZone.UpdaterWatchZone)

		for _, gameUser := range game.GetPlayers() {

			resultAction.watchPlayer[strconv.Itoa(gameUser.GetID())] = gameUser.GetWatchCoordinates()
			// расчет видимости на каждый экшен для каждого пользователя [user_ID]watch
			resultAction.WatchNode[strconv.Itoa(gameUser.GetID())] = watchZone.UpdateWatchZone(game, gameUser)
		}

		resultBattle = append(resultBattle, resultAction) // добавляем результат экшена
	}

	// завершительная часть боя, проигрых уже наложеных эффектов, сначала отнимаем все статы потом пополняем
	targetsUnit := make([]TargetUnit, 0)

	for _, q := range game.GetUnits() {
		for _, gameUnit := range q {
			if gameUnit.HP > 0 {
				for _, effect := range gameUnit.Effects {
					if effect != nil && effect.Type == "takes_away" && !effect.Used {
						powEnEffect(effect, gameUnit, &targetsUnit)
					}
				}
			}
		}
	}

	for _, q := range game.GetUnits() {
		for _, gameUnit := range q {
			for _, effect := range gameUnit.Effects {
				if gameUnit.HP > 0 {
					if effect != nil && effect.Type == "replenishes" && !effect.Used {
						powEnEffect(effect, gameUnit, &targetsUnit)
					}
				}
			}
		}
	}

	resultBattle = append(resultBattle, &ResultBattle{TargetUnits: targetsUnit})

	return
}
