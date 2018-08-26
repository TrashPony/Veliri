package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/equip"
	"../../mechanics/player"
	"../../mechanics/localGame"
	"../../mechanics/gameObjects/coordinate"
	"fmt"
)

func UseEquip(msg Message, ws *websocket.Conn) {
	fmt.Printf("%+v\n", msg)

	/*client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())
	playerEquip, findEquip := client.GetEquipByID(msg.EquipID)

	if findClient && findGame && !client.GetReady() && findEquip && !playerEquip.Used && (activeGame.Phase == "move" || activeGame.Phase == "targeting") {
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
	}*/
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
	Unit         *unit.Unit                                   `json:"unit"`
	AppliedEquip *equip.Equip                                 `json:"applied_equip"`
	ZoneEffect   map[string]map[string]*coordinate.Coordinate `json:"zone_effect"`
	XUse         int                                          `json:"x_use"`
	YUse         int                                          `json:"y_use"`
}
