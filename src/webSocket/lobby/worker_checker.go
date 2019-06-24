package lobby

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/blueWorks"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
	"github.com/satori/go.uuid"
	"time"
)

func WorkerChecker() {
	for {
		workers := blueWorks.BlueWorks.GetAll()

		//todo возможна проблема конкуретного дотупа
		//todo и вообще это не оптимально в планет трафика
		for _, user := range usersLobbyWs {
			// просто обновляет всем юзера таймер крафта
			userBase, _ := bases.Bases.Get(user.InBaseID)
			baseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

			lobbyPipe <- Message{
				Event:                    "WorkbenchStorage",
				UserID:                   user.GetID(),
				Storage:                  baseStorage,
				BlueWorks:                blueWorks.BlueWorks.GetByUserAndBase(user.GetID(), user.InBaseID),
				UserWorkSkillTimePercent: user.CurrentSkills["production_time"].Level * 5,
				Base:                     userBase,
			}
		}

		for _, work := range workers {
			// проверяет работы на готовность
			if time.Now().Unix() >= work.FinishTime.Unix() {

				bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)
				var item interface{}

				if bp.ItemType == "weapon" {
					weapon, _ := gameTypes.Weapons.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, weapon, bp.ItemType, weapon.ID, bp.Count, weapon.MaxHP, weapon.Size, weapon.MaxHP, false)
					item = weapon
				}
				if bp.ItemType == "equip" {
					equip, _ := gameTypes.Equips.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, equip, bp.ItemType, equip.ID, bp.Count, equip.MaxHP, equip.Size, equip.MaxHP, false)
					item = equip
				}
				if bp.ItemType == "detail" {
					detail, _ := gameTypes.Resource.GetDetailByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, detail, bp.ItemType, detail.TypeID, bp.Count, 1, detail.Size, 1, false)
					item = detail
				}
				if bp.ItemType == "ammo" {
					ammo, _ := gameTypes.Ammo.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, ammo, bp.ItemType, ammo.ID, bp.Count, 1, ammo.Size, 1, false)
					item = ammo
				}
				if bp.ItemType == "body" {
					body, _ := gameTypes.Bodies.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, body, bp.ItemType, body.ID, bp.Count, body.MaxHP, body.CapacitySize, body.MaxHP, false)
					item = body
				}
				if bp.ItemType == "boxes" {
					box, _ := gameTypes.Boxes.GetByID(bp.ItemId)
					storages.Storages.AddItem(work.UserID, work.BaseID, box, bp.ItemType, box.TypeID, bp.Count, 1, box.FoldSize, 1, false)
					item = box
				}

				user, _ := players.Users.Get(work.UserID)
				notifyUUID := uuid.Must(uuid.NewV4(), nil).String()
				base, _ := bases.Bases.Get(work.BaseID)
				mp, _ := maps.Maps.GetByID(base.MapID)

				user.NotifyQueue[notifyUUID] = &player.Notify{
					Name:  "craft",
					UUID:  notifyUUID,
					Event: "complete",
					Item:  &inventory.Slot{Item: item, Quantity: bp.Count, Type: bp.ItemType},
					Base:  base,
					Map:   mp.GetShortInfoMap(),
				}

				dbPlayer.UpdateUser(user)

				blueWorks.BlueWorks.Remove(work)
				wsInventory.UpdateStorage(work.UserID)
			}
		}
		time.Sleep(time.Second * 1) // проверяем каждую секунду
	}
}
