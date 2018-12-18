package base

type Base struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Q     int    `json:"q"`
	R     int    `json:"r"`
	MapID int    `json:"map_id"`
	RespQ int    `json:"resp_q"`
	RespR int    `json:"resp_r"`
}
