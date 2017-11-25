package field

import (
	"../../game/mechanics"
	"../../game/objects"
	"github.com/gorilla/websocket"
	"strconv"
)

func toGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]
	client, _ := usersFieldWs[ws]

	if !ok {
		Game = initGame(msg)
	}

	for _, userStat := range Game.getPlayers() {
		if userStat.Name == client.Login {
			var playersParam = FieldResponse{Event: "InitPlayer", UserName: client.Login, PlayerPrice: userStat.Price,
				GameStep: Game.Stat.Step, GamePhase: Game.Stat.Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока


		}
	}

	var mapParam = FieldResponse{Event: "InitMap", UserName: client.Login, NameMap: Game.MapInfo.Name, TypeMap: Game.MapInfo.Type, XMap: Game.MapInfo.Xsize, YMap: Game.MapInfo.Ysize}
	fieldPipe <- mapParam // отправляем параметры карты

	respawn := objects.GetRespawns(client.Id, msg.IdGame)
	client.Respawn = respawn

	permitCoordinates := mechanics.GetCoordinates(respawn.X, respawn.Y, 2)
	client.CreateZone = make(map[string]*objects.Coordinate)

	for _, coordinate := range permitCoordinates {
		if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
			client.CreateZone[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate
			openCoordinate(client.Login, coordinate.X, coordinate.Y)
		}
	}

	var respawnParameter = sendCoordinate{Event: "InitResp", UserName: client.Login, X: respawn.X, Y: respawn.Y}
	coordiante <- respawnParameter

    for _, xline := range Game.Coordinate {
    	for _, coordinate := range xline {
    		if coordinate.Type == "obstacle"{
    			var obstacle = sendCoordinate{Event: "InitObstacle", UserName: client.Login, X: coordinate.X, Y: coordinate.Y}
				coordiante <- obstacle
			}
		}
	}

	client.updateWatchZone(Game.getUnits())
	client.GameID = Game.Stat.Id // добавляем принадлежность игрока в игре
	SendWatchCoordinate(client)
}

func initGame(msg FieldMessage) (newGame *ActiveGame) {
	newGame = &ActiveGame{}

	gameStat := objects.GetGame(msg.IdGame)
	userStat := objects.GetUserStat(msg.IdGame)
	infoMap := objects.GetInfoMap(gameStat.IdMap)
	units := objects.GetAllUnits(msg.IdGame)
	coordinate := objects.GetMap(infoMap.Id)

	newGame.setPlayers(userStat) // добавляем параметры всех игроков к обьекту игры
	newGame.setStat(&gameStat)   // добавляем информацию об игре в обьект игры
	newGame.setInfoMap(&infoMap) // добавляем информацию об карте
	newGame.setUnits(units)      // добавляем имеющихся юнитов
	newGame.setMap(coordinate)	 // добавляем 1 слой карты отвечающий за фон текстур, препятсвия и расположение респаунов

	Games[newGame.Stat.Id] = newGame // добавляем новую игру в карту активных игор
	return
}
