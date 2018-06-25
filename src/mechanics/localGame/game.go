package localGame

import (
	"./map/gameMap"
	"../player"
	"../matherShip"
	"../unit"
)

type Game struct {
	Id          int
	Name        string
	MapID       int
	Step        int
	Phase       string
	Winner      string
	Map         *gameMap.Map
	players     []*player.Player
	unitStorage []*unit.Unit
	units       map[int]map[int]*unit.Unit
	MatherShips map[int]map[int]*matherShip.MatherShip
}

func (game *Game) SetMatherShips(matherShips map[int]map[int]*matherShip.MatherShip) {
	game.MatherShips = matherShips
}

func (game *Game) SetPlayers(players []*player.Player) {
	game.players = players
}

func (game *Game) SetMap(gameMap *gameMap.Map) {
	game.Map = gameMap
}

func (game *Game) SetUnits(unit map[int]map[int]*unit.Unit) {
	game.units = unit
}

func (game *Game) SetUnitsStorage(unit []*unit.Unit) {
	game.unitStorage = unit
}

func (game *Game) SetUnit(gameUnit *unit.Unit) {
	if game.units[gameUnit.X] != nil {
		game.units[gameUnit.X][gameUnit.Y] = gameUnit
	} else {
		game.units[gameUnit.X] = make(map[int]*unit.Unit)
		game.units[gameUnit.X][gameUnit.Y] = gameUnit
	}
}

func (game *Game) DelUnit(unit *unit.Unit) {
	delete(game.units[unit.X], unit.Y)
}

func (game *Game) GetMap() (mp *gameMap.Map) {
	return game.Map
}

func (game *Game) GetUnits() (units map[int]map[int]*unit.Unit) {
	return game.units
}

func (game *Game) GetUnitsStorage() (units []*unit.Unit) {
	return game.unitStorage
}

func (game *Game) GetUnit(x, y int) (unit *unit.Unit, find bool) {
	unit, find = game.units[x][y]
	return
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

func (game *Game) GetMatherShips() (matherShips map[int]map[int]*matherShip.MatherShip) {
	return game.MatherShips
}

func (game *Game) GetMatherShip(x, y int) (matherShip *matherShip.MatherShip, find bool) {
	matherShip, find = game.MatherShips[x][y]
	return
}

func (game *Game) GetStep() int {
	return game.Step
}

func (game *Game) GetPhase() string {
	return game.Phase
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