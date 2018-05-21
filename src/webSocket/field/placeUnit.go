package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/placePhase"
	"../../mechanics/unit"
	"../../mechanics/game"
	"../../mechanics/watchZone"
	"strconv"
)

func placeUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]
	actionGame, ok := Games[client.GetGameID()]

	if client.GetReady() == false {

		if !ok {
			delete(usersFieldWs, ws)
			return
		}

		storageUnit, find := client.GetUnitStorage(msg.UnitID)

		if find {
			_, find = client.GetCreateZone()[strconv.Itoa(msg.X)][strconv.Itoa(msg.Y)]

			if find {
				_, find := actionGame.GetUnit(msg.X, msg.Y)
				coordinate, _ := actionGame.Map.GetCoordinate(msg.X, msg.Y)

				if !find && coordinate.Type != "obstacle" {
					err := placePhase.PlaceUnit(storageUnit, msg.X, msg.Y, actionGame, client)
					if err == nil {
						ws.WriteJSON(PlaceUnit{Event: "PlaceUnit", Unit: storageUnit})
						UpdatePlaceHostilePlayers(actionGame, msg.X, msg.Y)
						return
					} else {
						ws.WriteJSON(ErrorMessage{Event: "Error", Error: "add to db"})
					}
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is busy"})
				}
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is not allow"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

func UpdatePlaceHostilePlayers(actionGame *game.Game, x, y int) {
	for _, player := range actionGame.GetPlayers() {

		_, find := player.GetWatchCoordinate(x, y)

		if find {
			updater := watchZone.UpdateWatchZone(actionGame, player)
			watchPipe <- Watch{Event: "UpdateWatchMap", UserName: player.GetLogin(), GameID: actionGame.Id, Update: updater}
		}
	}
}

type Watch struct {
	Event    string                      `json:"event"`
	UserName string                      `json:"user_name"`
	GameID   int                         `json:"game_id"`
	Update   *watchZone.UpdaterWatchZone `json:"update"`
}

type PlaceUnit struct {
	Event string     `json:"event"`
	Unit  *unit.Unit `json:"unit"`
}
