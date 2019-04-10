package field

import (
	localGameDB "github.com/TrashPony/Veliri/src/mechanics/db/localGame"
	gameUpdate "github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
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

		if find || LastLeave(activeGame, false) {
			if !LastLeave(activeGame, false) && activeGame.Phase == "targeting" {
				SendMessage(Message{Event: "leave"}, client.GetID(), activeGame.Id)
			} else {
				SendMessage(Message{Event: "softLeave"}, client.GetID(), activeGame.Id)
			}
		} else {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
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

			leave(client, activeGame, false)

		} else {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "not allow"}, client.GetID(), activeGame.Id)
		}
	}
}

func softFlee(client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())
	if findGame && LastLeave(activeGame, false) {
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

	if !soft {
		for _, unitSlot := range client.GetSquad().MatherShip.Units {
			if unitSlot.Unit != nil && unitSlot.Unit.OnMap {
				unitSlot.Unit.GameID = activeGame.Id
				localGameDB.AddLeaveUnit(unitSlot.Unit, client.GetID(), activeGame.Id)
				unitSlot.Unit = nil
			}
		}
	}

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
}
