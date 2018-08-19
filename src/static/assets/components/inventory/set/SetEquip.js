function SetEquip(equip, slot) {

    let msFunc = function (slotData) {
        inventorySocket.send(JSON.stringify({
            event: "SetMotherShipEquip",
            equip_id: Number(equip.id),
            inventory_slot: Number(slot),
            equip_slot:  Number(slotData.number_slot),
            equip_slot_type: Number(equip.type_slot)
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };
    EquipSlotMark("inventoryEquip", "inventoryEquipping", equip.type_slot, 5, msFunc);

    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        let unitSlot = JSON.parse(constructorUnit.slotData).number_slot;
        let unitFunc = function (slotData) {

            inventorySocket.send(JSON.stringify({
                event: "SetUnitEquip",
                equip_id: Number(equip.id),
                inventory_slot: Number(slot),
                equip_slot:  Number(slotData.number_slot),
                equip_slot_type: Number(equip.type_slot),
                unit_slot: Number(unitSlot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };
        EquipSlotMark("UnitEquip", "UnitEquip", equip.type_slot, 3, unitFunc);
    }
}

function EquipSlotMark(idPrefix, classPrefix, typeSlot, countSlots, func) {
    for (let i = 1; i <= countSlots; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + typeSlot);
        if (equipSlot.className === classPrefix + " active") {

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