package squad

import (
	"../../gameObjects/unit"
	"../../gameObjects/matherShip"
	"../../gameObjects/equip"
)

type Squad struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	MatherShip *matherShip.MatherShip `json:"mather_ship"`
	Units      map[int]*unit.Unit     `json:"units"`
	Equip      map[int]*equip.Equip   `json:"equip"`
}
