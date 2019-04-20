function SetEquip(equip, slot, source) {

    let msFunc = function (slotData) {
        inventorySocket.send(JSON.stringify({
            event: "SetMotherShipEquip",
            equip_id: Number(equip.id),
            inventory_slot: Number(slot),
            equip_slot: Number(slotData.number_slot),
            equip_slot_type: Number(equip.type_slot),
            source: source,
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    if (equip.applicable === "ore" || equip.applicable === "digger") {
        EquipSlotMark("inventoryEquip", "inventoryEquipping", equip.type_slot, 5, msFunc, true);
    } else {
        EquipSlotMark("inventoryEquip", "inventoryEquipping", equip.type_slot, 5, msFunc, false);
    }


    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        let unitSlot = JSON.parse(constructorUnit.slotData).number_slot;
        let unitFunc = function (slotData) {

            inventorySocket.send(JSON.stringify({
                event: "SetUnitEquip",
                equip_id: Number(equip.id),
                inventory_slot: Number(slot),
                equip_slot: Number(slotData.number_slot),
                equip_slot_type: Number(equip.type_slot),
                unit_slot: Number(unitSlot),
                source: source,
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        if (equip.applicable === "ore" || equip.applicable === "digger") {
            EquipSlotMark("UnitEquip", "UnitEquip", equip.type_slot, 3, unitFunc, true);
        } else {
            EquipSlotMark("UnitEquip", "UnitEquip", equip.type_slot, 3, unitFunc, false);
        }
    }
}

function EquipSlotMark(idPrefix, classPrefix, typeSlot, countSlots, func, mining) {
    for (let i = 1; i <= countSlots; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + typeSlot);

        if (equipSlot && $(equipSlot).hasClass("active") && !$(equipSlot).hasClass("weapon")) {

            if (!JSON.parse(equipSlot.slotData).mining && mining) {
                continue
            }

            equipSlot.className = classPrefix + " active select";
            equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
            equipSlot.style.cursor = "pointer";
            equipSlot.onmouseout = function () {
                for (let i = 1; i <= 40; i++) {
                    let cell = document.getElementById("inventory " + i + 6);
                    cell.className = "InventoryCell";
                }
            };

            equipSlot.onclick = function () {
                func(JSON.parse(equipSlot.slotData));
            }
        }
    }
}