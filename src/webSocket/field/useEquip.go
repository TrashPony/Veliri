package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/equip"
	"../../mechanics/player"
	"../../mechanics/localGame"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame/useEquip"
	"../../mechanics/gameObjects/detail"
)

func UseEquip(msg Message, ws *websocket.Conn) {

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

	if findUnit && findClient && findGame && !client.GetReady() && ok && equipSlot.Equip != nil {
		if equipSlot.Equip.Applicable == "map" {
			gameCoordinate, findCoordinate := client.GetWatchCoordinate(msg.X, msg.Y)
			if findCoordinate {
				effectCoordinates := useEquip.ToMap(gameUnit, gameCoordinate, activeGame, equipSlot, client)
				if effectCoordinates != nil {
					ws.WriteJSON(SendUseEquip{Event: "UseMapEquip",UseUnit: gameUnit, ZoneEffect: effectCoordinates, AppliedEquip: equipSlot.Equip, XUse: msg.X, YUse: msg.Y})
					updateUseMapEquipHostileUser(msg.X, msg.Y, client, activeGame, effectCoordinates, equipSlot.Equip)
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find coordinate"})
				}
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find coordinate"})
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
				equipPipe <- SendUseEquip{Event: "UseUnitEquip", UserName: user.GetLogin(), GameID: activeGame.Id, ZoneEffect: zoneEffect, AppliedEquip: playerEquip, XUse: xUse, YUse: yUse}
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
	XUse         int                                          `json:"x_use"`
	YUse         int                                          `json:"y_use"`
}
