package lobby

import (
	"../../mechanics/factories/blueWorks"
	"time"
	"../../mechanics/factories/gameTypes"
	"../../mechanics/factories/storages"
	"../storage"
)

func WorkerChecker() {
	for {
		workers := blueWorks.BlueWorks.GetAll()

		for _, work := range workers {
			if time.Now().Unix() >= work.FinishTime.Unix() {

				bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)

				if bp.ItemType == "weapon" {
					weapon, _ := gameTypes.Weapons.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, weapon, bp.ItemType, weapon.ID, bp.Count, weapon.MaxHP, weapon.Size, weapon.MaxHP)
				}
				if bp.ItemType == "equip" {
					equip, _ := gameTypes.Equips.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, equip, bp.ItemType, equip.ID, bp.Count, equip.MaxHP, equip.Size, equip.MaxHP)
				}
				if bp.ItemType == "detail" {
					detail, _ := gameTypes.Resource.GetDetailByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, detail, bp.ItemType, detail.TypeID, bp.Count, 1, detail.Size, 1)
				}
				if bp.ItemType == "ammo" {
					ammo, _ := gameTypes.Ammo.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, ammo, bp.ItemType, ammo.ID, bp.Count, 1, ammo.Size, 1)
				}
				if bp.ItemType == "body" {
					body, _ := gameTypes.Bodies.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, body, bp.ItemType, body.ID, bp.Count, body.MaxHP, body.CapacitySize, body.MaxHP)
				}
				if bp.ItemType == "boxes" {
					box, _ := gameTypes.Boxes.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, box, bp.ItemType, box.TypeID, bp.Count, 1, box.FoldSize, 1)
				}

				blueWorks.BlueWorks.Remove(work)
				storage.Updater(work.UserID)
			}
		}
		time.Sleep(time.Second) // проверяем каждую секунду
	}
}
