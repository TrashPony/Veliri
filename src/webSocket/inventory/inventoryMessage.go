package inventory

import (
	"../../mechanics/gameObjects/squad"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/matherShip"
	"../../mechanics/gameObjects/equip"
)

type Message struct {
	Event string `json:"event"`

	ChassisID int `json:"chassis"`
	WeaponID  int `json:"weapon"`
	TowerID   int `json:"tower"`
	BodyID    int `json:"body"`
	RadarID   int `json:"radar"`

	SquadName    string `json:"squad_name"`
	SquadID      int    `json:"squad_id"`
	UnitSlot     int    `json:"slot"`
	EquipSlot    int    `json:"equip_slot"`
	EquipID      int    `json:"equip_id"`
	MatherShipID int    `json:"mather_ship_id"`
}

type Response struct {
	Event        string `json:"event"`
	UserName     string `json:"user_name"`
	NameGame     string `json:"name_game"`
	IdGame       int    `json:"id_game"`
	PhaseGame    string `json:"phase_game"`
	StepGame     string `json:"step_game"`
	Ready        string `json:"ready"`
	NumOfPlayers string `json:"num_of_players"`
	Players      string `json:"players"`
	Creator      string `json:"creator"`
	NewUser      string `json:"new_user"`
	Error        string `json:"error"`
	Message      string `json:"message"`
	GameUser     string `json:"game_user"`

	MatherShips []matherShip.MatherShip `json:"mather_ships"`
	Equipping   []equip.Equip   `json:"equipping"`
	Unit        unit.Unit               `json:"unit"`

	Weapons []detail.Weapon `json:"weapons"`
	Bodies  []detail.Body   `json:"bodies"`

	Squads  []*squad.Squad `json:"squads"`
	Squad   *squad.Squad   `json:"squad"`
	SquadID int                `json:"squad_id"`

	EquipSlot int `json:"equip_slot"`
	UnitSlot  int `json:"slot"`
}
