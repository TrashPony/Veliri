package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func loadGame(msg Message, ws *websocket.Conn) {
	loadGame, ok := Games[msg.IdGame]
	newClient, _ := usersFieldWs[ws]

	if !ok {
		loadGame = game.InitGame(msg.IdGame)
		Games[loadGame.GetStat().Id] = loadGame // добавляем новую игру в карту активных игор
	}

	for _, player := range loadGame.GetPlayers() {
		if (newClient.GetLogin() == player.GetLogin()) && (newClient.GetID() == player.GetID()) {
			usersFieldWs[ws] = player
		}
	}

	var sendLoadGame = LoadGame{
		Event:              "LoadGame",
		UserName:           usersFieldWs[ws].GetLogin(),
		Ready:              usersFieldWs[ws].GetReady(),
		Equip:              usersFieldWs[ws].GetEquip(),
		Units:              usersFieldWs[ws].GetUnits(),
		HostileUnits:       usersFieldWs[ws].GetHostileUnits(),
		UnitStorage:        usersFieldWs[ws].GetUnitsStorage(),
		Map:                loadGame.GetMap(),
		GameInfo:           loadGame.GetStat(),
		MatherShip:         usersFieldWs[ws].GetMatherShip(),
		HostileMatherShips: usersFieldWs[ws].GetHostileMatherShips(),
		Watch:              usersFieldWs[ws].GetWatchCoordinates()}
	ws.WriteJSON(sendLoadGame)
}

type LoadGame struct {
	Event              string                                 `json:"event"`
	UserName           string                                 `json:"user_name"`
	Ready              bool                                   `json:"ready"`
	Equip              []*game.Equip                          `json:"equip"`
	Units              map[string]map[string]*game.Unit       `json:"units"`
	HostileUnits       map[string]map[string]*game.Unit       `json:"hostile_units"`
	UnitStorage        []*game.Unit                           `json:"unit_storage"`
	Map                *game.Map                              `json:"map"`
	GameInfo           *game.InfoGame                         `json:"game_info"`
	MatherShip         *game.MatherShip                       `json:"mather_ship"`
	HostileMatherShips map[string]map[string]*game.MatherShip `json:"hostile_mather_ships"`
	Watch              map[string]map[string]*game.Coordinate `json:"watch"`
}
