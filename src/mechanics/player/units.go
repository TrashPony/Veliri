package player

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/getlantern/deepcopy"
	"strconv"
)

func (client *Player) AddUnit(gameUnit *unit.Unit) {
	if client.units != nil {
		if client.units[strconv.Itoa(gameUnit.Q)] != nil {
			client.units[strconv.Itoa(gameUnit.Q)][strconv.Itoa(gameUnit.R)] = gameUnit
		} else {
			client.units[strconv.Itoa(gameUnit.Q)] = make(map[string]*unit.Unit)
			client.AddUnit(gameUnit)
		}
	} else {
		client.units = make(map[string]map[string]*unit.Unit)
		client.AddUnit(gameUnit)
	}
}

func (client *Player) AddHostileUnit(hostile *unit.Unit) {
	if client.hostileUnits != nil {
		if client.hostileUnits[strconv.Itoa(hostile.Q)] != nil {
			client.hostileUnits[strconv.Itoa(hostile.Q)][strconv.Itoa(hostile.R)] = hostile
		} else {
			client.hostileUnits[strconv.Itoa(hostile.Q)] = make(map[string]*unit.Unit)
			client.AddHostileUnit(hostile)
		}
	} else {
		client.hostileUnits = make(map[string]map[string]*unit.Unit)
		client.AddHostileUnit(hostile)
	}
}

//TODO метод костыль, потому что я долбаеб
func (client *Player) GetUnitsINTKEY() (units map[int]map[int]*unit.Unit) {

	units = make(map[int]map[int]*unit.Unit)
	for _, qLine := range client.GetUnits() {
		for _, clientUnit := range qLine {
			if units[clientUnit.Q] != nil {
				units[clientUnit.Q][clientUnit.R] = clientUnit
			} else {
				units[clientUnit.Q] = make(map[int]*unit.Unit)
				units[clientUnit.Q][clientUnit.R] = clientUnit
			}
		}
	}
	return units
}

func (client *Player) GetUnits() (units map[string]map[string]*unit.Unit) {
	return client.units
}

func (client *Player) SetUnits(units map[string]map[string]*unit.Unit) {
	client.units = units
}

func (client *Player) GetUnit(q, r int) (gameUnit *unit.Unit, find bool) {
	gameUnit, find = client.units[strconv.Itoa(q)][strconv.Itoa(r)]
	return
}

func (client *Player) DelUnit(gameUnit *unit.Unit, delSquad bool) {
	delete(client.units[strconv.Itoa(gameUnit.Q)], strconv.Itoa(gameUnit.R))

	if delSquad {
		for _, slot := range client.squad.MatherShip.Units {
			if slot.Unit != nil {
				if slot.Unit.Q == gameUnit.Q && slot.Unit.R == gameUnit.R && slot.Unit.ID == gameUnit.ID {
					slot.Unit = nil
				}
			}
		}
	}
}

func (client *Player) GetHostileUnits() (units map[string]map[string]*unit.Unit) {

	hostilesUnits := make(map[string]map[string]*unit.Unit)

	// удаляем всю информацию которую не должен видит другой юзер в игре
	for q, xLine := range client.hostileUnits {
		for r, hostile := range xLine {

			shortInfoUnit := shortHostileUnit(hostile)

			if hostilesUnits[q] != nil {
				hostilesUnits[q][r] = shortInfoUnit
			} else {
				hostilesUnits[q] = make(map[string]*unit.Unit)
				hostilesUnits[q][r] = shortInfoUnit
			}
		}
	}

	return client.hostileUnits
}

func (client *Player) GetHostileUnit(q, r int) (*unit.Unit, bool) {
	gameUnit, find := client.hostileUnits[strconv.Itoa(q)][strconv.Itoa(r)]

	shortInfoUnit := shortHostileUnit(gameUnit)

	return shortInfoUnit, find
}

func (client *Player) GetHostileUnitByID(id int) (gameUnit *unit.Unit, find bool) {
	for _, xLine := range client.GetHostileUnits() {
		for _, hostile := range xLine {
			if hostile.ID == id {

				shortInfoUnit := shortHostileUnit(hostile)

				return shortInfoUnit, true
			}
		}
	}
	return
}

func shortHostileUnit(gameUnit *unit.Unit) *unit.Unit {
	if gameUnit == nil {
		return nil
	}

	var shortInfoUnit unit.Unit
	deepcopy.Copy(&shortInfoUnit, &gameUnit)

	shortInfoUnit.Target = nil
	shortInfoUnit.Defend = false
	shortInfoUnit.Power = 0

	shortInfoUnit.Speed = shortInfoUnit.Body.Speed
	shortInfoUnit.MaxHP = shortInfoUnit.Body.MaxHP
	shortInfoUnit.Armor = shortInfoUnit.Body.Armor
	shortInfoUnit.EvasionCritical = shortInfoUnit.Body.EvasionCritical
	shortInfoUnit.VulToKinetics = shortInfoUnit.Body.VulToKinetics
	shortInfoUnit.VulToThermo = shortInfoUnit.Body.VulToThermo
	shortInfoUnit.VulToEM = shortInfoUnit.Body.VulToEM
	shortInfoUnit.VulToExplosion = shortInfoUnit.Body.VulToExplosion
	shortInfoUnit.RangeView = shortInfoUnit.Body.RangeView
	shortInfoUnit.Accuracy = shortInfoUnit.Body.Accuracy
	shortInfoUnit.MaxPower = shortInfoUnit.Body.MaxPower
	shortInfoUnit.RecoveryPower = shortInfoUnit.Body.RecoveryPower
	shortInfoUnit.RecoveryHP = shortInfoUnit.Body.RecoveryHP
	shortInfoUnit.WallHack = shortInfoUnit.Body.WallHack

	shortInfoUnit.Units = nil

	shortInfoUnit.Reload = nil

	shortInfoUnit.Body.EquippingI = nil
	shortInfoUnit.Body.EquippingII = nil
	shortInfoUnit.Body.EquippingIII = nil
	shortInfoUnit.Body.EquippingIV = nil
	shortInfoUnit.Body.EquippingV = nil
	shortInfoUnit.Body.ThoriumSlots = nil

	if weaponSlot := shortInfoUnit.GetWeaponSlot(); weaponSlot != nil && weaponSlot.Weapon != nil {
		weaponSlot.HP = 0
		weaponSlot.AmmoQuantity = 0
		weaponSlot.Ammo = nil
	}

	return &shortInfoUnit
}

func (client *Player) SetHostileUnits(units map[string]map[string]*unit.Unit) {
	client.hostileUnits = units
}

func (client *Player) DelHostileUnit(id int) {
	for x, xLine := range client.GetHostileUnits() {
		for y, hostile := range xLine {
			if hostile.ID == id {
				delete(client.hostileUnits[x], y)
			}
		}
	}
}

func (client *Player) AddUnitStorage(gameUnit *unit.Unit) {
	if client.unitStorage == nil {
		client.unitStorage = make([]*unit.Unit, 0)
	}
	client.unitStorage = append(client.unitStorage, gameUnit)
}

func (client *Player) GetUnitsStorage() (gameUnit []*unit.Unit) {
	return client.unitStorage
}

func (client *Player) RemoveUnitsStorage() {
	client.unitStorage = nil
}

func (client *Player) GetUnitStorage(id int) (storageUnit *unit.Unit, find bool) {
	for _, storageUnit := range client.GetUnitsStorage() {
		if id == storageUnit.ID {
			return storageUnit, true
		}
	}

	return
}

func (client *Player) DelUnitStorage(id int) (find bool) {
	for _, storageUnit := range client.GetUnitsStorage() {
		if id == storageUnit.ID {
			client.unitStorage = remove(client.GetUnitsStorage(), storageUnit)
			return true
		}
	}

	return
}

func (client *Player) SetMoveParamsMemoryUnit(idUnit int, move bool, actionPoint int) {
	memoryUnit, ok := client.memoryHostileUnits[strconv.Itoa(idUnit)]
	if ok {
		memoryUnit.Move = move
		memoryUnit.ActionPoints = actionPoint
		client.memoryHostileUnits[strconv.Itoa(idUnit)] = memoryUnit
	}
}

func (client *Player) AddNewMemoryHostileUnit(newUnit unit.Unit) {
	if client.memoryHostileUnits == nil {
		client.memoryHostileUnits = make(map[string]unit.Unit)
	}

	client.memoryHostileUnits[strconv.Itoa(newUnit.ID)] = newUnit
}

func (client *Player) DelMemoryHostileUnits(id int) {
	delete(client.memoryHostileUnits, strconv.Itoa(id))
}

func (client *Player) GetMemoryHostileUnits() map[string]unit.Unit {
	return client.memoryHostileUnits
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

func (client *Player) GetMoveUnit() *unit.Unit {
	for _, q := range client.GetUnits() {
		for _, gameUnit := range q {
			if gameUnit.Move {
				return gameUnit
			}
		}
	}

	for _, gameUnit := range client.GetUnitsStorage() {
		if gameUnit.Move {
			return gameUnit
		}
	}

	return nil
}
