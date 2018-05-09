package lobby

import (
	"../../DetailUnit"
	"../../lobby/Squad"
	"../../lobby"
)

type Message struct {
	Event     string `json:"event"`
	UserName  string `json:"user_name"`
	GameName  string `json:"game_name"`
	RespawnID int    `json:"respawn_id"`
	Message   string `json:"message"`

	MapID int `json:"map_id"`

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
	IdGame       int `json:"id_game"`
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

	User      *lobby.User   `json:"user"`
	GameUsers []*lobby.User `json:"game_users"`

	Respawn  *lobby.Respawn    `json:"respawn"`
	Respawns []*lobby.Respawn  `json:"respawns"`
	Game     *lobby.LobbyGames `json:"game"`
	Map      *lobby.Map        `json:"map"`

	MatherShips []Squad.MatherShip `json:"mather_ships"`
	Equipping   []Squad.Equipping  `json:"equipping"`
	Unit        Squad.Unit         `json:"unit"`

	Chassis []DetailUnit.Chassis `json:"chassis"`
	Weapons []DetailUnit.Weapon  `json:"weapons"`
	Towers  []DetailUnit.Tower   `json:"towers"`
	Bodies  []DetailUnit.Body    `json:"bodies"`
	Radars  []DetailUnit.Radar   `json:"radars"`

	Squads  []*Squad.Squad `json:"squads"`
	Squad   *Squad.Squad   `json:"squad"`
	SquadID int            `json:"squad_id"`

	EquipSlot int `json:"equip_slot"`
	UnitSlot  int `json:"slot"`

	DontEndGames []lobby.DontEndGames `json:"dont_end_games"`
}
