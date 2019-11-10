package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamicMapObject"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

var globalPipe = make(chan Message, 1)

type Message struct {
	// когда я забил х на эту структуру данных а теперь тут какое то адище
	IDSender      int                             // переменная не для данных а для отсылки сообщений
	IDUserSend    int                             // переменная не для данных а для отсылки сообщений
	IDMap         int                             // переменная не для данных а для отсылки сообщений
	Event         string                          `json:"event"`
	Map           *_map.Map                       `json:"map"`
	Error         string                          `json:"error"`
	Squad         *squad.Squad                    `json:"squad"`
	User          *player.Player                  `json:"user"`
	Bases         map[int]*base.Base              `json:"bases"`
	X             int                             `json:"x"`
	Y             int                             `json:"y"`
	ToX           float64                         `json:"to_x"`
	ToY           float64                         `json:"to_y"`
	Bullet        *unit.Bullet                    `json:"bullet"`
	PathUnit      *unit.PathUnit                  `json:"path_unit"`
	Path          []*unit.PathUnit                `json:"path"`
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
	MapID         int                             `json:"map_id"`
	Missions      map[string]*mission.Mission     `json:"missions"`
	MissionUUID   string                          `json:"mission_uuid"`
	UnitsID       []int                           `json:"units_id"`
	ShortUnits    map[int]*unit.ShortUnitInfo     `json:"short_units"`
	ShortUnit     *unit.ShortUnitInfo             `json:"short_unit"`
	Unit          *unit.Unit                      `json:"unit"`
	UnitID        int                             `json:"unit_id"`
	Equip         *detail.BodyEquipSlot           `json:"equip"`
	Weapon        *detail.Weapon                  `json:"weapon"`
	HighGravity   bool                            `json:"high_gravity"`
	Type          string                          `json:"type"`

	Color    string `json:"color"`
	RectSize int    `json:"rect_size"`

	Polygon collisions.Polygon `json:"polygon"`

	NeedCheckView bool                  `json:"-"` // если тру то надо проверить сообщение на дальность видимости/радара
	RadarMark     *squad.VisibleObjects `json:"radar_mark"`
	ActionObject  string                `json:"action_object"` // удалить/создать обьект
	ActionMark    string                `json:"action_mark"`   // удалить/создать/скрыть/раскрыть метку
	Object        interface{}           `json:"object"`        // обьект (ящик, транспорт, юнит и тд)
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

	if newPlayer.GetSquad() == nil {
		// значит игрок оказался не там ибо на глобалке невозмоно быть без отряда
		if newPlayer.LastBaseID > 0 {
			// отправляем его на ближаюшую базу
			IntoToBase(newPlayer, newPlayer.LastBaseID)
		} else {
			// TODO иначе на респаун его фракции
			IntoToBase(newPlayer, 1)
		}

		return
	}

	globalGame.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента

	println("WS global Сессия: login: " + login + " id: " + strconv.Itoa(id))
	go Reader(ws, newPlayer)

	if debug.Store.Move {
		DebugMoveWorker(newPlayer)
	}
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
			go DisconnectUser(globalGame.Clients.GetByWs(ws), false)
			return
		}

		// если игрок на базе то ему нельзя поднимать соеденение глобальной игры
		if user.InBaseID != 0 {

			if user.InBaseID != 0 {
				go SendMessage(Message{Event: "IntoToBase"})
			}

			DisconnectUser(user, false)
		}

		if msg.Event == "InitGame" {
			LoadGame(user, msg)
		}

		if msg.Event == "MoveTo" {
			go Move(user, msg, true)
		}

		if msg.Event == "PlaceUnit" {
			placeUnit(user, msg)
		}

		if msg.Event == "StopMove" {
			for _, id := range msg.UnitsID {
				stopMove(globalGame.Clients.GetUnitByID(id), true)
			}
		}

		if msg.Event == "ThrowItems" {
			throwItems(user, msg)
		}

		if msg.Event == "openBox" {
			openBox(user, msg)
		}

		if msg.Event == "placeNewBox" {
			placeNewBox(user, msg)
		}

		if msg.Event == "getItemsFromBox" || msg.Event == "getItemFromBox" {
			useBox(user, msg)
		}

		if msg.Event == "placeItemToBox" || msg.Event == "placeItemsToBox" {
			useBox(user, msg)
		}

		if msg.Event == "boxToBoxItem" || msg.Event == "boxToBoxItems" {
			boxToBox(user, msg)
		}

		if msg.Event == "evacuation" {
			evacuationUnit(user.GetSquad().MatherShip) // игрок может инициализировать эвакуацию только МС
		}

		if msg.Event == "updateThorium" {
			updateThorium(user, msg)
		}

		if msg.Event == "removeThorium" {
			removeThorium(user, msg)
		}

		if msg.Event == "AfterburnerToggle" {
			afterburnerToggle(user, msg)
		}

		if msg.Event == "startMining" {
			startMining(globalGame.Clients.GetUnitByID(msg.UnitID), msg, user)
		}

		if msg.Event == "SelectEquip" {
			SelectEquip(user, msg)
		}

		if msg.Event == "useDigger" {
			useDigger(user, msg)
		}

		if msg.Event == "Attack" {
			Attack(user, msg)
		}

		if msg.Event == "GetMissions" {
			go SendMessage(Message{Event: msg.Event, IDUserSend: user.GetID(), Missions: user.Missions, MissionUUID: user.SelectMission})
		}

		if msg.Event == "NewFormationPos" {
			NewFormationPos(user, msg)
		}

		if msg.Event == "SelectMission" {

			if msg.MissionUUID == "" {
				user.SelectMission = msg.MissionUUID
			} else {
				_, ok := user.Missions[msg.MissionUUID]
				if ok {
					user.SelectMission = msg.MissionUUID
				}
			}
			go SendMessage(Message{Event: "GetMissions", IDUserSend: user.GetID(), Missions: user.Missions, MissionUUID: user.SelectMission})
		}

		if msg.Event == "GetPortalPointToGlobalPath" {
			_, transitionPoints := maps.Maps.FindGlobalPath(user.GetSquad().MatherShip.MapID, msg.MapID)
			if len(transitionPoints) > 0 {
				go SendMessage(Message{
					Event:      "GetPortalPointToGlobalPath",
					IDUserSend: user.GetID(),
					Name:       msg.Name,
					X:          transitionPoints[0].X,
					Y:          transitionPoints[0].Y,
				})
			}
		}
	}
}

var i = 0

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
					if resp.NeedCheckView {
						// фильтр по дальности видимости
						newResp := CheckView(globalGame.Clients.GetByWs(ws), &resp)
						if newResp != nil {
							err = ws.WriteJSON(newResp)
						}
					} else {
						err = ws.WriteJSON(resp)
					}
				}

				if err != nil {
					go DisconnectUser(globalGame.Clients.GetByWs(ws), false)
					log.Printf("error global: %v", err)
				}
			}
			rLock.Unlock()
		default:
		}
	}
}
