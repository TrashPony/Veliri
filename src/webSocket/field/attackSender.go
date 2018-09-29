package field

import (
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/attackPhase"
	"../../mechanics/localGame/map/watchZone"
)

type AttackMessage struct {
	Event        string                      `json:"event"`
	UserName     string                      `json:"user_name"`
	GameID       int                         `json:"game_id"`
	ResultBattle []*attackPhase.ResultBattle `json:"result_battle"`
	ResultEquip  []*attackPhase.ResultEquip  `json:"result_equip"`
	WatchNode    *watchZone.UpdaterWatchZone `json:"watch_node"`
}

func attack(activeGame *localGame.Game) {

	resultBattle, resultEquip := attackPhase.AttackPhase(activeGame)

	for _, player := range activeGame.GetPlayers() {
		// todo препроцесинг данных, не все пользователи видят весь бой
		attack := AttackMessage{Event: "AttackPhase", UserName: player.GetLogin(), GameID: player.GetGameID(),
			ResultBattle: resultBattle, ResultEquip: resultEquip, WatchNode: watchZone.UpdateWatchZone(activeGame, player)}
		attackPipe <- attack
	}
}
