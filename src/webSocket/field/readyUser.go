package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/userReady"
	"github.com/gorilla/websocket"
)

func Ready(ws *websocket.Conn) {
	client := localGame.Clients.GetByWs(ws)

	if client != nil {
		activeGame, findGame := games.Games.Get(client.GetGameID())

		if findGame {
			userReady.UserReady(client)

			changePhase := CheckAllReady(activeGame)

			if !changePhase {
				SendMessage(UserReady{Event: "Ready", Ready: client.GetReady()}, client.GetID(), activeGame.Id)
				QueueSender(activeGame)
			}
		}
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
			GamePhase: actionGame.Phase,
		}
		SendMessage(phaseInfo, client.GetID(), actionGame.Id)
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
