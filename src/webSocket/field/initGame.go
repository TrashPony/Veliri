package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
	"strconv"
)

func InitGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]
	client, _ :=usersFieldWs[ws]

	if !ok {
		var newGame ActiveGame
		gameStat := objects.GetGame(msg.IdGame)
		userStat := objects.GetUserStat(msg.IdGame)
		mp := objects.GetMap(gameStat.IdMap)
		units := objects.GetAllUnits(msg.IdGame)

		newGame.addPlayers(userStat) // добавляем параметры всех игроков к обьекту игры
		newGame.addStat(&gameStat)   // добавляем информацию об игре в обьект игры
		newGame.addMap(&mp)          // добавляем информацию об карте
		newGame.addUnits(units)

		Games[newGame.Stat.Id] = &newGame   // добавляем новую игру в карту активных игор
		Game = &newGame
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

	var respawnParameter= sendCoordinate{Event: "InitResp", UserName: client.Login, X: respawn.X, Y: respawn.Y}
	coordiante <- respawnParameter

	client.updateWatchZone(Game.getUnits())
	client.GameID = Game.Stat.Id // добавляем принадлежность игрока в игре
	SendWatchCoordinate(client)
}
