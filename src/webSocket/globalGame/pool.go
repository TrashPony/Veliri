package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/maps"
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/map"
	"../../mechanics/gameObjects/squad"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersGlobalWs = make(map[*websocket.Conn]*player.Player)

type Message struct {
	Event    string                `json:"event"`
	Map      *_map.Map             `json:"map"`
	Error    string                `json:"error"`
	Squad    *squad.Squad          `json:"squad"`
	User     *player.Player        `json:"user"`
	Bases    map[int]*base.Base    `json:"bases"`
	ToX      float64               `json:"to_x"`
	ToY      float64               `json:"to_y"`
	PathUnit globalGame.PathUnit   `json:"path_unit"`
	Path     []globalGame.PathUnit `json:"path"`
	BaseID   int                   `json:"base_id"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersGlobalWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersGlobalWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS global Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	mutex.Unlock()

	Reader(ws)
}

func Reader(ws *websocket.Conn) {

	stopMove := make(chan bool)
	move := false

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			println(err.Error())
			utils.DelConn(ws, &usersGlobalWs, err)
			break
		}

		if msg.Event == "InitGame" {
			mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
			user := usersGlobalWs[ws]

			if user.GetSquad().GlobalX == 0 && user.GetSquad().GlobalY == 0 {
				x, y := globalGame.GetXYCenterHex(user.GetSquad().Q, user.GetSquad().R)
				user.GetSquad().GlobalX = x
				user.GetSquad().GlobalY = y
			}

			if find && user != nil && user.InBaseID == 0 {
				ws.WriteJSON(Message{
					Event: msg.Event,
					Map:   mp, User: user,
					Squad: user.GetSquad(),
					Bases: bases.Bases.GetBasesByMap(usersGlobalWs[ws].GetSquad().MapID),
					// todo отдача других игроков
				})
			} else {
				ws.WriteJSON(Message{Event: "Error", Error: "no allow"})
			}
		}

		if msg.Event == "MoveTo" {
			mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
			user := usersGlobalWs[ws]

			if find && user.InBaseID == 0 {
				if move {
					stopMove <- true // останавливаем прошлое движение
				}

				path := globalGame.MoveTo(user, msg.ToX, msg.ToY, mp)

				err := ws.WriteJSON(Message{Event: "PreviewPath", Path: path})
				if err != nil {
					println(err.Error())
				}

				go MoveUserMS(ws, msg, user, path, stopMove, &move)
				move = true
			}
		}

		if msg.Event == "IntoToBase" {

			user := usersGlobalWs[ws]

			if move {
				stopMove <- true // останавливаем прошлое движение
			}

			intoBase, _ := bases.Bases.Get(msg.BaseID)
			x, y := globalGame.GetXYCenterHex(intoBase.Q, intoBase.R)

			dist := ((user.GetSquad().GlobalX - x) * (user.GetSquad().GlobalX - x)) +
				((user.GetSquad().GlobalY - y) * (user.GetSquad().GlobalX - y))

			if dist < 400*400 { // 400*400 это 400 пикселей, выбрано рандомно
				user.InBaseID = intoBase.ID
				bases.UserIntoBase(user.GetID(), intoBase.ID)
				user.GetSquad().GlobalX = 0
				user.GetSquad().GlobalY = 0

				ws.WriteJSON(Message{Event: "IntoToBase"})
			} else {
				ws.WriteJSON(Message{Event: "Error", Error: "no allow"})
			}
		}
	}
}
