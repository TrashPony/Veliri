package main

import (
	"net/http"
	"mux-master"
	"./auth"
	"log"
)

func main() {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("src/static/")) // эээ файловый сервер для раздачи статики который нехуя не работает так как разные библиотеки
	router.Handle("/", fs) // гуглить mux fileServer
	router.HandleFunc("/ws", auth.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	go auth.HandleMessages() // запускается гарутина для рассылки сообщений, гуглить гарутины

	router.HandleFunc("/", nil)
	router.HandleFunc("/login", auth.Login) // если заходят на /login то отрабатывает функция auth.Login
	log.Println("http server started on :8080")
	http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
}



