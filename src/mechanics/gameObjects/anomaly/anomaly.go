package anomaly

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
)

type Anomaly struct {
	q     int
	r     int
	MapID int `json:"map_id"`
	power int // сила это растояние на котором можно увидить аномалию на сканере со сканером нулевого радиуса, общая дальность сила аномалии + радиус сканера
	Type  int

	/* награда */
	box      *boxInMap.Box
	resource *resource.Map
	text     *dialog.Dialog
}

func (a *Anomaly) GetQ() int {
	return a.q
}

func (a *Anomaly) SetQ(q int) {
	a.q = q
}

func (a *Anomaly) GetR() int {
	return a.r
}

func (a *Anomaly) SetR(r int) {
	a.r = r
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
