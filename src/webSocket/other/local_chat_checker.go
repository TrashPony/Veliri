package other

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

func LocalChatChecker() {
	// воркет сделит что бы у игроков был открыть нужный локальный чат, и обновляет в нем игроков

	for {

		for _, localGroup := range chats.Groups.GetAllLocalGroups() {
			update := false
			for id := range localGroup.Users {
				update = checkUserOnline(localGroup, id, true)
			}

			if update {
				SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil,
					getUsersInChatGroup(localGroup, false), true, nil, nil, nil, nil)
			}
		}

		users, mx := chat.Clients.GetAllConnects()

		// делаем копию карты что бы не вызвать дедлок или рантайм ошибку конкурентного чтения записи.
		fakeUsers := make(map[*websocket.Conn]*player.Player)
		for key, value := range users {
			fakeUsers[key] = value
		}

		mx.Unlock()

		for _, client := range fakeUsers {

			_, id := getLocalChat(client)

			for localID, localGroup := range chats.Groups.GetAllLocalGroups() {
				// если текущий локальный чат не совподает с прошлым, то удаляем игрока и обновляем у всех список юзеров
				if id != localID && localGroup.CheckUserInGroup(client.GetID()) {
					delete(localGroup.Users, client.GetID())
					SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil,
						getUsersInChatGroup(localGroup, false), true, nil, nil, nil, nil)
				}

				// если текущий пользователь не находится в группе или он там офлайн притом что это его группа то обновляем статус онлайн у всех
				if id == localID && !localGroup.CheckUserInGroup(client.GetID()) {
					localGroup.Users[client.GetID()] = true
					SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil,
						getUsersInChatGroup(localGroup, false), true, nil, nil, nil, nil)
				}
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func getLocalChat(client *player.Player) (*chatGroup.Group, string) {
	if client.InBaseID > 0 {
		// игрок на базе
		return chats.Groups.GetLocalGroup("base:" + strconv.Itoa(client.InBaseID)), "base:" + strconv.Itoa(client.InBaseID)
	}

	// игрок на глобальной карте
	return chats.Groups.GetLocalGroup("map:" + strconv.Itoa(client.GetSquad().MatherShip.MapID)), "map:" + strconv.Itoa(client.GetSquad().MatherShip.MapID)
}
