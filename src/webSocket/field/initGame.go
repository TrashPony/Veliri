package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func toGame(msg Message, ws *websocket.Conn) {
	loadGame, ok := Games[msg.IdGame]
	newClient, _ := usersFieldWs[ws]

	if !ok {
		loadGame = game.InitGame(msg.IdGame)
		Games[loadGame.GetStat().Id] = loadGame // добавляем новую игру в карту активных игор
	}

	for _, player := range loadGame.GetPlayers() {
		if (newClient.GetLogin() == player.GetLogin()) && (newClient.GetID() == player.GetID()) {
			newClient = player
		}
	}

	var sendLoadGame = LoadGame{
		Event:              "LoadGame",
		UserName:           newClient.GetLogin(),
		Ready:              newClient.GetReady(),
		Equip:              newClient.GetEquip(),
		Units:              newClient.GetUnits(),
		HostileUnits:       newClient.GetHostileUnits(),
		UnitStorage:        newClient.GetUnitsStorage(),
		Map:                loadGame.GetMap(),
		GameInfo:           loadGame.GetStat(),
		MatherShip:         newClient.GetMatherShip(),
		HostileMatherShips: newClient.GetHostileMatherShips(),
		Watch:              newClient.GetWatchCoordinates()}
	ws.WriteJSON(sendLoadGame)
}
