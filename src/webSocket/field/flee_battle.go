package field

import (
	localGameDB "github.com/TrashPony/Veliri/src/mechanics/db/localGame"
	gameUpdate "github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
	"time"
)

// виды выхода из боя
// 1) мгноменно с потерей
// 2) с ожиданием 15 сек но без потерь

func initFlee(msg Message, client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())
	if findGame {

		gameZone := activeGame.GetGameZone(client)
		_, find := gameZone[strconv.Itoa(client.GetSquad().MatherShip.Q)][strconv.Itoa(client.GetSquad().MatherShip.R)]

		// !find - значит мы вышли за зону игры и можем ливнуть
		if !find || LastLeave(activeGame, false) || !activeGame.FindUserHostile(client) {

			// если игрок последний или у него нет врагов ему доступно софт лив
			if LastLeave(activeGame, false) || !activeGame.FindUserHostile(client) {

				SendMessage(Message{Event: "softLeave"}, client.GetID(), activeGame.Id)

			} else {
				// иначе он может ливнуть только в фазе таргетинга с потерей всех юнитов
				if activeGame.Phase == "targeting" {
					SendMessage(Message{Event: "leave"}, client.GetID(), activeGame.Id)
				} else {
					SendMessage(ErrorMessage{Event: "Error", Error: "not allow"}, client.GetID(), activeGame.Id)
				}
			}

		} else {
			SendMessage(ErrorMessage{Event: "Error", Error: "not allow"}, client.GetID(), activeGame.Id)
		}
	}
}

func LastLeave(activeGame *localGame.Game, message bool) bool {
	activePlayersCount := 0
	for _, user := range activeGame.GetPlayers() {
		if !user.Leave {
			activePlayersCount++
		}
	}

	if activePlayersCount < 2 {
		if message {
			for _, user := range activeGame.GetPlayers() {
				if !user.Leave {
					SendMessage(Message{Event: "softLeave"}, user.GetID(), activeGame.Id)
				}
			}
		}
		return true
	}
	return false
}

func fleeBattle(msg Message, client *player.Player) {

	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findGame && activeGame.Phase == "targeting" {

		gameZone := activeGame.GetGameZone(client)
		_, find := gameZone[strconv.Itoa(client.GetSquad().MatherShip.Q)][strconv.Itoa(client.GetSquad().MatherShip.R)]

		if find || LastLeave(activeGame, false) {

			if activeGame.CheckEndGame() {
				go EndGame(activeGame)
			}

			leave(client, activeGame, false)

		} else {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
		}
	}
}

func softFlee(client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())
	// если игрок последний или он достиг мира со всеми другими то он может легко ливнуть в софте
	if findGame && (LastLeave(activeGame, false) || !activeGame.FindUserHostile(client)) {
		go softFleeTimer(client, activeGame)
	}
}

func softFleeTimer(client *player.Player, activeGame *localGame.Game) {
	client.ToLeave = true

	for i := 15; i > 0; i-- {
		time.Sleep(1 * time.Second)
		SendMessage(Message{Event: "timeToLeave", Seconds: i}, client.GetID(), activeGame.Id)
	}

	leave(client, activeGame, true)
	client.ToLeave = false
}

func leave(client *player.Player, activeGame *localGame.Game, soft bool) {
	// когда игрок ливает из боя все юниты которые на карте остаются на поле
	// боя как болванки и больше не принадлежат отряду и хранятся в отдельной таблице
	// оставлять игрока в игре со статусом leave true

	if soft {
		// если софт лив то всех юнитов кладем в трюм
		for _, unitSlot := range client.GetSquad().MatherShip.Units {
			if unitSlot.Unit != nil {
				unitSlot.Unit.OnMap = false
				SendAllMessage(Message{Event: "LeaveUnit", UnitID: unitSlot.Unit.ID}, activeGame)
			}
		}
	} else {
		// иначе они остаются в игре
		for _, unitSlot := range client.GetSquad().MatherShip.Units {
			if unitSlot.Unit != nil && unitSlot.Unit.OnMap {
				unitSlot.Unit.GameID = activeGame.Id
				localGameDB.AddLeaveUnit(unitSlot.Unit, client.GetID(), activeGame.Id)
				unitSlot.Unit = nil
			}
		}
	}

	SendAllMessage(Message{Event: "LeaveUnit", UnitID: client.GetSquad().MatherShip.ID}, activeGame)

	client.GetSquad().InGame = false
	update.Squad(client.GetSquad(), true)

	client.Leave = true
	gameUpdate.Player(client)

	// заменяем игрока на фейка т.к. тот уже не игре
	activeGame.SetFakePlayer(client)

	// проверяем что игроков поврежнему больше 2х
	LastLeave(activeGame, true)

	// отправляем вышедшему игроку что он может переходить в глобальную игру
	SendMessage(Message{Event: "toGlobal"}, client.GetID(), activeGame.Id)

	// обновления тумана войны у всех игроков т.к. союзник мог ливнуть и видимость поменялась
	for _, user := range activeGame.GetPlayers() {
		SendMessage(Message{Event: "UpdateWatchMap", Update: watchZone.UpdateWatchZone(activeGame, user)}, user.GetID(), activeGame.Id)
	}
}
