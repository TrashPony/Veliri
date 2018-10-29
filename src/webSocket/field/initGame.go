package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/equip"
	"../../mechanics/gameObjects/map"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame/initGame"
	"github.com/gorilla/websocket"
)

func loadGame(msg Message, ws *websocket.Conn) {
	loadGame, ok := Games.Get(msg.IdGame)
	newClient, _ := usersFieldWs[ws]

	if !ok {
		loadGame = initGame.InitGame(msg.IdGame)
		Games.Add(loadGame.Id, loadGame) // добавляем новую игру в карту активных игор
	}

	player := loadGame.GetPlayer(newClient.GetID(), newClient.GetLogin())

	if player != nil {

		usersFieldWs[ws] = player

		var sendLoadGame = LoadGame{
			Event:        "LoadGame",
			UserName:     usersFieldWs[ws].GetLogin(),
			Ready:        usersFieldWs[ws].GetReady(),
			Units:        usersFieldWs[ws].GetUnits(),
			HostileUnits: usersFieldWs[ws].GetHostileUnits(),
			UnitStorage:  usersFieldWs[ws].GetUnitsStorage(),
			Map:          loadGame.GetMap(),
			Watch:        usersFieldWs[ws].GetWatchCoordinates(),
			GameStep:     loadGame.GetStep(),
			GamePhase:    loadGame.GetPhase()}
		ws.WriteJSON(sendLoadGame)

		if loadGame.Phase == "move" {
			ws.WriteJSON(Move{Event: "QueueMove", UserName: player.GetLogin(), GameID: loadGame.Id, Move: player.Move})
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "error"})
	}
}

type LoadGame struct {
	Event        string                                       `json:"event"`
	UserName     string                                       `json:"user_name"`
	Ready        bool                                         `json:"ready"`
	Equip        []*equip.Equip                               `json:"equip"`
	Units        map[string]map[string]*unit.Unit             `json:"units"`
	HostileUnits map[string]map[string]*unit.Unit             `json:"hostile_units"`
	UnitStorage  []*unit.Unit                                 `json:"unit_storage"`
	Map          *_map.Map                                    `json:"map"`
	Watch        map[string]map[string]*coordinate.Coordinate `json:"watch"`
	GameStep     int                                          `json:"game_step"`
	GamePhase    string                                       `json:"game_phase"`
}
