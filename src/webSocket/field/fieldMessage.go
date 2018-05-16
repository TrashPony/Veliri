package field

import (
	"../../mechanics/unit"
	"../../mechanics/watchZone"
	"../../mechanics/coordinate"
)

type Move struct {
	Event     string                                 `json:"event"`
	UserName  string                                 `json:"user_name"`
	Unit      *unit.Unit                             `json:"unit"`
	PathNodes []coordinate.Coordinate                `json:"path_nodes"`
	WatchNode map[string]*watchZone.UpdaterWatchZone `json:"watch_node"`
	Error     string                                 `json:"error"`
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

type ErrorMessage struct {
	Event string `json:"event"`
	Error string `json:"error"`
}
