package attackPhase

import (
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"math/rand"
	"time"
)

type QueueAttack struct {
	ActionUnit     *unit.Unit             `json:"action_unit"`
	ItemInitiative int                    `json:"item_initiative"`
	WeaponSlot     *detail.BodyWeaponSlot `json:"weapon_slot"`
	EquipSlot      *detail.BodyEquipSlot  `json:"equip_slot"`
	sort           bool
}

func createQueueAttack(Units map[int]map[int]*unit.Unit) (sortUnits []*QueueAttack) {

	preloadQueue := make([]*QueueAttack, 0)

	for _, xLine := range Units { // закидываем в очередь все оружия всех юнитов
		for _, gameUnit := range xLine {
			for _, weaponSlot := range gameUnit.Body.Weapons {
				if weaponSlot.Weapon != nil && weaponSlot.AmmoQuantity > 0 && gameUnit.Target != nil {
					preloadQueue = append(preloadQueue, &QueueAttack{ActionUnit: gameUnit, ItemInitiative: weaponSlot.Weapon.Initiative, WeaponSlot: weaponSlot})
				}
			}
		}
	}

	var equipCreate = func(equip map[int]*detail.BodyEquipSlot, gameUnit *unit.Unit) {
		for _, equipSlot := range equip {
			if equipSlot.Equip != nil && !equipSlot.Used && equipSlot.Target != nil {
				preloadQueue = append(preloadQueue, &QueueAttack{ActionUnit: gameUnit, ItemInitiative: equipSlot.Equip.Initiative, EquipSlot: equipSlot})
			}
		}
	}

	for _, xLine := range Units { // закидываем в очередь все снаряжения всех юнитов
		for _, gameUnit := range xLine {
			equipCreate(gameUnit.Body.EquippingI, gameUnit)
			equipCreate(gameUnit.Body.EquippingII, gameUnit)
			equipCreate(gameUnit.Body.EquippingIII, gameUnit)
			equipCreate(gameUnit.Body.EquippingIV, gameUnit)
			equipCreate(gameUnit.Body.EquippingV, gameUnit)
		}
	}

	sortUnits = sortQueueAttack(&preloadQueue)

	return
}

func sortQueueAttack(preloadQueue *[]*QueueAttack) (sortUnits []*QueueAttack) {

	var checkSortItems = func(preloadQueue *[]*QueueAttack) bool {
		// проверяем есть ли еще не отсортированые итемы
		for _, item := range *preloadQueue {
			if !item.sort {
				return true
			}
		}
		return false
	}

	for checkSortItems(preloadQueue) {
		sortUnits = append(sortUnits, getMaxInitiativeItem(preloadQueue))
	}

	return sortUnits
}

func getMaxInitiativeItem(preloadQueue *[]*QueueAttack) *QueueAttack {

	maxInitiative := 0
	var maxItem *QueueAttack

	for _, item := range *preloadQueue {
		if item.ItemInitiative > maxInitiative && !item.sort {
			maxInitiative = item.ItemInitiative
			maxItem = item
		}
	}

	equalInitiative := make([]*QueueAttack, 0)

	for _, item := range *preloadQueue {
		if item.ItemInitiative == maxItem.ItemInitiative && !item.sort {
			equalInitiative = append(equalInitiative, item)
		}
	}

	if len(equalInitiative) > 1 {
		maxItem = randomInitiativeSort(equalInitiative)
	}

	maxItem.sort = true
	return maxItem
}

func randomInitiativeSort(items []*QueueAttack) *QueueAttack {
	//Генератор случайных чисел обычно нужно рандомизировать перед использованием, иначе, он, действительно,
	// будет выдавать одну и ту же последовательность.
	rand.Seed(time.Now().UnixNano())
	numberUnit := rand.Intn(len(items))

	return items[numberUnit]
}
