package main

import (
	"github.com/TrashPony/Veliri/src/auth"
	globalGameGenerators "github.com/TrashPony/Veliri/src/mechanics/globalGame/generators"
	"github.com/TrashPony/Veliri/src/uploadFiles"
	"github.com/TrashPony/Veliri/src/webSocket"
	"github.com/TrashPony/Veliri/src/webSocket/chat"
	"github.com/TrashPony/Veliri/src/webSocket/field"
	"github.com/TrashPony/Veliri/src/webSocket/globalGame"
	"github.com/TrashPony/Veliri/src/webSocket/globalGame/ai"
	"github.com/TrashPony/Veliri/src/webSocket/lobby"
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

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/"))) // раздача статичный файлов

	globalGameGenerators.GenerateObjectsMap()

	go lobby.ReposeSender()
	go lobby.WorkerChecker()

	go chat.CommonChatSender()
	go chat.UserOnlineChecker()
	go chat.LocalChatChecker()
	go chat.NotifyWorker()

	go field.Sender()

	go globalGame.MoveSender()

	go ai.AnomaliesLife()   // запускает работу аномалий на карте
	go ai.SkyGenerator()    // запускает генерацию облаков на картах, небо тоже немножко аи)
	go ai.HandlersLife()    // мониторинг входов выходов секторов
	go ai.EvacuationsLife() // простенький аи для эвакуаторов на базах
	go ai.InitAI()          // запускает ботов

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", router) // запускает веб сервер на 8080 порту
	if err != nil {
		log.Panic(err)
	}
}
