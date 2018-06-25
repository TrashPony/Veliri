package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/localGame"
	"../../mechanics/localGame/userReady"
	"../../mechanics/gameObjects/unit"
)

func Ready(ws *websocket.Conn) {

	client := usersFieldWs[ws]
	activeGame, _ := Games.Get(client.GetGameID())

	changePhase := userReady.UserReady(client, activeGame)

	if changePhase {

		ChangePhase(activeGame)

		if activeGame.Phase == "attack" {
			// todo бой
		}
	} else {
		ws.WriteJSON(UserReady{Event: "Ready", Ready: client.GetReady()})
	}
}

type UserReady struct {
	Event string `json:"event"`
	Ready bool   `json:"ready"`
}

func ChangePhase(actionGame *localGame.Game) {
	for _, client := range actionGame.GetPlayers() {
		phaseInfo := PhaseInfo{
			Event:     "ChangePhase",
			UserName:  client.GetLogin(),
			GameID:    actionGame.Id,
			Ready:     client.GetReady(),
			Units:     client.GetUnits(),
			GameStep:  actionGame.Step,
			GamePhase: actionGame.Phase}

		phasePipe <- phaseInfo
	}
}

type PhaseInfo struct {
	Event     string                           `json:"event"`
	UserName  string                           `json:"user_name"`
	GameID    int                              `json:"game_id"`
	Ready     bool                             `json:"ready"`
	Units     map[string]map[string]*unit.Unit `json:"units"`
	GameStep  int                              `json:"game_step"`
	GamePhase string                           `json:"game_phase"`
}
