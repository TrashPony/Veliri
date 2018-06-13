package effect

type Effect struct {
	ID          int    `json:"id"`
	TypeID      int    `json:"type_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	StepsTime   int    `json:"steps_time"`
	Parameter   string `json:"parameter"`
	Quantity    int    `json:"quantity"`
	Percentages bool   `json:"percentages"`
	Forever     bool   `json:"forever"`
}
