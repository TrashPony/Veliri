function SetEquip(equip, slot) {
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