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

		if client.InBaseID == 0 {
			return
		}

		storage, _ := storages.Storages.Get(client.GetID(), client.InBaseID)

		findMS := false

		// ищем тело в сторедже на базе
		if storage != nil {
			for _, slot := range storage.Slots {
				if slot.Type == "body" {
					body, _ := gameTypes.Bodies.GetByID(slot.ItemID) // MS
					if body.MotherShip {
						findMS = true
						break
					}
				}
			}
		}

		// если тела нет то надо выдать игроку стандартный набор снаряжения
		if !findMS {
			// ms ки
			addBodyToStorage(5, client)
			addBodyToStorage(6, client)
			addBodyToStorage(7, client)

			// юниты
			addBodyToStorage(3, client)
			addBodyToStorage(4, client)
			addBodyToStorage(1, client)

			addEquip(1, client)
			addEquip(2, client)
			addEquip(3, client)
			addEquip(4, client)
			addEquip(5, client)
			addEquip(6, client)
			addEquip(7, client)
			addEquip(8, client)

			addWeapon(1, client)
			addWeapon(2, client)
			addWeapon(3, client)
			addWeapon(4, client)
			addWeapon(5, client)
			addWeapon(6, client)

			ammoMissile, _ := gameTypes.Ammo.GetByID(5) // weapon
			storages.Storages.AddItem(client.GetID(), client.InBaseID, ammoMissile, "ammo",
				ammoMissile.ID, 50, 1, ammoMissile.Size*float32(50), 1)

			enrichedThorium, _ := gameTypes.Resource.GetRecycledByID(1) // топляк
			storages.Storages.AddItem(client.GetID(), client.InBaseID, enrichedThorium, "recycle",
				enrichedThorium.TypeID, 500, 1, enrichedThorium.Size*float32(50), 1)
		}
	}
}

func addBodyToStorage(id int, client *player.Player) {
	body, _ := gameTypes.Bodies.GetByID(id) // MS
	storages.Storages.AddItem(client.GetID(), client.InBaseID, body, "body",
		body.ID, 1, body.MaxHP, body.CapacitySize*float32(1), body.MaxHP)
}

func addWeapon(id int, client *player.Player) {
	smallMissile, _ := gameTypes.Weapons.GetByID(id) // weapon
	storages.Storages.AddItem(client.GetID(), client.InBaseID, smallMissile, "weapon",
		smallMissile.ID, 1, smallMissile.MaxHP, smallMissile.Size*float32(3), smallMissile.MaxHP)
}

func addEquip(id int, client *player.Player) {
	miningLaser, _ := gameTypes.Equips.GetByID(id) // mining laser
	storages.Storages.AddItem(client.GetID(), client.InBaseID, miningLaser, "equip",
		miningLaser.ID, 1, miningLaser.MaxHP, miningLaser.Size*float32(1), miningLaser.MaxHP)
}
