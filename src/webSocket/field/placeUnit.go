package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics"
	"../../mechanics/game"
	"../../mechanics/watchZone"
)

func placeUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]
	actionGame, ok := Games[client.GetGameID()]

	if !ok {
		delete(usersFieldWs, ws)
		return
	}

	storageUnit, find := client.GetUnitStorage(msg.UnitID)

	if find {
		for _, coordinate := range client.GetCreateZone() {
			if coordinate.X == msg.X && coordinate.Y == msg.Y {

				_, find := actionGame.GetUnit(msg.X, msg.Y)
				coordinate, _ := actionGame.Map.GetCoordinate(msg.X, msg.Y)

				if !find && coordinate.Type != "obstacle" {

					err := mechanics.PlaceUnit(storageUnit, msg.X, msg.Y, actionGame, client)

					if err == nil {
						UpdatePlaceHostilePlayers(actionGame, msg.X, msg.Y)
						return
					} else {
						ws.WriteJSON(ErrorMessage{Event: "Error", Error: "add to db"} )
					}
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is busy"} )
				}
			}
		}
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is not allow"} )
	}
}

func UpdatePlaceHostilePlayers(actionGame *game.Game, x, y int) {
	for _, player := range actionGame.GetPlayers() {

		_, find := player.GetWatchCoordinate(x, y)

		if find {
			updater := watchZone.UpdateWatchZone(actionGame, player)
			watchPipe <- Watch{Event: "UpdateWatchMap", UserName: player.GetLogin(), Update: updater}
		}
	}
}

type Watch struct {
	Event    string                      `json:"event"`
	UserName string                      `json:"user_name"`
	Update   *watchZone.UpdaterWatchZone `json:"update"`
}
