package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
	"strconv"
)

func InitGame(msg FieldMessage, ws *websocket.Conn) {

	Game, ok := Games[msg.IdGame]

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
		if userStat.Name == usersFieldWs[ws].Login {
			var playersParam = FieldResponse{Event: "InitPlayer", UserName: usersFieldWs[ws].Login, PlayerPrice: userStat.Price,
				GameStep: Game.Stat.Step, GamePhase: Game.Stat.Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока
		}
	}

	var mapParam = FieldResponse{Event: "InitMap", UserName: usersFieldWs[ws].Login, NameMap: Game.Map.Name, TypeMap: Game.Map.Type, XMap: Game.Map.Xsize, YMap: Game.Map.Ysize}
	fieldPipe <- mapParam // отправляем параметры карты

	respawn := objects.GetRespawns(usersFieldWs[ws].Id, msg.IdGame)
	usersFieldWs[ws].Respawn = respawn
	permitCoordinates := mechanics.GetCoordinates(respawn.X, respawn.Y, 2)
	usersFieldWs[ws].CreateZone = make(map[string]*objects.Coordinate)

	for i := 0; i < len(permitCoordinates); i++ {
		if !(permitCoordinates[i].X == respawn.X && permitCoordinates[i].Y == respawn.Y) {
			usersFieldWs[ws].CreateZone[strconv.Itoa(permitCoordinates[i].X)+":"+strconv.Itoa(permitCoordinates[i].Y)] = permitCoordinates[i]
			var emptyCoordinates= Coordinate{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: permitCoordinates[i].X, Y: permitCoordinates[i].Y}
			coordiante <- emptyCoordinates
		}
	}

	var respawnParameter= FieldResponse{Event: "InitResp", UserName: usersFieldWs[ws].Login, RespawnX: respawn.X, RespawnY: respawn.Y}
	fieldPipe <- respawnParameter

	client, _ :=usersFieldWs[ws]
	client.getAllWatchObject(Game.getUnits())
	client.GameID = Game.Stat.Id // добавляем принадлежность игрока в игре
	SendWatchCoordinate(usersFieldWs[ws])
}
