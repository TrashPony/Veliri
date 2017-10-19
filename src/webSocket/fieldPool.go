package webSocket

import (
	"log"
	"websocket-master"
	"../game"
	"../game/initGame"
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

			units := initGame.GetAllUnits(msg.IdGame)
			for i := 0; i < len(units); i++ {
				sendPermissionCoordinates(msg.IdGame, LoginWs(ws, &usersFieldWs), units[i])
			}

			respawn := initGame.GetRespawns(IdWs(ws, &usersFieldWs), msg.IdGame)
			var respawnParametr = FieldResponse{Event: "InitResp", UserName: LoginWs(ws, &usersFieldWs), RespawnX:strconv.Itoa(respawn.X), RespawnY:strconv.Itoa(respawn.Y)}
			FieldPipe <- respawnParametr
		}

		if msg.Event == "CreateUnit" {
			var resp FieldResponse
			// 1) надо проверить возможно ли его туда поставить например в зависимости от респауна
			unit, price, createError := game.CreateUnit(msg.IdGame, strconv.Itoa(IdWs(ws, &usersFieldWs)), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), PlayerPrice: strconv.Itoa(price), X: strconv.Itoa(unit.X), Y: strconv.Itoa(unit.Y), TypeUnit: unit.NameType}
				FieldPipe <- resp
				sendPermissionCoordinates(msg.IdGame, LoginWs(ws, &usersFieldWs), unit)
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), ErrorType: createError.Error()}
				FieldPipe <- resp
			}
		}

		if msg.Event == "MouseOver" {
			var resp FieldResponse
			unit, errUnitParams := initGame.GetXYUnits(msg.IdGame, msg.X, msg.Y)
			if errUnitParams == nil {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: strconv.Itoa(unit.Hp),
					UnitAction: strconv.FormatBool(unit.Action) , Target: strconv.Itoa(unit.Target), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
					Init: strconv.Itoa(unit.Init), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
					AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
				FieldPipe <- resp
			}
		}

		if msg.Event == "Ready" {
			var resp FieldResponse
			phase := game.UserReady(IdWs(ws, &usersFieldWs), msg.IdGame)
			resp = FieldResponse{Event:msg.Event, UserName:LoginWs(ws, &usersFieldWs), Phase:phase}
			FieldPipe <- resp
		}
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

func sendPermissionCoordinates(idGame string, login string, unit initGame.Unit) {
	units := initGame.GetAllUnits(idGame)
	if login == unit.NameUser {
		PermissCoordinates := game.GetCoordinates(unit)
		for c := 0; c < len(PermissCoordinates); c++ {
			var emptyCoordinates= FieldResponse{Event: "emptyCoordiantes", UserName: login, X: strconv.Itoa(PermissCoordinates[c].X), Y: strconv.Itoa(PermissCoordinates[c].Y)}
			FieldPipe <- emptyCoordinates
			for j := 0; j < len(units); j++ {
				if (PermissCoordinates[c].X == units[j].X) && (PermissCoordinates[c].Y == units[j]	.Y) {
					var unitsParametr = FieldResponse{Event: "InitUnit", UserName: login, TypeUnit: units[j].NameType, UserOwned: units[j].NameUser,
						HP: strconv.Itoa(units[j].Hp), UnitAction: strconv.FormatBool(units[j].Action), Target: strconv.Itoa(units[j].Target), X: strconv.Itoa(units[j].X), Y: strconv.Itoa(units[j].Y)}
					FieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
				}
			}
		}
	}
}

