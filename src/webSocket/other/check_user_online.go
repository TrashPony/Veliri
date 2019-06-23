package other

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

func checkUserOnline(group *chatGroup.Group, id int, local bool) bool {

	chatUser := chat.Clients.GetByID(id)

	// если игрок онайн но игрок уже отключился от сети то говорим что он не онлайн и обновляем у всех список
	if group.Users[id] && chatUser == nil {
		if local {
			delete(group.Users, id)
		} else {
			group.Users[id] = false
		}
		return true
	}

	// если игрока был офлайн и стал онлайн то обновляем у всех список пользователей
	if !group.Users[id] && chatUser != nil {
		group.Users[id] = true
		return true
	}

	return false
}

func UserOnlineChecker() {
	for {
		for _, group := range chats.Groups.GetAllGroups() {
			update := false
			for id := range group.Users {
				update = checkUserOnline(group, id, false)
			}

			if update {
				SendMessage("UpdateUsers", nil, 0, group.ID, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func getUsersInChatGroup(group *chatGroup.Group, all bool) []*player.ShortUserInfo {
	users := make([]*player.ShortUserInfo, 0)

	for id := range group.Users {
		chatUser := chat.Clients.GetByID(id)

		if chatUser != nil {
			if all || group.Users[id] {
				users = append(users, chatUser.GetShortUserInfo(false))
			}
		} else {
			group.Users[id] = false
		}
	}

	return users
}
