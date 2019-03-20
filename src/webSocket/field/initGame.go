package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/initGame"
	"github.com/gorilla/websocket"
)

func loadGame(msg Message, ws *websocket.Conn) {
	// TODO смотреть в какой игре состоит пользователь и загружать ее а не присылать ее сообщением
	IDGame := 0

	loadGame, ok := games.Games.Get(IDGame)
	newClient, _ := usersFieldWs[ws]

	if !ok {
		loadGame = initGame.InitGame(IDGame)
		games.Games.Add(loadGame.Id, loadGame) // добавляем новую игру в карту активных игор
	}

	player := loadGame.GetPlayer(newClient.GetID(), newClient.GetLogin())

	if player != nil {

		usersFieldWs[ws] = player

		var sendLoadGame = LoadGame{
			Event:             "LoadGame",
			UserName:          usersFieldWs[ws].GetLogin(),
			Ready:             usersFieldWs[ws].GetReady(),
			Units:             usersFieldWs[ws].GetUnits(),
			HostileUnits:      usersFieldWs[ws].GetHostileUnits(),
			MemoryHostileUnit: usersFieldWs[ws].GetMemoryHostileUnits(),
			UnitStorage:       usersFieldWs[ws].GetUnitsStorage(),
			Map:               loadGame.GetMap(),
			Watch:             usersFieldWs[ws].GetWatchCoordinates(),
			GameStep:          loadGame.GetStep(),
			GamePhase:         loadGame.GetPhase()}
		ws.WriteJSON(sendLoadGame)

		if loadGame.Phase == "move" {
			UserQueueSender(newClient, loadGame)
		}

	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "error"})
	}
}

type LoadGame struct {
	Event             string                                       `json:"event"`
	UserName          string                                       `json:"user_name"`
	Ready             bool                                         `json:"ready"`
	Equip             []*equip.Equip                               `json:"equip"`
	Units             map[string]map[string]*unit.Unit             `json:"units"`
	HostileUnits      map[string]map[string]*unit.Unit             `json:"hostile_units"`
	MemoryHostileUnit map[string]unit.Unit                         `json:"memory_hostile_unit"`
	UnitStorage       []*unit.Unit                                 `json:"unit_storage"`
	Map               *_map.Map                                    `json:"map"`
	Watch             map[string]map[string]*coordinate.Coordinate `json:"watch"`
	GameStep          int                                          `json:"game_step"`
	GamePhase         string                                       `json:"game_phase"`
}
