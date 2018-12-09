package get

import (
	inv "../../gameObjects/inventory"
	"database/sql"
	"log"
)

func FillInventory(inventory *inv.Inventory, rows *sql.Rows) {
	for rows.Next() {

		var inventorySlot = inv.Slot{}
		var slot int

		err := rows.Scan(&slot, &inventorySlot.Type, &inventorySlot.ItemID, &inventorySlot.Quantity, &inventorySlot.HP)
		if err != nil {
			log.Fatal("scan inventory slots " + err.Error())
		}

		if inventorySlot.Type == "weapon" {
			weapon := Weapon(inventorySlot.ItemID)
			inventorySlot.Item = weapon
			inventorySlot.Size = weapon.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = weapon.MaxHP

			inventory.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "ammo" {
			ammo := Ammo(inventorySlot.ItemID)
			inventorySlot.Item = ammo
			inventorySlot.Size = ammo.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = 1 // у аммо нет хп

			inventory.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "equip" {
			equip := TypeEquip(inventorySlot.ItemID)
			inventorySlot.Item = equip
			inventorySlot.Size = equip.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = equip.MaxHP

			inventory.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "body" {
			body := Body(inventorySlot.ItemID)
			inventorySlot.Item = body
			inventorySlot.Size = body.CapacitySize * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = body.MaxHP

			inventory.Slots[slot] = &inventorySlot
		}
	}
}
