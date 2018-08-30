package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/localGame/Phases/targetPhase"
	"../../mechanics/localGame/useEquip"
	"../../mechanics/gameObjects/equip"
	"../../mechanics/player"
	"../../mechanics/localGame"
)

func UseUnitEquip(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games.Get(client.GetGameID())

	ok := false
	equipSlot := &detail.BodyEquipSlot{}

	if msg.EquipType == 3 {
		equipSlot, ok = gameUnit.Body.EquippingIII[msg.NumberSlot]
	}

	if msg.EquipType == 2 {
		equipSlot, ok = gameUnit.Body.EquippingII[msg.NumberSlot]
	}

	if findClient && findUnit && findGame && ok && equipSlot.Equip != nil {
		if !client.GetReady() && !gameUnit.UseEquip && !equipSlot.Used {

			var targetUnits []*unit.Unit

			if equipSlot.Equip.Applicable == "my_units" {
				targetUnits = targetPhase.GetEquipMyUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)
			}

			if equipSlot.Equip.Applicable == "hostile_units" {
				targetUnits = targetPhase.GetEquipHostileUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)
			}

			if equipSlot.Equip.Applicable == "myself" {
				// todo на себя
			}

			if equipSlot.Equip.Applicable == "all" {
				// todo  и свои и чужие но не карта GetEquipAllUnitTarget
			}

			for _, targetUnit := range targetUnits {
				if targetUnit.X == msg.ToX && targetUnit.Y == msg.ToY {
					err := useEquip.ToUnit(gameUnit, targetUnit, equipSlot, client)
					if err != nil {
						ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
					} else {
						ws.WriteJSON(SendUseEquip{Event: "UseUnitEquip", UseUnit:gameUnit, ToUnit: targetUnit, AppliedEquip: equipSlot.Equip})
						updateUseUnitEquipHostileUser(client, activeGame, gameUnit, targetUnit, equipSlot.Equip)
					}
				}
			}
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
		}
	}
}

func updateUseUnitEquipHostileUser(client *player.Player, activeGame *localGame.Game, gameUnit, targetUnit *unit.Unit, playerEquip *equip.Equip) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(targetUnit.X, targetUnit.Y)
			if watch {
				equipPipe <- SendUseEquip{Event: "UseUnitEquip", UserName: user.GetLogin(), GameID: activeGame.Id, UseUnit:gameUnit, ToUnit: targetUnit, AppliedEquip: playerEquip}
			}
		}
	}
}