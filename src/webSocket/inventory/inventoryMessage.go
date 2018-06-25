package inventory

import (
	"../../inventory"
	"../../detailUnit"
)

type Message struct {
	Event    string `json:"event"`

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

	MatherShips []inventory.MatherShip `json:"mather_ships"`
	Equipping   []inventory.Equipping  `json:"equipping"`
	Unit        inventory.Unit         `json:"unit"`

	Weapons []detailUnit.Weapon  `json:"weapons"`
	Bodies  []detailUnit.Body    `json:"bodies"`

	Squads  []*inventory.Squad `json:"squads"`
	Squad   *inventory.Squad   `json:"squad"`
	SquadID int            `json:"squad_id"`

	EquipSlot int `json:"equip_slot"`
	UnitSlot  int `json:"slot"`
}
