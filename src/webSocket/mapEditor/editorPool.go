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
			// TODO  увеличить высоту координате
			addHeightCoordinate(msg, ws)
		}

		if msg.Event == "subtractHeightCoordinate" {
			// TODO уменьшить высоту координате
			subtractHeightCoordinate(msg, ws)
		}

		if msg.Event == "placeCoordinate" {
			// TODO  просмот замена координаты на тип в id
		}

		if msg.Event == "placeTerrain" {
			// TODO взять терейн из координаты по id, взятие координа на карте если там есть обьект
			// TODO посмотреть есть ли в бд такая координата земля+обьект, если есть ставим ее, если нет создаем новую и ставим ее
			// TODO все мето данные радиуса, разрешений берем из координаты на которой стоит обьект (обьект всезда задает все параметры)
		}

		if msg.Event == "placeObjects" {
			// TODO взять обьект из координаты по id, взять координату на карте и взять от туда терейн
			// TODO посмотреть есть ли в бд такая координата земля+обьект, если есть ставим ее, если нет создаем новую и ставим ее
			// TODO все мето данные радиуса, разрешений берем из координаты на которой стоит обьект (обьект всезда задает все параметры)
		}

		if msg.Event == "placeAnimate" {
			// TODO взять анимацию из координаты по id, взять координату на карте и взять от туда терейн
			// TODO посмотреть есть ли в бд такая координата земля+анимацию, если есть ставим ее, если нет создаем новую и ставим ее
			// TODO все мето данные радиуса, разрешений берем из координаты на которой стоит анимацию (анимацию всезда задает все параметры)
		}

		if msg.Event == "loadNewTypeCoordinate" {
			// TODO проверка на существование такого же файла что бы случайно не затереть старый
		}

		// ---------------------------- //
		if msg.Event == "addStartRow" {
			// todo перенести все обьекты +1 к R
			// todo добавить строку
		}

		if msg.Event == "addEndRow" {
			// todo просто увеличиваем высоту карты на 1
		}

		if msg.Event == "removeStartRow" {
			// todo удаляем все координаты из таблицы конструктор карты по первой строке,
			// todo смещаем все обьекты вниз на -1 R,
			// todo уменьшаем высоту карты на 1
		}

		if msg.Event == "removeEndRow" {
			// todo удаляем все координаты из таблицы конструктор карты по поледней строке,
			// todo уменьшаем высоту карты на 1
		}

		// ---------------------------- //
		if msg.Event == "addStartColumn" {
			// todo перенести все обьекты +1 к Q
			// todo добавить столбец
		}

		if msg.Event == "addEndColumn" {
			// todo просто увеличиваем длинну карты на 1
		}

		if msg.Event == "removeStartColumn" {
			// todo удаляем все координаты из таблицы конструктор карты по первому столбцу,
			// todo смещаем все обьекты вправо на +1 Q,
			// todo уменьшаем длинну карты на 1
		}

		if msg.Event == "removeEndColumn" {
			// todo удаляем все координаты из таблицы конструктор карты по поледнему столбцу,
			// todo уменьшаем высоту карты на 1
		}
	}
}
