package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

// TODO отключить игрока если он отключился
func RecoveryPowerWorker(user *player.Player) {

	user.GetSquad().RecoveryPowerWork = true
	defer func() {
		user.GetSquad().RecoveryPowerWork = false
	}()

	for {

		if user.GetSquad().RecoveryPowerExit {
			user.GetSquad().RecoveryPowerExit = false
			return
		}

		update := false
		if user != nil && user.GetSquad() != nil && user.GetSquad().MatherShip != nil {
			if user.GetSquad().MatherShip.Power < user.GetSquad().MatherShip.MaxPower {

				update = true
				// заправляем реактор машина за счет тория в реакторе, с указаной эффективностью
				efficiency := user.GetSquad().MatherShip.GetReactorEfficiency()
				user.GetSquad().MatherShip.WorkOutPower(float32(user.GetSquad().MatherShip.RecoveryPower / 100))

				if user.GetSquad().MatherShip.Power+int(float64(user.GetSquad().MatherShip.RecoveryPower)*(float64(efficiency)/100)) >= user.GetSquad().MatherShip.MaxPower {
					user.GetSquad().MatherShip.Power = user.GetSquad().MatherShip.MaxPower
				} else {
					user.GetSquad().MatherShip.Power += int(float64(user.GetSquad().MatherShip.RecoveryPower) * (float64(efficiency) / 100))
				}
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.Power < unitSlot.Unit.MaxPower && !unitSlot.Unit.OnMap {

					update = true
					// если у какогото юнита который сейчас находится в трюме нехватает энергии заправляем ее за счет тория МС
					efficiency := user.GetSquad().MatherShip.GetReactorEfficiency()
					user.GetSquad().MatherShip.WorkOutPower(float32(unitSlot.Unit.RecoveryPower / 100))

					if unitSlot.Unit.Power+int(float64(unitSlot.Unit.RecoveryPower)*(float64(efficiency)/100)) >= unitSlot.Unit.MaxPower {
						unitSlot.Unit.Power = unitSlot.Unit.MaxPower
					} else {
						unitSlot.Unit.Power += int(float64(unitSlot.Unit.RecoveryPower) * (float64(efficiency) / 100))
					}
				}
			}
		}

		if update {
			go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), Unit: user.GetSquad().MatherShip,
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
			go SendMessage(Message{Event: "FillSquadBlock", IDUserSend: user.ID, IDMap: user.GetSquad().MatherShip.MapID, Squad: user.GetSquad(), Bot: user.Bot})
		}

		time.Sleep(time.Second)
	}
}
