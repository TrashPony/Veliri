package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame/Phases/targetPhase"
)

func SelectEquip(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
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
		if !client.GetReady() && !gameUnit.UseEquip {
			if equipSlot.Equip.Applicable == "map" {
				ws.WriteJSON(EquipMapCoordinate{Event: "GetEquipMapTargets", Unit: gameUnit,
					EquipSlot: equipSlot, Targets: targetPhase.GetEquipAllTargetZone(gameUnit, equipSlot.Equip, activeGame)})
			}

			if equipSlot.Equip.Applicable == "my_units" {
				ws.WriteJSON(EquipTargetCoordinate{Event: "GetEquipMyUnitTargets", Unit: gameUnit, EquipSlot: equipSlot,
					Units: targetPhase.GetEquipMyUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)})
			}

			if equipSlot.Equip.Applicable == "hostile_units" {
				ws.WriteJSON(EquipTargetCoordinate{Event: "GetEquipHostileUnitTargets", Unit: gameUnit, EquipSlot: equipSlot,
					Units: targetPhase.GetEquipHostileUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)})
			}

			if equipSlot.Equip.Applicable == "myself" {
				ws.WriteJSON(EquipTargetCoordinate{Event: "GetEquipMySelfTarget", Unit: gameUnit, EquipSlot: equipSlot})
			}

			if equipSlot.Equip.Applicable == "all" {
				// todo  и свои и чужие но не карта GetEquipAllUnitTarget
			}
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
		}
	}
}

type EquipTargetCoordinate struct {
	Event     string                `json:"event"`
	Unit      *unit.Unit            `json:"unit"`
	Units     []*unit.Unit          `json:"units"`
	EquipSlot *detail.BodyEquipSlot `json:"equip_slot"`
}

type EquipMapCoordinate struct {
	Event     string                                       `json:"event"`
	Unit      *unit.Unit                                   `json:"unit"`
	EquipSlot *detail.BodyEquipSlot                        `json:"equip_slot"`
	Targets   map[string]map[string]*coordinate.Coordinate `json:"targets"`
}
