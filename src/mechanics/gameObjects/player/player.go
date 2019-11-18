package player

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/getlantern/deepcopy"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

type Player struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	email string

	credits int

	ScientificPoints int `json:"scientific_points"`
	AttackPoints     int `json:"attack_points"`
	ProductionPoints int `json:"production_points"`

	squad  *squad.Squad   // отряд игрока
	squads []*squad.Squad // не активные отряды которые ждут игрока на безах

	Training   int `json:"training"`
	openDialog *dialog.Dialog

	LobbyReady bool

	InBaseID           int `json:"-"` // ид базы в которой сидит игрок
	LastBaseID         int `json:"-"` // последняя база которую посетил игрок
	LastBaseEfficiency int `json:"-"` // эффективность базы которая была зафисирована в последний раз, сделала что бы не спамить статусом базы

	/* мета для ботов */
	Bot          bool `json:"bot"`      // переменная говорит что это не игрок))
	Behavior     int  `json:"behavior"` // тип поведения бота
	fakeWS       *websocket.Conn
	UUID         string                   `json:"uuid"`
	GlobalPath   []*coordinate.Coordinate `json:"global_path"`   // маршрут через сектора, тут лежат координаты переходов, входов на базы
	CurrentPoint int                      `json:"current_point"` // номер ячейку куда надо пиздовать
	Leave        bool                     `json:"leave"`
	ToLeave      bool                     `json:"to_leave"`
	LocalPact    int                      `json:"local_pact"`
	Fraction     string                   `json:"fraction"`
	avatarIcon   string                   // ава в base64
	Biography    string                   `json:"biography"`
	Title        string                   `json:"title"`

	// [name]Skill
	CurrentSkills map[string]*skill.Skill     `json:"current_skills"`
	Missions      map[string]*mission.Mission `json:"missions"`
	// uuid мисси которую отслеживает игрок, от этого зависит что будет отображатся на мини карте и в блоке заданий на фронте
	SelectMission string `json:"select_mission"`

	NotifyQueue map[string]*Notify `json:"notify_queue"`

	UserInterface map[string]map[string]*Window `json:"user_interface"` // resolution, window_id, state
	StoryEpisode  int                           `json:"story_episode"`

	DebugMoveMessage []interface{}

	// запомненные динамические обьекты на карте [map_id][x][y]
	MemoryDynamicObjects   map[int]map[int]map[int]*dynamic_map_object.Object `json:"memory_dynamic_objects"`
	memoryDynamicObjectsMX sync.Mutex
}

type ShortUserInfo struct {
	// структура которая описываем минимальный набор данных для отображение и взаимодействия,
	// что бы другие игроки не палили трюмы, фиты и дронов без спец оборудования
	UserID   string `json:"user_id"`
	SquadID  string `json:"squad_id"`
	UserName string `json:"user_name"`

	Biography string `json:"biography"`
	Title     string `json:"title"`
	Fraction  string `json:"fraction"`

	MotherShip *unit.ShortUnitInfo `json:"mother_ship"`
}

type Notify struct {
	Name    string             `json:"name"`
	UUID    string             `json:"uuid"`
	Event   string             `json:"event"`
	Send    bool               `json:"send"`
	Data    interface{}        `json:"data"`
	Destroy bool               `json:"destroy"`
	Count   int                `json:"count"`
	Price   int                `json:"price"`
	Item    *inventory.Slot    `json:"item"`
	Base    *base.Base         `json:"base"`
	Map     *_map.ShortInfoMap `json:"map"`
}

func (client *Player) ReloadNotify() {
	for _, notify := range client.NotifyQueue {
		notify.Send = false
	}
}

// текущие положение интерфейса пользователя
type Window struct {
	Left   int  `json:"left"`
	Top    int  `json:"top"`
	Height int  `json:"height"`
	Width  int  `json:"width"`
	Open   bool `json:"open"`
}

func (client *Player) AddDynamicObject(object *dynamic_map_object.Object, mapID int) {
	client.memoryDynamicObjectsMX.Lock()
	defer client.memoryDynamicObjectsMX.Unlock()

	var memoryObj dynamic_map_object.Object // создаем копию обьекта
	err := deepcopy.Copy(&memoryObj, &object)
	if err != nil {
		println(err.Error())
	}

	if client.MemoryDynamicObjects == nil {
		client.MemoryDynamicObjects = make(map[int]map[int]map[int]*dynamic_map_object.Object)
	}

	if client.MemoryDynamicObjects[mapID] == nil {
		client.MemoryDynamicObjects[mapID] = make(map[int]map[int]*dynamic_map_object.Object)
	}

	if client.MemoryDynamicObjects[mapID][object.X] == nil {
		client.MemoryDynamicObjects[mapID][object.X] = make(map[int]*dynamic_map_object.Object)
	}

	client.MemoryDynamicObjects[mapID][object.X][object.Y] = &memoryObj
}

func (client *Player) RemoveDynamicObject(object *dynamic_map_object.Object, mapID int) {
	client.memoryDynamicObjectsMX.Lock()
	defer client.memoryDynamicObjectsMX.Unlock()

	delete(client.MemoryDynamicObjects[mapID][object.X], object.Y)
}

func (client *Player) GetMapDynamicObject(mapID, x, y int) *dynamic_map_object.Object {
	client.memoryDynamicObjectsMX.Lock()
	defer client.memoryDynamicObjectsMX.Unlock()

	return client.MemoryDynamicObjects[mapID][x][y]
}

func (client *Player) GetMapDynamicObjectByID(mapID, id int) *dynamic_map_object.Object {
	client.memoryDynamicObjectsMX.Lock()
	defer client.memoryDynamicObjectsMX.Unlock()

	if client.MemoryDynamicObjects[mapID] != nil {
		for _, x := range client.MemoryDynamicObjects[mapID] {
			for _, obj := range x {
				if obj.ID == id {
					return obj
				}
			}
		}
	}

	return nil
}

func (client *Player) GetMapDynamicObjects(mapID int) map[int]map[int]*dynamic_map_object.Object {
	client.memoryDynamicObjectsMX.Lock()
	defer client.memoryDynamicObjectsMX.Unlock()

	return client.MemoryDynamicObjects[mapID]
}

func (client *Player) GetShortUserInfo(squad bool) *ShortUserInfo {
	var hostile ShortUserInfo

	if client.Bot {
		hostile.UserID = client.UUID
	} else {
		hostile.UserID = strconv.Itoa(client.GetID())
	}

	hostile.Fraction = client.Fraction
	hostile.UserName = client.GetLogin()
	hostile.Biography = client.Biography
	hostile.Title = client.Title

	if !squad {
		return &hostile
	}

	if client.Bot {
		hostile.SquadID = client.UUID
	} else {
		hostile.SquadID = strconv.Itoa(client.GetSquad().ID)
	}

	if client.GetSquad() != nil && client.GetSquad().MatherShip != nil {
		hostile.MotherShip = client.GetSquad().MatherShip.GetShortInfo()
	}

	return &hostile
}

func (client *Player) SetFakeWS(ws *websocket.Conn) {
	client.fakeWS = ws
}

func (client *Player) GetFakeWS() *websocket.Conn {
	return client.fakeWS
}

func (client *Player) GetOpenDialog() *dialog.Dialog {
	return client.openDialog
}

func (client *Player) SetOpenDialog(newDialog *dialog.Dialog) {
	client.openDialog = newDialog
}

func (client *Player) SetLogin(login string) {
	client.Login = login
}

func (client *Player) GetLogin() (login string) {
	return client.Login
}

func (client *Player) SetID(id int) {
	client.ID = id
}

func (client *Player) GetID() (id int) {
	return client.ID
}

func (client *Player) GetSquad() *squad.Squad {
	if client != nil && client.squad != nil {
		return client.squad
	}
	return nil
}

func (client *Player) SetSquad(squad *squad.Squad) {
	client.squad = squad
}

func (client *Player) GetSquads() []*squad.Squad {
	return client.squads
}

func (client *Player) GetSquadsByID(ID int) *squad.Squad {
	for _, userSquad := range client.squads {
		if userSquad != nil && userSquad.ID == ID {
			return userSquad
		}
	}

	return nil
}

func (client *Player) RemoveSquadsByID(ID int) {
	for i, userSquad := range client.squads {
		if userSquad.ID == ID {
			client.squads[i] = nil
		}
	}
}

func (client *Player) GetSquadsByBaseID(BaseID int) []*squad.Squad {
	squads := make([]*squad.Squad, 0)

	for _, userSquad := range client.squads {
		if client.GetSquad() != nil {
			if userSquad != nil && userSquad.BaseID == BaseID && client.GetSquad().ID != userSquad.ID {
				squads = append(squads, userSquad)
			}
		} else {
			if userSquad != nil && userSquad.BaseID == BaseID {
				squads = append(squads, userSquad)
			}
		}
	}

	return squads
}

func (client *Player) SetSquads(squads []*squad.Squad) {
	client.squads = squads
}

func (client *Player) SetEmail(email string) {
	client.email = email
}

func (client *Player) GetEmail() string {
	return client.email
}

func (client *Player) SetCredits(credits int) {
	client.credits = credits
}

func (client *Player) GetCredits() int {
	return client.credits
}

func (client *Player) UpSkill(id int) (*skill.Skill, bool) {
	if skillType := gameTypes.Skills.GetByID(id); skillType != nil {
		currentSkill := client.CurrentSkills[skillType.Name]
		var needPoints int

		//todo Я не очень умный :С
		if currentSkill.Level == 0 {
			needPoints = 100
		}

		if currentSkill.Level == 1 {
			needPoints = 200
		}

		if currentSkill.Level == 2 {
			needPoints = 400
		}

		if currentSkill.Level == 3 {
			needPoints = 800
		}

		if currentSkill.Level == 4 {
			needPoints = 1600
		}

		if currentSkill.Level > 4 {
			return nil, false
		}

		if currentSkill.Type == "scientific" {
			if client.ScientificPoints >= needPoints {

				client.ScientificPoints -= needPoints
				client.CurrentSkills[skillType.Name].Level++
				return client.CurrentSkills[skillType.Name], true
			}
		}

		if currentSkill.Type == "attack" {
			if client.AttackPoints >= needPoints {

				client.AttackPoints -= needPoints
				client.CurrentSkills[skillType.Name].Level++
				return client.CurrentSkills[skillType.Name], true
			}
		}

		if currentSkill.Type == "production" {
			if client.ProductionPoints >= needPoints {

				client.ProductionPoints -= needPoints
				client.CurrentSkills[skillType.Name].Level++
				return client.CurrentSkills[skillType.Name], true
			}
		}
	}
	return nil, false
}

func (client *Player) GetAvatar() string {
	return client.avatarIcon
}

func (client *Player) SetAvatar(avatar string) {
	client.avatarIcon = avatar
}
