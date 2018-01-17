package lobby

import (
	"log"
	"github.com/gorilla/websocket"
	"errors"
	"../../lobby"
)

func StartNewGame(msg Message, ws *websocket.Conn)  {
	game, errGetName := lobby.GetGame(msg.GameName)
	if errGetName != nil {
		log.Panic(errGetName) //TODO no found this game
	} // список игроков которым надо разослать данные взятые из обьекта игры
	if len(game.Users) > 1 {
		var readyAll = true
		for _, ready := range game.Users {
			if !ready {
				readyAll = false
				break
			}
		}
		if readyAll {
			id, success := lobby.StartNewGame(msg.GameName)
			if success {
				lobby.DelLobbyGame(usersLobbyWs[ws].Login) // удаляем обьект игры из лоби, ищем его по имени создателя ¯\_(ツ)_/¯
				for user := range game.Users {
					var resp = Response{Event: msg.Event, UserName: user, IdGame: id}
					lobbyPipe <- resp
				}
			} else {
				var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Error: errors.New("error ad to DB").Error()}
				lobbyPipe <- resp
			}
		} else {
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Error: errors.New("PlayerNotReady").Error()}
			lobbyPipe <- resp
		}
	} else {
		var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Error: errors.New("Players < 2").Error()}
		lobbyPipe <- resp
	}
}
