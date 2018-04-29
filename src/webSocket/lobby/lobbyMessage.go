package lobby

import (
	"../../lobby/DetailUnit"
	"../../lobby/Squad"
)

type Message struct {
	Event    string `json:"event"`
	MapName  string `json:"map_name"`
	UserName string `json:"user_name"`
	GameName string `json:"game_name"`
	Respawn  string `json:"respawn"`
	Message  string `json:"message"`

	SquadName    string `json:"squad_name"`
	SquadID      int    `json:"squad_id"`
	UnitSlot     int    `json:"slot"`
	EquipSlot    int    `json:"equip_slot"`
	EquipID      int    `json:"equip_id"`
	MatherShipID int    `json:"mather_ship_id"`

	ChassisID int `json:"chassis"`
	WeaponID  int `json:"weapon"`
	TowerID   int `json:"tower"`
	BodyID    int `json:"body"`
	RadarID   int `json:"radar"`
}

type Response struct {
	Event        string `json:"event"`
	UserName     string `json:"user_name"`
	NameGame     string `json:"name_game"`
	IdGame       string `json:"id_game"`
	PhaseGame    string `json:"phase_game"`
	StepGame     string `json:"step_game"`
	Ready        string `json:"ready"`
	NameMap      string `json:"name_map"`
	NumOfPlayers string `json:"num_of_players"`
	Players      string `json:"players"`
	Creator      string `json:"creator"`
	NewUser      string `json:"new_user"`
	GameUser     string `json:"game_user"`
	Error        string `json:"error"`
	Respawn      string `json:"respawn"`
	RespawnName  string `json:"respawn_name"`
	Message      string `json:"message"`

	MatherShips []Squad.MatherShip `json:"mather_ships"`
	Equipping   []Squad.Equipping  `json:"equipping"`

	Chassis []DetailUnit.Chassis `json:"chassis"`
	Weapons []DetailUnit.Weapon  `json:"weapons"`
	Towers  []DetailUnit.Tower   `json:"towers"`
	Bodies  []DetailUnit.Body    `json:"bodies"`
	Radars  []DetailUnit.Radar   `json:"radars"`

	Squads []*Squad.Squad `json:"squads"`
	Squad  *Squad.Squad   `json:"squad"`

	EquipSlot int `json:"equip_slot"`
	UnitSlot  int `json:"slot"`
}
