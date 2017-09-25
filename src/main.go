package main

import (
	"net/http"
	"mux-master"
	"log"
	"./auth"
	"./webSocket"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login", auth.Login) // если заходят на /login то отрабатывает функция auth.Login
	router.HandleFunc("/wsLobby", auth.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src/static/"))) // раздача статичный файлов
	go webSocket.LobbySender() // запускается гарутина для рассылки сообщений, гуглить гарутины
	log.Println("http server started on :8080")
	http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
}



