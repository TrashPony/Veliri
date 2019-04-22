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
		" fraction " +
		" " +
		"FROM chats")
	if err != nil {
		log.Fatal("get all chats " + err.Error())
	}
	defer rows.Close()

	allChats := make(map[int]*chatGroup.Group)

	for rows.Next() {
		var gameChat chatGroup.Group

		err := rows.Scan(&gameChat.ID, &gameChat.Name, &gameChat.Public, &gameChat.Password, &gameChat.Fraction)
		if err != nil {
			log.Fatal("get scan all chats " + err.Error())
		}

		getUsersChat(&gameChat)
		allChats[gameChat.ID] = &gameChat
	}

	return allChats
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
