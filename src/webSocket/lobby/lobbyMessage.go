package lobby

import (
	"../../mechanics/gameObjects/coordinate"
	LocalMap "../../mechanics/gameObjects/map"
	"../../mechanics/gameObjects/squad"
	"../../mechanics/lobby"
	"../../mechanics/lobby/notFinishedGames"
	"../../mechanics/player"
)

type Message struct {
	Event     string `json:"event"`
	UserName  string `json:"user_name"`
	GameName  string `json:"game_name"`
	RespawnID int    `json:"respawn_id"`
	Message   string `json:"message"`
	GameID    int    `json:"game_id"`

	MapID int `json:"map_id"`
}

type Response struct {
	Event        string `json:"event"`
	UserName     string `json:"user_name"`
	NameGame     string `json:"name_game"`
	IdGame       int    `json:"id_game"`
	PhaseGame    string `json:"phase_game"`
	StepGame     string `json:"step_game"`
	Ready        bool   `json:"ready"`
	NumOfPlayers string `json:"num_of_players"`
	Players      string `json:"players"`
	Creator      string `json:"creator"`
	NewUser      string `json:"new_user"`
	Error        string `json:"error"`
	Message      string `json:"message"`
	GameUser     string `json:"game_user"`

	User      *player.Player                 `json:"user"`
	GameUsers map[string]*player.Player      `json:"game_users"`
	Squad     *squad.Squad                   `json:"squad"`
	Respawn   *coordinate.Coordinate         `json:"respawn"`
	Respawns  map[int]*coordinate.Coordinate `json:"respawns"`
	Game      *lobby.Game                    `json:"game"`
	Map       *LocalMap.Map                  `json:"map"`

	DontEndGames []notFinishedGames.NotFinishedGames `json:"dont_end_games"`
}
