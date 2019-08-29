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
	mx        sync.RWMutex
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
	c.mx.Lock()
	defer c.mx.Unlock()

	c.connectMX.Lock()
	defer c.connectMX.Unlock()

	for ws, client := range c.users {
		if !client.Bot && client.GetID() == newClient.GetID() {
			delete(c.users, ws) // удаляем его из активных подключений
			ws.Close()
		}
	}

	c.users[newWS] = newClient
	c.connects[newWS] = gameConnect{ID: newClient.GetID(), Bot: newClient.Bot, MapID: newClient.GetSquad().MapID}

	// мазершип всегда сразу на карте
	c.units[newClient.GetSquad().MatherShip.ID] = newClient.GetSquad().MatherShip
}

func (c *wsUsers) GetAllShortUnits() map[int]*unit.ShortUnitInfo {
	// этого метода хвати и для колизий
	c.unitsMX.Lock()
	defer c.unitsMX.Unlock()

	shortUnits := make(map[int]*unit.ShortUnitInfo)

	for _, gameUnit := range c.units {
		shortUnits[gameUnit.ID] = gameUnit.GetShortInfo()
	}

	return shortUnits
}

func (c *wsUsers) GetByWs(ws *websocket.Conn) *player.Player {
	c.mx.Lock()
	defer c.mx.Unlock()

	user := c.users[ws]
	return user
}

func (c *wsUsers) GetBySquadId(id int) *player.Player {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, client := range c.users {
		if client.GetSquad().ID == id {
			return client
		}
	}

	return nil
}

func (c *wsUsers) GetById(id int) *player.Player {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, client := range c.users {
		if client.GetID() == id {
			return client
		}
	}
	return nil
}

func (c *wsUsers) GetBotByUUID(uuid string) *player.Player {
	c.mx.Lock()
	defer c.mx.Unlock()

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
	c.mx.Lock()
	return c.users, &c.mx
}

func (c *wsUsers) DelClientByID(id int) {
	c.mx.Lock()

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
	c.mx.Unlock()

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
