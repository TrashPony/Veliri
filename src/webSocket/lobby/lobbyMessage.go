package lobby

import "../../lobby"

type Message struct {
	Event    string `json:"event"`
	MapName  string `json:"map_name"`
	UserName string `json:"user_name"`
	GameName string `json:"game_name"`
	Respawn  string `json:"respawn"`
	Message  string `json:"message"`
}

type Response struct {
	Event        string             `json:"event"`
	UserName     string             `json:"user_name"`
	NameGame     string             `json:"name_game"`
	IdGame       string             `json:"id_game"`
	PhaseGame    string             `json:"phase_game"`
	StepGame     string             `json:"step_game"`
	Ready        string             `json:"ready"`
	NameMap      string             `json:"name_map"`
	NumOfPlayers string             `json:"num_of_players"`
	Players      string             `json:"players"`
	Creator      string             `json:"creator"`
	NewUser      string             `json:"new_user"`
	GameUser     string             `json:"game_user"`
	Error        string             `json:"error"`
	Respawn      string             `json:"respawn"`
	RespawnName  string             `json:"respawn_name"`
	Message      string             `json:"message"`
	MatherShips  []lobby.MatherShip `json:"mather_ships"`
	Chassis      []lobby.Chassis    `json:"chassis"`
	Weapons      []lobby.Weapon     `json:"weapons"`
	Towers       []lobby.Tower      `json:"towers"`
	Bodies       []lobby.Body       `json:"bodies"`
	Radars       []lobby.Radar      `json:"radars"`
}
