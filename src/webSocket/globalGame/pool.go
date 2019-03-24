package globalGame

import (
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
	// когда я забил х на эту структуру данных а теперь тут какое то адище
	IDSender      int
	IDUserSend    int
	IDMap         int
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
	OtherUser     *player.ShortUserInfo           `json:"other_user"`
	OtherUsers    []*player.ShortUserInfo         `json:"other_users"`
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
	Anomaly       *_map.Anomalies                 `json:"anomaly"`
	DynamicObject *dynamicMapObject.DynamicObject `json:"dynamic_object"`
	BoxPassword   int                             `json:"box_password"`
	Reservoir     *resource.Map                   `json:"reservoir"`
	Cloud         *Cloud                          `json:"cloud"`
	ToSquadID     string                          `json:"to_squad_id"`
	Bot           bool                            `json:"bot"`
}

type Cloud struct {
	Name     string  `json:"name"`
	Speed    int     `json:"speed"`
	Alpha    float64 `json:"alpha"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	Angle    int     `json:"angle"`
	Uuid     string  `json:"uuid"`
	SizeMapX int     `json:"-"`
	SizeMapY int     `json:"-"`
	IDMap    int     `json:"-"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	globalGame.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента

	println("WS global Сессия: login: " + login + " id: " + strconv.Itoa(id))
	go Reader(ws, newPlayer)

}

func SendMessage(msg Message) {
	globalPipe <- msg
}

/*
если случается дедлок то скорее всего кто то закрыл соеденение а мьютекс там закрывается не через девер и он остается закрытым
ненадо так

UPD нельзя делать дефер в функции которыя шлет в pipe месаги, это верный путь к дедлоку т.к. на другом конце канала опять происходит лок
	Выносить все где есть defer .Unlock в отдельные функции
*/

func Reader(ws *websocket.Conn, user *player.Player) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			println(err.Error())
			go DisconnectUser(globalGame.Clients.GetByWs(ws), ws, false)
			return
		}

		// если игрок на базе или в локальной игре то ему нельзя поднимать соеденение глобальной игры
		if user.InBaseID != 0 || user.GetSquad().InGame {
			// todo перенаправлять на нужный сервис а не тупо кикать
			DisconnectUser(user, ws, false)
		}

		if msg.Event == "InitGame" {
			LoadGame(ws, msg)
		}

		if msg.Event == "MoveTo" {
			go Move(ws, msg)
		}

		if msg.Event == "StopMove" {
			stopMove(globalGame.Clients.GetByWs(ws), true)
		}

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

		if msg.Event == "Attack" {
			startLocalGame(ws, msg)
		}

		if msg.Event == "OpenDialog" {

		}

		if msg.Event == "Ask" {

		}
	}
}

func MoveSender() {
	for {
		select {
		case resp := <-globalPipe:

			usersGlobalWs, rLock := globalGame.Clients.GetAllConnects()
			for ws, client := range usersGlobalWs {
				// боты не получают сообщений они и так все знают
				if client.Bot || resp.Bot || ws == nil {
					continue
				}

				var err error

				// получают все кроме отправителя в пределах карты
				if client.ID != resp.IDSender && resp.IDSender > 0 && resp.IDUserSend == 0 && client.MapID == resp.IDMap {
					err = ws.WriteJSON(resp)
				}

				// получает только отправитель
				if client.ID == resp.IDUserSend && resp.IDUserSend > 0 && resp.IDSender == 0 {
					err = ws.WriteJSON(resp)
				}

				// получают все в пределах карты
				if resp.IDSender == 0 && resp.IDUserSend == 0 && client.MapID == resp.IDMap {
					err = ws.WriteJSON(resp)
				}

				if err != nil {
					go DisconnectUser(globalGame.Clients.GetByWs(ws), ws, false)
					log.Printf("error: %v", err)
				}
			}
			rLock.Unlock()
		default:
		}
	}
}
