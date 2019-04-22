package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/remove"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func SquadDamage(user *player.Player, damage int, ws *websocket.Conn) {
	// 1 наносим урон корпусу
	user.GetSquad().MatherShip.HP -= damage

	// наносим урон 2м рандомным эквипам
	for i := 0; i < 2; i++ {
		equipmentSlot := user.GetSquad().MatherShip.Body.GetRandomEquip()
		if equipmentSlot != nil && equipmentSlot.Equip != nil {
			equipmentSlot.HP -= damage * 3
			if equipmentSlot.HP < 0 {
				equipmentSlot.HP = 0
			}
		}
	}

	go SendMessage(Message{Event: "DamageSquad", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot, Squad: user.GetSquad()})
	if user.GetSquad().MatherShip.HP <= 0 {
		// останавливаем движение, Обязательно! иначае в методе move, приложение упадет на всех возможных проверках
		stopMove(user, true)
		go SendMessage(Message{Event: "DeadSquad", OtherUser: user.GetShortUserInfo(true), IDMap: user.GetSquad().MapID})
		// время для проигрыша анимации например
		time.Sleep(2 * time.Second)
		// удаляем отряд из игры
		remove.Squad(user.GetSquad())
		// отнимание всего отряда и инвентаря в трюме
		user.SetSquad(nil)
		// тащим юзера в последнюю посещенную им базу
		IntoToBase(user, user.LastBaseID, ws)
	}
}
