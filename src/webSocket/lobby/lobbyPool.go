package lobby

import (
	"log"
	"websocket-master"
	"../../lobby"
	"strconv"
	"sync"
	"errors"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var lobbyPipe = make(chan LobbyResponse)

func AddNewUser(ws *websocket.Conn, login string, id int)  {
	CheckDoubleLogin(login, &usersLobbyWs)
	usersLobbyWs[ws] = &Clients{Login:login, Id:id} // Регистрируем нового Клиента
	print("WS lobby Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	LobbyReader(ws)
}
func LobbyReader(ws *websocket.Conn)  {
	for {
		var msg LobbyMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps= lobby.GetMapList()
			for _, Map := range maps {
				var resp = LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, NameMap: Map.Name, NumOfPlayers: strconv.Itoa(Map.Respawns)}
				lobbyPipe <- resp // Отправляет сообщение в тред
			}
		}

		if msg.Event == "GameView" {
			games := lobby.GetLobbyGames()
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, NameGame: game.Name, NameMap: game.Map, Creator: game.Creator,
				Players: strconv.Itoa(len(game.Users)), NumOfPlayers:strconv.Itoa(len(game.Respawns))}
				lobbyPipe <- resp
			}
		}

		if msg.Event == "DontEndGamesList" {
			games := lobby.GetDontEndGames(usersLobbyWs[ws].Login)
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, NameGame: game.Name, IdGame: game.Id, PhaseGame: game.Phase, StepGame: game.Step, Ready: game.Ready}
				lobbyPipe <- resp
			}
		}

		if msg.Event == "JoinToLobbyGame" {
			var resp LobbyResponse
			game, errGetName := lobby.GetGame(msg.GameName)
			if errGetName != nil {
				log.Panic(errGetName)
			}
			err := lobby.JoinToLobbyGame(msg.GameName, usersLobbyWs[ws].Login)
			if err != nil {
				resp = LobbyResponse{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Login, NameGame: msg.GameName, Error: err.Error()}
				lobbyPipe <- resp
			} else {
				//все кто в лоби получают сообщение о том что подключился новйы игрок
				resp = LobbyResponse{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Login, NameGame: msg.GameName}
				lobbyPipe <- resp

				for user := range game.Users {
					if user != usersLobbyWs[ws].Login {
						resp = LobbyResponse{Event: "NewUser", UserName: user, NewUser: usersLobbyWs[ws].Login}
						lobbyPipe <- resp
					}
				}

				// игрок получает список всех игроков в лоби и их респауны
				for user, ready := range game.Users {
					if user != usersLobbyWs[ws].Login {
						var respown string
						if ready {
							for respawns := range game.Respawns {
								if game.Respawns[respawns] == user {
									respown = respawns.Name
								}
							}
						}
						resp = LobbyResponse{Event: "JoinToLobby", UserName: usersLobbyWs[ws].Login, GameUser: user, Ready: strconv.FormatBool(ready), RespawnName:respown}
						lobbyPipe <- resp
					}
				}
				RefreshLobbyGames(ws) // обновляет кол-во игроков и их характиристики в неигровом лоби
			}
		}

		if msg.Event == "CreateLobbyGame" {
			lobby.CreateNewLobbyGame(msg.GameName, msg.MapName, usersLobbyWs[ws].Login)
			var resp = LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, NameGame: msg.GameName}
			lobbyPipe <- resp

			RefreshLobbyGames(ws)
		}

		if msg.Event == "Ready" {
			var resp LobbyResponse
			game, errGetName := lobby.GetGame(msg.GameName)
			if errGetName != nil {
				log.Panic(errGetName)
			}
			respName, errRespawn := lobby.SetRespawnUser(msg.GameName, usersLobbyWs[ws].Login, msg.Respawn)

			if errRespawn == nil {
				lobby.UserReady(msg.GameName, usersLobbyWs[ws].Login)

				for user:= range game.Users {
					resp = LobbyResponse{Event: msg.Event, UserName: user, GameUser: usersLobbyWs[ws].Login, Ready: strconv.FormatBool(game.Users[usersLobbyWs[ws].Login]), RespawnName: respName}
					lobbyPipe <- resp
				}
			} else {
				resp = LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, GameUser: usersLobbyWs[ws].Login, Error: errRespawn.Error()}
				lobbyPipe <- resp
			}
		}

		if msg.Event == "StartNewGame" {
			game, errGetName := lobby.GetGame(msg.GameName)
			if errGetName != nil {
				log.Panic(errGetName)
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
							var resp = LobbyResponse{Event: msg.Event, UserName: user, IdGame: id}
							lobbyPipe <- resp
						}
					} else {
						var resp= LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Error: errors.New("error ad to DB").Error()}
						lobbyPipe <- resp
					}
				} else {
					var resp = LobbyResponse{Event: msg.Event, UserName:usersLobbyWs[ws].Login,  Error: errors.New("PlayerNotReady").Error()}
					lobbyPipe <- resp
				}
			} else {
				var resp = LobbyResponse{Event: msg.Event, UserName:usersLobbyWs[ws].Login,  Error: errors.New("Players < 2").Error()}
				lobbyPipe <- resp
			}
		}
		if msg.Event == "Respawn"{
			games := lobby.GetLobbyGames()
			user := usersLobbyWs[ws].Login
			for _, game := range games {
				for player := range game.Users {
					if user == player {
						for respawn := range game.Respawns {
							if game.Respawns[respawn] == "" {
								var resp= LobbyResponse{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Respawn: strconv.Itoa(respawn.Id) , RespawnName:respawn.Name}
								lobbyPipe <- resp
							}
						}
						break
					}
				}
			}
		}
		if msg.Event == "Logout" {
			ws.Close()
		}
		if msg.Event == "InitLobby"{
			var resp = LobbyResponse{Event: msg.Event, UserName:usersLobbyWs[ws].Login}
			lobbyPipe <- resp
		}
	}
}

func LobbyReposeSender() {
	for {
		resp := <-lobbyPipe
		mutex.Lock()
		for ws, client := range usersLobbyWs {
			if client.Login == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					lobby.DelLobbyGame(client.Login)
					ws.Close()
					delete(usersLobbyWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id int
}