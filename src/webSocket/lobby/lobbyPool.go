package lobby

import (
	"../../lobby"
	"../../lobby/Squad"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var lobbyPipe = make(chan Response)

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersLobbyWs)
	usersLobbyWs[ws] = &Clients{Login: login, Id: id} // Регистрируем нового Клиента
	print("WS lobby Сессия: ")                        // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	go NewLobbyUser(login, &usersLobbyWs)
	go SentOnlineUser(login, &usersLobbyWs)

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps = lobby.GetMapList()
			for _, Map := range maps {
				var resp = Response{Event: msg.Event, Map: Map}
				ws.WriteJSON(resp)
			}
		}

		if msg.Event == "GameView" {
			games := lobby.GetLobbyGames()
			for _, game := range games {
				var resp = Response{Event: msg.Event, Game: game}
				ws.WriteJSON(resp)
			}
		}

		if msg.Event == "CreateLobbyGame" {
			game := lobby.CreateNewLobbyGame(msg.GameName, msg.MapID, usersLobbyWs[ws].Login)

			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, Game: game}
			ws.WriteJSON(resp)

			RefreshLobbyGames(usersLobbyWs[ws].Login)
		}

		if msg.Event == "JoinToLobbyGame" {
			JoinToLobbyGame(msg, ws)
		}

		if msg.Event == "InitLobby" {
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login}
			ws.WriteJSON(resp)
		}

		if msg.Event == "Respawn" {
			games := lobby.GetLobbyGames()
			user := usersLobbyWs[ws].Login
			for _, game := range games {
				for player := range game.Users {
					if user == player {
						for _, respawn := range game.Respawns {
							if respawn.UserName == "" {
								var resp = Response{Event: msg.Event, Respawn: respawn}
								ws.WriteJSON(resp)
							}
						}
					}
				}
			}
		}

		if msg.Event == "Ready" {
			Ready(msg, ws)
		}

		if msg.Event == "StartNewGame" {
			StartNewGame(msg, ws)
		}

		/*if msg.Event == "DontEndGamesList" {
			games := lobby.GetDontEndGames(usersLobbyWs[ws].Login)
			for _, game := range games {
				var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, NameGame: game.Name, IdGame: game.Id, PhaseGame: game.Phase, StepGame: game.Step, Ready: game.Ready}
				ws.WriteJSON(resp)
			}
		}//todo
		*/

		if msg.Event == "AddNewSquad" || msg.Event == "SelectSquad" || msg.Event == "SelectMatherShip" || msg.Event == "DeleteSquad" {
			SquadSettings(ws, msg)
		}

		if msg.Event == "GetMatherShips" || msg.Event == "GetListSquad" || msg.Event == "GetDetailOfUnits" || msg.Event == "GetEquipping" {
			GetDetailSquad(ws, msg)
		}

		if msg.Event == "AddUnit" || msg.Event == "ReplaceUnit" || msg.Event == "RemoveUnit" {
			UnitSquad(ws, msg)
		}

		if msg.Event == "AddEquipment" || msg.Event == "ReplaceEquipment" || msg.Event == "RemoveEquipment" {
			EquipSquad(ws, msg)
		}

		if msg.Event == "UnitConstructor" {
			UnitConstructor(ws, msg)
		}

		if msg.Event == "Logout" {
			ws.Close()
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

type Clients struct {
	Login  string
	Id     int
	Squad  *Squad.Squad
	Squads []*Squad.Squad
}
