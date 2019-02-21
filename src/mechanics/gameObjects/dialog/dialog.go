package dialog

type Dialog struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Pages      map[int]Page `json:"pages"` // все страницы диалога
	Picture    string       `json:"picture"`
	AccessType string       `json:"access_type"`
}

type Page struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Text   string `json:"text"` // текст страницы
	Asc    []Ask  `json:"asc"`  // варианты отетов
}

type Ask struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`    // текст ответа
	ToPage     int    `json:"to_page"` // страница на которую ведет ответ
	typeAction string // функция которая выолнается при выборе этого варианта ответа
}

func (a *Ask) SetAction(action string) {
	a.typeAction = action
}
