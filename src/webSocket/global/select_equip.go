package global

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"

func SelectEquip(user *player.Player, msg Message) {
	if user.GetSquad() != nil {
		selectUnit := user.GetSquad().GetUnitByID(msg.UnitID)
		if selectUnit != nil {
			equip := selectUnit.Body.GetEquip(msg.TypeSlot, msg.Slot)
			if equip != nil && equip.Equip != nil {

				if equip.Equip.Applicable == "digger" {
					go SendMessage(Message{Event: "InitDigger", IDUserSend: user.GetID(), TypeSlot: msg.TypeSlot,
						Slot: msg.Slot, IDMap: selectUnit.MapID, ShortUnit: selectUnit.GetShortInfo(), Equip: equip})
				}

				if equip.Equip.Applicable == "ore" {
					go SendMessage(Message{Event: "InitMiningOre", IDUserSend: user.GetID(), TypeSlot: msg.TypeSlot,
						Slot: msg.Slot, IDMap: selectUnit.MapID, ShortUnit: selectUnit.GetShortInfo(), Equip: equip})
				}

			}
		}
	}
}
