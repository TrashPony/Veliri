package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/maps"
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/map"
	"../../mechanics/gameObjects/squad"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersGlobalWs = make(map[*websocket.Conn]*player.Player)

const HexagonHeight  = 111 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth  = 100
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

type Message struct {
	Event string             `json:"event"`
	Map   *_map.Map          `json:"map"`
	Error string             `json:"error"`
	Squad *squad.Squad       `json:"squad"`
	User  *player.Player     `json:"user"`
	Bases map[int]*base.Base `json:"bases"`
	ToX   int                `json:"to_x"`
	ToY   int                `json:"to_y"`
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
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			println(err.Error())
			utils.DelConn(ws, &usersGlobalWs, err)
			break
		}

		/*
		TODO Механика глоабльной карты не продуманая часть:
			реал тайм рпг каждый клик игрока мгновенно просчитывается на бекенде

			Сервер знает что игрок находится в позиции (10, 10); клиент говорит: «Я хочу подвинуться на единицу вправо».
			Сервер обновляет позицию игрока на (11, 10), производя все необходимые проверки, а затем отвечает игроку: «Вы на (11, 10)»:

		*/

		if msg.Event == "InitGame" {
			mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
			user := usersGlobalWs[ws]

			// определяем положение отряда на пиксельной сети, через глобальную гексову сеть
			if user.GetSquad().R%2 != 0 {
				user.GetSquad().GlobalX = HexagonWidth + (HorizontalOffset * user.GetSquad().Q)
			} else {
				user.GetSquad().GlobalX = HexagonWidth/2 + (HorizontalOffset * user.GetSquad().Q)
			}
			user.GetSquad().GlobalY = HexagonHeight/2 + (user.GetSquad().R * VerticalOffset)

			if find && user != nil {
				// todo отдача других игроков
				ws.WriteJSON(Message{
					Event: msg.Event,
					Map:   mp, User: user,
					Squad: user.GetSquad(),
					Bases: bases.Bases.GetBasesByMap(usersGlobalWs[ws].GetSquad().MapID),
				})
			} else {
				ws.WriteJSON(Message{Event: "Error", Error: "map not find"})
			}
		}

		if msg.Event == "MoveTo" {
			// todo игрок двигает мс в точку ToX ToY
		}

		if msg.Event == "IntoToBase" {
			// todo игрок заходит на базу
		}
	}
}
