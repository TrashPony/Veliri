package field

import (
	"../../game/objects"
	"github.com/gorilla/websocket"
)

func toGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]
	client, _ := usersFieldWs[ws]

	if !ok {
		Game = initGame(msg)
	}

	for _, userStat := range Game.getPlayers() {
		if userStat.Name == client.Login {

			client.setRespawn(Game.structure[userStat.RespX][userStat.RespY])

			var playersParam = FieldResponse{Event: "InitPlayer", UserName: client.Login, PlayerPrice: userStat.Price,
				GameStep: Game.stat.Step, GamePhase: Game.stat.Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока
		}
	}

	var mapParam = FieldResponse{Event: "InitMap", UserName: client.Login, NameMap: Game.mapInfo.Name, TypeMap: Game.mapInfo.Type, XMap: Game.mapInfo.Xsize, YMap: Game.mapInfo.Ysize}
	fieldPipe <- mapParam // отправляем параметры карты

    for _, xline := range Game.coordinate {
    	for _, coordinate := range xline {
    		if coordinate.Type == "obstacle"{
    			var obstacle = sendCoordinate{Event: "InitObstacle", UserName: client.Login, X: coordinate.X, Y: coordinate.Y}
				coordiante <- obstacle
			}
		}
	}

	client.updateWatchZone(Game.getUnits(), Game.getStructure())
	client.GameID = Game.stat.Id // добавляем принадлежность игрока в игре
}

func initGame(msg FieldMessage) (newGame *ActiveGame) {
	newGame = &ActiveGame{}

	gameStat := objects.GetGame(msg.IdGame)
	userStat := objects.GetUserStat(msg.IdGame)
	infoMap := objects.GetInfoMap(gameStat.IdMap)
	units := objects.GetAllUnits(msg.IdGame)
	coordinate := objects.GetMap(infoMap.Id)
	structure := objects.GetAllStrcuture(msg.IdGame)

	newGame.setPlayers(userStat)     // добавляем параметры всех игроков к обьекту игры
	newGame.setStat(&gameStat)       // добавляем информацию об игре в обьект игры
	newGame.setInfoMap(&infoMap)     // добавляем информацию об карте
	newGame.setUnits(units)          // добавляем имеющихся юнитов
	newGame.setMap(coordinate)	     // добавляем 1 слой карты отвечающий за фон текстур, препятсвия и расположение респаунов
    newGame.setStructure(structure)  // добавляем в игру все структуры на карте

	Games[newGame.stat.Id] = newGame // добавляем новую игру в карту активных игор
	return
}
