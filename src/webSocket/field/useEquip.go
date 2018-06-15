package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/useEquip"
	"../../mechanics/unit"
	"../../mechanics/equip"
	"../../mechanics/player"
	"../../mechanics/game"
	"../../mechanics/coordinate"
)

func UseEquip(msg Message, ws *websocket.Conn) {
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games[client.GetGameID()]
	playerEquip, findEquip := client.GetEquipByID(msg.EquipID)

	if findClient && findGame && !client.GetReady() && findEquip && !playerEquip.Used {
		if playerEquip.Applicable == "map" {
			gameCoordinate, findCoordinate := client.GetWatchCoordinate(msg.X, msg.Y)
			if findCoordinate {
				effectCoordinates := useEquip.ToMap(gameCoordinate, activeGame, playerEquip, client)
				ws.WriteJSON(SendUseEquip{Event: "UseMapEquip", ZoneEffect: effectCoordinates, AppliedEquip: playerEquip, XUse: msg.X, YUse: msg.Y})
				updateUseMapEquipHostileUser(msg.X, msg.Y, client, activeGame, effectCoordinates, playerEquip)
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find coordinate"})
			}
		} else {

			gameUnit := EquipApplicable(playerEquip, client, msg.X, msg.Y)

			if gameUnit != nil {
				useEquip.ToUnit(gameUnit, playerEquip, client)
				ws.WriteJSON(SendUseEquip{Event: "UseUnitEquip", Unit: gameUnit, AppliedEquip: playerEquip})
				updateUseUnitEquipHostileUser(client, activeGame, gameUnit, playerEquip)
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find unit"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
	}
}

func EquipApplicable(playerEquip *equip.Equip, client *player.Player, x, y int) *unit.Unit {
	if playerEquip.Applicable == "my_units" {
		gameUnit, findUnit := client.GetUnit(x, y)
		if findUnit {
			return gameUnit
		}
	}

	if playerEquip.Applicable == "hostile_units" {
		gameUnit, findUnit := client.GetHostileUnit(x, y)
		if findUnit {
			return gameUnit
		}
	}

	if playerEquip.Applicable == "all" {
		gameUnit, findUnit := client.GetUnit(x, y)
		if findUnit {
			return gameUnit
		} else {
			gameUnit, findUnit := client.GetHostileUnit(x, y)
			if findUnit {
				return gameUnit
			}
		}
	}

	return nil
}

func updateUseUnitEquipHostileUser(client *player.Player, activeGame *game.Game, gameUnit *unit.Unit, playerEquip *equip.Equip) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(gameUnit.X, gameUnit.Y)
			if watch {
				equipPipe <- SendUseEquip{Event: "UseUnitEquip", UserName: user.GetLogin(), GameID: activeGame.Id, Unit: gameUnit, AppliedEquip: playerEquip}
			}
		}
	}
}

func updateUseMapEquipHostileUser(xUse, yUse int, client *player.Player, activeGame *game.Game, zoneEffect map[string]map[string]*coordinate.Coordinate, playerEquip *equip.Equip) {
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
	Unit         *unit.Unit                                   `json:"unit"`
	AppliedEquip *equip.Equip                                 `json:"applied_equip"`
	ZoneEffect   map[string]map[string]*coordinate.Coordinate `json:"zone_effect"`
	XUse         int                                          `json:"x_use"`
	YUse         int                                          `json:"y_use"`
}
