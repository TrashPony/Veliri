package main

import (
	"github.com/TrashPony/Veliri/src/auth"
	"github.com/TrashPony/Veliri/src/end_points"
	globalGameGenerators "github.com/TrashPony/Veliri/src/mechanics/globalGame/generators"
	"github.com/TrashPony/Veliri/src/webSocket"
	"github.com/TrashPony/Veliri/src/webSocket/global"
	"github.com/TrashPony/Veliri/src/webSocket/global/ai"
	"github.com/TrashPony/Veliri/src/webSocket/lobby"
	"github.com/TrashPony/Veliri/src/webSocket/other"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/login", auth.Login) // если заходят на /login то отрабатывает функция auth.Login

	router.HandleFunc("/upload", end_points.Upload)                        // метод для загрузки файлов на сервер
	router.HandleFunc("/avatar", end_points.GetUserAvatar)                 // метод для взятия аватарок игроков
	router.HandleFunc("/chat_group_avatar", end_points.GetChatGroupAvatar) // метод для взятия аватарок чат груп
	router.HandleFunc("/get_picture_dialog", end_points.GetPictureDialog)  // метод для взятия аватарок чат груп

	router.HandleFunc("/registration", auth.Registration)
	router.HandleFunc("/wsLobby", webSocket.HandleConnections)
	router.HandleFunc("/wsInventory", webSocket.HandleConnections)
	router.HandleFunc("/wsField", webSocket.HandleConnections)
	router.HandleFunc("/wsGlobal", webSocket.HandleConnections)
	router.HandleFunc("/wsChat", webSocket.HandleConnections)
	router.HandleFunc("/wsMarket", webSocket.HandleConnections)
	router.HandleFunc("/wsStorage", webSocket.HandleConnections)

	router.HandleFunc("/wsMapEditor", webSocket.HandleConnections)
	router.HandleFunc("/wsDialogEditor", webSocket.HandleConnections)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/"))) // раздача статичный файлов

	globalGameGenerators.GenerateObjectsMap()     // заполняем карты динамическими обьектами
	globalGameGenerators.UpdateMapZoneCollision() // заполняем зоны проходимости на карте

	go lobby.ReposeSender()
	go lobby.WorkerChecker()
	go lobby.BaseStatusSender()

	go other.ReposeSender()
	go other.UserOnlineChecker()
	go other.LocalChatChecker()
	go other.NotifyWorker()

	go global.MoveSender()

	go ai.AnomaliesLife() // запускает работу аномалий на карте
	go ai.SkyGenerator()  // запускает генерацию облаков на картах, небо тоже немножко аи)
	go ai.HandlersLife()  // мониторинг входов выходов секторов

	go ai.EvacuationsLife() // простенький аи для эвакуаторов на базах
	// TODO go ai.InitAI() // запускает ботов

	port := "8080"
	log.Println("http server started on :" + port)
	err := http.ListenAndServe(":"+port, router) // запускает веб сервер на 8080 порту
	if err != nil {
		log.Panic(err)
	}
}
