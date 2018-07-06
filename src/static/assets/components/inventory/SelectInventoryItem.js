function SelectInventoryItem(e) {

    DestroyInventoryTip();
    DestroyInventoryClickEvent();

    InventoryTip(this.slot.item, e.clientX, e.clientY);

    if (this.slot.type === "body") {
        SelectInventoryBody(this.slot.item, this.number);
    }

    if (this.slot.type === "weapon") {
        SelectInventoryWeapon(this.slot.item, this.number);
    }

    if (this.slot.type === "equip") {
        SelectInventoryEquip(this.slot.item, this.number)
    }

    if (this.slot.type === "ammo") {
        SelectInventoryAmmo(this.slot.item, this.number)
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
    for (let i = 1; i <= 5; i++) {
        let equipSlot = document.getElementById("inventoryEquip" + Number(i) + 3); // оружие всегда ствиться в 3 слоты по диз-доку
        if (equipSlot.className === "inventoryEquipping active weapon") {

            equipSlot.className = "inventoryEquipping active select";
            equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
            equipSlot.style.cursor = "pointer";
            equipSlot.onmouseout = null;

            equipSlot.onclick = function () {
                inventorySocket.send(JSON.stringify({
                    event: "SetMotherShipWeapon",
                    weapon_id: Number(weapon.id),
                    inventory_slot: Number(slot),
                    equip_slot: this.slot.number_slot
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}

function SelectInventoryEquip(equip, slot) {
    for (let i = 1; i <= 5; i++) {
        let equipSlot = document.getElementById("inventoryEquip" + Number(i) + equip.type_slot); // оружие всегда ствиться в 3 слоты по диз-доку
        if (equipSlot.className === "inventoryEquipping active") {

            equipSlot.className = "inventoryEquipping active select";
            equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
            equipSlot.style.cursor = "pointer";
            equipSlot.onmouseout = null;

            equipSlot.onclick = function () {
                inventorySocket.send(JSON.stringify({
                    event: "SetMotherShipEquip",
                    equip_id: Number(equip.id),
                    inventory_slot: Number(slot),
                    equip_slot: this.slot.number_slot,
                    equip_slot_type: Number(equip.type_slot)
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}

function SelectInventoryAmmo(ammo, slot) {

    let ammoCells = document.getElementsByClassName("inventoryAmmoCell");

    for (let i = 0; i < ammoCells.length; i++) {

        ammoCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        ammoCells[i].style.cursor = "pointer";
        ammoCells[i].onmouseout = null;
        ammoCells[i].onclick = function (event) {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            inventorySocket.send(JSON.stringify({
                event: "SetMotherShipAmmo",
                ammo_id: Number(ammo.id),
                inventory_slot: Number(slot),
                equip_slot: this.slot.number_slot,
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    }
}