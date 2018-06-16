package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics"
	"../../mechanics/unit"
	"../../mechanics/equip"
	"../../mechanics/gameMap"
	"../../mechanics/matherShip"
	"../../mechanics/coordinate"
)

func loadGame(msg Message, ws *websocket.Conn) {
	loadGame, ok := Games.Get(msg.IdGame)
	newClient, _ := usersFieldWs[ws]

	if !ok {
		loadGame = mechanics.InitGame(msg.IdGame)
		Games.Add(loadGame.Id, loadGame) // добавляем новую игру в карту активных игор
		println(loadGame)
	}

	player := loadGame.GetPlayer(newClient.GetID(), newClient.GetLogin())

	if player != nil {

		usersFieldWs[ws] = player

		var sendLoadGame = LoadGame{
			Event:              "LoadGame",
			UserName:           usersFieldWs[ws].GetLogin(),
			Ready:              usersFieldWs[ws].GetReady(),
			Equip:              usersFieldWs[ws].GetEquips(),
			Units:              usersFieldWs[ws].GetUnits(),
			HostileUnits:       usersFieldWs[ws].GetHostileUnits(),
			UnitStorage:        usersFieldWs[ws].GetUnitsStorage(),
			Map:                loadGame.GetMap(),
			MatherShip:         usersFieldWs[ws].GetMatherShip(),
			HostileMatherShips: usersFieldWs[ws].GetHostileMatherShips(),
			Watch:              usersFieldWs[ws].GetWatchCoordinates(),
			GameStep:           loadGame.GetStep(),
			GamePhase:          loadGame.GetPhase()}
		ws.WriteJSON(sendLoadGame)
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "error"})
	}
}

type LoadGame struct {
	Event              string                                       `json:"event"`
	UserName           string                                       `json:"user_name"`
	Ready              bool                                         `json:"ready"`
	Equip              []*equip.Equip                               `json:"equip"`
	Units              map[string]map[string]*unit.Unit             `json:"units"`
	HostileUnits       map[string]map[string]*unit.Unit             `json:"hostile_units"`
	UnitStorage        []*unit.Unit                                 `json:"unit_storage"`
	Map                *gameMap.Map                                 `json:"map"`
	MatherShip         *matherShip.MatherShip                       `json:"mather_ship"`
	HostileMatherShips map[string]map[string]*matherShip.MatherShip `json:"hostile_mather_ships"`
	Watch              map[string]map[string]*coordinate.Coordinate `json:"watch"`
	GameStep           int                                          `json:"game_step"`
	GamePhase          string                                       `json:"game_phase"`
}
