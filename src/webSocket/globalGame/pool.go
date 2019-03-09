package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamicMapObject"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

var globalPipe = make(chan Message, 1)

type Message struct {
	idSender      int
	idUserSend    int
	idMap         int
	Event         string                          `json:"event"`
	Map           *_map.Map                       `json:"map"`
	Error         string                          `json:"error"`
	Squad         *squad.Squad                    `json:"squad"`
	User          *player.Player                  `json:"user"`
	Bases         map[int]*base.Base              `json:"bases"`
	X             int                             `json:"x"`
	Y             int                             `json:"y"`
	Q             int                             `json:"q"`
	R             int                             `json:"r"`
	ToX           float64                         `json:"to_x"`
	ToY           float64                         `json:"to_y"`
	PathUnit      squad.PathUnit                  `json:"path_unit"`
	Path          []squad.PathUnit                `json:"path"`
	BaseID        int                             `json:"base_id"`
	OtherUser     *hostileMS                      `json:"other_user"`
	OtherUsers    []*hostileMS                    `json:"other_users"`
	ThrowItems    []inventory.Slot                `json:"throw_items"`
	Boxes         []*boxInMap.Box                 `json:"boxes"`
	Box           *boxInMap.Box                   `json:"box"`
	BoxID         int                             `json:"box_id"`
	ToBoxID       int                             `json:"to_box_id"`
	TypeSlot      int                             `json:"type_slot"`
	Name          string                          `json:"name"`
	Slot          int                             `json:"slot"`
	Slots         []int                           `json:"slots"`
	Size          float32                         `json:"size"`
	Inventory     *inventory.Inventory            `json:"inventory"`
	InventorySlot int                             `json:"inventory_slot"`
	TransportID   int                             `json:"transport_id"`
	ThoriumSlots  map[int]*detail.ThoriumSlot     `json:"thorium_slots"`
	ThoriumSlot   int                             `json:"thorium_slot"`
	Afterburner   bool                            `json:"afterburner"`
	Credits       int                             `json:"credits"`
	Experience    int                             `json:"experience"`
	Seconds       int                             `json:"seconds"`
	Count         int                             `json:"count"`
	Coordinates   []*coordinate.Coordinate        `json:"coordinates"`
	Radius        int                             `json:"radius"`
	Anomalies     []globalGame.VisibleAnomaly     `json:"anomalies"`
	DynamicObject *dynamicMapObject.DynamicObject `json:"dynamic_object"`
	BoxPassword   int                             `json:"box_password"`
	Reservoir     *resource.Map                   `json:"reservoir"`
	Cloud         *cloud                          `json:"cloud"`
	Bot           bool                            `json:"bot"`
}

type hostileMS struct {
	// структура которая описываем минимальный набор данных для отображение и взаимодействия,
	// что бы другие игроки не палили трюмы, фиты и дронов без спец оборудования
	SquadID    string       `json:"squad_id"`
	UserName   string       `json:"user_name"`
	X          int          `json:"x"`
	Y          int          `json:"y"`
	Q          int          `json:"q"`
	R          int          `json:"r"`
	BodyName   string       `json:"body_name"`
	WeaponName string       `json:"weapon_name"`
	Rotate     int          `json:"rotate"`
	Body       *detail.Body `json:"body"`
}

func GetShortUserInfo(user *player.Player) *hostileMS {
	var hostile hostileMS

	if user == nil || user.GetSquad() == nil || user.GetSquad().MatherShip == nil || user.GetSquad().MatherShip.Body == nil {
		return nil
	}

	if user.Bot {
		hostile.SquadID = user.UUID
	} else {
		hostile.SquadID = strconv.Itoa(user.GetSquad().ID)
	}

	hostile.UserName = user.GetLogin()
	hostile.X = user.GetSquad().GlobalX
	hostile.Y = user.GetSquad().GlobalY
	hostile.Q = user.GetSquad().Q
	hostile.R = user.GetSquad().R
	hostile.Rotate = user.GetSquad().MatherShip.Rotate
	hostile.BodyName = user.GetSquad().MatherShip.Body.Name

	hostile.Body, _ = gameTypes.Bodies.GetByID(user.GetSquad().MatherShip.Body.ID)

	if user.GetSquad().MatherShip.GetWeaponSlot() != nil && user.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {
		hostile.WeaponName = user.GetSquad().MatherShip.GetWeaponSlot().Weapon.Name
	}

	return &hostile
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	globalGame.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента

	print("WS global Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	go Reader(ws)
}

func sendMessage(msg Message) {
	globalPipe <- msg
}

/*
если случается дедлок то скорее всего кто то закрыл соеденение а мьютекс там закрывается не через девер и он остается закрытым
ненадо так

UPD нельзя делать дефер в функции которыя шлет в pipe месаги, это верный путь к дедлоку т.к. на другом конце канала опять происходит лок
	Выносить все где есть defer .Unlock в отдельные функции

// TODO еще раз проверить где используется Clients.GetAll()
*/

func Reader(ws *websocket.Conn) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			println(err.Error())
			go DisconnectUser(globalGame.Clients.GetByWs(ws), ws)
			return
		}

		if msg.Event == "InitGame" {
			loadGame(ws, msg)
		}

		if msg.Event == "MoveTo" {
			go move(ws, msg)
		}

		if msg.Event == "StopMove" {
			stopMove(globalGame.Clients.GetByWs(ws), true)
		}

		//if msg.Event == "IntoToBase" {
		//	intoToBase(ws, msg)
		//}

		if msg.Event == "ThrowItems" {
			throwItems(ws, msg)
		}

		if msg.Event == "openBox" {
			openBox(ws, msg)
		}

		if msg.Event == "placeNewBox" {
			placeNewBox(ws, msg)
		}

		if msg.Event == "getItemsFromBox" || msg.Event == "getItemFromBox" {
			useBox(ws, msg)
		}

		if msg.Event == "placeItemToBox" || msg.Event == "placeItemsToBox" {
			useBox(ws, msg)
		}

		if msg.Event == "boxToBoxItem" || msg.Event == "boxToBoxItems" {
			boxToBox(ws, msg)
		}

		if msg.Event == "evacuation" {
			evacuationSquad(ws)
		}

		if msg.Event == "updateThorium" {
			updateThorium(ws, msg)
		}

		if msg.Event == "removeThorium" {
			removeThorium(ws, msg)
		}

		if msg.Event == "AfterburnerToggle" {
			afterburnerToggle(ws, msg)
		}

		if msg.Event == "startMining" {
			startMining(ws, msg)
		}

		if msg.Event == "SelectDigger" {
			selectDigger(ws, msg)
		}

		if msg.Event == "useDigger" {
			useDigger(ws, msg)
		}

		if msg.Event == "OpenDialog" {

		}

		if msg.Event == "Ask" {

		}
	}
}

func MoveSender() {
	for {
		// TODO тут происходит деадлок, если при большом трафике постоянно жмакать ф5
		select {
		case resp := <-globalPipe:

			usersGlobalWs, rLock := globalGame.Clients.GetAllConnects()
			for ws, client := range usersGlobalWs {
				// боты не получают сообщений они и так все знают
				if client.Bot || resp.Bot || ws == nil {
					continue
				}

				var err error

				// получают все кроме отправителя
				if client.ID != resp.idSender && resp.idSender > 0 && resp.idUserSend == 0 && client.MapID == resp.idMap {
					err = ws.WriteJSON(resp)
				}

				// получает только отправитель
				if client.ID == resp.idUserSend && resp.idUserSend > 0 && resp.idSender == 0 && client.MapID == resp.idMap {
					err = ws.WriteJSON(resp)
				}

				// получают все
				if resp.idSender == 0 && resp.idUserSend == 0 && client.MapID == resp.idMap {
					err = ws.WriteJSON(resp)
				}

				if err != nil {
					go DisconnectUser(globalGame.Clients.GetByWs(ws), ws)
					log.Printf("error: %v", err)
					//panic(err)
				}
			}
			rLock.Unlock()
		default:
		}
	}
}
