package dialog

type Dialog struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Pages      map[int]*Page `json:"pages"` // все страницы диалога
	AccessType string        `json:"access_type"`
}

type Page struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	Name    string `json:"name"`
	Text    string `json:"text"` // текст страницы
	Asc     []Ask  `json:"asc"`  // варианты отетов
	Picture string `json:"picture"`
}

type Ask struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`    // текст ответа
	ToPage     int    `json:"to_page"` // страница на которую ведет ответ
	typeAction string // функция которая выолнается при выборе этого варианта ответа
}

func (a *Ask) GetAction() string {
	return a.typeAction
}

func (a *Ask) SetAction(action string) {
	a.typeAction = action
}
