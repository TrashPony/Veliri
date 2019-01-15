package dialog

type Dialog struct {
	Pages []Page `json:"pages"` // все страницы диалога
}

type Page struct {
	Text string `json:"text"` // текст страницы
	Asc  []Ask  `json:"asc"`  // варианты отетов
}

type Ask struct {
	Text   string `json:"text"`    // текст ответа
	ToPage int    `json:"to_page"` // страница на которую ведет ответ
	action func()				   // функция которая выолнается при выборе этого варианта ответа
}

func (a *Ask)GetAction()  {

}