function SelectInventoryItem(e) {
    DestroyInventoryTip();
    InventoryTip(this.slot.item, e.clientX, e.clientY);

    if (this.slot.type === "body") {
        SelectInventoryBody(this.slot.item, this.number);
    }

    if (this.slot.type === "weapon") {
        SelectInventoryWeapon(this.slot.item, this.number);
    }

    if (this.slot.type === "ammo") {
        console.log(this.slot.item);
        console.log(this.slot.quantity);
    }

    if (this.slot.type === "equip") {
        console.log(this.slot.item);
        console.log(this.slot.quantity);
    }
}

function SelectInventoryBody(body, slot) {
    if (body.mother_ship) {
        let shipIcon = document.getElementById("UnitIcon");
        shipIcon.className = "UnitIconSelect";
        shipIcon.onclick = function () {

            inventorySocket.send(JSON.stringify({
                event: "SetMotherShipBody",
                id_body: Number(body.id),
                inventory_slot: Number(slot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    }
}

function SelectInventoryWeapon(weapon, slot) {
    if (weapon) {
        for (let i = 1; i <= 5; i++ ) {
            let equipSlot = document.getElementById("inventoryEquip" + Number(i) + 3);
            if (equipSlot.className === "inventoryEquipping active weapon") {
                equipSlot.className = "inventoryEquipping active select";

                equipSlot.onclick = function () {
                    inventorySocket.send(JSON.stringify({
                        event: "SetMotherShipWeapon",
                        id_body: Number(weapon.id),
                        inventory_slot: Number(slot),
                        equip_slot: this.slot.number_slot
                    }));

                    DestroyInventoryClickEvent();
                    DestroyInventoryTip();
                }
            }
        }
    }
}