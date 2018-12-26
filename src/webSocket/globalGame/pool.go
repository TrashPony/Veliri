package globalGame

import (
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/base"
	"../../mechanics/gameObjects/box"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/gameObjects/map"
	"../../mechanics/gameObjects/squad"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersGlobalWs = make(map[*websocket.Conn]*player.Player)
var globalPipe = make(chan Message)

type Message struct {
	idSender      int
	idUserSend    int
	Event         string                      `json:"event"`
	Map           *_map.Map                   `json:"map"`
	Error         string                      `json:"error"`
	Squad         *squad.Squad                `json:"squad"`
	User          *player.Player              `json:"user"`
	Bases         map[int]*base.Base          `json:"bases"`
	X             int                         `json:"x"`
	Y             int                         `json:"y"`
	ToX           float64                     `json:"to_x"`
	ToY           float64                     `json:"to_y"`
	PathUnit      globalGame.PathUnit         `json:"path_unit"`
	Path          []globalGame.PathUnit       `json:"path"`
	BaseID        int                         `json:"base_id"`
	OtherUser     *hostileMS                  `json:"other_user"`
	OtherUsers    []*hostileMS                `json:"other_users"`
	ThrowItems    []inventory.Slot            `json:"throw_items"`
	Boxes         []*box.Box                  `json:"boxes"`
	Box           *box.Box                    `json:"box"`
	BoxID         int                         `json:"box_id"`
	Slot          int                         `json:"slot"`
	Size          float32                     `json:"size"`
	Inventory     *inventory.Inventory        `json:"inventory"`
	InventorySlot int                         `json:"inventory_slot"`
	TransportID   int                         `json:"transport_id"`
	ThoriumSlots  map[int]*detail.ThoriumSlot `json:"thorium_slots"`
	ThoriumSlot   int                         `json:"thorium_slot"`
	Afterburner   bool                        `json:"afterburner"`
}

type hostileMS struct {
	// структура которая описываем минимальный набор данных для отображение и взаимодействия,
	// что бы другие игроки не палили трюмы, фиты и дронов без спец оборудования
	SquadID    int    `json:"squad_id"`
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

	hostile.SquadID = user.GetSquad().ID
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
	evacuation := false

	for {

		if evacuation {
			continue
		}

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

		if msg.Event == "evacuation" {
			evacuationSquad(ws, msg, stopMove, &moveChecker, &evacuation)
		}

		if msg.Event == "updateThorium" {
			updateThorium(ws, msg, stopMove, &moveChecker)
		}

		if msg.Event == "removeThorium" {
			removeThorium(ws, msg, stopMove, &moveChecker)
		}

		if msg.Event == "AfterburnerToggle" {
			afterburnerToggle(ws, msg, stopMove, &moveChecker)
		}
	}
}

func MoveSender() {
	for {
		resp := <-globalPipe
		mutex.Lock()
		for ws, client := range usersGlobalWs {

			var err error

			// получают все кроме отправителя
			if client.GetID() != resp.idSender && resp.idSender > 0 && resp.idUserSend == 0 {
				err = ws.WriteJSON(resp)
			}

			// получает только отправитель
			if client.GetID() == resp.idUserSend && resp.idUserSend > 0 && resp.idSender == 0 {
				err = ws.WriteJSON(resp)
			}

			// получают все
			if resp.idSender == 0 && resp.idUserSend == 0 {
				err = ws.WriteJSON(resp)
			}

			if err != nil {
				DisconnectUser(usersGlobalWs[ws])
				log.Printf("error: %v", err)
				ws.Close()
				delete(usersGlobalWs, ws)
			}
		}
		mutex.Unlock()
	}
}