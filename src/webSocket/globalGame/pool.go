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

type Message struct {
	Event string             `json:"event"`
	Map   *_map.Map          `json:"map"`
	Error string             `json:"error"`
	Squad *squad.Squad       `json:"squad"`
	User  *player.Player     `json:"user"`
	Bases map[int]*base.Base `json:"bases"`
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

					1) реал тайм рпг каждый клик игрока МГНОВЕННО просчитывается на бекенде (сложно, небезопасно)

					2) псевдопошаговая рпг это когда за ход береться не готовность игроков, а время. Например 1 секунда
					игрок тыкнул куда хочет идти, дожидается пока закончится текущий ход (меньше секунды), начинается новый ход
					юнит проходит ровно столько пикселей за ход, сколько ему позволяет скорость, начинается новый ход юнит
					продолжает движение. (безопасно, сложно сделать движения бесшовными)

			Сервер знает что игрок находится в позиции (10, 10); клиент говорит: «Я хочу подвинуться на единицу вправо».
			Сервер обновляет позицию игрока на (11, 10), производя все необходимые проверки, а затем отвечает игроку: «Вы на (11, 10)»:

		*/

		if msg.Event == "InitGame" {
			mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
			user := usersGlobalWs[ws]
			// todo отдача других игроков
			if find && user != nil {
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
			// todo игрок двигает мс в точку
		}

		if msg.Event == "IntoToBase" {
			// todo игрок заходит на базу
		}
	}
}
