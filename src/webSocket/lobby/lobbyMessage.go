package lobby

import (
	"../../lobby"
)

type Message struct {
	Event     string `json:"event"`
	UserName  string `json:"user_name"`
	GameName  string `json:"game_name"`
	RespawnID int    `json:"respawn_id"`
	Message   string `json:"message"`

	MapID int `json:"map_id"`
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

	User      *lobby.User   `json:"user"`
	GameUsers []*lobby.User `json:"game_users"`

	Respawn  *lobby.Respawn    `json:"respawn"`
	Respawns []*lobby.Respawn  `json:"respawns"`
	Game     *lobby.LobbyGames `json:"game"`
	Map      *lobby.Map        `json:"map"`

	DontEndGames []lobby.DontEndGames `json:"dont_end_games"`
}
