package field

import (
	"websocket-master"
	"../../game/mechanics"
	"strconv"
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
					fieldPipe <- resp
					for _, clients := range usersFieldWs{
						if clients.Login == userStat.Name && clients.GameStat.Id == usersFieldWs[ws].GameStat.Id {
							clients.GameStat.Phase = phase
							if phase == "move" {
								resp = FieldResponse{Event: msg.Event, UserName: userStat.Name, Phase: phase, GameStep:clients.GameStat.Step + 1}
								clients.GameStat.Step += 1
							}
							for _, unit := range clients.Units {
								unit.Action = true
								var unitsParametr = InitUnit{Event: "InitUnit", UserName: clients.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
									HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y}
								initUnit <- unitsParametr // отправляем параметры каждого юнита отдельно
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