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
			playerParams, idMap := initGame.InitGame(msg.IdGame, IdWs(ws, &usersFieldWs)); // отправляет параметры игрока
			var playersParam= FieldResponse{Event: "InitPlayer", UserName: LoginWs(ws, &usersFieldWs), PlayerPrice: playerParams[0], GameStep: playerParams[1], GamePhase: playerParams[2], UserReady: playerParams[3]}
			FieldPipe <- playersParam

			x, y := initGame.GetMap(idMap) // отправляем параметры карты это конечно пиздец >_<
			var mapParam= FieldResponse{Event: "InitMap", UserName: LoginWs(ws, &usersFieldWs), XMap: strconv.Itoa(x), YMap: strconv.Itoa(y)}
			FieldPipe <- mapParam

			units := initGame.GetUnitList(msg.IdGame)
			var unitsParametr = FieldResponse{Event: "InitUnit", UserName: LoginWs(ws, &usersFieldWs), TypeUnit: units[0], UserId: units[1], HP: units[2], UnitAction: units[3],
				Target: units[4], X: units[5], Y: units[6]}
			FieldPipe <- unitsParametr
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
		if msg.Event == "MouseOver" {
			var resp FieldResponse
			success, unitParams := initGame.GetUnit(msg.IdGame, msg.X, msg.Y)
			if success {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), TypeUnit: unitParams[0], UserId: unitParams[1], HP: unitParams[2],
					UnitAction: unitParams[3], Target: unitParams[4], Damage: unitParams[5], MoveSpeed: unitParams[6], Init: unitParams[7], RangeAttack: unitParams[8],
					RangeView: unitParams[9], AreaAttack: unitParams[10], TypeAttack: unitParams[11]}
			}
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

