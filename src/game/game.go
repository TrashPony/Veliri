package game

import (
	"log"
)

type Game struct {
	gameMap    *Map
	stat       *InfoGame
	players    []*UserStat
	units      map[int]map[int]*Unit
	structure  map[int]map[int]*Structure
}

func (game *Game) SetStructure(structure map[int]map[int]*Structure) {
	game.structure = structure
}

func (game *Game) SetPlayers(players []*UserStat) {
	game.players = players
}

func (game *Game) SetInfoMap(gameMap *Map) {
	game.gameMap = gameMap
}

func (game *Game) SetStat(stat *InfoGame) {
	game.stat = stat
}

func (game *Game) SetUnits(unit map[int]map[int]*Unit) {
	game.units = unit
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
	return game.gameMap
}

func (game *Game) GetUnits() (units map[int]map[int]*Unit) {
	return game.units
}

func (game *Game) GetPlayers() (Players []*UserStat) {
	return game.players
}

func (game *Game) GetStructure() (Structure map[int]map[int]*Structure) {
	return game.structure
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
	userStat := GetUserStat(idGAme)
	Map := GetMap(gameStat.IdMap)
	units := GetAllUnits(idGAme)
	structure := GetAllStrcuture(idGAme)

	newGame.SetPlayers(userStat)     // добавляем параметры всех игроков к обьекту игры
	newGame.SetStat(&gameStat)       // добавляем информацию об игре в обьект игры
	newGame.SetInfoMap(&Map)         // добавляем информацию об карте
	newGame.SetUnits(units)          // добавляем имеющихся юнитов
	//newGame.setMap(Map.OneLayerMap)// добавляем 1 слой карты отвечающий за фон текстур, препятсвия и расположение респаунов
	newGame.SetStructure(structure)  // добавляем в игру все структуры на карте

	return
}

