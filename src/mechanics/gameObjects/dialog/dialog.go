package dialog

import (
	"strings"
)

type Dialog struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Pages      map[int]*Page `json:"pages"` // все страницы диалога
	AccessType string        `json:"access_type"`
	Fraction   string        `json:"fraction"`
	Type       string        `json:"type"`
	Mission    string        `json:"mission"` // говорит какая миссия начнется при ивенте асепт_миссион
}

func (d *Dialog) GetPageByType(typePage string) *Page {
	for _, page := range d.Pages {
		if page.Type == typePage {
			return page
		}
	}
	return nil
}

func (d *Dialog) ProcessingDialogText(userName, BaseName, ToBaseName, ToSectorName string) {
	// %UserName% %BaseName% %ToBaseName%

	for _, page := range d.Pages {
		page.Text = strings.Replace(page.Text, "%UserName%", userName, -1)
		page.Text = strings.Replace(page.Text, "%BaseName%", BaseName, -1)
		page.Text = strings.Replace(page.Text, "%ToBaseName%", ToBaseName, -1)
		page.Text = strings.Replace(page.Text, "%ToSectorName%", ToSectorName, -1)

		for _, asc := range page.Asc {
			asc.Text = strings.Replace(asc.Text, "%UserName%", userName, -1)
			asc.Text = strings.Replace(asc.Text, "%BaseName%", BaseName, -1)
			asc.Text = strings.Replace(asc.Text, "%ToBaseName%", ToBaseName, -1)
			asc.Text = strings.Replace(asc.Text, "%ToSectorName%", ToSectorName, -1)
		}
	}
}

type Page struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	Name    string `json:"name"`
	Text    string `json:"text"` // текст страницы
	Asc     []Ask  `json:"asc"`  // варианты отетов
	Picture string `json:"picture"`
	Type    string `json:"type"`
}

type Ask struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`        // текст ответа
	ToPage     int    `json:"to_page"`     // страница на которую ведет ответ
	TypeAction string `json:"type_action"` // функция которая выолнается при выборе этого варианта ответа
}
