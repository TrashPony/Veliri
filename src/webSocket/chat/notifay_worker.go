package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
)

// бот который оповещает игроков об изменениях на них панеле новостей и заметок
// все имущество на базах
// на панеле отображаются активные задания
// сюда падают нотификация если была совершена сделка торговли или завершился крафт

func NotifayWorker() {
	users, mx := chat.Clients.GetAllConnects()
	// делаем копию карты что бы не вызвать дедлок или рантайм ошибку конкурентного чтения записи.
	fakeUsers := make(map[*websocket.Conn]*player.Player)
	for key, value := range users {
		fakeUsers[key] = value
	}

	mx.Unlock()
}
