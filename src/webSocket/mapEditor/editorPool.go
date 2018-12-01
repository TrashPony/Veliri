package mapEditor

import (
	"../../mechanics/db/get"
	"../../mechanics/db/mapEditor"
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

	NewIDType int `json:"new_id_type"`
	OldIDType int `json:"old_id_type"`

	TerrainName string `json:"terrain_name"`
	ObjectName  string `json:"object_name"`
	AnimateName string `json:"animate_name"`

	Move   bool `json:"move"`
	Watch  bool `json:"watch"`
	Attack bool `json:"attack"`

	Scale  int  `json:"scale"`
	Shadow bool `json:"shadow"`

	Rotate int `json:"rotate"`

	Radius int `json:"radius"`

	CountSprites int `json:"count_sprites"`
	XSize        int `json:"x_size"`
	YSize        int `json:"y_size"`
}

type Response struct {
	Event           string                   `json:"event"`
	Map             gameMap.Map              `json:"map"`
	Maps            []gameMap.Map            `json:"maps"`
	TypeCoordinates []*coordinate.Coordinate `json:"type_coordinates"`
	Success         bool                     `json:"success"`
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
			ws.WriteJSON(Response{Event: msg.Event, TypeCoordinates: get.AllTypeCoordinate()})
		}

		if msg.Event == "addHeightCoordinate" {
			mapEditor.ChangeHeightCoordinate(msg.ID, msg.Q, msg.R, 1)
			selectMap(msg, ws)
		}

		if msg.Event == "subtractHeightCoordinate" {
			mapEditor.ChangeHeightCoordinate(msg.ID, msg.Q, msg.R, -1)
			selectMap(msg, ws)
		}

		if msg.Event == "placeCoordinate" {
			mapEditor.PlaceCoordinate(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeTerrain" {
			mapEditor.PlaceTerrain(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeObjects" {
			mapEditor.PlaceObject(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeAnimate" {
			mapEditor.PlaceAnimate(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "loadNewTypeTerrain" {
			ws.WriteJSON(Response{Event: "loadNewTypeTerrain", Success: mapEditor.CreateNewTerrain(msg.TerrainName)})
		}

		if msg.Event == "loadNewTypeObject" {
			success := mapEditor.CreateNewObject(msg.ObjectName, msg.AnimateName, msg.Move, msg.Watch,
				msg.Attack, msg.Radius, msg.Scale, msg.Shadow, msg.CountSprites, msg.XSize, msg.YSize)
			ws.WriteJSON(Response{Event: "loadNewTypeObject", Success: success})
		}

		if msg.Event == "deleteType" {
			mapEditor.DeleteType(msg.IDType)
			selectMap(msg, ws)
		}

		if msg.Event == "replaceCoordinateType" {
			mapEditor.ReplaceType(msg.OldIDType, msg.NewIDType)
			selectMap(msg, ws)
		}

		if msg.Event == "changeType" {
			mapEditor.ChangeType(msg.IDType, msg.Scale, msg.Shadow, msg.Move, msg.Watch, msg.Attack)
			selectMap(msg, ws)
		}

		if msg.Event == "rotateObject" {
			mapEditor.Rotate(msg.ID, msg.Q, msg.R, msg.Rotate)
			selectMap(msg, ws)
		}

		// ---------------------------- //
		if msg.Event == "addStartRow" {
			mapEditor.AddStartRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "addEndRow" {
			mapEditor.AddEndRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeStartRow" {
			mapEditor.RemoveStartRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeEndRow" {
			mapEditor.RemoveEndRow(msg.ID)
			selectMap(msg, ws)
		}

		// ---------------------------- //
		if msg.Event == "addStartColumn" {
			mapEditor.AddStartColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "addEndColumn" {
			mapEditor.AddEndColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeStartColumn" {
			mapEditor.RemoveStartColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeEndColumn" {
			mapEditor.RemoveEndColumn(msg.ID)
			selectMap(msg, ws)
		}
	}
}
