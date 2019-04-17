package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/movePhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/getlantern/deepcopy"
	"strconv"
	"time"
)

type Move struct {
	Event             string                                       `json:"event"`
	UserName          string                                       `json:"user_name"`
	GameID            int                                          `json:"game_id"`
	Unit              *unit.Unit                                   `json:"unit"`
	Path              []*movePhase.TruePatchNode                   `json:"path"`
	Error             string                                       `json:"error"`
	Move              bool                                         `json:"move"`
	MemoryHostileUnit map[string]unit.Unit                         `json:"memory_hostile_unit"`
	GameZone          map[string]map[string]*coordinate.Coordinate `json:"game_zone"`
}

func MoveUnit(msg Message, client *player.Player) {
	var event string

	activeGame, findGame := games.Games.Get(client.GetGameID())
	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)

	if !findUnit {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)

		if !gameUnit.Body.MotherShip && msg.ToQ == client.GetSquad().MatherShip.Q && msg.ToR == client.GetSquad().MatherShip.R {
			event = "ToMC"
		}

	} else {
		event = "SelectStorageUnit"
	}

	if findUnit && findGame {
		if !client.GetReady() && gameUnit.ActionPoints > 0 {

			moveCoordinate := movePhase.GetMoveCoordinate(gameUnit, client, activeGame, event)
			_, find := moveCoordinate[strconv.Itoa(msg.ToQ)][strconv.Itoa(msg.ToR)]

			if find {

				path := movePhase.InitMove(gameUnit, msg.ToQ, msg.ToR, client, activeGame, event)
				if event != "ToMC" {
					client.DelUnitStorage(gameUnit.ID)
				} else {
					SendMessage(LoadGame{Event: "UpdateUnitStorage", UnitStorage: client.GetUnitsStorage()}, client.GetID(), activeGame.Id)
				}

				SendMessage(
					Move{
						Event:    msg.Event,
						Unit:     gameUnit,
						UserName: client.GetLogin(),
						Path:     path,
					},
					client.GetID(),
					activeGame.Id,
				)

				updateWatchHostileUser(client, activeGame, gameUnit, path, event)
				QueueSender(activeGame)
			} else {
				SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
			}
		} else {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "unit already move"}, client.GetID(), activeGame.Id)
		}
	} else {
		SendMessage(ErrorMessage{Event: msg.Event, Error: "not found unit"}, client.GetID(), activeGame.Id)
	}
}

// воркер который следит за тем что бы ходы менялись каждый n секунд
func timerMoveUnits(game *localGame.Game) {
	for {

		moveUnit := game.GetMoveUnit()
		forceMove := true

		if game.Phase == "move" && moveUnit != nil && !game.End {

			for i := 60; i > 0; i-- {

				if !moveUnit.Move {
					// если юнит больше не мове значит им походили
					forceMove = false
					break
				}

				time.Sleep(1 * time.Second)
				SendAllMessage(Message{Event: "timeToChangePhase", Seconds: i}, game)
			}

			if forceMove && !game.End {
				movePhase.SkipMove(moveUnit, game)
				client := game.GetUserByName(moveUnit.Owner)
				SendMessage(Move{Event: "UpdateUnit", Unit: moveUnit, UserName: client.GetLogin()}, client.GetID(), game.Id)

				QueueSender(game)
			}

		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func QueueSender(game *localGame.Game) {

	allReady := true

	for _, user := range game.GetPlayers() {
		for _, q := range user.GetUnits() {

			for _, gameUnit := range q {
				SendMessage(
					Move{Event: "QueueMove",
						UserName: user.GetLogin(),
						GameID:   game.Id,
						Unit:     gameUnit,
					},
					user.GetID(),
					game.Id,
				)
			}

			if !user.Ready {
				allReady = false
			}
		}

		for _, gameUnit := range user.GetUnitsStorage() {
			SendMessage(
				Move{
					Event:    "QueueMove",
					UserName: user.GetLogin(),
					GameID:   game.Id, Unit: gameUnit,
				},
				user.GetID(),
				game.Id,
			)
		}
		SendMessage(
			Move{Event: "UpdateMemoryUnit",
				UserName:          user.GetLogin(),
				GameID:            game.Id,
				MemoryHostileUnit: user.GetMemoryHostileUnits(),
			},
			user.GetID(),
			game.Id,
		)
	}

	if allReady {
		CheckAllReady(game)
	}
}

func SkipMoveUnit(msg Message, client *player.Player) {

	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)
	if gameUnit == nil {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
	}

	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findUnit && findGame {
		movePhase.SkipMove(gameUnit, activeGame)
		SendMessage(Move{Event: "UpdateUnit", Unit: gameUnit, UserName: client.GetLogin()}, client.GetID(), activeGame.Id)
		QueueSender(activeGame)
	} else {
		SendMessage(ErrorMessage{Event: "MoveUnit", Error: "not found unit or game or player"}, client.GetID(), activeGame.Id)
	}
}

func updateWatchHostileUser(client *player.Player, activeGame *localGame.Game, gameUnit *unit.Unit, pathNodes []*movePhase.TruePatchNode, event string) {
	// глубокое копирование pathNodes для каждого клиента
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {

			// если ходим мс, то у всех игроков обновляется игровая зона независимо от того видят они его или нет
			if gameUnit.Body.MotherShip {
				SendMessage(
					Move{Event: "UpdateGameZone", UserName: user.GetLogin(), GameID: activeGame.Id, GameZone: activeGame.GetGameZone(user)},
					user.GetID(),
					activeGame.Id,
				)
			}

			// глубокое копирование pathNodes для каждого клиента что бы не менять исходные данные
			var path []*movePhase.TruePatchNode
			deepcopy.Copy(&path, &pathNodes)

			// пытаемся взять юнита по начальной координате
			_, okGetUnit := user.GetHostileUnitByID(gameUnit.ID)

			// если юзер видит юнита то удаляем его со строго места
			if okGetUnit {
				user.DelHostileUnit(gameUnit.ID)
			}

			// пытаемся взять юнита по конечной координате
			_, okGetEndQR := user.GetWatchCoordinate(gameUnit.Q, gameUnit.R)

			// если конечная точка пути видима и он не погрузился в мс то добавляем юнита
			if okGetEndQR && gameUnit.OnMap {
				user.AddHostileUnit(gameUnit)
			}

			send := false
			okSecondNode := false
			okEarlyNode := false
			// тут происходит формирование пути для пользователя который может видеть не весь путь юнита
			for i, pathNode := range path {
				pathNode.WatchNode = nil

				firstNode, okFirstNode := user.GetWatchCoordinate(pathNode.PathNode.Q, pathNode.PathNode.R)

				if len(path) > i+1 {
					_, okSecondNode = user.GetWatchCoordinate(path[i+1].PathNode.Q, path[i+1].PathNode.R)
				}
				if 0 < i {
					_, okEarlyNode = user.GetWatchCoordinate(path[i-1].PathNode.Q, path[i-1].PathNode.R)
				}

				if event == "SelectStorageUnit" && i == 0 && okFirstNode {
					// если машина выходит из мазер шипа говорим что как бе он вышел из тумана что бы она появилась на карте)
					firstNode.Type = "outFog"
				}

				// если юзер не видит координату то скрваем ее
				if !okFirstNode {
					var fakeNode coordinate.Coordinate

					if okSecondNode {
						fakeNode.Type = "outFog"
						fakeNode.Q = pathNode.PathNode.Q
						fakeNode.R = pathNode.PathNode.R
					} else {
						if (okGetUnit && i == 0) || okEarlyNode {
							fakeNode.Type = "inToFog"
							fakeNode.Q = pathNode.PathNode.Q
							fakeNode.R = pathNode.PathNode.R
						} else {
							fakeNode.Type = "hide"
							fakeNode.Q = 0
							fakeNode.R = 0
						}
					}

					pathNode.PathNode = &fakeNode
				} else {
					send = true
				}
			}

			// отправляем только тем кто видит хотя бы 1 клетку пути
			if send || okGetUnit {

				if !okGetUnit { // добавляем нового вражеского юнита тем юзерам кто его не видел до этого
					user.AddNewMemoryHostileUnit(*gameUnit)
					update.Player(user)
				}

				SendMessage(
					Move{
						Event:    "HostileUnitMove",
						Unit:     gameUnit,
						UserName: user.GetLogin(),
						GameID:   activeGame.Id,
						Path:     path,
					},
					user.GetID(),
					activeGame.Id,
				)
			}
		}
	}
}
