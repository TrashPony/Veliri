package end_points

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"net/http"
	"strconv"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		userID := r.URL.Query().Get("user_id")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Cache-Control", "max-age=60")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Vary", "Accept-Encoding")

		id, err := strconv.Atoi(userID)
		if err != nil {
			//bot
			w.Write([]byte(`{"avatar": "data:image/png;base64,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"}`))
		} else {
			user, ok := players.Users.Get(id)
			if ok {
				w.Write([]byte(`{"avatar": "` + user.GetAvatar() + `"}`))
			}
		}
	}
}
