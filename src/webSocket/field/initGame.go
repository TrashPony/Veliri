package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"time"
)

func toGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]
	client, _ := usersFieldWs[ws]

	if !ok {
		Game = game.InitGame(msg.IdGame)
		Games[Game.GetStat().Id] = Game // добавляем новую игру в карту активных игор
	}

	var playersParam = FieldResponse{Event: "InitPlayer", UserName: client.GetLogin(), Users: Game.GetPlayers()}
	ws.WriteJSON(playersParam) // отправляет параметры игрока

	var mapParam = FieldResponse{Event: "InitMap", UserName: client.GetLogin(), Map: Game.GetMap()}
	ws.WriteJSON(mapParam) // отправляем параметры карты

	time.Sleep(1000 * time.Millisecond) // todo хз зачем

	UpdateWatchZone(client, Game, nil)
	client.SetGameID(Game.GetStat().Id) // добавляем принадлежность игрока в игре
}
