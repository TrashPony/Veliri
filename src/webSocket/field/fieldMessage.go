package field

import (
	"../../game"
	"strconv"
)

type FieldResponse struct {
	Event       string `json:"event"`
	UserName    string `json:"user_name"`
	PlayerPrice int    `json:"player_price"`
	GameStep    int    `json:"game_step"`
	GamePhase   string `json:"game_phase"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	ToX         int    `json:"to_x"`
	ToY         int    `json:"to_y"`
	XMap        int    `json:"x_map"`
	YMap        int    `json:"y_map"`
	TypeMap     string `json:"type_map"`
	NameMap     string `json:"name_map"`
	TypeUnit    string `json:"type_unit"`
	ErrorType   string `json:"error_type"`
	Phase       string `json:"phase"`
	UserReady   bool `json:"user_ready"`
	UserOwned   string `json:"user_owned"`
	Error       string `json:"error"`
}

type InitUnit struct {
	Event       string `json:"event"`
	UserName    string `json:"user_name"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	TypeUnit    string `json:"type_unit"`
	UserOwned   string `json:"user_owned"`
	HP          int    `json:"hp"`
	UnitAction  string `json:"unit_action"`
	Target      string `json:"target"`
	Damage      string `json:"damage"`
	MoveSpeed   string `json:"move_speed"`
	Init        string `json:"init"`
	RangeAttack string `json:"range_attack"`
	RangeView   string `json:"range_view"`
	AreaAttack  string `json:"area_attack"`
	TypeAttack  string `json:"type_attack"`
	Error       string `json:"error"`
}

func (msg *InitUnit) initUnit(unit *game.Unit, login string) {
	if unit.Target == nil {
		var unitsParams = InitUnit{Event: "InitUnit", UserName: login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
			HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: "", X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
		initUnit <- unitsParams
	} else {
		var unitsParametr = InitUnit{Event: "InitUnit", UserName: login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
			HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target.X) + ":" + strconv.Itoa(unit.Target.Y), X: unit.X, Y: unit.Y}
		initUnit <- unitsParametr
	}
}

type Move struct {
	Event       string `json:"event"`
	UserName    string `json:"user_name"`
	UnitX		int    `json:"unit_x"`
	UnitY		int	   `json:"unit_y"`
	PathNodes   []game.Coordinate `json:"path_nodes"`
}

type InitStructure struct {
	Event      		 string `json:"event"`
	UserName  		 string `json:"user_name"`
	X      		     int    `json:"x"`
	Y          		 int    `json:"y"`
	TypeStructure    string `json:"type_structure"`
	UserOwned   	 string `json:"user_owned"`
	Error            string `json:"error"`
}

func (msg *InitStructure) initStructure(structure *game.Structure, login string) {
	var structureParams = InitStructure{Event: "InitStructure", UserName: login, TypeStructure: structure.Type, UserOwned: structure.NameUser, X: structure.X, Y: structure.Y} // остылаем событие добавления юнита
	initStructure <- structureParams
}

type sendCoordinate struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

func openCoordinate(login string, x, y int) {
	resp := sendCoordinate{Event: "OpenCoordinate", UserName: login, X: x, Y: y}
	coordinate <- resp
}

func closeCoordinate(login string, x, y int) {
	resp := sendCoordinate{Event: "DellCoordinate", UserName: login, X: x, Y: y}
	coordinate <- resp
}

type FieldMessage struct {
	Event    string `json:"event"`
	IdGame   int    `json:"id_game"`
	IdUnit   string `json:"id_unit"`
	IdTarget string `json:"id_target"`
	TypeUnit string `json:"type_unit"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	ToX      int    `json:"to_x"`
	ToY      int    `json:"to_y"`
	TargetX  int    `json:"target_x"`
	TargetY  int    `json:"target_y"`
}
