package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"log"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*game.Player) {
	for ws, client := range *usersWs {
		if client.GetLogin() == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*game.Player, err error) {
	log.Printf("error: %v", err)
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func subtraction(slice1 []*game.Coordinate, slice2 []*game.Coordinate) (ab []game.Coordinate) {
	mb := map[game.Coordinate]bool{}
	for _, x := range slice2 {
		mb[*x] = true
	}
	for _, x := range slice1 {
		if _, ok := mb[*x]; !ok {
			ab = append(ab, *x)
		}
	}
	return ab
}

func ActionGameUser(players []*game.UserStat) (activeUser []*game.Player) {
	for _, clients := range usersFieldWs { // TODO в обьект игры сразу инициализировать всех игроков
		add := false
		for _, userStat := range players {
			if clients.GetLogin() == userStat.Name && clients.GetGameID() == userStat.IdGame {
				add = true
			}
		}
		if add {
			activeUser = append(activeUser, clients)
		}
	}
	return
}