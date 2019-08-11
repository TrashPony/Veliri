package chats

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"

	"log"
)

func Chats() map[int]*chatGroup.Group {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" public," +
		" password," +
		" fraction," +
		" private," +
		" private_key," +
		" avatar," +
		" greetings," +
		" user_create," +
		" user_id_create" +
		" " +
		"FROM chats")
	if err != nil {
		log.Fatal("get all chats " + err.Error())
	}
	defer rows.Close()

	allChats := make(map[int]*chatGroup.Group)

	for rows.Next() {
		var gameChat chatGroup.Group
		var avatar string
		var password string

		err := rows.Scan(&gameChat.ID, &gameChat.Name, &gameChat.Public, &password, &gameChat.Fraction,
			&gameChat.Private, &gameChat.PrivateKey, &avatar, &gameChat.Greetings, &gameChat.UserCreate, &gameChat.UserIdCreate)
		if err != nil {
			log.Fatal("get scan all chats " + err.Error())
		}

		gameChat.SetAvatar(avatar)
		gameChat.SetPassword(password)

		getUsersChat(&gameChat)
		gameChat.History = make([]*chatGroup.Message, 0)
		allChats[gameChat.ID] = &gameChat
	}

	return allChats
}

func AddNewGroup(gameChat *chatGroup.Group) int {
	id := 0
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
		"chats "+
		"(name, public, password, fraction, private, private_key, avatar, greetings, user_create, user_id_create) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) "+
		"RETURNING id",
		gameChat.Name, gameChat.Public, gameChat.GetPassword(), gameChat.Fraction, gameChat.Private, gameChat.PrivateKey,
		gameChat.GetAvatar(), gameChat.Greetings, gameChat.UserCreate, gameChat.UserIdCreate).Scan(&id)

	if err != nil {
		log.Fatal("add new chat group " + err.Error())
	}

	return id
}

func RemoveGroup(idChat int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM users_in_chat WHERE id_chat = $1",
		idChat)
	if err != nil {
		log.Fatal("remove all user in chat" + err.Error())
	}

	_, err = dbConnect.GetDBConnect().Exec("DELETE FROM chats WHERE id = $1",
		idChat)
	if err != nil {
		log.Fatal("remove chat group" + err.Error())
	}
}

func getUsersChat(gameChat *chatGroup.Group) {
	gameChat.Users = make(map[int]bool)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" id_user "+
		" "+
		"FROM users_in_chat WHERE id_chat=$1", gameChat.ID)
	if err != nil {
		log.Fatal("get all users in chats " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var userID int

		err := rows.Scan(&userID)
		if err != nil {
			log.Fatal("get all users " + err.Error())
		}

		gameChat.Users[userID] = false // при поднятие сервера все игроки не онлайн
	}
}

func AddUserInChat(idChat, idUser int) {
	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO users_in_chat (id_chat, id_user) VALUES ($1, $2)",
		idChat, idUser)
	if err != nil {
		log.Fatal("add new user in chat" + err.Error())
	}
}

func RemoveUserInChat(idChat, idUser int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM users_in_chat WHERE id_chat = $1 AND id_user=$2",
		idChat, idUser)
	if err != nil {
		log.Fatal("remove user in chat" + err.Error())
	}
}
