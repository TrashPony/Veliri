package webSocket

import (
	"log"
	"websocket-master"
	"../game"
	"../game/initGame"
	"../game/createUnit"
	"strconv"
)

func FieldReader(ws *websocket.Conn)  {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs , err)
			break
		}

		if msg.Event == "InitGame" {
			gameStat := initGame.GetGame(msg.IdGame)
			userStat := initGame.GetUserStat(msg.IdGame, IdWs(ws, &usersFieldWs))
			var playersParam= FieldResponse{Event: "InitPlayer", UserName: LoginWs(ws, &usersFieldWs), PlayerPrice: strconv.Itoa(userStat.Price),
				GameStep: strconv.Itoa(gameStat.Step), GamePhase: gameStat.Phase, UserReady: userStat.Ready}
			FieldPipe <- playersParam // отправляет параметры игрока

			mp := initGame.GetMap(gameStat.IdMap)
			var mapParam= FieldResponse{Event: "InitMap", UserName: LoginWs(ws, &usersFieldWs), NameMap: mp.Name, TypeMap: mp.Type, XMap: strconv.Itoa(mp.Xsize), YMap: strconv.Itoa(mp.Ysize)}
			FieldPipe <- mapParam // отправляем параметры карты

			units := initGame.GetUnits(msg.IdGame)

			for i := 0; i < len(units); i++ {
				if LoginWs(ws, &usersFieldWs) == units[i].NameUser {
					PermissCoordinates := game.GetCoordinates(units[i])
					for c := 0; c < len(PermissCoordinates); c++ {
						for j := 0; j < len(units); j++ {
							if (PermissCoordinates[c].X == units[j].X) && (PermissCoordinates[c].Y == units[j].Y) {
								var unitsParametr = FieldResponse{Event: "InitUnit", UserName: LoginWs(ws, &usersFieldWs), TypeUnit: units[j].NameType, UserOwned: units[j].NameUser,
									HP: strconv.Itoa(units[j].Hp), UnitAction: strconv.FormatBool(units[j].Action), Target: strconv.Itoa(units[j].Target), X: strconv.Itoa(units[j].X), Y: strconv.Itoa(units[j].Y) }
								FieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
							}
						}
					}
				}
			}
			//InitResp
			respawn := initGame.GetRespawns(IdWs(ws, &usersFieldWs), msg.IdGame)
			var respawnParametr = FieldResponse{Event: "InitResp", UserName: LoginWs(ws, &usersFieldWs), RespawnX:strconv.Itoa(respawn.X), RespawnY:strconv.Itoa(respawn.Y)}
			FieldPipe <- respawnParametr
		}

		if msg.Event == "CreateUnit" {
			var resp FieldResponse
			// 1) надо проверить возможно ли его туда поставить например в зависимости от респауна
			success, price := createUnit.CreateUnit(msg.IdGame, strconv.Itoa(IdWs(ws, &usersFieldWs)), msg.TypeUnit, msg.X, msg.Y)

			if success {
				resp = FieldResponse{Event:msg.Event, UserName:LoginWs(ws, &usersFieldWs),PlayerPrice: strconv.Itoa(price), X:msg.X, Y:msg.Y, TypeUnit:msg.TypeUnit}
				FieldPipe <- resp
			} else {
				if price == 1 {
					resp = FieldResponse{Event:msg.Event, UserName:LoginWs(ws, &usersFieldWs), ErrorType:"busy"}
					FieldPipe <- resp
				}
				if price == 2 {
					resp = FieldResponse{Event:msg.Event, UserName:LoginWs(ws, &usersFieldWs), ErrorType:"noMany"}
					FieldPipe <- resp
				}
			}
		}

		if msg.Event == "Ready" {
			var resp FieldResponse
			phase := game.UserReady(IdWs(ws, &usersFieldWs), msg.IdGame)
			resp = FieldResponse{Event:msg.Event, UserName:LoginWs(ws, &usersFieldWs), Phase:phase}
			FieldPipe <- resp
		}

		/*if msg.Event == "MouseOver" {
			var resp FieldResponse
			success, unitParams := initGame.GetUnit(msg.IdGame, msg.X, msg.Y)
			if success {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), TypeUnit: unitParams[0], UserId: unitParams[1], HP: unitParams[2],
					UnitAction: unitParams[3], Target: unitParams[4], Damage: unitParams[5], MoveSpeed: unitParams[6], Init: unitParams[7], RangeAttack: unitParams[8],
					RangeView: unitParams[9], AreaAttack: unitParams[10], TypeAttack: unitParams[11]}
				FieldPipe <- resp
			}
		}*/
	}
}

func FieldReposeSender() {
	for {
		resp := <-FieldPipe
		for client := range usersFieldWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					client.ws.Close()
					delete(usersFieldWs, client)
				}
			}
		}
	}
}

