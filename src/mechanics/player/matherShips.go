package player

import (
	"strconv"
	"../gameObjects/matherShip"
)

func (client *Player) AddMatherShips(ship *matherShip.MatherShip) {
	client.matherShip = ship
}

func (client *Player) AddHostileMatherShip(ship *matherShip.MatherShip) {
	if client.hostileMatherShips != nil {
		if client.hostileMatherShips[strconv.Itoa(ship.X)] != nil {
			client.hostileMatherShips[strconv.Itoa(ship.X)][strconv.Itoa(ship.Y)] = ship
		} else {
			client.hostileMatherShips[strconv.Itoa(ship.X)] = make(map[string]*matherShip.MatherShip)
			client.AddHostileMatherShip(ship)
		}
	} else {
		client.hostileMatherShips = make(map[string]map[string]*matherShip.MatherShip)
		client.AddHostileMatherShip(ship)
	}
}

func (client *Player) GetMatherShip() (*matherShip.MatherShip) {
	return client.matherShip
}

func (client *Player) SetMatherShip(ship *matherShip.MatherShip) () {
	client.matherShip = ship
}

func (client *Player) GetHostileMatherShips() (ship map[string]map[string]*matherShip.MatherShip) {
	return client.hostileMatherShips
}

func (client *Player) SetHostileMatherShips(ship map[string]map[string]*matherShip.MatherShip) () {
	client.hostileMatherShips = ship
}

func (client *Player) GetHostileMatherShip(x, y int) (ship *matherShip.MatherShip, find bool) {
	ship, find = client.hostileMatherShips[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}