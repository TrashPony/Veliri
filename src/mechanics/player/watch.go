package player

import (
	"../gameObjects/coordinate"
	"strconv"
)

func (client *Player) AddCoordinate(gameCoordinate *coordinate.Coordinate) { // Todo AddWatchCoordinate
	if client.watch != nil {
		if client.watch[strconv.Itoa(gameCoordinate.Q)] != nil {
			client.watch[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
		} else {
			client.watch[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*coordinate.Coordinate)
			client.AddCoordinate(gameCoordinate)
		}
	} else {
		client.watch = make(map[string]map[string]*coordinate.Coordinate)
		client.AddCoordinate(gameCoordinate)
	}
}

func (client *Player) DelWatchCoordinate(q, r int) {
	delete(client.watch[strconv.Itoa(q)], strconv.Itoa(r))
}

func (client *Player) GetWatchCoordinates() (coordinates map[string]map[string]*coordinate.Coordinate) {
	return client.watch
}

func (client *Player) SetWatchCoordinates(coordinates map[string]map[string]*coordinate.Coordinate) {
	client.watch = coordinates
}

func (client *Player) GetWatchCoordinate(q, r int) (gameCoordinate *coordinate.Coordinate, find bool) {
	gameCoordinate, find = client.watch[strconv.Itoa(q)][strconv.Itoa(r)]
	return
}
