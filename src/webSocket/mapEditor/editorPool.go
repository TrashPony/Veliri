package mapEditor

import (
	"../../mechanics/gameObjects/coordinate"
	gameMap "../../mechanics/gameObjects/map"
	"../../mechanics/player"
	"../../mechanics/players"
	"../utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersWs = make(map[*websocket.Conn]*player.Player)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS mapEditor Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	mutex.Unlock()

	Reader(ws)
}

type Message struct {
	Event string `json:"event"`
	ID    int    `json:"id"`
	Q     int    `json:"q"`
	R     int    `json:"r"`

	IDType int `json:"id_type"`

	TerrainName string `json:"terrain_name"`
	ObjectName  string `json:"object_name"`
	AnimateName string `json:"animate_name"`

	Move   bool `json:"move"`
	Watch  bool `json:"watch"`
	Attack bool `json:"attack"`

	Radius int `json:"radius"`
}

type Response struct {
	Event           string                   `json:"event"`
	Map             gameMap.Map              `json:"map"`
	Maps            []gameMap.Map            `json:"maps"`
	TypeCoordinates []*coordinate.Coordinate `json:"type_coordinates"`
	Error           string                   `json:"error"`
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			println(err.Error())
			utils.DelConn(ws, &usersWs, err)
			break
		}

		if msg.Event == "getMapList" {
			getMapList(msg, ws)
		}

		if msg.Event == "SelectMap" {
			selectMap(msg, ws)
		}

		if msg.Event == "getAllTypeCoordinate" {
			getAllCoordinate(msg, ws)
		}

		if msg.Event == "addHeightCoordinate" {
			addHeightCoordinate(msg, ws)
		}

		if msg.Event == "subtractHeightCoordinate" {
			subtractHeightCoordinate(msg, ws)
		}

		if msg.Event == "placeCoordinate" {

		}

		if msg.Event == "loadNewTypeCoordinate" {
			// TODO проверка на существование такого же файла что бы случайно не затереть старый
		}

		// ---------------------------- //
		if msg.Event == "addStartRow" {

		}

		if msg.Event == "addEndRow" {

		}

		if msg.Event == "addStartRow" {

		}

		if msg.Event == "removeEndRow" {

		}

		// ---------------------------- //
		if msg.Event == "addStartColumn" {

		}

		if msg.Event == "addEndColumn" {

		}

		if msg.Event == "removeStartColumn" {

		}

		if msg.Event == "removeEndColumn" {

		}
	}
}
