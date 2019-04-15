package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/initGame"
	"github.com/gorilla/websocket"
)

func loadGame(msg Message, ws *websocket.Conn) {
	client := localGame.Clients.GetByWs(ws)

	if client != nil {

		loadGame, ok := games.Games.GetPlayerID(client.GetID())
		if !ok {
			loadGame = initGame.InitGame(client.GetID())
			games.Games.Add(loadGame.Id, loadGame) // добавляем новую игру в карту активных игор
			go timerMoveUnits(loadGame) // активируем таймер для юнитов
		}

		// берется заного игрок что бы проверить нашлась игра или нет
		player := loadGame.GetPlayer(client.GetID(), client.GetLogin())
		if player != nil && !player.Leave {
			var sendLoadGame = LoadGame{
				Event:             "LoadGame",
				UserName:          player.GetLogin(),
				Ready:             player.GetReady(),
				Units:             player.GetUnits(),
				HostileUnits:      player.GetHostileUnits(),
				MemoryHostileUnit: player.GetMemoryHostileUnits(),
				UnitStorage:       player.GetUnitsStorage(),
				Map:               loadGame.GetMap(),
				Watch:             player.GetWatchCoordinates(),
				GameStep:          loadGame.GetStep(),
				GamePhase:         loadGame.GetPhase(),
				GameZone:          loadGame.GetGameZone(player),
			}
			SendMessage(sendLoadGame, player.GetID(), loadGame.Id)

			if moveUnit := player.GetMoveUnit(); moveUnit != nil && loadGame.Phase == "move" {
				SendMessage(Move{Event: "QueueMove", UserName: client.GetLogin(), GameID: loadGame.Id, Unit: moveUnit}, client.GetID(), loadGame.Id)
			}

		} else {
			SendMessage(ErrorMessage{Event: "Error", Error: "error"}, client.GetID(), loadGame.Id)
		}
	} else {
		SendMessage(ErrorMessage{Event: "Error", Error: "error"}, client.GetID(), 0)
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
	GameZone          map[string]map[string]*coordinate.Coordinate `json:"game_zone"`
}
