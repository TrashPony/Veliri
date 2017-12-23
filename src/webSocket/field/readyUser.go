package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	phase, err, phaseChange := game.UserReady(usersFieldWs[ws].GetID(), msg.IdGame)
	client := usersFieldWs[ws]
	activeGame := Games[client.GetGameID()]
	activeGame.SetUserReady(client.GetLogin(), "true") // TODO коректоно обновлять статус готовности игрока
 	players := activeGame.GetPlayers()
	activeUser := ActionGameUser(players)

	if phase != "" { // TODO
		activeGame.GetStat().Phase = phase
	}

	if phase == "attack" {

		attack(activeGame, activeUser, msg, phase)

		phaseChange = true
		phase, _ = game.PhaseСhange(msg.IdGame)
	}

	if err != nil {
		resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Error: err.Error()}
		fieldPipe <- resp
		return
	}

	if 0 == len(usersFieldWs[ws].GetUnits()) {
		resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Error: "not units"}
		fieldPipe <- resp
		// TODO добавить окончание игры
		return
	}

	if phaseChange {
		for _, player := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase}
			fieldPipe <- resp
			activeGame.GetStat().Phase = phase

			if phase == "move" {
				resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase, GameStep: activeGame.GetStat().Step + 1}
				activeGame.GetStat().Step += 1
			}

			for yLine := range player.GetUnits() { // TODO Нахера?
				for _, unit := range player.GetUnits()[yLine] {
					unit.Action = true

					if phase == "move" {
						unit.Target = nil
					}

					var unitsParameter InitUnit
					unitsParameter.initUnit(unit, player.GetLogin())
				}
			}
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), Phase: phase}
		fieldPipe <- resp
	}
}

