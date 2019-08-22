package dialogEditor

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"log"
)

func UpdateDialog(updatedDialog *dialog.Dialog, oldDialog *dialog.Dialog) {

	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE dialogs SET name = $1, type = $2, access_type = $3, fraction = $4 WHERE id = $5",
		updatedDialog.Name, updatedDialog.Type, updatedDialog.AccessType, updatedDialog.Fraction, updatedDialog.ID)
	if err != nil {
		log.Fatal("update dialog main info" + err.Error())
	}

	DeleteOldPageAndAsc(oldDialog, tx)
	AddPages(updatedDialog, tx)

	err = tx.Commit()
	if err != nil {
		log.Fatal("update dialog: " + err.Error())
	}
}

func AddDialog(newDialog *dialog.Dialog) {
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

func DeleteDialog(deleteDialog *dialog.Dialog) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	DeleteOldPageAndAsc(deleteDialog, tx)

	_, err = tx.Exec("DELETE FROM dialogs WHERE id=$1",
		deleteDialog.ID)
	if err != nil {
		log.Fatal("delete dialog" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("delete dialog: " + err.Error())
	}
}

func DeleteOldPageAndAsc(oldDialog *dialog.Dialog, tx *sql.Tx) {
	// из за лени я просту удаляю все старые страницы и ответы, после добавляю новые))
	for _, page := range oldDialog.Pages {
		_, err := tx.Exec("DELETE FROM dialog_asc WHERE id_page=$1",
			page.ID)
		if err != nil {
			log.Fatal("delete old asc dialog" + err.Error())
		}
	}

	_, err := tx.Exec("DELETE FROM dialog_pages WHERE id_dialog=$1",
		oldDialog.ID)
	if err != nil {
		log.Fatal("delete old page dialog" + err.Error())
	}
}

func AddPages(updatedDialog *dialog.Dialog, tx *sql.Tx) {
	for _, page := range updatedDialog.Pages {

		err := tx.QueryRow("INSERT INTO dialog_pages (id_dialog, type, number, name, text, picture, picture_replics, picture_explores, picture_reverses) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
			updatedDialog.ID, page.Type, page.Number, page.Name, page.Text, page.GetPicture("main"),
			page.GetPicture("Replics"), page.GetPicture("Explores"), page.GetPicture("Reverses")).Scan(&page.ID)
		if err != nil {
			log.Fatal("add new page dialog " + err.Error())
		}

		AddAsks(page, tx)
	}
}

func AddAsks(page *dialog.Page, tx *sql.Tx) {
	for i, asc := range page.Asc {
		err := tx.QueryRow("INSERT INTO dialog_asc (id_page, to_page, name, text, type_action) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
			page.ID, asc.ToPage, asc.Name, asc.Text, asc.TypeAction).Scan(&page.Asc[i].ID)

		if err != nil {
			log.Fatal("add new asc dialog " + err.Error())
		}
	}
}
