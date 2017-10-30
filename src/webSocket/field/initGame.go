package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
	"strconv"
)

func InitGame(msg FieldMessage, ws *websocket.Conn)  {
	gameStat := objects.GetGame(msg.IdGame)
	userStat := objects.GetUserStat(msg.IdGame)
	usersFieldWs[ws].Players = userStat // добавляем параметры всех игроков к обьекту пользователя
	usersFieldWs[ws].GameStat = gameStat // добавляем информацию об игре
	for _, userStat := range usersFieldWs[ws].Players {
		if userStat.Name == usersFieldWs[ws].Login {
			var playersParam = FieldResponse{Event: "InitPlayer", UserName: usersFieldWs[ws].Login, PlayerPrice: userStat.Price,
				GameStep: gameStat.Step, GamePhase: gameStat.Phase, UserReady: userStat.Ready}
			fieldPipe <- playersParam // отправляет параметры игрока
		}
	}
	mp := objects.GetMap(gameStat.IdMap)
	usersFieldWs[ws].Map = mp
	var mapParam= FieldResponse{Event: "InitMap", UserName: usersFieldWs[ws].Login, NameMap: mp.Name, TypeMap: mp.Type, XMap: mp.Xsize, YMap: mp.Ysize}
	fieldPipe <- mapParam // отправляем параметры карты


	respawn := objects.GetRespawns(usersFieldWs[ws].Id, msg.IdGame)
	usersFieldWs[ws].Respawn = respawn
	permitCoordinates := mechanics.GetCoordinates(respawn.X, respawn.Y, 2)

	for i := 0; i < len(permitCoordinates); i++ {
		if  !(permitCoordinates[i].X == respawn.X && permitCoordinates[i].Y == respawn.Y) {
			usersFieldWs[ws].CreateZone = append(usersFieldWs[ws].CreateZone, permitCoordinates[i])
			var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: usersFieldWs[ws].Login, X: permitCoordinates[i].X, Y: permitCoordinates[i].Y}
			fieldPipe <- emptyCoordinates
		}
	}


	var respawnParametr = FieldResponse{Event: "InitResp", UserName: usersFieldWs[ws].Login, RespawnX: respawn.X, RespawnY: respawn.Y}
	fieldPipe <- respawnParametr

	units := objects.GetAllUnits(msg.IdGame)
	usersFieldWs[ws].Units = make(map[string]*objects.Unit)
	usersFieldWs[ws].HostileUnits = make(map[string]*objects.Unit)
	for _, unit := range units {
		var err error
		client, ok := usersFieldWs[ws]
		if ok {
			unit.Watch, unit.WatchUnit, err = PermissionCoordinates(*client, unit, units)

			if err != nil {
				continue
			}

			for _, hostile := range unit.WatchUnit {
				if hostile.NameUser != usersFieldWs[ws].Login {
					usersFieldWs[ws].HostileUnits[strconv.Itoa(hostile.X) + ":" + strconv.Itoa(hostile.Y)] = hostile
				} else {
					continue
				}
			}

			usersFieldWs[ws].Units[strconv.Itoa(unit.X) + ":" + strconv.Itoa(unit.Y)] = unit
			SendWatchCoordinate(ws, unit)
		}
	}
}
