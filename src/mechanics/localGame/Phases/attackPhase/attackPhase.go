package attackPhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../localGame/Phases/movePhase"
	"../../../localGame/map/watchZone"
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
	AttackUnit  unit.Unit                   `json:"attack_unit"`
	TargetUnits []TargetUnit                `json:"targets_units"` // юниты на которых воздействует действие
	WeaponSlot  *detail.BodyWeaponSlot      `json:"weapon_slot"`   // Чем воздействуем (если оружием то EquipSlot == nil)
	EquipSlot   *detail.BodyEquipSlot       `json:"equip_slot"`    // Чем воздействуем (если снарягой то WeaponSlot == nil)
	Target      *coordinate.Coordinate      `json:"target"`        // куда летит снаряд, действие
	WatchNode   *watchZone.UpdaterWatchZone `json:"watch_node"`    // если юнит переместился, то надо обновить зону выдимости
	Error       string                      `json:"error"`
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
		if item.ActionUnit.HP > 0 {
			if item.WeaponSlot != nil {
				// firearms может пулять куда угодно
				if item.WeaponSlot.Weapon.Type == "firearms" {
					targetCoordinate, ok := game.Map.GetCoordinate(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
					if ok {
						resultBattle = append(resultBattle, InitAttack(item.ActionUnit, targetCoordinate, game))
					}
				}
				// laser и missile только в юнитов
				if item.WeaponSlot.Weapon.Type == "laser" || item.WeaponSlot.Weapon.Type == "missile" {
					_, ok := game.GetUnit(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
					if ok {
						targetCoordinate, _ := game.Map.GetCoordinate(item.ActionUnit.Target.Q, item.ActionUnit.Target.R)
						resultBattle = append(resultBattle, InitAttack(item.ActionUnit, targetCoordinate, game))
					}
				}
			} else {
				if item.EquipSlot.HP > 0 {
					if item.EquipSlot.Equip.Applicable == "my_units" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok && targetUnit.Owner == item.ActionUnit.Owner {
							resultBattle = append(resultBattle, ToUnit(item.ActionUnit, targetUnit, item.EquipSlot))
						}
					}

					if item.EquipSlot.Equip.Applicable == "hostile_units" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok && targetUnit.Owner != item.ActionUnit.Owner {
							resultBattle = append(resultBattle, ToUnit(item.ActionUnit, targetUnit, item.EquipSlot))
						}
					}

					if item.EquipSlot.Equip.Applicable == "all" {
						targetUnit, ok := game.GetUnit(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok {
							resultBattle = append(resultBattle, ToUnit(item.ActionUnit, targetUnit, item.EquipSlot))
						}
					}

					if item.EquipSlot.Equip.Applicable == "myself" {
						resultBattle = append(resultBattle, ToUnit(item.ActionUnit, item.ActionUnit, item.EquipSlot))
					}

					if item.EquipSlot.Equip.Applicable == "myself_move" {
						resultBattle = append(resultBattle, MoveEquip(item.ActionUnit, game, item.EquipSlot))
					}

					if item.EquipSlot.Equip.Applicable == "map" {
						targetCoordinate, ok := game.Map.GetCoordinate(item.EquipSlot.Target.Q, item.EquipSlot.Target.R)
						if ok {
							resultBattle = append(resultBattle, ToMap(item.ActionUnit, targetCoordinate, game, item.EquipSlot))
						}
					}
				} else {
					resultBattle = append(resultBattle, &ResultBattle{Error: "equip breaking"})
				}
			}
		} else {
			resultBattle = append(resultBattle, &ResultBattle{Error: "unit is dead"})
		}
	}

	// завершительная часть боя, проигрых уже наложеных эффектов, сначала отнимаем все статы потом пополняем
	targetsUnit := make([]TargetUnit, 0)

	for _, q := range game.GetUnits() {
		for _, gameUnit := range q {
			if gameUnit.HP > 0 {
				for _, effect := range gameUnit.Effects {
					if effect.Type == "takes_away" && !effect.Used {
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
					if effect.Type == "replenishes" && !effect.Used {
						powEnEffect(effect, gameUnit, &targetsUnit)
					}
				}
			}
		}
	}

	resultBattle = append(resultBattle, &ResultBattle{TargetUnits: targetsUnit})

	return
}
