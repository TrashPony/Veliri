package localGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/getlantern/deepcopy"
)

type Game struct {
	Id          int
	Name        string
	MapID       int
	Step        int
	Phase       string
	Winner      string
	Map         *_map.Map
	players     []*player.Player
	unitStorage []*unit.Unit
	units       map[int]map[int]*unit.Unit
}

func (game *Game) SetPlayers(players []*player.Player) {
	game.players = players
}

func (game *Game) SetFakePlayer(realPlayer *player.Player) {
	// функция для замены юзеров болванками если игрок ливнул, погибнул в игре
	for i, user := range game.players {
		if user.GetID() == realPlayer.GetID() {
			var fakePlayer player.Player
			err := deepcopy.Copy(&fakePlayer, &user)
			if err != nil {
				println(err.Error())
			}
			fakePlayer.Leave = true
			game.players[i] = &fakePlayer
		}
	}
}

func (game *Game) SetMap(gameMap *_map.Map) {
	game.Map = gameMap
}

func (game *Game) SetUnits(unit map[int]map[int]*unit.Unit) {
	game.units = unit
}

func (game *Game) SetUnitsStorage(unit []*unit.Unit) {
	game.unitStorage = unit
}

func (game *Game) SetUnit(gameUnit *unit.Unit) {
	if game.units[gameUnit.Q] != nil {
		game.units[gameUnit.Q][gameUnit.R] = gameUnit
	} else {
		game.units[gameUnit.Q] = make(map[int]*unit.Unit)
		game.units[gameUnit.Q][gameUnit.R] = gameUnit
	}
}

func (game *Game) DelUnit(unit *unit.Unit) {
	delete(game.units[unit.Q], unit.R)
}

func (game *Game) GetMap() (mp *_map.Map) {
	return game.Map
}

func (game *Game) GetUnits() (units map[int]map[int]*unit.Unit) {
	return game.units
}

func (game *Game) GetUnitsStorage() (units []*unit.Unit) {
	return game.unitStorage
}

func (game *Game) GetUnit(q, r int) (unit *unit.Unit, find bool) {
	unit, find = game.units[q][r]
	return
}

func (game *Game) GetUnitByID(id int) (unit *unit.Unit, find bool) {
	for _, q := range game.units {
		for _, gameUnit := range q {
			if unit.ID == id {
				return gameUnit, true
			}
		}
	}
	return nil, false
}

func (game *Game) GetPlayers() (Players []*player.Player) {
	return game.players
}

func (game *Game) GetPlayer(id int, login string) (Players *player.Player) {
	for i, gamePlayer := range game.players {
		if (login == gamePlayer.GetLogin()) && (id == gamePlayer.GetID()) {
			return game.players[i]
		}
	}

	return nil
}

func (game *Game) GetGameZone(player *player.Player) map[string]map[string]*coordinate.Coordinate {
	// зона игры от лица пользователя.
	zone := make(map[string]map[string]*coordinate.Coordinate)

	for _, gamePlayer := range game.players {

		if gamePlayer.Leave {
			continue
		}

		if gamePlayer.GetID() != player.GetID() && gamePlayer != nil && gamePlayer.GetSquad() != nil &&
			gamePlayer.GetSquad().MatherShip != nil && gamePlayer.GetSquad().MatherShip.Body != nil {

			center, find := game.Map.GetCoordinate(gamePlayer.GetSquad().MatherShip.Q, gamePlayer.GetSquad().MatherShip.R)
			if find {
				watchZone := coordinate.GetCoordinatesRadius(center, gamePlayer.GetSquad().MatherShip.RangeView*2)
				for _, coor := range watchZone {
					Phases.AddCoordinate(zone, coor)
				}
			}
		}
	}

	return zone
}

func (game *Game) GetStep() int {
	return game.Step
}

func (game *Game) GetPhase() string {
	return game.Phase
}

func (game *Game) AddUnitToStorage(gameUnit *unit.Unit) {
	game.unitStorage = append(game.unitStorage, gameUnit)
}

func (game *Game) DelUnitStorage(id int) (find bool) {
	for _, storageUnit := range game.GetUnitsStorage() {
		if id == storageUnit.ID {
			game.unitStorage = remove(game.GetUnitsStorage(), storageUnit)
			return true
		}
	}

	return
}

func remove(units []*unit.Unit, item *unit.Unit) []*unit.Unit {
	for i, v := range units {
		if v == item {
			copy(units[i:], units[i+1:])
			units[len(units)-1] = nil // обнуляем "хвост"
			units = units[:len(units)-1]
		}
	}
	return units
}
