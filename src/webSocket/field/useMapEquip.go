package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/gameObjects/equip"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/useEquip"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
)

func UseEquip(msg Message, ws *websocket.Conn) {

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

	if findUnit && findClient && findGame && !client.GetReady() && ok && equipSlot.Equip != nil {
		if equipSlot.Equip.Applicable == "map" {
			gameCoordinate, findCoordinate := activeGame.Map.GetCoordinate(msg.TargetQ, msg.TargetR)
			if findCoordinate {
				effectCoordinates, err := useEquip.ToMap(gameUnit, gameCoordinate, activeGame, equipSlot, client)
				if err == nil {
					ws.WriteJSON(SendUseEquip{Event: "UseMapEquip", UseUnit: gameUnit, ZoneEffect: effectCoordinates, AppliedEquip: equipSlot.Equip, QUse: msg.Q, RUse: msg.R})
					updateUseMapEquipHostileUser(msg.Q, msg.R, client, activeGame, effectCoordinates, equipSlot.Equip)
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: err.Error()})
				}
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find game coordinate"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
	}
}

func updateUseMapEquipHostileUser(xUse, yUse int, client *player.Player, activeGame *localGame.Game, zoneEffect map[string]map[string]*coordinate.Coordinate, playerEquip *equip.Equip) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(xUse, yUse)
			if watch {
				equipPipe <- SendUseEquip{Event: "UseUnitEquip", UserName: user.GetLogin(), GameID: activeGame.Id, ZoneEffect: zoneEffect, AppliedEquip: playerEquip, QUse: xUse, RUse: yUse}
			}
		}
	}
}

type SendUseEquip struct {
	Event        string                                       `json:"event"`
	UserName     string                                       `json:"user_name"`
	GameID       int                                          `json:"game_id"`
	UseUnit      *unit.Unit                                   `json:"use_unit"`
	ToUnit       *unit.Unit                                   `json:"to_unit"`
	AppliedEquip *equip.Equip                                 `json:"applied_equip"`
	ZoneEffect   map[string]map[string]*coordinate.Coordinate `json:"zone_effect"`
	QUse         int                                          `json:"q_use"`
	RUse         int                                          `json:"r_use"`
}
