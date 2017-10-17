package webSocket

import (
	"log"
	"websocket-master"
	"../DB_info"
	"strconv"
	"sync"
	"errors"
)

// пайп доп. читать в документации
var mutex = &sync.Mutex{}
func LobbyReader(ws *websocket.Conn)  {
	for {
		var msg LobbyMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps= DB_info.GetMapList()
			for _, Map := range maps {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameMap: Map.Name, NumOfPlayers: strconv.Itoa(Map.Respawns)}
				LobbyPipe <- resp // Отправляет сообщение в тред
			}
		}

		if msg.Event == "GameView" {
			games := DB_info.GetLobbyGames()
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameGame: game.Name, NameMap: game.Map, Creator: game.Creator,
				Players: strconv.Itoa(len(game.Users)), NumOfPlayers:strconv.Itoa(len(game.Respawns))}
				LobbyPipe <- resp
			}
		}

		if msg.Event == "DontEndGamesList" {
			games := DB_info.GetDontEndGames(LoginWs(ws, &usersLobbyWs))
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameGame: game.Name, IdGame: game.Id, PhaseGame: game.Phase, StepGame: game.Step, Ready: game.Ready}
				LobbyPipe <- resp
			}
		}

		if msg.Event == "JoinToLobbyGame" {
			var resp LobbyResponse
			game, errGetName := DB_info.GetGame(msg.GameName)
			if errGetName != nil {
				log.Panic(errGetName)
			}
			err := DB_info.JoinToLobbyGame(msg.GameName, LoginWs(ws, &usersLobbyWs))
			if err != nil {
				log.Panic(err)
				resp = LobbyResponse{Event: "initLobbyGame", UserName: LoginWs(ws, &usersLobbyWs), NameGame: msg.GameName, Error: err.Error()}
			} else {
				//все кто в лоби получают сообщение о том что подключился новйы игрок
				resp = LobbyResponse{Event: "initLobbyGame", UserName: LoginWs(ws, &usersLobbyWs), NameGame: msg.GameName}
				LobbyPipe <- resp

				for user := range game.Users {
					if user != LoginWs(ws, &usersLobbyWs) {
						resp = LobbyResponse{Event: "NewUser", UserName: user, NewUser: LoginWs(ws, &usersLobbyWs)}
						LobbyPipe <- resp
					}
				}

				// игрок получает список всех игроков в лоби и их респауны
				for user, ready := range game.Users {
					if user != LoginWs(ws, &usersLobbyWs) {
						var respown int
						if ready {
							for respawns := range game.Respawns {
								if game.Respawns[respawns] == user {
									respown = respawns.Id
								}
							}
						}
						resp = LobbyResponse{Event: "JoinToLobby", UserName: LoginWs(ws, &usersLobbyWs), GameUser: user, Ready: strconv.FormatBool(ready), Respawn:strconv.Itoa(respown)}
						LobbyPipe <- resp
					}
				}
				RefreshLobbyGames(ws) // обновляет кол-во игроков и их характиристики в неигровом лоби
			}
		}

		if msg.Event == "CreateLobbyGame" {
			DB_info.CreateNewLobbyGame(msg.GameName, msg.MapName, LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameGame: msg.GameName}
			LobbyPipe <- resp

			RefreshLobbyGames(ws)
		}

		if msg.Event == "Ready" {
			var resp LobbyResponse
			game, errGetName := DB_info.GetGame(msg.GameName)
			if errGetName != nil {
				log.Panic(errGetName)
			}
			respId, errRespawn := DB_info.SetRespawnUser(msg.GameName, LoginWs(ws, &usersLobbyWs), msg.Respawn)

			if errRespawn == nil {
				DB_info.UserReady(msg.GameName, LoginWs(ws, &usersLobbyWs))

				for user:= range game.Users {
					resp = LobbyResponse{Event: msg.Event, UserName: user, GameUser: LoginWs(ws, &usersLobbyWs), Ready: strconv.FormatBool(game.Users[LoginWs(ws, &usersLobbyWs)]), Respawn: respId}
					LobbyPipe <- resp
				}
			} else {
				resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), GameUser: LoginWs(ws, &usersLobbyWs), Error: errRespawn.Error()}
				LobbyPipe <- resp
			}
		}

		if msg.Event == "StartNewGame" {
			game, errGetName := DB_info.GetGame(msg.GameName)
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
					id, success := DB_info.StartNewGame(msg.GameName)
					if success {
						DB_info.DelLobbyGame(LoginWs(ws, &usersLobbyWs)) // удаляем обьект игры из лоби, ищем его по имени создателя ¯\_(ツ)_/¯
						for user := range game.Users {
							var resp = LobbyResponse{Event: msg.Event, UserName: user, IdGame: id}
							LobbyPipe <- resp
						}
					} else {
						var resp= LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), Error: errors.New("error ad to DB").Error()}
						LobbyPipe <- resp
					}
				} else {
					var resp = LobbyResponse{Event: msg.Event, UserName:LoginWs(ws, &usersLobbyWs),  Error: errors.New("PlayerNotReady").Error()}
					LobbyPipe <- resp
				}
			} else {
				var resp = LobbyResponse{Event: msg.Event, UserName:LoginWs(ws, &usersLobbyWs),  Error: errors.New("Players < 2").Error()}
				LobbyPipe <- resp
			}
		}
		if msg.Event == "Respawn"{
			games := DB_info.GetLobbyGames()
			user := LoginWs(ws, &usersLobbyWs)
			for _, game := range games {
				for player := range game.Users {
					if user == player {
						for respawn := range game.Respawns {
							if game.Respawns[respawn] == "" {
								var resp= LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), Respawn: strconv.Itoa(respawn.Id)}
								LobbyPipe <- resp
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
			var resp = LobbyResponse{Event: msg.Event, UserName:LoginWs(ws, &usersLobbyWs)}
			LobbyPipe <- resp
		}
	}
}

func LobbyReposeSender() {
	for {
		resp := <-LobbyPipe
		mutex.Lock()
		for client := range usersLobbyWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					DB_info.DelLobbyGame(client.login)
					client.ws.Close()
					delete(usersLobbyWs, client)
				}
			}
		}
		mutex.Unlock()
	}
}