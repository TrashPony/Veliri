package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"log"
)

func Dialogs() map[int]dialog.Dialog {
	dialogs := make(map[int]dialog.Dialog)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name, " +
		" access_type" +
		" " +
		"FROM dialogs ")
	if err != nil {
		log.Fatal("get all dialogs " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		dialogType := dialog.Dialog{}
		err := rows.Scan(&dialogType.ID, &dialogType.Name, &dialogType.AccessType)
		if err != nil {
			log.Fatal("get scan dialogs " + err.Error())
		}

		getDialogPages(&dialogType)

		dialogs[dialogType.ID] = dialogType
	}

	return dialogs
}

func getDialogPages(gameDialog *dialog.Dialog) {
	gameDialog.Pages = make(map[int]dialog.Page)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" id,"+
		" name,"+
		" text,"+
		" picture, "+
		" number"+
		" "+
		"FROM dialog_pages "+
		"WHERE id_dialog=$1", gameDialog.ID)
	if err != nil {
		log.Fatal("get all dialog pages " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		page := dialog.Page{}
		err := rows.Scan(&page.ID, &page.Name, &page.Text, &page.Picture, &page.Number)
		if err != nil {
			log.Fatal("get scan all dialog pages " + err.Error())
		}

		getPageAsk(&page)
		gameDialog.Pages[page.Number] = page
	}
}

func getPageAsk(page *dialog.Page) {
	page.Asc = make([]dialog.Ask, 0)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" id,"+
		" to_page,"+
		" name,"+
		" text,"+
		" type_action"+
		" "+
		"FROM dialog_asc "+
		"WHERE id_page=$1", page.ID)
	if err != nil {
		log.Fatal("get all page ask " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		ask := dialog.Ask{}
		action := ""

		err := rows.Scan(&ask.ID, &ask.ToPage, &ask.Name, &ask.Text, &action)
		if err != nil {
			log.Fatal("get scan all page ask " + err.Error())
		}

		ask.SetAction(action)
		page.Asc = append(page.Asc, ask)
	}
}
