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
	router.HandleFunc("/wsLobby", webSocket.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	router.HandleFunc("/wsField", webSocket.HandleConnections)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src/static/"))) // раздача статичный файлов
	go webSocket.LobbyReposeSender()// запускается гарутина для рассылки сообщений, гуглить гарутины
	go webSocket.FieldReposeSender()
	log.Println("http server started on :8080")
	http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
}



