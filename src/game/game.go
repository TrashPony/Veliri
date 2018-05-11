package game

import (
	"log"
)

type Game struct {
	Map         *Map
	stat        *InfoGame
	players     []*Player
	notGameUnit []*Unit
	units       map[int]map[int]*Unit
	MatherShips map[int]map[int]*MatherShip
}

func (game *Game) SetMatherShips(matherShips map[int]map[int]*MatherShip) {
	game.MatherShips = matherShips
}

func (game *Game) SetPlayers(players []*Player) {
	game.players = players
}

func (game *Game) SetMap(gameMap *Map) {
	game.Map = gameMap
}

func (game *Game) SetStat(stat *InfoGame) {
	game.stat = stat
}

func (game *Game) SetUnits(unit map[int]map[int]*Unit) {
	game.units = unit
}

func (game *Game) SetNotGameUnits(unit []*Unit) {
	game.notGameUnit = unit
}

func (game *Game) SetUnit(unit *Unit) {
	if game.units[unit.X] != nil {
		game.units[unit.X][unit.Y] = unit
	} else {
		game.units[unit.X] = make(map[int]*Unit)
		game.units[unit.X][unit.Y] = unit
	}
}

func (game *Game) DelUnit(unit *Unit) {
	delete(game.units[unit.X], unit.Y)
}

func (game *Game) GetMap() (mp *Map) {
	return game.Map
}

func (game *Game) GetUnits() (units map[int]map[int]*Unit) {
	return game.units
}

func (game *Game) GetNotGameUnits() (units []*Unit) {
	return game.notGameUnit
}

func (game *Game) GetUnit(x, y int) (unit *Unit, find bool) {
	unit, find = game.units[x][y]
	return
}

func (game *Game) GetPlayers() (Players []*Player) {
	return game.players
}

func (game *Game) GetMatherShips() (matherShips map[int]map[int]*MatherShip) {
	return game.MatherShips
}

func (game *Game) GetMatherShip(x, y int) (matherShip *MatherShip, find bool) {
	matherShip, find = game.MatherShips[x][y]
	return
}

func (game *Game) GetStat() (stat *InfoGame) {
	return game.stat
}

type InfoGame struct {
	Id     int
	Name   string
	IdMap  int
	Step   int
	Phase  string
	Winner string
}

func GetInfoGame(idGame int) InfoGame {

	rows, err := db.Query("Select * FROM action_games WHERE id=$1", idGame)
	if err != nil {
		log.Fatal("Error GetInfo Game")
	}
	defer rows.Close()

	var game InfoGame

	for rows.Next() {
		err := rows.Scan(&game.Id, &game.Name, &game.IdMap, &game.Step, &game.Phase, &game.Winner)
		if err != nil {
			log.Fatal("Error GetInfo Game")
		}
	}

	return game
}

func InitGame(idGAme int) (newGame *Game) {
	newGame = &Game{}

	gameStat := GetInfoGame(idGAme)
	Map := GetMap(gameStat.IdMap)
	units, notGameUnits := GetAllUnits(idGAme)
	matherShips := GetMatherShips(idGAme)

	newGame.SetStat(&gameStat)          // добавляем информацию об игре в обьект игры
	newGame.SetMap(&Map)                // добавляем информацию об карте
	newGame.SetUnits(units)             // добавляем имеющихся юнитов
	newGame.SetNotGameUnits(notGameUnits)
	newGame.SetMatherShips(matherShips) // добавляем в игру все структуры на карте

	userStat := GetPlayer(newGame)
	newGame.SetPlayers(userStat) // добавляем параметры всех игроков к обьекту игры

	return
}
