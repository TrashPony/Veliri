package player

import (
	"strconv"
	"../coordinate"
)

func (client *Player) AddCoordinate(gameCoordinate *coordinate.Coordinate) { // Todo AddWatchCoordinate
	if client.watch != nil {
		if client.watch[strconv.Itoa(gameCoordinate.X)] != nil {
			client.watch[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
		} else {
			client.watch[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
			client.AddCoordinate(gameCoordinate)
		}
	} else {
		client.watch = make(map[string]map[string]*coordinate.Coordinate)
		client.AddCoordinate(gameCoordinate)
	}
}

func (client *Player) DelWatchCoordinate(x, y int) {
	delete(client.watch[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) GetWatchCoordinates() (coordinates map[string]map[string]*coordinate.Coordinate) {
	return client.watch
}

func (client *Player) SetWatchCoordinates(coordinates map[string]map[string]*coordinate.Coordinate) () {
	client.watch = coordinates
}

func (client *Player) GetWatchCoordinate(x, y int) (gameCoordinate *coordinate.Coordinate, find bool) {
	gameCoordinate, find = client.watch[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}
