package field

import (
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/attackPhase"
)

type AttackMessage struct {
	Event        string                     `json:"event"`
	UserName     string                     `json:"user_name"`
	GameID       int                        `json:"game_id"`
	ResultBattle []attackPhase.ResultBattle `json:"result_battle"`
}

func attack(activeGame *localGame.Game) {

	resultBattle := attackPhase.AttackPhase(activeGame)

	for _, player := range activeGame.GetPlayers() {
		attack := AttackMessage{Event: "AttackPhase", UserName: player.GetLogin(), GameID: player.GetGameID(),
			ResultBattle: dataPreparation(resultBattle)}
		attackPipe <- attack
	}
}

func dataPreparation(resultBattle []*attackPhase.ResultBattle) []attackPhase.ResultBattle {
	watchResultBattle := make([]attackPhase.ResultBattle, 0)

	for _, actionBattle := range resultBattle {
		// todo не все пользователи видят весь бой, надо фильтровал
		watchResultBattle = append(watchResultBattle, *actionBattle)
	}

	return watchResultBattle
}
