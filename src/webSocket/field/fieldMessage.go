package field

type FieldResponse struct {
	Event    	  string `json:"event"`
	UserName	  string `json:"user_name"`
	PlayerPrice	  string `json:"player_price"`
	GameStep 	  string `json:"game_step"`
	GamePhase 	  string `json:"game_phase"`
    X      		  string `json:"x"`
    Y			  string `json:"y"`
	XMap 	  	  string `json:"x_map"`
	YMap	  	  string `json:"y_map"`
	TypeMap       string `json:"type_map"`
	NameMap       string `json:"name_map"`
	TypeUnit 	  string `json:"type_unit"`
	ErrorType	  string `json:"error_type"`
	Phase		  string `json:"phase"`
	UserReady	  string `json:"user_ready"`
	UserOwned 	  string `json:"user_owned"`
	HP 			  string `json:"hp"`
	UnitAction	  string `json:"unit_action"`
	Target 		  string `json:"target"`
	Damage		  string `json:"damage"`
	MoveSpeed	  string `json:"move_speed"`
	Init		  string `json:"init"`
	RangeAttack	  string `json:"range_attack"`
	RangeView	  string `json:"range_view"`
	AreaAttack    string `json:"area_attack"`
	TypeAttack	  string `json:"type_attack"`
	RespawnX       string `json:"respawn_x"`
	RespawnY       string `json:"respawn_y"`
	Error		 string   `json:"error"`
}


type FieldMessage struct {
	Event    string `json:"event"`
	IdGame   string `json:"id_game"`
	IdUnit	 string `json:"id_unit"`
	IdTarget string	`json:"id_target"`
	TypeUnit string `json:"type_unit"`
	X        string	`json:"x"`
	Y 	     string	`json:"y"`
	ToX	 string `json:"to_x"`
	ToY	 string `json:"to_y"`
}