function UpdateStorage(inventory) {

    $('#storage').droppable({
        drop: function (event, ui) {

            $('.ui-selected').removeClass('ui-selected');

            let draggable = ui.draggable;
            if (draggable.data("slotData").parent === "squadInventory") {
                if (draggable.data("selectedItems") !== undefined) {
                    inventorySocket.send(JSON.stringify({
                        event: "itemsToStorage",
                        inventory_slots: draggable.data("selectedItems").slotsNumbers,
                    }));
                } else {
                    inventorySocket.send(JSON.stringify({
                        event: "itemToStorage",
                        inventory_slot: Number(draggable.data("slotData").number)
                    }));
                }
            } else if (draggable.data("slotData").parent === "Constructor") {
                inventorySocket.send(JSON.stringify({
                    event: draggable.data("slotData").event,
                    equip_slot: Number(draggable.data("slotData").equipSlot),
                    equip_slot_type: Number(draggable.data("slotData").equipType),
                    unit_slot: Number(draggable.data("slotData").unitSlot),
                    destination: "storage",
                }));
            } else if (draggable.data("slotData").parent === "storage") {
            }
        }
    });

    let inventoryStorage = $('#inventoryStorage');
    inventoryStorage.empty();
    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item !== null) {
            let cell = document.createElement("div");
            cell.className = "InventoryCell active";
            CreateInventoryCell(cell, inventory.slots[i], i, "storage");
            cell.onclick = SelectInventoryItem;
            cell.onmousemove = StorageOverTip;
            cell.onmouseout = OffTip;
            cell.source = 'storage';
            inventoryStorage.append(cell);
        }
    }
}


function StorageOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true, false);
    }
}