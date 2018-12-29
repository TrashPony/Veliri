package base

type Base struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Q             int                `json:"q"`
	R             int                `json:"r"`
	MapID         int                `json:"map_id"`
	RespQ         int                `json:"resp_q"`
	RespR         int                `json:"resp_r"`
	Transports    map[int]*Transport `json:"transports"`
	Defenders     map[int]*Transport `json:"defenders"`
	GravityRadius int                `json:"gravity_radius"`
}

type Transport struct {
	ID      int  `json:"id"`
	X       int  `json:"x"`
	Y       int  `json:"y"`
	Job     bool `json:"job"`      /* на задание он или нет */
	Down    bool `json:"down"`     /* на земле он или нет */
	SquadID bool `json:"squad_id"` /* ид того кого он тащит */
}

func (b *Base) CreateTransports(count int) {
	b.Transports = make(map[int]*Transport)

	for i := 0; i < count; i++ {
		b.Transports[i] = &Transport{ID: i, Down: true}
	}
}

func (b *Base) GetFreeTransport() *Transport {
	for _, transport := range b.Transports {
		if !transport.Job {
			return transport
		}
	}
	return nil
}

type defender struct {
	// TODO
}

func (b *Base) CreateDefenders(count int) {
	for i := 0; i < count; i++ {
		// TODO
	}
}
