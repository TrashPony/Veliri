package field

import (
	"strconv"
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
)

func InitGame(msg FieldMessage, ws *websocket.Conn)  {
	gameStat := objects.GetGame(msg.IdGame)
	userStat := objects.GetUserStat(msg.IdGame)
	usersFieldWs[ws].Players = userStat // добавляем параметры всех игроков к обьекту пользователя
	for _, userStat := range usersFieldWs[ws].Players {
		if userStat.Name == usersFieldWs[ws].Login {
			var playersParam = FieldResponse{Event: "InitPlayer", UserName: usersFieldWs[ws].Login, PlayerPrice: strconv.Itoa(userStat.Price),
				GameStep: strconv.Itoa(gameStat.Step), GamePhase: gameStat.Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока
		}
	}
	mp := objects.GetMap(gameStat.IdMap)
	var mapParam= FieldResponse{Event: "InitMap", UserName: usersFieldWs[ws].Login, NameMap: mp.Name, TypeMap: mp.Type, XMap: strconv.Itoa(mp.Xsize), YMap: strconv.Itoa(mp.Ysize)}
	fieldPipe <- mapParam // отправляем параметры карты


	respawn := objects.GetRespawns(usersFieldWs[ws].Id, msg.IdGame)
	usersFieldWs[ws].Respawn = respawn
	permitCoordinates := mechanics.GetCoordinates(respawn.X, respawn.Y, 2)

	for i := 0; i < len(permitCoordinates); i++ {
		if  !(permitCoordinates[i].X == respawn.X && permitCoordinates[i].Y == respawn.Y) {
			usersFieldWs[ws].PermittedCoordinates = append(usersFieldWs[ws].PermittedCoordinates, permitCoordinates[i])
			usersFieldWs[ws].CreateZone = append(usersFieldWs[ws].CreateZone, permitCoordinates[i])
			var emptyCoordinates= FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: strconv.Itoa(permitCoordinates[i].X), Y: strconv.Itoa(permitCoordinates[i].Y)}
			fieldPipe <- emptyCoordinates
		}
	}


	var respawnParametr = FieldResponse{Event: "InitResp", UserName: usersFieldWs[ws].Login, RespawnX: strconv.Itoa(respawn.X), RespawnY: strconv.Itoa(respawn.Y)}
	fieldPipe <- respawnParametr

	units := objects.GetAllUnits(msg.IdGame)
	for i := 0; i < len(units); i++ {
		unitPermissCoordinate := sendPermissionCoordinates(msg.IdGame, ws, units[i])
		usersFieldWs[ws].Units = append(usersFieldWs[ws].Units, units[i])
		for j := 0; j < len(unitPermissCoordinate); j++ {
			usersFieldWs[ws].PermittedCoordinates = append(usersFieldWs[ws].PermittedCoordinates, unitPermissCoordinate[j])
		}
	}
}
