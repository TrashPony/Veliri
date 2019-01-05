package main

import (
	"./auth"
	"./uploadFiles"
	"./webSocket"
	"./webSocket/chat"
	"./webSocket/field"
	"./webSocket/globalGame"
	"./webSocket/lobby"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/login", auth.Login) // если заходят на /login то отрабатывает функция auth.Login

	router.HandleFunc("/upload", uploadFiles.Upload) // метод для загрузки файлов на сервер

	router.HandleFunc("/registration", auth.Registration)
	router.HandleFunc("/wsLobby", webSocket.HandleConnections)
	router.HandleFunc("/wsInventory", webSocket.HandleConnections)
	router.HandleFunc("/wsMapEditor", webSocket.HandleConnections)
	router.HandleFunc("/wsField", webSocket.HandleConnections)
	router.HandleFunc("/wsGlobal", webSocket.HandleConnections)
	router.HandleFunc("/wsChat", webSocket.HandleConnections)
	router.HandleFunc("/wsMarket", webSocket.HandleConnections)
	router.HandleFunc("/wsStorage", webSocket.HandleConnections)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src/static/"))) // раздача статичный файлов

	go lobby.ReposeSender()
	go chat.CommonChatSender()

	go field.MoveSender()
	go field.UnitSender()
	go field.PhaseSender()
	go field.AttackSender()
	go globalGame.MoveSender()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
	if err != nil {
		log.Panic(err)
	}
}
