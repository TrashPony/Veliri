package field

import (
	"websocket-master"
	"../../game/mechanics"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	if 0 < len(usersFieldWs[ws].Units) {
		phase, err, phaseChange := mechanics.UserReady(usersFieldWs[ws].Id, msg.IdGame)
		if err != nil {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Error: err.Error()}
		} else {
			if phaseChange {
				for _, userStat := range usersFieldWs[ws].Players {
					resp = FieldResponse{Event: msg.Event, UserName: userStat.Name, Phase: phase}
					fieldPipe <- resp // TODO: надо обновить всем информацию в соеденения типо фаза, ход и тд
					for _, clients := range usersFieldWs{
						if clients.Login == userStat.Name {
							clients.GameStat.Phase = phase
							if phase == "move" {
								clients.GameStat.Step += 1
							}
							break
						}
					}
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: phase}
				fieldPipe <- resp
			}
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Error: "not units"}
		fieldPipe <- resp
	}
}