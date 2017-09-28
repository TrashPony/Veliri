package webSocket


type LobbyMessage struct {
	Event    string `json:"event"`
	MapName  string `json:"map_name"`
	UserName string `json:"user_name"`
	GameName string `json:"game_name"`
}

type LobbyResponse struct {
	Event    	  string `json:"event"`
	UserName	  string
	ResponseNameGame  string `json:"response_name_game"`
	ResponseNameMap   string `json:"response_name_map"`
	ResponseNameUser  string `json:"response_name_user"`
	ResponseNameUser2 string `json:"response_name_user_2"`
}
