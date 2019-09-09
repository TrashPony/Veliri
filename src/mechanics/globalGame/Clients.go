package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/gorilla/websocket"
	"sync"
)

type wsUsers struct {
	users     map[*websocket.Conn]*player.Player // карта игроков которые онлайн
	units     map[int]*unit.Unit                 // карта юнитов в игре (юнити и мсы)
	connects  map[*websocket.Conn]gameConnect    // специальная карта для быстрой отправки сообщений.
	usersMX   sync.RWMutex
	connectMX sync.RWMutex
	unitsMX   sync.RWMutex
}

type gameConnect struct {
	ID    int  `json:"id"`
	Bot   bool `json:"bot"`
	MapID int  `json:"map_id"`
}

var Clients = NewClientsStore()

func NewClientsStore() *wsUsers {
	return &wsUsers{
		users:    make(map[*websocket.Conn]*player.Player),
		connects: make(map[*websocket.Conn]gameConnect),
		units:    make(map[int]*unit.Unit),
	}
}

func (c *wsUsers) AddNewClient(newWS *websocket.Conn, newClient *player.Player) {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	c.connectMX.Lock()
	defer c.connectMX.Unlock()

	c.unitsMX.Lock()
	defer c.unitsMX.Unlock()

	for ws, client := range c.users {
		if !client.Bot && client.GetID() == newClient.GetID() {
			delete(c.users, ws) // удаляем его из активных подключений
			ws.Close()
		}
	}

	c.users[newWS] = newClient
	c.connects[newWS] = gameConnect{ID: newClient.GetID(), Bot: newClient.Bot, MapID: newClient.GetSquad().MatherShip.MapID}

	// мазершип всегда сразу на карте
	if newClient.Bot {
		// у ботов отрицательныйх идиник
		newClient.GetSquad().MatherShip.ID = c.getBotID()
	}

	GetPlaceCoordinate(newClient.GetSquad().MatherShip, c.GetAllShortUnits(newClient.GetSquad().MatherShip.MapID, false))
	c.units[newClient.GetSquad().MatherShip.ID] = newClient.GetSquad().MatherShip

	for _, unitSlot := range newClient.GetSquad().MatherShip.Units {
		if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap {
			// юнит на карте
			GetPlaceCoordinate(unitSlot.Unit, c.GetAllShortUnits(unitSlot.Unit.MapID, false))

			if unitSlot.Unit.ID == 0 {
				unitSlot.Unit.ID = c.getBotID()
			}

			c.units[unitSlot.Unit.ID] = unitSlot.Unit
		}
	}
}

func (c *wsUsers) getBotID() int {
	botID := -1

	for {
		_, ok := c.units[botID]
		if !ok {
			return botID
		} else {
			botID--
		}
	}
}

func (c *wsUsers) PlaceUnit(newUnit *unit.Unit) {
	c.unitsMX.Lock()
	defer c.unitsMX.Unlock()

	c.units[newUnit.ID] = newUnit
}

func (c *wsUsers) GetAllShortUnits(mapID int, lock bool) map[int]*unit.ShortUnitInfo {
	if lock {
		c.unitsMX.Lock()
		defer c.unitsMX.Unlock()
	}

	shortUnits := make(map[int]*unit.ShortUnitInfo)

	for _, gameUnit := range c.units {
		if gameUnit.MapID == mapID {
			shortUnits[gameUnit.ID] = gameUnit.GetShortInfo()
		}
	}

	return shortUnits
}

func (c *wsUsers) GetUnitByID(id int) *unit.Unit {
	c.unitsMX.Lock()
	defer c.unitsMX.Unlock()
	return c.units[id]
}

func (c *wsUsers) RemoveUnitByID(id int) {
	c.unitsMX.Lock()
	defer c.unitsMX.Unlock()
	delete(c.units, id)
}

func (c *wsUsers) GetUserByUnitId(unitID int) *player.Player {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	for _, user := range c.users {
		if user.GetSquad() != nil && user.GetSquad().MatherShip != nil && user.GetSquad().GetUnitByID(unitID) != nil {
			return user
		}
	}

	return nil
}

func (c *wsUsers) GetByWs(ws *websocket.Conn) *player.Player {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	user := c.users[ws]
	return user
}

func (c *wsUsers) GetBySquadId(id int) *player.Player {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	for _, client := range c.users {
		if client.GetSquad().ID == id {
			return client
		}
	}

	return nil
}

func (c *wsUsers) GetById(id int) *player.Player {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	for _, client := range c.users {
		if client.GetID() == id {
			return client
		}
	}
	return nil
}

func (c *wsUsers) GetBotByUUID(uuid string) *player.Player {
	c.usersMX.Lock()
	defer c.usersMX.Unlock()

	for _, client := range c.users {
		if client.Bot && client.UUID == uuid {
			return client
		}
	}
	return nil
}

func (c *wsUsers) GetAllConnects() (map[*websocket.Conn]gameConnect, *sync.RWMutex) {
	c.connectMX.Lock()
	return c.connects, &c.connectMX
}

func (c *wsUsers) GetAll() (map[*websocket.Conn]*player.Player, *sync.RWMutex) {
	/*
		DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!!
		Кто бы ты небыл, всегда! всегда Закрывай этот ебучий мьютекс дефером, где бы ты его не вызвал!
		DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!! DANGER!!!
	*/
	c.usersMX.Lock()
	return c.users, &c.usersMX
}

func (c *wsUsers) DelClientByID(id int) {
	c.usersMX.Lock()

	var ws *websocket.Conn
	for userWS, user := range c.connects {
		if user.ID == id {
			ws = userWS
		}
	}

	if c.users[ws] != nil && !c.users[ws].Bot {
		delete(c.users, ws)
		if ws != nil {
			ws.Close()
		}
	}
	c.usersMX.Unlock()

	c.connectMX.Lock()
	_, ok := c.connects[ws]
	if ok && !c.connects[ws].Bot {
		delete(c.connects, ws)
		if ws != nil {
			ws.Close()
		}
	}
	c.connectMX.Unlock()
}
