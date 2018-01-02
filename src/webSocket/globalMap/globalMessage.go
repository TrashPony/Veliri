package globalMap

type GlobalMessage struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	Character string `json:"character"`
	X 			 int	`json:"x"`
	Y 			 int	`json:"y"`
	Rotate		 int	`json:"rotate"`
}

type GlobalResponse struct {
	Event        string `json:"event"`
	UserName     string `json:"user_name"`
	X 			 int	`json:"x"`
	Y 			 int	`json:"y"`
	Character    string `json:"character"`
	Rotate		 int	`json:"rotate"`
}

