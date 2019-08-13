package dialogEditor

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"log"
)

func UpdateDialog(updatedDialog *dialog.Dialog) {

	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE dialogs SET name = $1, type = $2, access_type = $3, fraction = $4 WHERE id = $5",
		updatedDialog.Name, updatedDialog.Type, updatedDialog.AccessType, updatedDialog.Fraction, updatedDialog.ID)
	if err != nil {
		log.Fatal("update dialog main info" + err.Error())
	}

	// из за лени я просту удаляю все старые страницы и добавляю новые))
	_, err = tx.Exec("DELETE FROM dialog_pages WHERE id_dialog=$1",
		updatedDialog.ID)
	if err != nil {
		log.Fatal("delete old page dialog" + err.Error())
	}

	AddPages(updatedDialog, tx)

	err = tx.Commit()
	if err != nil {
		log.Fatal("update dialog: " + err.Error())
	}
}

func AddDialog(newDialog *dialog.Dialog)  {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	err = tx.QueryRow("INSERT INTO dialogs (name, type, access_type, fraction) "+
		"VALUES ($1, $2, $3, $4) RETURNING id",
		newDialog.Name, newDialog.Type, newDialog.AccessType, newDialog.Fraction).Scan(&newDialog.ID)
	if err != nil {
		log.Fatal("add new dialog " + err.Error())
	}

	AddPages(newDialog, tx)

	err = tx.Commit()
	if err != nil {
		log.Fatal("insert dialog: " + err.Error())
	}
}

func AddPages(updatedDialog *dialog.Dialog, tx *sql.Tx) {
	for _, page := range updatedDialog.Pages {

		_, err := tx.Exec("DELETE FROM dialog_asc WHERE id_page=$1",
			page.ID)
		if err != nil {
			log.Fatal("delete old asc dialog" + err.Error())
		}

		err = tx.QueryRow("INSERT INTO dialog_pages (id_dialog, type, number, name, text, picture) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			updatedDialog.ID, page.Type, page.Number, page.Name, page.Text, page.Picture).Scan(&page.ID)
		if err != nil {
			log.Fatal("add new page dialog " + err.Error())
		}

		AddAsks(page, tx)
	}
}

func AddAsks(page *dialog.Page, tx *sql.Tx) {
	for _, asc := range page.Asc {
		err := tx.QueryRow("INSERT INTO dialog_asc (id_page, to_page, name, text, type_action) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
			page.ID, asc.ToPage, asc.Name, asc.Text, asc.TypeAction).Scan(&asc.ID)
		if err != nil {
			log.Fatal("add new asc dialog " + err.Error())
		}
	}
}
