package webSocket

import (
	"websocket-master"
	"log"
	"../DB_info"
	"strconv"
	"../game/initGame"
)


func LoginWs(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients) (string)  {
	login := (*usersWs)[ws].login
	return login
}

func IdWs(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients) (int)  {
	id := (*usersWs)[ws].id
	return id
}

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients)  {
	for ws, client  := range *usersWs {
		if client.login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	login := (*usersWs)[ws].login

	delGame, users := DB_info.DelLobbyGame(login)
	diconnect, nameGame := DB_info.DisconnectLobbyGame(login)
	if delGame || diconnect {
		if delGame {
			DiconnectLobby(users)
		}
		if diconnect {
			RefreshUsersList(nameGame)
		}
		RefreshLobbyGames(ws)
	}
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func RefreshUsersList(nameGame string)  {
	games := DB_info.GetLobbyGames()
	for _, game := range games {
		if game.Name == nameGame {
			for player, ready := range game.Users {
				var refresh = LobbyResponse{Event: "DelUser", UserName: player}
				LobbyPipe <- refresh
				var respown int
				if ready {
					for respawns := range game.Respawns {
						if game.Respawns[respawns] == player {
							respown = respawns.Id
						}
					}
				}
				refresh = LobbyResponse{Event: "UserRefresh", UserName: player, GameUser: player, Ready: strconv.FormatBool(ready), Respawn:strconv.Itoa(respown)}
				LobbyPipe <- refresh
			}
			break
		}
	}
}

func DiconnectLobby(users map[string]bool)  {
	for client := range users {
		var refresh = LobbyResponse{Event: "DisconnectLobby",  UserName: client}
		LobbyPipe <- refresh
	}
}
func RefreshLobbyGames(ws *websocket.Conn)  {
	login := (usersLobbyWs)[ws].login
	games := DB_info.GetLobbyGames()
	for _, client  := range usersLobbyWs { // TODO: сильно затратная операция, над сделать что бы отсылалась только новая игра а не обновляляся весь список заного
		if client.login != login {
			var refresh = LobbyResponse{Event: "GameRefresh",  UserName: client.login}
			LobbyPipe <- refresh
			for _, game := range games {
				var resp = LobbyResponse{Event: "GameView", UserName: client.login, NameGame: game.Name, NameMap: game.Map, Creator: game.Creator,
					Players: strconv.Itoa(len(game.Users)), NumOfPlayers: strconv.Itoa(len(game.Respawns))}
				LobbyPipe <- resp
			}
		}
	}
}

func subtraction(slice1 []initGame.Coordinate, slice2 []initGame.Coordinate) []initGame.Coordinate  {
	mb := map[initGame.Coordinate]bool{}
	for _, x := range slice2 {
		mb[x] = true
	}
	ab := []initGame.Coordinate{}
	for _, x := range slice1 {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func difference(slice1 []initGame.Coordinate, slice2 []initGame.Coordinate) []initGame.Coordinate {
	var diff []initGame.Coordinate

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
