package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
)

func SelectEquip(msg Message, client *player.Player) {

	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	ok := false
	equipSlot := &detail.BodyEquipSlot{}

	if msg.EquipType == 3 {
		equipSlot, ok = gameUnit.Body.EquippingIII[msg.NumberSlot]
	}

	if msg.EquipType == 2 {
		equipSlot, ok = gameUnit.Body.EquippingII[msg.NumberSlot]
	}

	if findUnit && findGame && ok && equipSlot.Equip != nil {
		if !client.GetReady() {

			if equipSlot.Equip.Applicable == "map" {
				SendMessage(
					EquipMapCoordinate{
						Event:     "GetEquipMapTargets",
						Unit:      gameUnit,
						EquipSlot: equipSlot,
						Targets:   targetPhase.GetEquipAllTargetZone(gameUnit, equipSlot.Equip, activeGame, client),
					},
					client.GetID(),
					activeGame.Id,
				)
			}

			if equipSlot.Equip.Applicable == "my_units" {
				SendMessage(
					EquipTargetCoordinate{
						Event:     "GetEquipMyUnitTargets",
						Unit:      gameUnit,
						EquipSlot: equipSlot,
						Units:     targetPhase.GetEquipMyUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client),
					},
					client.GetID(),
					activeGame.Id,
				)
			}

			if equipSlot.Equip.Applicable == "hostile_units" {
				SendMessage(
					EquipTargetCoordinate{
						Event:     "GetEquipHostileUnitTargets",
						Unit:      gameUnit,
						EquipSlot: equipSlot,
						Units:     targetPhase.GetEquipHostileUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client),
					},
					client.GetID(),
					activeGame.Id,
				)
			}

			if equipSlot.Equip.Applicable == "myself" {
				SendMessage(
					EquipTargetCoordinate{
						Event:     "GetEquipMySelfTarget",
						Unit:      gameUnit,
						EquipSlot: equipSlot,
					},
					client.GetID(),
					activeGame.Id,
				)
			}

			if equipSlot.Equip.Applicable == "all" {
				// todo и свои и чужие но не карта GetEquipAllUnitTarget
			}

			if equipSlot.Equip.Applicable == "mining" {
				// todo выбрана копающая хуйня для добычи ресурсов
			}

		} else {
			SendMessage(ErrorMessage{Event: "Error", Error: "you ready"}, client.GetID(), activeGame.Id)
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
