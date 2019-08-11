package end_points

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"net/http"
	"strconv"
)

func GetChatGroupAvatar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		groupID := r.URL.Query().Get("chat_group_id")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Cache-Control", "max-age=60")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Vary", "Accept-Encoding")

		id, err := strconv.Atoi(groupID)
		if err != nil {
			return
		} else {
			group := chats.Groups.GetGroup(id)
			if group != nil {
				w.Write([]byte(`{"avatar": "` + group.GetAvatar() + `"}`))
			}
		}
	}
}
