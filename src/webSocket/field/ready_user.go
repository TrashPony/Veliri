package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/userReady"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
	"time"
)

func Ready(client *player.Player) {
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

func CheckAllReady(activeGame *localGame.Game) bool {

	changePhase := userReady.AllReady(activeGame)

	if changePhase {

		ChangePhase(activeGame)

		if activeGame.Phase == "targeting" {
			go timerTargetingPhase(activeGame)
		}

		if activeGame.Phase == "attack" {
			attack(activeGame)
		}
	}

	return changePhase
}

// принудительная смена фазы по таймеру
func timerTargetingPhase(game *localGame.Game) {
	currentStep := game.Step

	for i := 60; i > 0; i-- {

		if currentStep != game.Step || game.Phase != "targeting" {
			// если шаг или фаза изменились то значит смена произошла уже с помощью игроков
			return
		}

		time.Sleep(1 * time.Second)
		SendAllMessage(Message{Event: "timeToChangePhase", Seconds: i}, game)
	}

	for _, user := range game.GetPlayers() {
		userReady.UserReady(user)
	}

	CheckAllReady(game)
}

type UserReady struct {
	Event string `json:"event"`
	Ready bool   `json:"ready"`
}

func ChangePhase(actionGame *localGame.Game) {
	for _, client := range actionGame.GetPlayers() {

		if client.Leave {
			continue
		}

		// проверяем игровую зону, если игрок вышел из зоны боевых действий то он может выйти из боя.
		gameZone := actionGame.GetGameZone(client)
		_, find := gameZone[strconv.Itoa(client.GetSquad().MatherShip.Q)][strconv.Itoa(client.GetSquad().MatherShip.R)]

		phaseInfo := PhaseInfo{
			Event:      "ChangePhase",
			UserName:   client.GetLogin(),
			GameID:     actionGame.Id,
			Ready:      client.GetReady(),
			Units:      client.GetUnits(),
			GameStep:   actionGame.Step,
			GamePhase:  actionGame.Phase,
			FleeBattle: find,
		}
		SendMessage(phaseInfo, client.GetID(), actionGame.Id)
	}
}

type PhaseInfo struct {
	Event      string                           `json:"event"`
	UserName   string                           `json:"user_name"`
	GameID     int                              `json:"game_id"`
	Ready      bool                             `json:"ready"`
	Units      map[string]map[string]*unit.Unit `json:"units"`
	GameStep   int                              `json:"game_step"`
	GamePhase  string                           `json:"game_phase"`
	FleeBattle bool                             `json:"flee_battle"`
}
