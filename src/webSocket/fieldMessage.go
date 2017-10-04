package webSocket

type FieldMessage struct {
	Event    string `json:"event"`
	IdGame   string `json:"id_game"`
	IdUnit	 string `json:"id_unit"`
	IdTarget string	`json:"id_target"`
	TypeUnit string `json:"type_unit"`
	X        string	`json:"x"`
	Y 	 string	`json:"y"`
}

type FieldResponse struct {
	Event    	  string `json:"event"`
	UserName	  string
	PlayerPrice	  string `json:"player_price"`
	GameStep 	  string `json:"game_step"`
	GamePhase 	  string `json:"game_phase"`
        X      		  string `json:"x"`
        Y		  string `json:"y"`
	XMap 	  	  string `json:"x_map"`
	YMap	  	  string `json:"y_map"`
	TypeUnit 	  string `json:"type_unit"`
	ErrorType	  string `json:"error_type"`
}
