package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"time"
)

func toGame(msg Message, ws *websocket.Conn) {
	Game, ok := Games[msg.IdGame]
	client, _ := usersFieldWs[ws]

	if !ok {
		Game = game.InitGame(msg.IdGame)
		Games[Game.GetStat().Id] = Game // добавляем новую игру в карту активных игор
	}


	for _, player := range Game.GetPlayers() {
		if player.ID == client.GetID() {
			client.SetEquip(player.Equip) // добавляем эквип
			var playersParam = Response{Event: "InitPlayer", UserName: client.GetLogin(), User: player}
			ws.WriteJSON(playersParam) // отправляет параметры игроков
		}
	}

	var mapParam = Response{Event: "InitMap", UserName: client.GetLogin(), Map: Game.GetMap()}
	ws.WriteJSON(mapParam) // отправляем параметры карты

	time.Sleep(1000 * time.Millisecond) // todo хз зачем

	UpdateWatchZone(client, Game, nil) // добавляем всех юнитов и мазершип в игрока
	client.SetGameID(Game.GetStat().Id) // добавляем принадлежность игрока в игре
}
