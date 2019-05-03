package squad_inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/get"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
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

			getAmmo(1, client)
			getAmmo(2, client)
			getAmmo(3, client)
			getAmmo(4, client)
			getAmmo(5, client)
			getAmmo(6, client)

			getResource(1, client) // топливо
			getResource(2, client)
			getResource(3, client)
			getResource(4, client)
			getResource(5, client)
			getResource(6, client)

			getBox(1, client)
			getBox(2, client)
			getBox(3, client)

			for i := 1; i < 36; i++ {
				getBlueprints(i, client)
			}
		}
	}
}

func addBodyToStorage(id int, client *player.Player) {
	body, _ := gameTypes.Bodies.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, body, "body",
		body.ID, 1, body.MaxHP, body.CapacitySize*float32(1), body.MaxHP, false)
}

func addWeapon(id int, client *player.Player) {
	weapon, _ := gameTypes.Weapons.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, weapon, "weapon",
		weapon.ID, 1, weapon.MaxHP, weapon.Size*float32(1), weapon.MaxHP, false)
}

func addEquip(id int, client *player.Player) {
	equip, _ := gameTypes.Equips.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, equip, "equip",
		equip.ID, 1, equip.MaxHP, equip.Size*float32(1), equip.MaxHP, false)
}

func getAmmo(id int, client *player.Player) {
	ammo, _ := gameTypes.Ammo.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, ammo, "ammo",
		ammo.ID, 50, 1, ammo.Size*float32(50), 1, false)
}

func getResource(id int, client *player.Player) {
	res, _ := gameTypes.Resource.GetRecycledByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, res, "recycle",
		res.TypeID, 500, 1, res.Size*float32(500), 1, false)
}

func getBox(id int, client *player.Player) {
	box, _ := gameTypes.Boxes.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, box, "recycle",
		box.TypeID, 2, 1, box.FoldSize*float32(2), 1, false)
}

func getBlueprints(id int, client *player.Player) {
	bp, _ := gameTypes.BluePrints.GetByID(id)
	storages.Storages.AddItem(client.GetID(), client.InBaseID, bp, "blueprints",
		bp.ID, 2, 1, 0, 1, false)
}
