package end_points

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"net/http"
	"strconv"
)

func GetPictureDialog(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		dialogID := r.URL.Query().Get("dialog_page_id")
		userID := r.URL.Query().Get("user_id")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Cache-Control", "max-age=60")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Vary", "Accept-Encoding")

		id, err := strconv.Atoi(dialogID)
		if err != nil {
			return
		}

		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			return
		}

		dialogPage := gameTypes.Dialogs.GetDialogPageByID(id)

		if userIDint == -1 {
			//это для редактора что бы отдать все картинки
			jsonData, _ := json.Marshal(dialogPage.GetAllPicture())
			w.Write(jsonData)
		} else {
			user, ok := players.Users.Get(userIDint)
			if dialogPage != nil && ok {
				w.Write([]byte(`{"picture": "` + dialogPage.GetPicture(user.Fraction) + `"}`))
			}
		}
	}
}
