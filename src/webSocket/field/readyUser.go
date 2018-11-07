package field

import (
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/userReady"
	"github.com/gorilla/websocket"
)

func Ready(ws *websocket.Conn) {
	client := usersFieldWs[ws]
	activeGame, _ := Games.Get(client.GetGameID())

	userReady.UserReady(client)

	changePhase := CheckAllReady(activeGame)

	if !changePhase {
		ws.WriteJSON(UserReady{Event: "Ready", Ready: client.GetReady()})
		QueueSender(activeGame)
	}
}

func CheckAllReady(activeGame *localGame.Game) bool {

	changePhase := userReady.AllReady(activeGame)

	if changePhase {

		ChangePhase(activeGame)

		if activeGame.Phase == "attack" {
			attack(activeGame)
		}
	}

	return changePhase
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
