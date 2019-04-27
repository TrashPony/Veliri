package skill

type Skill struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Specification   string `json:"specification"`
	ExperiencePoint int    `json:"experience_point"`
	Type            string `json:"type"`
	Level           int    `json:"level"`
	Icon            string `json:"icon"`
}
