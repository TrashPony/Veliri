package squadInventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/get"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func GetInventory(client *player.Player) {
	squads, err := get.UserSquads(client.GetID())
	if err != nil {
		println("error, get Squads")
		log.Fatal(err)
	}

	if len(squads) > 0 {
		client.SetSquads(squads)

		for _, activeSquad := range squads {
			if activeSquad.Active {
				client.SetSquad(activeSquad)
				return
			}
		}
	} else {

		storage, _ := storages.Storages.Get(client.GetID(), client.InBaseID)

		findMS := false

		// ищем тело в сторедже на базе
		for _, slot := range storage.Slots {
			if slot.Type == "body" {
				body, _ := gameTypes.Bodies.GetByID(slot.ItemID) // MS
				if body.MotherShip {
					findMS = true
					break
				}
			}
		}

		// если тела нет то надо выдать игроку стандартный набор снаряжения
		if !findMS {
			msBody, _ := gameTypes.Bodies.GetByID(2) // MS
			storages.Storages.AddItem(client.GetID(), client.InBaseID, msBody, "body",
				msBody.ID, 1, msBody.MaxHP, msBody.CapacitySize*float32(1), msBody.MaxHP)

			lightTank, _ := gameTypes.Bodies.GetByID(3) // L. tank
			storages.Storages.AddItem(client.GetID(), client.InBaseID, lightTank, "body",
				lightTank.ID, 3, lightTank.MaxHP, lightTank.CapacitySize*float32(3), lightTank.MaxHP)

			miningLaser, _ := gameTypes.Equips.GetByID(6) // mining laser
			storages.Storages.AddItem(client.GetID(), client.InBaseID, miningLaser, "equip",
				miningLaser.ID, 1, miningLaser.MaxHP, miningLaser.Size*float32(1), miningLaser.MaxHP)

			smallMissile, _ := gameTypes.Weapons.GetByID(5) // weapon
			storages.Storages.AddItem(client.GetID(), client.InBaseID, smallMissile, "weapon",
				smallMissile.ID, 3, smallMissile.MaxHP, smallMissile.Size*float32(3), smallMissile.MaxHP)

			ammoMissile, _ := gameTypes.Ammo.GetByID(5) // weapon
			storages.Storages.AddItem(client.GetID(), client.InBaseID, ammoMissile, "ammo",
				ammoMissile.ID, 50, 1, ammoMissile.Size*float32(50), 1)

			enrichedThorium, _ := gameTypes.Resource.GetRecycledByID(1) // топляк
			storages.Storages.AddItem(client.GetID(), client.InBaseID, enrichedThorium, "recycle",
				enrichedThorium.TypeID, 500, 1, enrichedThorium.Size*float32(50), 1)
		}
	}
}
