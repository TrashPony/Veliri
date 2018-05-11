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

	var mapParam = LoadGame{
		Event:        "LoadGame",
		Ready:        newClient.GetReady(),
		Equip:        newClient.GetEquip(),
		Units:        newClient.GetUnits(),
		NotGameUnits: newClient.GetNotGameUnits(),
		Map:          loadGame.GetMap(),
		GameInfo:     loadGame.GetStat(),
		MatherShip:   newClient.GetMatherShip()}
	ws.WriteJSON(mapParam)
}
