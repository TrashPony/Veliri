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

	var mapParam = FieldResponse{Event: "InitMap", UserName: client.Login, NameMap: Game.Map.Name, TypeMap: Game.Map.Type, XMap: Game.Map.Xsize, YMap: Game.Map.Ysize}
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

	client.updateWatchZone(Game.getUnits())
	client.GameID = Game.Stat.Id // добавляем принадлежность игрока в игре
	SendWatchCoordinate(client)
}

func initGame(msg FieldMessage) (newGame *ActiveGame) {
	gameStat := objects.GetGame(msg.IdGame)
	userStat := objects.GetUserStat(msg.IdGame)
	infoMap := objects.GetInfoMap(gameStat.IdMap)
	units := objects.GetAllUnits(msg.IdGame)

	newGame.setPlayers(userStat) // добавляем параметры всех игроков к обьекту игры
	newGame.setStat(&gameStat)   // добавляем информацию об игре в обьект игры
	newGame.setInfoMap(&infoMap) // добавляем информацию об карте
	newGame.setUnits(units)      // добавляем имеющихся юнитов

	Games[newGame.Stat.Id] = newGame // добавляем новую игру в карту активных игор
	return
}
