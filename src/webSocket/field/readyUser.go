package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics"
)

func Ready(msg Message, ws *websocket.Conn) {

	client := usersFieldWs[ws]
	activeGame := Games[client.GetGameID()]

	mechanics.UserReady(client, activeGame)

 	/*players := activeGame.GetPlayers()
	activeUser := ActionGameUser(players)

	if phase != "" { // если произошла смена фазы то устанавливаем ее
		activeGame.GetStat().Phase = phase
	}

	if phase == "attack" {

		attack(activeGame, activeUser, msg, phase)

		phaseChange = true
		phase, _ = game.PhaseСhange(activeGame)
	}

	if err != nil {
		resp := Response{Event: msg.Event, UserName: client.GetLogin(), Error: err.Error()}
		fieldPipe <- resp
		return
	}

	if 0 == len(usersFieldWs[ws].GetUnits()) {
		resp := Response{Event: msg.Event, UserName: client.GetLogin(), Error: "not units"}
		fieldPipe <- resp
		// TODO добавить окончание игры
		return
	}

	if phaseChange {
		for _, player := range activeUser {

			// обновляем статус игроков в памяти
			activeGame.SetUserReady(player.GetLogin(), false)

			resp := Response{Event: msg.Event, UserName: player.GetLogin(), Phase: phase}
			fieldPipe <- resp
			activeGame.GetStat().Phase = phase

			if phase == "move" {
				resp = Response{Event: msg.Event, UserName: player.GetLogin(), Phase: phase, GameStep: activeGame.GetStat().Step + 1}
				activeGame.GetStat().Step += 1
			}

			for yLine := range player.GetUnits() { // TODO Нахера?
				for _, unit := range player.GetUnits()[yLine] {
					unit.Action = true

					if phase == "move" {
						unit.Target = nil
					}

					var unitsParameter InitUnit
					unitsParameter.initUnit("InitUnit", unit, player.GetLogin())
				}
			}
		}
	} else {
		resp := Response{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), Phase: phase}
		fieldPipe <- resp
	}*/
}

