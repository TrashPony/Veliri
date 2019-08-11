package chatGroup

import "time"

type Group struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Local    bool         `json:"local"`
	Public   bool         `json:"public"` /* публичный чат любой может войти */
	Users    map[int]bool `json:"users"`  /* [id] online */
	password string
	Secure   bool       `json:"secure"`
	History  []*Message `json:"history"`
	Fraction string     `json:"fraction"`
	Private  bool       `json:"private"` // приватные чаты 1 на 1, живут до тех пор пока кто то не отписался

	avatar       string
	Greetings    string `json:"greetings"`
	UserCreate   bool   `json:"user_create"`
	UserIdCreate int    `json:"user_id_create"`

	// ключь для приватного чата, что бы если 1 игрок вышел, а потом начал заного не создавать новую группу, если
	// выйдут оба конечно группа будет удалена
	PrivateKey string `json:"private_key"`
}

func (group *Group) SetPassword(password string) {
	if password != "" {
		group.Secure = true
	}
	group.password = password
}

func (group *Group) GetPassword() string {
	return group.password
}

func (group *Group) SetAvatar(avatar string) {
	group.avatar = avatar
}

func (group *Group) GetAvatar() string {
	return group.avatar
}

func (group *Group) CheckUserInGroup(userID int) bool {
	for id := range group.Users {
		if userID == id && group.Users[id] {
			return true
		}
	}
	return false
}

type Message struct {
	UserName string    `json:"user_name"`
	UserID   string    `json:"user_id"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	System   bool      `json:"system"` // системыне сообщения это сообщения которая пишет в чат бекенд, но он не очень общителен ;С
}
