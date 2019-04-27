package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

// отправляет на фронт все боеприписы которые лежат в инвентаре и подходят к текущему оружию юнита,
// а так же смотрим что бы юнит был в пределах досягаемости
func initAmmoReload(msg Message, client *player.Player) {
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findUnit && findGame && activeGame.Phase == "targeting" && gameUnit.GetWeaponSlot() != nil && gameUnit.GetWeaponSlot().Weapon != nil {
		reloadZone := coordinate.GetCoordinatesRadius(
			&coordinate.Coordinate{
				Q: client.GetSquad().MatherShip.Q,
				R: client.GetSquad().MatherShip.R,
			},
			client.GetSquad().MatherShip.GetWatchZone(),
		)

		find := false
		for _, coor := range reloadZone {
			if coor.Q == msg.Q && coor.R == msg.R {
				find = true
			}
		}

		ammoSlots := make(map[int]*inventory.Slot)

		if find || gameUnit.Body.MotherShip {
			for numberSlot, slot := range client.GetSquad().Inventory.Slots {
				if slot.Type == "ammo" {
					ammo, _ := gameTypes.Ammo.GetByID(slot.ItemID)
					if ammo.Type == gameUnit.GetWeaponSlot().Weapon.Type && ammo.StandardSize == gameUnit.GetWeaponSlot().Weapon.StandardSize {
						ammoSlots[numberSlot] = slot
					}
				}
			}
			SendMessage(Message{Event: msg.Event, AmmoSlots: ammoSlots, Q: msg.Q, R: msg.R}, client.GetID(), activeGame.Id)
		} else {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
		}
	}
}

// перезарядка возможно в дальности обзора мса (вылетает мини пиздюк и перекидывает боезапас юнитом Х) )
// перезарядка возможно только в фазе таргетинга и не дает стрелять, однако использовать эквип можно
func ammoReload(msg Message, client *player.Player) {

	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())
	slot, findSlot := client.GetSquad().Inventory.Slots[msg.Slot]

	if findUnit && findGame && findSlot && activeGame.Phase == "targeting" && gameUnit.GetWeaponSlot() != nil && gameUnit.GetWeaponSlot().Weapon != nil {

		ammo, _ := gameTypes.Ammo.GetByID(slot.ItemID)

		if ammo.Type == gameUnit.GetWeaponSlot().Weapon.Type && ammo.StandardSize == gameUnit.GetWeaponSlot().Weapon.StandardSize {
			// снимаем цель у юнита
			gameUnit.Target = nil
			// делаем экшон на перезарядку в фазе атаки
			gameUnit.Reload = &unit.ReloadAction{AmmoID: ammo.ID, InventorySlot: msg.Slot}

			SendMessage(Message{Event: msg.Event, Accept: true, Q: msg.Q, R: msg.R}, client.GetID(), activeGame.Id)
		}
	} else {
		SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
	}
}
