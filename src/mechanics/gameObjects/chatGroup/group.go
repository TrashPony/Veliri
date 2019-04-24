package chatGroup

import "time"

type Group struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Local    bool         `json:"local"`
	Public   bool         `json:"public"` /* публичный чат любой может войти */
	Users    map[int]bool `json:"users"`  /* [id] online */
	Password string       `json:"password"`
	History  []*Message   `json:"history"`
	Fraction string       `json:"fraction"`
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
	UserName   string    `json:"user_name"`
	AvatarIcon string    `json:"avatar_icon"`
	Message    string    `json:"message"`
	Time       time.Time `json:"time"`
}
