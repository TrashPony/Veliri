package globalGame

import (
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/box"
	"../../mechanics/gameObjects/inventory"
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
	Event      string                `json:"event"`
	Map        *_map.Map             `json:"map"`
	Error      string                `json:"error"`
	Squad      *squad.Squad          `json:"squad"`
	User       *player.Player        `json:"user"`
	Bases      map[int]*base.Base    `json:"bases"`
	X          int                   `json:"x"`
	Y          int                   `json:"y"`
	ToX        float64               `json:"to_x"`
	ToY        float64               `json:"to_y"`
	PathUnit   globalGame.PathUnit   `json:"path_unit"`
	Path       []globalGame.PathUnit `json:"path"`
	BaseID     int                   `json:"base_id"`
	OtherUser  *hostileMS            `json:"other_user"`
	OtherUsers []*hostileMS          `json:"other_users"`
	ThrowItems []inventory.Slot      `json:"throw_items"`
	Boxes      []*box.Box            `json:"boxes"`
	Box        *box.Box              `json:"box"`
	BoxID      int                   `json:"box_id"`
	Slot       int                   `json:"slot"`
	Size       float32               `json:"size"`
	Inventory  *inventory.Inventory  `json:"inventory"`
}

type hostileMS struct {
	// структура которая описываем минимальный набор данных для отображение и взаимодействия,
	// что бы другие игроки не палили трюмы, фиты и дронов без спец оборудования
	UserName   string `json:"user_name"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Q          int    `json:"q"`
	R          int    `json:"r"`
	BodyName   string `json:"body_name"`
	WeaponName string `json:"weapon_name"`
	Rotate     int    `json:"rotate"`
}

func GetShortUserInfo(user *player.Player) *hostileMS {
	var hostile hostileMS

	hostile.UserName = user.GetLogin()
	hostile.X = user.GetSquad().GlobalX
	hostile.Y = user.GetSquad().GlobalY
	hostile.Q = user.GetSquad().Q
	hostile.R = user.GetSquad().R
	hostile.Rotate = user.GetSquad().MatherShip.Rotate
	hostile.BodyName = user.GetSquad().MatherShip.Body.Name

	if user.GetSquad().MatherShip.GetWeaponSlot() != nil && user.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {
		hostile.WeaponName = user.GetSquad().MatherShip.GetWeaponSlot().Weapon.Name
	}

	return &hostile
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
	moveChecker := false

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			println(err.Error())
			DisconnectUser(usersGlobalWs[ws])
			utils.DelConn(ws, &usersGlobalWs, err)
			break
		}

		if msg.Event == "InitGame" {
			loadGame(ws, msg)
		}

		if msg.Event == "MoveTo" {
			move(ws, msg, stopMove, &moveChecker)
		}

		if msg.Event == "IntoToBase" {
			intoToBase(ws, msg, stopMove, &moveChecker)
		}

		if msg.Event == "ThrowItems" {
			throwItems(ws, msg)
		}

		if msg.Event == "openBox" {
			openBox(ws, msg, stopMove, &moveChecker)
		}

		if msg.Event == "getItemFromBox" {
			useBox(ws, msg)
		}

		if msg.Event == "placeItemToBox" {
			useBox(ws, msg)
		}
	}
}
