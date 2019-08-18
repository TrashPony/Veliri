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
	// %UserName% %BaseName% %ToBaseName%

	if d == nil {
		return
	}

	// TODO
	//if d.Fraction == "All" {
	//	d.Fraction = userFraction
	//	for _, page := range d.Pages {
	//		if page.Picture == "" {
	//			page.Picture = strings.ToLower(userFraction) + "_logo"
	//		}
	//	}
	//}

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
