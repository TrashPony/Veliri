package webSocket

import (
	"log"
	"websocket-master"
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
			playerParams, idMap := initGame.InitGame(msg.IdGame); // отправляет параметры игрока
			var playersParam = FieldResponse{Event:"InitPlayer",UserName:LoginWs(ws, &usersFieldWs), PlayerPrice:playerParams[0], GameStep:playerParams[1], GamePhase:playerParams[2]}
			FieldPipe <- playersParam

			x, y := initGame.GetMap(idMap) // отправляем параметры карты это конечно пиздец >_<
			var mapParam = FieldResponse{Event:"InitMap",UserName:LoginWs(ws, &usersFieldWs), XMap: strconv.Itoa(x), YMap: strconv.Itoa(y)}
			FieldPipe <- mapParam

			if(playerParams[2] != "Init"){ // если игроки еще не начали играть значить и юнитов нет
				// тут надо возвращать параметры юнитов и их расположение на карте
			}
		}
		if msg.Event == "CreateUnit" {
			// 1) надо проверить возможно ли его туда поставить например в зависимости от респауна
			createUnit.CreateUnit(msg.IdGame, strconv.Itoa(IdWs(ws, &usersFieldWs)), msg.TypeUnit, msg.X, msg.Y)

			var resp = FieldResponse{Event:msg.Event,UserName:LoginWs(ws, &usersFieldWs), X:msg.X, Y:msg.Y, TypeUnit:msg.TypeUnit}
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

