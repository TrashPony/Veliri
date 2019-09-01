package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/remove"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"time"
)

func SquadDamage(user *player.Player, damage int, damageUnit *unit.Unit) {
	// 1 наносим урон корпусу
	damageUnit.HP -= damage

	// наносим урон 2м рандомным эквипам
	for i := 0; i < 2; i++ {
		equipmentSlot := damageUnit.Body.GetRandomEquip()
		if equipmentSlot != nil && equipmentSlot.Equip != nil {
			equipmentSlot.HP -= damage * 3
			if equipmentSlot.HP < 0 {
				equipmentSlot.HP = 0
			}
		}
	}

	// говорим всем урон получаемый юнитом
	go SendMessage(Message{Event: "DamageUnit", IDMap: damageUnit.MapID, ShortUnit: damageUnit.GetShortInfo(), Bot: user.Bot})

	// владельцу полное состояние отряда
	go SendMessage(Message{Event: "FillSquadBlock", IDUserSend: user.GetID(), IDMap: damageUnit.MapID, Squad: user.GetSquad(), Bot: user.Bot})

	// если умер мс то весь отряд умирает
	if damageUnit.Body.MotherShip && damageUnit.HP <= 0 {
		// останавливаем движение, Обязательно! иначае в методе move, приложение упадет на всех возможных проверках
		stopMove(damageUnit, true)
		go SendMessage(Message{Event: "DeadSquad", OtherUser: user.GetShortUserInfo(true), IDMap: damageUnit.MapID})
		// время для проигрыша анимации например
		time.Sleep(2 * time.Second)
		// удаляем отряд из игры
		remove.Squad(user.GetSquad())
		// отнимание всего отряда и инвентаря в трюме
		user.SetSquad(nil)
		// тащим юзера в последнюю посещенную им базу
		IntoToBase(user, user.LastBaseID)
	} else {
		if damageUnit.HP <= 0 {
			// останавливаем движение, Обязательно! иначае в методе move, приложение упадет на всех возможных проверках
			stopMove(damageUnit, true)
			// todo удаляем юнита и обновляем в бд
			go SendMessage(Message{Event: "DeadUnit", IDMap: damageUnit.MapID, Unit: damageUnit})
		}
	}
}
