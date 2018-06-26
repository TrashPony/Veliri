package main

import (
	"./auth"
	"./webSocket"
	"./webSocket/field"
	"./webSocket/lobby"
	"./webSocket/chat"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", auth.Login) // если заходят на /login то отрабатывает функция auth.Login
	router.HandleFunc("/registration", auth.Registration)
	router.HandleFunc("/wsLobby", webSocket.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	router.HandleFunc("/wsField", webSocket.HandleConnections)
	router.HandleFunc("/wsGlobal", webSocket.HandleConnections)
	router.HandleFunc("/wsChat", webSocket.HandleConnections)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src/static/"))) // раздача статичный файлов

	go lobby.ReposeSender() // запускается гарутина для рассылки сообщений, гуглить гарутины
	go chat.CommonChatSender()

	go field.WatchSender()
	go field.MoveSender()
	go field.UnitSender()
	go field.PhaseSender()
	go field.EquipSender()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
	if err != nil {
		log.Panic(err)
	}
}
