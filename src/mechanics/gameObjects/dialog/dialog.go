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

func (d *Dialog) ProcessingDialogText(userName, BaseName, ToBaseName, ToSectorName, userFraction string) {

	if d == nil {
		return
	}

	for _, page := range d.Pages {
		page.Text = ProcessingText(page.Text, userName, BaseName, ToBaseName, ToSectorName, userFraction)
		for _, asc := range page.Asc {
			asc.Text = ProcessingText(asc.Text, userName, BaseName, ToBaseName, ToSectorName, userFraction)
		}
	}
}

func ProcessingText(text, userName, BaseName, ToBaseName, ToSectorName, userFraction string) string {
	importantlyWrapperStart := "<span class=\"importantly\">"
	importantlyWrapperEnd := "</span>"

	text = strings.Replace(text, "%UserName%", importantlyWrapperStart+userName+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "%BaseName%", importantlyWrapperStart+BaseName+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "%ToBaseName%", importantlyWrapperStart+ToBaseName+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "%ToSectorName%", importantlyWrapperStart+ToSectorName+importantlyWrapperEnd, -1)

	text = strings.Replace(text, "Replics", importantlyWrapperStart+"Replics"+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "Explores", importantlyWrapperStart+"Explores"+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "Reverses", importantlyWrapperStart+"Reverses"+importantlyWrapperEnd, -1)

	text = strings.Replace(text, "Veliri-5", importantlyWrapperStart+"Veliri-5"+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "Veliri", importantlyWrapperStart+"Veliri"+importantlyWrapperEnd, -1)
	text = strings.Replace(text, "Veliri", importantlyWrapperStart+"Veliri"+importantlyWrapperEnd, -1)

	return text
}

type Page struct {
	ID              int    `json:"id"`
	Number          int    `json:"number"`
	Name            string `json:"name"`
	Text            string `json:"text"` // текст страницы
	Asc             []Ask  `json:"asc"`  // варианты отетов
	picture         string
	pictureReplics  string
	pictureExplores string
	pictureReverses string
	Type            string `json:"type"`
}

func (p *Page) SetPictures(mainPicture, pictureReplics, pictureExplores, pictureReverses string) {
	p.picture = mainPicture
	p.pictureReplics = pictureReplics
	p.pictureExplores = pictureExplores
	p.pictureReverses = pictureReverses
}

func (p *Page) SetPicture(picture, typePic string) {
	if typePic == "main" {
		p.picture = picture
	}

	if typePic == "Replics" {
		p.pictureReplics = picture
	}

	if typePic == "Explores" {
		p.pictureExplores = picture
	}

	if typePic == "Reverses" {
		p.pictureReverses = picture
	}
}

func (p *Page) GetPicture(typePic string) string {
	// если в диалог есть только главная картинка то значин в диалоге нет разделения на фракции
	if p.pictureReplics == "" && p.pictureExplores == "" && p.pictureReverses == "" {
		return p.picture
	}

	if typePic == "main" {
		return p.picture
	}

	if typePic == "Replics" {
		return p.pictureReplics
	}

	if typePic == "Explores" {
		return p.pictureExplores
	}

	if typePic == "Reverses" {
		return p.pictureReverses
	}

	return ""
}

func (p *Page) GetAllPicture() map[string]string {
	pictures := make(map[string]string)
	if p.picture != "" {
		pictures["main"] = p.picture
	}
	if p.pictureReplics != "" {
		pictures["Replics"] = p.pictureReplics
	}
	if p.pictureExplores != "" {
		pictures["Explores"] = p.pictureExplores
	}
	if p.pictureReverses != "" {
		pictures["Reverses"] = p.pictureReverses
	}

	return pictures
}

type Ask struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`        // текст ответа
	ToPage     int    `json:"to_page"`     // страница на которую ведет ответ
	TypeAction string `json:"type_action"` // функция которая выолнается при выборе этого варианта ответа
}
