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
	"math"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

var usersGlobalWs = make(map[*websocket.Conn]*player.Player)

const HexagonHeight = 111 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 100
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

type Message struct {
	Event string              `json:"event"`
	Map   *_map.Map           `json:"map"`
	Error string              `json:"error"`
	Squad *squad.Squad        `json:"squad"`
	User  *player.Player      `json:"user"`
	Bases map[int]*base.Base  `json:"bases"`
	ToX   float64             `json:"to_x"`
	ToY   float64             `json:"to_y"`
	Path  globalGame.PathUnit `json:"path"`
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

			user := usersGlobalWs[ws]

			//path := make([]globalGame.PathUnit, 0)

			forecastX := float64(user.GetSquad().GlobalX)
			forecastY := float64(user.GetSquad().GlobalY)
			speed := user.GetSquad().MatherShip.Speed * 3
			rotate := user.GetSquad().MatherShip.Rotate

			diffRotate := 0 // разница между углом цели и носа корпуса
			dist := 900.0

			for {
				dist = math.Sqrt(((forecastX - msg.ToX) * (forecastX - msg.ToX)) + ((forecastY - msg.ToY) * (forecastY - msg.ToY)))
				if dist < 10 {
					break
				}

				radRotate := float64(rotate) * math.Pi / 180
				stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
				stopY := float64(speed) * math.Sin(radRotate)

				forecastX = forecastX + stopX // находим новое положение корпуса на глобальной карте
				forecastY = forecastY + stopY

				for i := 0; i < speed; i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
					needRad := math.Atan2(msg.ToY-forecastY, msg.ToX-forecastX)
					needRotate := int(needRad * 180 / 3.14) // находим какой угол необходимо принять телу

					newRotate := globalGame.RotateUnit(&rotate, &needRotate)

					if rotate >= needRotate {
						diffRotate = rotate - needRotate
					} else {
						diffRotate = needRotate - rotate
					}

					if diffRotate != 0 { // если разница есть то поворачиваем корпус
						rotate += newRotate
					} else {
						break
					}
				}

				ws.WriteJSON(Message{Event: msg.Event, Path: globalGame.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100}})

				user.GetSquad().MatherShip.Rotate = rotate
				user.GetSquad().GlobalX = int(forecastX)
				user.GetSquad().GlobalY = int(forecastY)

				time.Sleep(100 * time.Millisecond)
			}

		}

		if msg.Event == "IntoToBase" {
			// todo игрок заходит на базу
		}
	}
}
