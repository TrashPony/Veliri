package anomaly

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
)

type Anomaly struct {
	x     int
	y     int
	MapID int `json:"map_id"`
	power int // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type  int

	/* награда */
	box      *boxInMap.Box
	resource *resource.Map
	text     *dialog.Dialog
}

func (a *Anomaly) GetX() int {
	return a.x
}

func (a *Anomaly) SetX(x int) {
	a.x = x
}

func (a *Anomaly) GetY() int {
	return a.y
}

func (a *Anomaly) SetY(y int) {
	a.y = y
}

func (a *Anomaly) GetPower() int {
	return a.power
}

func (a *Anomaly) SetPower(power int) {
	a.power = power
}

func (a *Anomaly) GetLoot() (*boxInMap.Box, *resource.Map, *dialog.Dialog) {
	return a.box, a.resource, a.text
}

func (a *Anomaly) SetBox(box *boxInMap.Box) {
	a.box = box
}

func (a *Anomaly) SetRes(res *resource.Map) {
	a.resource = res
}

func (a *Anomaly) SetDialog(dialog *dialog.Dialog) {
	a.text = dialog
}
