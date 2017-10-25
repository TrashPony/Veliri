package lobby

import (
	"websocket-master"
	"log"
	"../../lobby"
	"strconv"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients)  {
	for ws, client  := range *usersWs {
		if client.Login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	login := (*usersWs)[ws].Login

	delGame, users := lobby.DelLobbyGame(login)
	diconnect, nameGame := lobby.DisconnectLobbyGame(login)
	if delGame || diconnect {
		if delGame {
			DiconnectLobby(users)
		}
		if diconnect {
			RefreshUsersList(nameGame)
		}
		RefreshLobbyGames(ws) //TODO: // вываливается экзепшен при выходе главного игрока из лоби игры когда в игре кто то есть
	}
	delete(*usersWs, ws) // удаляем его из активных подключений
}

func RefreshUsersList(nameGame string)  {
	games := lobby.GetLobbyGames()
	for _, game := range games {
		if game.Name == nameGame {
			for player, ready := range game.Users {
				var refresh = LobbyResponse{Event: "DelUser", UserName: player}
				lobbyPipe <- refresh
				var respown int
				if ready {
					for respawns := range game.Respawns {
						if game.Respawns[respawns] == player {
							respown = respawns.Id
						}
					}
				}
				refresh = LobbyResponse{Event: "UserRefresh", UserName: player, GameUser: player, Ready: strconv.FormatBool(ready), Respawn:strconv.Itoa(respown)}
				lobbyPipe <- refresh
			}
			break
		}
	}
}

func DiconnectLobby(users map[string]bool)  {
	for client := range users {
		var refresh = LobbyResponse{Event: "DisconnectLobby",  UserName: client}
		lobbyPipe <- refresh
	}
}
func RefreshLobbyGames(ws *websocket.Conn)  {
	login := (usersLobbyWs)[ws].Login // TODO: // вываливается экзепшен при выходе главного игрока из лоби игры когда в игре кто то есть
	games := lobby.GetLobbyGames()
	for _, client  := range usersLobbyWs {
		if client.Login != login {
			var refresh = LobbyResponse{Event: "GameRefresh",  UserName: client.Login}
			lobbyPipe <- refresh
			for _, game := range games {
				var resp = LobbyResponse{Event: "GameView", UserName: client.Login, NameGame: game.Name, NameMap: game.Map, Creator: game.Creator,
					Players: strconv.Itoa(len(game.Users)), NumOfPlayers: strconv.Itoa(len(game.Respawns))}
				lobbyPipe <- resp
			}
		}
	}
}