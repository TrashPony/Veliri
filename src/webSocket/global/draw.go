package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

func CreateRect(color string, startX, startY int, rectSize, mapID int, user *player.Player) {
	go SendMessage(Message{Event: "CreateRect", Color: color, RectSize: rectSize,
		X: int(startX), Y: int(startY), IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})

	time.Sleep(20 * time.Millisecond)
}

func ClearVisiblePath(mapID int, user *player.Player) {
	go SendMessage(Message{Event: "ClearPath", IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})
}

func CreateLine(color string, X, Y, ToX, ToY int, rectSize, mapID int, user *player.Player) {
	go SendMessage(Message{Event: "CreateLine", Color: color, RectSize: rectSize,
		X: int(X), Y: int(Y), ToX: float64(ToX), ToY: float64(ToY), IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})

	//time.Sleep(20 * time.Millisecond)
}
