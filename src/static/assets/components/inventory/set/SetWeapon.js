function SetWeapon(weapon, slot) {

    let msFunc = function (slotData) {
        inventorySocket.send(JSON.stringify({
            event: "SetMotherShipWeapon",
            weapon_id: Number(weapon.id),
            inventory_slot: Number(slot),
            equip_slot: Number(slotData.number_slot)
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };
    WeaponSlotMark("inventoryEquip", "inventoryEquipping", 5, msFunc);

    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        let unitSlot = JSON.parse(constructorUnit.slotData).number_slot;
        let unitFunc = function (slotData) {

            inventorySocket.send(JSON.stringify({
                event: "SetUnitWeapon",
                weapon_id: Number(weapon.id),
                inventory_slot: Number(slot),
                equip_slot: Number(slotData.number_slot),
                unit_slot: Number(unitSlot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };
        WeaponSlotMark("UnitEquip", "UnitEquip", 3, unitFunc);
    }
}

function WeaponSlotMark(idPrefix, classPrefix, countSlots, func) {
    for (let i = 1; i <= countSlots; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + 3); // оружие всегда ствиться в 3 слоты по диз-доку
        if (equipSlot.className === classPrefix + " active weapon") {

            equipSlot.className = classPrefix + " active select";
            equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
            equipSlot.style.cursor = "pointer";
            equipSlot.onmouseout = null;

            equipSlot.onclick = function () {
                func(JSON.parse(equipSlot.slotData));
            }
        }
    }
}

