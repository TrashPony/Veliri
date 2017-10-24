package webSocket

import (
	"log"
	"websocket-master"
	"../game"
	"../game/initGame"
	"strconv"
)

func FieldReader(ws *websocket.Conn)  {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs , err)
			break
		}

		if msg.Event == "InitGame" {
			InitGame(msg, ws)
		}

		if msg.Event == "SelectCoordinateCreate" {
			SelectCoordinateCreate(ws)
		}

		if msg.Event == "CreateUnit" {
			CreateUnit(msg, ws)
		}

		if msg.Event == "MouseOver" {
			MouseOver(msg, ws)
		}

		if msg.Event == "Ready" {
			Ready(msg, ws)
		}
	}
}

func FieldReposeSender() {
	for {
		resp := <-FieldPipe // TODO : разделить пайп на множество под каждую фазу
		for ws, client := range usersFieldWs {
			if client.login == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
	}
}

func sendPermissionCoordinates(idGame string, ws *websocket.Conn, unit initGame.Unit) ([]initGame.Coordinate) {
	units := initGame.GetAllUnits(idGame)
	var allCoordinate []initGame.Coordinate
	login := usersFieldWs[ws].login
	respawn := usersFieldWs[ws].Respawn

	if login == unit.NameUser {
		PermissCoordinates := game.GetCoordinates(unit.X, unit.Y, unit.WatchZone)
		for i := 0; i < len(PermissCoordinates); i++ {
			allCoordinate = append(allCoordinate, PermissCoordinates[i])
			if !(PermissCoordinates[i].X == respawn.X && PermissCoordinates[i].Y == respawn.Y) {
				var emptyCoordinates = FieldResponse{Event: "emptyCoordinate", UserName: login, X: strconv.Itoa(PermissCoordinates[i].X), Y: strconv.Itoa(PermissCoordinates[i].Y)}
				FieldPipe <- emptyCoordinates
			}
			for j := 0; j < len(units); j++ {
				if (PermissCoordinates[i].X == units[j].X) && (PermissCoordinates[i].Y == units[j].Y) {
					var unitsParametr = FieldResponse{Event: "InitUnit", UserName: login, TypeUnit: units[j].NameType, UserOwned: units[j].NameUser,
						HP: strconv.Itoa(units[j].Hp), UnitAction: strconv.FormatBool(units[j].Action), Target: strconv.Itoa(units[j].Target), X: strconv.Itoa(units[j].X), Y: strconv.Itoa(units[j].Y)}
					FieldPipe <- unitsParametr // отправляем параметры каждого юнита отдельно
				}
			}
		}
	}
	return allCoordinate
}

func InitGame(msg FieldMessage, ws *websocket.Conn)  {
	gameStat := initGame.GetGame(msg.IdGame)
	userStat := initGame.GetUserStat(msg.IdGame)
	usersFieldWs[ws].Players = userStat // добавляем параметры всех игроков к обьекту пользователя
	for _, userStat := range usersFieldWs[ws].Players {
		if userStat.Name == LoginWs(ws, &usersFieldWs) {
			var playersParam = FieldResponse{Event: "InitPlayer", UserName: LoginWs(ws, &usersFieldWs), PlayerPrice: strconv.Itoa(userStat.Price),
				GameStep: strconv.Itoa(gameStat.Step), GamePhase: gameStat.Phase, UserReady: userStat.Ready}
			FieldPipe <- playersParam // отправляет параметры игрока
		}
	}
	mp := initGame.GetMap(gameStat.IdMap)
	var mapParam= FieldResponse{Event: "InitMap", UserName: LoginWs(ws, &usersFieldWs), NameMap: mp.Name, TypeMap: mp.Type, XMap: strconv.Itoa(mp.Xsize), YMap: strconv.Itoa(mp.Ysize)}
	FieldPipe <- mapParam // отправляем параметры карты


	respawn := initGame.GetRespawns(IdWs(ws, &usersFieldWs), msg.IdGame)
	usersFieldWs[ws].Respawn = respawn
	permitCoordinates := game.GetCoordinates(respawn.X, respawn.Y, 2)

	for i := 0; i < len(permitCoordinates); i++ {
		if  !(permitCoordinates[i].X == respawn.X && permitCoordinates[i].Y == respawn.Y) {
			usersFieldWs[ws].permittedCoordinates = append(usersFieldWs[ws].permittedCoordinates, permitCoordinates[i])
			usersFieldWs[ws].CreateZone = append(usersFieldWs[ws].CreateZone, permitCoordinates[i])
			var emptyCoordinates= FieldResponse{Event: "emptyCoordinate", UserName: LoginWs(ws, &usersFieldWs), X: strconv.Itoa(permitCoordinates[i].X), Y: strconv.Itoa(permitCoordinates[i].Y)}
			FieldPipe <- emptyCoordinates
		}
	}


	var respawnParametr = FieldResponse{Event: "InitResp", UserName: LoginWs(ws, &usersFieldWs), RespawnX: strconv.Itoa(respawn.X), RespawnY: strconv.Itoa(respawn.Y)}
	FieldPipe <- respawnParametr

	units := initGame.GetAllUnits(msg.IdGame)
	for i := 0; i < len(units); i++ {
		unitPermissCoordinate := sendPermissionCoordinates(msg.IdGame, ws, units[i])
		usersFieldWs[ws].Units = append(usersFieldWs[ws].Units, units[i])
		for j := 0; j < len(unitPermissCoordinate); j++ {
			usersFieldWs[ws].permittedCoordinates = append(usersFieldWs[ws].permittedCoordinates, unitPermissCoordinate[j])
		}
	}
}

func SelectCoordinateCreate(ws *websocket.Conn)  {
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	units := usersFieldWs[ws].Units
	unitsCoordinate := initGame.GetUnitsCoordinate(units)
	responseCoordinate := subtraction(coordinates, unitsCoordinate)

	for i := 0; i < len(responseCoordinate); i++ {
		if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
			var createCoordinates= FieldResponse{Event: "SelectCoordinateCreate", UserName: LoginWs(ws, &usersFieldWs), X: strconv.Itoa(responseCoordinate[i].X), Y: strconv.Itoa(responseCoordinate[i].Y)}
			FieldPipe <- createCoordinates
		}
	}

	for i := 0; i < len(coordinates); i++ {

	}
}
func CreateUnit(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].CreateZone
	respawn := usersFieldWs[ws].Respawn
	var errorMsg bool = true
	for i := 0; i < len(coordinates); i++ {
		if strconv.Itoa(coordinates[i].X) == msg.X && strconv.Itoa(coordinates[i].Y) == msg.Y &&
			!(msg.X == strconv.Itoa(respawn.X) && msg.Y == strconv.Itoa(respawn.Y)){
			errorMsg = false
			unit, price, createError := game.CreateUnit(msg.IdGame, strconv.Itoa(IdWs(ws, &usersFieldWs)), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {
				for _, userStat := range usersFieldWs[ws].Players {
					if userStat.Name == LoginWs(ws, &usersFieldWs) {
						resp = FieldResponse{Event: msg.Event, UserName: userStat.Name, PlayerPrice: strconv.Itoa(price), X: strconv.Itoa(unit.X), Y: strconv.Itoa(unit.Y), TypeUnit: unit.NameType, UserOwned: LoginWs(ws, &usersFieldWs)}
						FieldPipe <- resp
					}
				}
				unitPermissCoordinate := sendPermissionCoordinates(msg.IdGame, ws, unit)
				for i := 0; i < len(unitPermissCoordinate); i++ {
					usersFieldWs[ws].permittedCoordinates = append(usersFieldWs[ws].permittedCoordinates, unitPermissCoordinate[i])
				}
				usersFieldWs[ws].Units = append(usersFieldWs[ws].Units, unit)
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), X: msg.X, Y: msg.Y, ErrorType: createError.Error()}
				FieldPipe <- resp
			}
			break
		}
	}
	if errorMsg {
		resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), X: msg.X, Y: msg.Y, ErrorType: "not allow"}
		FieldPipe <- resp
	}
}

func MouseOver(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse
	coordinates := usersFieldWs[ws].permittedCoordinates
	units := usersFieldWs[ws].Units
	for i := 0; i < len(coordinates); i++{
		if strconv.Itoa(coordinates[i].X) == msg.X && strconv.Itoa(coordinates[i].Y) == msg.Y {
			for j := 0; j < len(units); j++ {
				if coordinates[i].X == units[j].X && coordinates[i].Y == units[j].Y {
					unit, errUnitParams := initGame.GetXYUnits(msg.IdGame, msg.X, msg.Y)
					if errUnitParams == nil {
						resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), TypeUnit: unit.NameType, UserOwned: unit.NameUser, HP: strconv.Itoa(unit.Hp),
							UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), Damage: strconv.Itoa(unit.Damage), MoveSpeed: strconv.Itoa(unit.MoveSpeed),
							Init: strconv.Itoa(unit.Init), RangeAttack: strconv.Itoa(unit.RangeAttack), RangeView: strconv.Itoa(unit.WatchZone),
							AreaAttack: strconv.Itoa(unit.AreaAttack), TypeAttack: unit.TypeAttack}
						FieldPipe <- resp
					}
				}
			}
		}
	}
}

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	if 0 < len(usersFieldWs[ws].Units) {
		phase, err, phaseChange := game.UserReady(IdWs(ws, &usersFieldWs), msg.IdGame)
		if err != nil {
			// TODO : обработать ошибку
		} else {
			if phaseChange {
				for _, userStat := range usersFieldWs[ws].Players {
					resp = FieldResponse{Event: msg.Event, UserName: userStat.Name, Phase: phase}
					FieldPipe <- resp
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), Phase: phase}
				FieldPipe <- resp
			}
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: LoginWs(ws, &usersFieldWs), Error: "not units"}
		FieldPipe <- resp
	}
}

