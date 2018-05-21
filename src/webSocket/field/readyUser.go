package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics"
	"../../mechanics/game"
	"../../mechanics/unit"
)

func Ready(ws *websocket.Conn) {

	client := usersFieldWs[ws]
	activeGame := Games[client.GetGameID()]

	changePhase := mechanics.UserReady(client, activeGame)

	if changePhase {

		ChangePhase(activeGame)

		if activeGame.Phase == "attack" {
			// todo бой
		}
	} else {
		ws.WriteJSON(UserReady{Ready: client.GetReady()})
	}
}

type UserReady struct {
	Ready bool `json:"ready"`
}

func ChangePhase(actionGame *game.Game) {
	for _, client := range actionGame.GetPlayers() {
		phaseInfo := PhaseInfo{
			Event:     "ChangePhase",
			UserName:  client.GetLogin(),
			GameID:    actionGame.Id,
			Ready:     client.GetReady(),
			Units:     client.GetUnits(),
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
	GamePhase string                           `json:"game_phase"`
}
