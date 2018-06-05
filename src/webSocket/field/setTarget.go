package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/targetPhase"
	"../../mechanics/unit"
	"../../mechanics/player"
	"../../mechanics/game"
	"strconv"
)

func SetTarget(msg Message, ws *websocket.Conn) {
	
	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games[client.GetGameID()]

	if findClient && findUnit && findGame && !client.GetReady() {

		targetCoordinate := targetPhase.GetTargetCoordinate(gameUnit, client, activeGame)
		_, find := targetCoordinate[strconv.Itoa(msg.ToX)][strconv.Itoa(msg.ToY)]

		if find {
			targetPhase.SetTarget(gameUnit, activeGame, msg.ToX, msg.ToY)
			ws.WriteJSON(Target{Event: "SetTarget", Unit: gameUnit})
			updateTargetHostileUser(client, activeGame, gameUnit)
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
		}
	}
}

func updateTargetHostileUser(client *player.Player, activeGame *game.Game, gameUnit *unit.Unit) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(gameUnit.X, gameUnit.Y)
			if watch {
				targetPipe <- Target{Event: "SetTarget", UserName: user.GetLogin(),GameID: activeGame.Id, Unit: gameUnit}
			}
		}
	}
}

type Target struct {
	Event    string     `json:"event"`
	UserName string     `json:"user_name"`
	GameID   int        `json:"game_id"`
	Unit     *unit.Unit `json:"unit"`
}
