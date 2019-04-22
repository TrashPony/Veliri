package player

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/gorilla/websocket"
	"strconv"
)

type Player struct {
	id              int
	login           string
	email           string
	credits         int
	experiencePoint int

	squad  *squad.Squad   // отряд игрока
	squads []*squad.Squad // не активные отряды которые ждут игрока на безах

	unitStorage []*unit.Unit // юниты которы находяться не на поле игры в трюме мса

	watch              map[string]map[string]*coordinate.Coordinate // map[X]map[Y] координаты которые видит пользватель
	units              map[string]map[string]*unit.Unit             // map[X]map[Y] свои юниты представленные ввиде карты на поле
	hostileUnits       map[string]map[string]*unit.Unit             // map[X]map[Y] вражеские юниты которы видно в настоящее время
	memoryHostileUnits map[string]unit.Unit                         // Юниты которые видел и запомнил пользователь за всю игру

	gameID int
	Ready  bool

	Training   int `json:"training"`
	openDialog *dialog.Dialog

	LobbyReady bool
	Respawn    *coordinate.Coordinate

	InBaseID   int // ид базы в которой сидит игрок
	LastBaseID int // последняя база которую посетил игрок

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
	AvatarIcon   string                   `json:"avatar_icon"` // путь к аватару
}

type ShortUserInfo struct {
	// структура которая описываем минимальный набор данных для отображение и взаимодействия,
	// что бы другие игроки не палили трюмы, фиты и дронов без спец оборудования
	SquadID    string       `json:"squad_id"`
	UserName   string       `json:"user_name"`
	X          int          `json:"x"`
	Y          int          `json:"y"`
	Q          int          `json:"q"`
	R          int          `json:"r"`
	Rotate     int          `json:"rotate"`
	Body       *detail.Body `json:"body"`
	AvatarIcon string       `json:"avatar_icon"` // путь к аватару
	Fraction   string       `json:"fraction"`
}

func (client *Player) GetShortUserInfo() *ShortUserInfo {
	var hostile ShortUserInfo

	hostile.AvatarIcon = client.AvatarIcon
	hostile.Fraction = client.Fraction

	if client == nil || client.GetSquad() == nil || client.GetSquad().MatherShip == nil || client.GetSquad().MatherShip.Body == nil {
		return nil
	}

	if client.Bot {
		hostile.SquadID = client.UUID
	} else {
		hostile.SquadID = strconv.Itoa(client.GetSquad().ID)
	}

	hostile.UserName = client.GetLogin()
	hostile.X = client.GetSquad().GlobalX
	hostile.Y = client.GetSquad().GlobalY
	hostile.Q = client.GetSquad().Q
	hostile.R = client.GetSquad().R
	hostile.Rotate = client.GetSquad().MatherShip.Rotate

	hostile.Body, _ = gameTypes.Bodies.GetByID(client.GetSquad().MatherShip.Body.ID)

	if client.GetSquad().MatherShip.GetWeaponSlot() != nil && client.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {
		for _, weaponSlot := range hostile.Body.Weapons {
			if weaponSlot != nil {
				weaponSlot.Weapon, _ = gameTypes.Weapons.GetByID(client.GetSquad().MatherShip.GetWeaponSlot().Weapon.ID)
			}
		}
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

func (client *Player) SetRespawn(respawn *coordinate.Coordinate) {
	client.Respawn = respawn
}

func (client *Player) GetRespawn() (respawn *coordinate.Coordinate) {
	return client.Respawn
}

func (client *Player) SetLogin(login string) {
	client.login = login
}

func (client *Player) GetLogin() (login string) {
	return client.login
}

func (client *Player) SetID(id int) {
	client.id = id
}

func (client *Player) GetID() (id int) {
	return client.id
}

func (client *Player) SetGameID(id int) {
	client.gameID = id
}

func (client *Player) GetGameID() (id int) {
	return client.gameID
}

func (client *Player) SetReady(ready bool) {
	client.Ready = ready
}

func (client *Player) GetReady() bool {
	return client.Ready
}

func (client *Player) SetLobbyReady(ready bool) {
	client.LobbyReady = ready
}

func (client *Player) GetLobbyReady() bool {
	return client.LobbyReady
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

func (client *Player) SetExperiencePoint(experiencePoint int) {
	client.experiencePoint = experiencePoint
}

func (client *Player) GetExperiencePoint() int {
	return client.experiencePoint
}
