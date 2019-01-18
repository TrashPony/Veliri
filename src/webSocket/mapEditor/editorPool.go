package mapEditor

import (
	"../../mechanics/db/get"
	"../../mechanics/db/mapEditor"
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/coordinate"
	gameMap "../../mechanics/gameObjects/map"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"../../mechanics/factories/maps"
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
	Speed  int `json:"speed"`

	Radius int `json:"radius"`

	CountSprites int `json:"count_sprites"`
	XSize        int `json:"x_size"`
	YSize        int `json:"y_size"`
	XOffset      int `json:"x_offset"`
	YOffset      int `json:"y_offset"`
}

type Response struct {
	Event           string                   `json:"event"`
	Map             gameMap.Map              `json:"map"`
	Maps            map[int]gameMap.Map      `json:"maps"`
	TypeCoordinates []*coordinate.Coordinate `json:"type_coordinates"`
	Success         bool                     `json:"success"`
	Bases           map[int]*base.Base       `json:"bases"`
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

			mapChange, _ := maps.Maps.GetByID(msg.ID)

			coordinateMap, ok := mapChange.GetCoordinate(msg.Q, msg.R)
			if ok {
				coordinateMap.Level++
			}

			go mapEditor.ChangeHeightCoordinate(msg.ID, msg.Q, msg.R, 1)
			selectMap(msg, ws)
		}

		if msg.Event == "subtractHeightCoordinate" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)

			coordinateMap, ok := mapChange.GetCoordinate(msg.Q, msg.R)
			if ok {
				coordinateMap.Level--
			}

			go mapEditor.ChangeHeightCoordinate(msg.ID, msg.Q, msg.R, -1)
			selectMap(msg, ws)
		}

		if msg.Event == "placeCoordinate" {
			//TODO
			mapEditor.PlaceCoordinate(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeTerrain" {
			//TODO
			mapEditor.PlaceTerrain(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeObjects" {
			//TODO
			mapEditor.PlaceObject(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "placeAnimate" {
			//TODO
			mapEditor.PlaceAnimate(msg.ID, msg.IDType, msg.Q, msg.R)
			selectMap(msg, ws)
		}

		if msg.Event == "loadNewTypeTerrain" {
			//TODO
			ws.WriteJSON(Response{Event: "loadNewTypeTerrain", Success: mapEditor.CreateNewTerrain(msg.TerrainName)})
		}

		if msg.Event == "loadNewTypeObject" {
			//TODO
			success := mapEditor.CreateNewObject(msg.ObjectName, msg.AnimateName, msg.Move, msg.Watch,
				msg.Attack, msg.Radius, msg.Scale, msg.Shadow, msg.CountSprites, msg.XSize, msg.YSize)
			ws.WriteJSON(Response{Event: "loadNewTypeObject", Success: success})
		}

		if msg.Event == "deleteType" {
			//TODO
			mapEditor.DeleteType(msg.IDType)
			selectMap(msg, ws)
		}

		if msg.Event == "replaceCoordinateType" {
			//TODO
			mapEditor.ReplaceType(msg.OldIDType, msg.NewIDType)
			selectMap(msg, ws)
		}

		if msg.Event == "changeType" {
			//TODO
			mapEditor.ChangeType(msg.IDType, msg.Scale, msg.Shadow, msg.Move, msg.Watch, msg.Attack, msg.Radius)
			selectMap(msg, ws)
		}

		if msg.Event == "rotateObject" {
			//TODO
			mapEditor.Rotate(msg.ID, msg.Q, msg.R, msg.Rotate, msg.Speed, msg.XOffset, msg.YOffset)
			selectMap(msg, ws)
		}

		// ---------------------------- //
		if msg.Event == "addStartRow" {
			//TODO
			mapEditor.AddStartRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "addEndRow" {
			//TODO
			mapEditor.AddEndRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeStartRow" {
			//TODO
			mapEditor.RemoveStartRow(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeEndRow" {
			//TODO
			mapEditor.RemoveEndRow(msg.ID)
			selectMap(msg, ws)
		}

		// ---------------------------- //
		if msg.Event == "addStartColumn" {
			//TODO
			mapEditor.AddStartColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "addEndColumn" {
			//TODO
			mapEditor.AddEndColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeStartColumn" {
			//TODO
			mapEditor.RemoveStartColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "removeEndColumn" {
			//TODO
			mapEditor.RemoveEndColumn(msg.ID)
			selectMap(msg, ws)
		}

		if msg.Event == "addOverTexture" {
			//TODO
		}

		if msg.Event == "removeOverTexture" {
			//TODO
		}
	}
}
