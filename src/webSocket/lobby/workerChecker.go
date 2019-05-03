package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/blueWorks"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
	"time"
)

func WorkerChecker() {
	for {
		workers := blueWorks.BlueWorks.GetAll()

		for _, user := range usersLobbyWs {
			// просто обновляет всем юзера таймер крафта
			baseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

			lobbyPipe <- Message{
				Event:     "WorkbenchStorage",
				UserID:    user.GetID(),
				Storage:   baseStorage,
				BlueWorks: blueWorks.BlueWorks.GetByUserAndBase(user.GetID(), user.InBaseID),
			}
		}

		for _, work := range workers {
			// проверяет работы на готовность
			if time.Now().Unix() >= work.FinishTime.Unix() {

				bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)

				if bp.ItemType == "weapon" {
					weapon, _ := gameTypes.Weapons.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, weapon, bp.ItemType, weapon.ID, bp.Count, weapon.MaxHP, weapon.Size, weapon.MaxHP, false)
				}
				if bp.ItemType == "equip" {
					equip, _ := gameTypes.Equips.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, equip, bp.ItemType, equip.ID, bp.Count, equip.MaxHP, equip.Size, equip.MaxHP, false)
				}
				if bp.ItemType == "detail" {
					detail, _ := gameTypes.Resource.GetDetailByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, detail, bp.ItemType, detail.TypeID, bp.Count, 1, detail.Size, 1, false)
				}
				if bp.ItemType == "ammo" {
					ammo, _ := gameTypes.Ammo.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, ammo, bp.ItemType, ammo.ID, bp.Count, 1, ammo.Size, 1, false)
				}
				if bp.ItemType == "body" {
					body, _ := gameTypes.Bodies.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, body, bp.ItemType, body.ID, bp.Count, body.MaxHP, body.CapacitySize, body.MaxHP, false)
				}
				if bp.ItemType == "boxes" {
					box, _ := gameTypes.Boxes.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, box, bp.ItemType, box.TypeID, bp.Count, 1, box.FoldSize, 1, false)
				}

				blueWorks.BlueWorks.Remove(work)
				wsInventory.UpdateStorage(work.UserID)
			}
		}
		time.Sleep(time.Second) // проверяем каждую секунду
	}
}
