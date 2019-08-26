package map_editor

import (
	bdMap "github.com/TrashPony/Veliri/src/mechanics/db/maps"
	"github.com/TrashPony/Veliri/src/mechanics/db/maps/mapEditor"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	gameMap "github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
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

	ToQ         int    `json:"to_q"`
	ToR         int    `json:"to_r"`
	ToBaseID    int    `json:"to_base_id"`
	ToMapID     int    `json:"to_map_id"`
	TypeHandler string `json:"type_handler"`

	IDType int `json:"id_type"`

	NewIDType int `json:"new_id_type"`
	OldIDType int `json:"old_id_type"`

	TerrainName string `json:"terrain_name"`
	ObjectName  string `json:"object_name"`
	AnimateName string `json:"animate_name"`
	TextureName string `json:"texture_name"`

	Move   bool `json:"move"`
	Watch  bool `json:"watch"`
	Attack bool `json:"attack"`

	Scale  int  `json:"scale"`
	Shadow bool `json:"shadow"`

	Rotate int `json:"rotate"`
	Speed  int `json:"speed"`

	Radius int `json:"radius"`

	XShadowOffset   int `json:"x_shadow_offset"`
	YShadowOffset   int `json:"y_shadow_offset"`
	ShadowIntensity int `json:"shadow_intensity"`

	CountSprites int    `json:"count_sprites"`
	XSize        int    `json:"x_size"`
	YSize        int    `json:"y_size"`
	XOffset      int    `json:"x_offset"`
	YOffset      int    `json:"y_offset"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	ToX          int    `json:"to_x"`
	ToY          int    `json:"to_y"`
	Color        string `json:"color"`
}

type Response struct {
	Event           string                        `json:"event"`
	Map             gameMap.Map                   `json:"map"`
	Maps            map[int]*gameMap.ShortInfoMap `json:"maps"`
	TypeCoordinates []*coordinate.Coordinate      `json:"type_coordinates"`
	Success         bool                          `json:"success"`
	Bases           map[int]*base.Base            `json:"bases"`
	Error           string                        `json:"error"`
	EntryToSector   []*coordinate.Coordinate      `json:"entry_to_sector"`
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
			getMapList(msg, ws) // +
		}

		if msg.Event == "SelectMap" {
			selectMap(msg, ws) // +
		}

		if msg.Event == "getAllTypeCoordinate" { // +
			ws.WriteJSON(Response{Event: msg.Event, TypeCoordinates: bdMap.AllTypeCoordinate()})
		}

		if msg.Event == "addHeightCoordinate" {
			heightCoordinate(msg, ws, 1) // +
		}

		if msg.Event == "subtractHeightCoordinate" { // +
			heightCoordinate(msg, ws, -1)
		}

		if msg.Event == "placeCoordinate" || msg.Event == "placeTerrain" || msg.Event == "placeObjects" || msg.Event == "placeAnimate" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)

			if msg.Event == "placeCoordinate" || msg.Event == "placeTerrain" {
				coordinateMap.TexturePriority = mapChange.GetMaxPriorityTexture()
				coordinateMap.TexturePriority++
			}

			if msg.Event == "placeCoordinate" || msg.Event == "placeObjects" || msg.Event == "placeAnimate" {
				coordinateMap.ObjectPriority = mapChange.GetMaxPriorityObject()
				coordinateMap.ObjectPriority++
			}

			mapEditor.PlaceCoordinate(coordinateMap, mapChange, msg.IDType)
		}

		if msg.Event == "loadNewTypeTerrain" {
			//ws.WriteJSON(Response{Event: "loadNewTypeTerrain", Success: mapEditor.CreateNewTerrain(msg.TerrainName)})
		}

		if msg.Event == "loadNewTypeObject" {
			//success := mapEditor.CreateNewObject(msg.ObjectName, msg.AnimateName, msg.Move, msg.Watch, msg.Attack, msg.CountSprites, msg.XSize, msg.YSize)
			//ws.WriteJSON(Response{Event: "loadNewTypeObject", Success: success})
		}

		if msg.Event == "deleteType" {
			//mapEditor.DeleteType(msg.IDType)
			//selectMap(msg, ws)
		}

		if msg.Event == "replaceCoordinateType" {
			//mapEditor.ReplaceType(msg.OldIDType, msg.NewIDType)
			//selectMap(msg, ws)
		}

		if msg.Event == "changeType" {
			mapEditor.ChangeType(msg.IDType, msg.Move, msg.Watch, msg.Attack)
			selectMap(msg, ws)
		}

		if msg.Event == "rotateObject" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)

			coordinateMap.ObjRotate = msg.Rotate
			coordinateMap.AnimationSpeed = msg.Speed
			coordinateMap.XOffset = msg.XOffset
			coordinateMap.YOffset = msg.YOffset
			coordinateMap.XShadowOffset = msg.XShadowOffset
			coordinateMap.YShadowOffset = msg.YShadowOffset
			coordinateMap.ShadowIntensity = msg.ShadowIntensity
			coordinateMap.Scale = msg.Scale
			coordinateMap.Shadow = msg.Shadow

			mapEditor.UpdateMapCoordinate(coordinateMap, mapChange)
		}

		// TODO ---------------------------- //
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
		// TODO --------------------------- //

		if msg.Event == "addOverTexture" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.PlaceTextures(coordinateMap, mapChange, msg.TextureName)
		}

		if msg.Event == "removeOverTexture" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.RemoveTextures(coordinateMap, mapChange)
			selectMap(msg, ws)
		}

		if msg.Event == "addTransport" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.PlaceTransport(coordinateMap, mapChange)
			selectMap(msg, ws)
		}

		if msg.Event == "removeTransport" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.RemoveTransport(coordinateMap, mapChange)
			selectMap(msg, ws)
		}

		if msg.Event == "addHandler" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.PlaceHandler(coordinateMap, mapChange, msg.ToQ, msg.ToR, msg.ToBaseID, msg.ToMapID, msg.TypeHandler)
			selectMap(msg, ws)
		}

		if msg.Event == "removeHandler" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)
			mapEditor.RemoveHandler(coordinateMap, mapChange)
			selectMap(msg, ws)
		}

		if msg.Event == "addGeoData" {
			mapEditor.AddGeoData(msg.X, msg.Y, msg.Radius, msg.ID)
		}

		if msg.Event == "removeGeoData" {
			mapEditor.RemoveGeoData(msg.ID)
		}

		if msg.Event == "addBeam" {
			mapEditor.AddBeam(msg.X, msg.Y, msg.ToX, msg.ToY, msg.ID, msg.Color)
		}

		if msg.Event == "removeBeam" {

		}

		if msg.Event == "toBack" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)

			coordinateMap.TexturePriority = 0
			coordinateMap.ObjectPriority = 0

			mapEditor.UpdateMapCoordinate(coordinateMap, mapChange)
			selectMap(msg, ws)
		}

		if msg.Event == "toFront" {
			mapChange, _ := maps.Maps.GetByID(msg.ID)
			coordinateMap, _ := mapChange.GetCoordinate(msg.Q, msg.R)

			coordinateMap.TexturePriority = mapChange.GetMaxPriorityTexture()
			coordinateMap.TexturePriority++

			coordinateMap.ObjectPriority = mapChange.GetMaxPriorityObject()
			coordinateMap.ObjectPriority++

			mapEditor.UpdateMapCoordinate(coordinateMap, mapChange)
		}
	}
}
