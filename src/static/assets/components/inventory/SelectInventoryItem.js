function SelectInventoryItem(e) {

    DestroyInventoryTip();
    DestroyInventoryClickEvent();

    InventoryTip(JSON.parse(this.slotData).item, e.clientX, e.clientY);

    if (JSON.parse(this.slotData).type === "body") {
        SelectInventoryBody(JSON.parse(this.slotData).item, this.number);
    }

    if (JSON.parse(this.slotData).type === "weapon") {
        SelectInventoryWeapon(JSON.parse(this.slotData).item, this.number);
    }

    if (JSON.parse(this.slotData).type === "equip") {
        SelectInventoryEquip(JSON.parse(this.slotData).item, this.number)
    }

    if (JSON.parse(this.slotData).type === "ammo") {
        SelectInventoryAmmo(JSON.parse(this.slotData).item, this.number)
    }
}

function SelectInventoryBody(body, slot) {
    if (body.mother_ship) {
        let shipIcon = document.getElementById("MSIcon");

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
                    equip_slot: Number(JSON.parse(this.slotData).number_slot)
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
                    equip_slot: Number(JSON.parse(this.slotData).number_slot),
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
                equip_slot: Number(JSON.parse(this.slotData).number_slot),
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    }
}