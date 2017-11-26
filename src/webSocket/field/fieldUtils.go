package field

import (
	"../../game/objects"
	"github.com/gorilla/websocket"
	"log"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients) {
	for ws, client := range *usersWs {
		if client.Login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func subtraction(slice1 []*objects.Coordinate, slice2 []*objects.Coordinate) (ab []objects.Coordinate) {
	mb := map[objects.Coordinate]bool{}
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

func ActionGameUser(players []*objects.UserStat) (activeUser []*Clients) {
	for _, clients := range usersFieldWs {
		add := false
		for _, userStat := range players {
			if clients.Login == userStat.Name && clients.GameID == userStat.IdGame {
				add = true
			}
		}
		if add {
			activeUser = append(activeUser, clients)
		}
	}
	return
}
