package field

import (
	"../../mechanics/unit"
	"../../mechanics/watchZone"
	"../../mechanics/matherShip"
	"../../mechanics/coordinate"
)

type Response struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`

	GameStep  int    `json:"game_step"`
	GamePhase string `json:"game_phase"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	ToX       int    `json:"to_x"`
	ToY       int    `json:"to_y"`
	XMap      int    `json:"x_map"`
	YMap      int    `json:"y_map"`
	TypeMap   string `json:"type_map"`
	NameMap   string `json:"name_map"`
	TypeUnit  string `json:"type_unit"`
	Phase     string `json:"phase"`
	UserReady bool   `json:"user_ready"`
	UserOwned string `json:"user_owned"`
	Error     string `json:"error"`
}

type InitUnit struct {
	Event    string     `json:"event"`
	UserName string     `json:"user_name"`
	Unit     *unit.Unit `json:"unit"`
}

func (msg *InitUnit) initUnit(event string, unit *unit.Unit, login string) {
	if unit.Target == nil {
		var unitsParams = InitUnit{Event: event, UserName: login, Unit: unit} // остылаем событие добавления юнита
		initUnit <- unitsParams
	} else {
		var unitsParams = InitUnit{Event: event, UserName: login, Unit: unit}
		initUnit <- unitsParams
	}
}

type Move struct {
	Event     string                                 `json:"event"`
	UserName  string                                 `json:"user_name"`
	Unit      *unit.Unit                             `json:"unit"`
	PathNodes []coordinate.Coordinate                `json:"path_nodes"`
	WatchNode map[string]*watchZone.UpdaterWatchZone `json:"watch_node"`
	Error     string                                 `json:"error"`
}

type InitStructure struct {
	Event     string                 `json:"event"`
	UserName  string                 `json:"user_name"`
	Structure *matherShip.MatherShip `json:"structure"`
}

func (msg *InitStructure) initMatherShip(event string, structure *matherShip.MatherShip, login string) {
	var matherShipParameter = InitStructure{Event: event, UserName: login, Structure: structure} // остылаем событие добавления юнита
	initStructure <- matherShipParameter
}

type sendCoordinate struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

func openCoordinate(login string, x, y int) {
	resp := sendCoordinate{Event: "OpenCoordinate", UserName: login, X: x, Y: y}
	senderCoordinate <- resp
}

func closeCoordinate(login string, x, y int) {
	resp := sendCoordinate{Event: "DellCoordinate", UserName: login, X: x, Y: y}
	senderCoordinate <- resp
}

type Message struct {
	Event    string `json:"event"`
	IdGame   int    `json:"id_game"`
	UnitID   int    `json:"unit_id"`
	IdTarget string `json:"id_target"`
	TypeUnit string `json:"type_unit"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	ToX      int    `json:"to_x"`
	ToY      int    `json:"to_y"`
	TargetX  int    `json:"target_x"`
	TargetY  int    `json:"target_y"`
}
