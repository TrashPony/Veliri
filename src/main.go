package main

import (
	"./auth"
	"./webSocket"
	"./webSocket/field"
	"./webSocket/lobby"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", auth.Login)                    // если заходят на /login то отрабатывает функция auth.Login
	router.HandleFunc("/registration", auth.Registration)
	router.HandleFunc("/wsLobby", webSocket.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	router.HandleFunc("/wsField", webSocket.HandleConnections)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src/static/"))) // раздача статичный файлов
	go lobby.LobbyReposeSender()                                               // запускается гарутина для рассылки сообщений, гуглить гарутины
	go field.FieldReposeSender()
	go field.InitUnitSender()
	go field.CoordinateSender()
	go field.InitStructureSender()
	go lobby.CommonChatSender()
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
	if err != nil {
		log.Panic(err)
	}
}
