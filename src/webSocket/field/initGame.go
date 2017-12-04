package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func toGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]
	client, _ := usersFieldWs[ws]

	if !ok {
		Game = game.InitGame(msg.IdGame)
		Games[Game.GetStat().Id] = Game // добавляем новую игру в карту активных игор
	}

	structures := Game.GetStructure()

	for _, userStat := range Game.GetPlayers() {
		if userStat.Name == client.GetLogin() {

			client.SetRespawn(structures[userStat.RespX][userStat.RespY])

			var playersParam = FieldResponse{Event: "InitPlayer", UserName: client.GetLogin(), PlayerPrice: userStat.Price,
				GameStep: Game.GetStat().Step, GamePhase: Game.GetStat().Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока
		}
	}

	var mapParam = FieldResponse{Event: "InitMap", UserName: client.GetLogin(), NameMap: Game.GetMap().Name, TypeMap: Game.GetMap().Type, XMap: Game.GetMap().Xsize, YMap: Game.GetMap().Ysize}
	fieldPipe <- mapParam // отправляем параметры карты

    for _, xline := range Game.GetMap().OneLayerMap {
    	for _, coordinate := range xline {
    		if coordinate.Type == "obstacle"{
    			var obstacle = sendCoordinate{Event: "InitObstacle", UserName: client.GetLogin(), X: coordinate.X, Y: coordinate.Y}
				coordiante <- obstacle
			}
		}
	}

	client.updateWatchZone(Game)
	client.SetGameID = Game.GetStat().Id // добавляем принадлежность игрока в игре
}