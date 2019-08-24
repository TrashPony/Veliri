package other

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
	"time"
)

// бот который оповещает игроков об изменениях на них панеле новостей и заметок
// все имущество на базах
// на панеле отображаются активные задания
// сюда падают нотификация если была совершена сделка торговли или завершился крафт

func NotifyWorker() {
	for {
		users, mx := chat.Clients.GetAllConnects()
		// делаем копию карты что бы не вызвать дедлок или рантайм ошибку конкурентного чтения записи.
		fakeUsers := make(map[*websocket.Conn]*player.Player)
		for key, value := range users {
			fakeUsers[key] = value
		}
		mx.Unlock()

		for _, user := range fakeUsers {
			for _, notify := range user.NotifyQueue {
				if notify != nil && !notify.Send {

					SendMessage("newNotify", nil, user.GetID(), 0, nil, nil,
						nil, false, nil, nil, notify, nil)
					notify.Send = true

					if notify.Name == "mission" && notify.Event == "complete" {
						//delete(user.NotifyQueue, notify.UUID)
					}
				}
			}
		}
		time.Sleep(time.Second)
	}
}

func DeleteNotify(user *player.Player, uuid string) bool {
	notify, ok := user.NotifyQueue[uuid]
	if ok && notify != nil {
		if notify.Name == "mission" && notify.Event != "complete" {
			// незавершенные миссии нельзя удалить с панели // todo проверять наличие миссии у игрока
		} else {
			delete(user.NotifyQueue, notify.UUID)
			dbPlayer.UpdateUser(user)
			return true
		}
	}

	return false
}
