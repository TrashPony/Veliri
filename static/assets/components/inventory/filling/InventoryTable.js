function InventoryTable(inventoryItems) {

    $('#Inventory').droppable({
        drop: function (event, ui) {

            $('.ui-selected').removeClass('ui-selected');

            let draggable = ui.draggable;
            if (draggable.data("slotData").parent === "squadInventory") {
            } else if (draggable.data("slotData").parent.split(':')[0] === "box") {

                let boxID = draggable.data("slotData").parent.split(':')[1];

                if (draggable.data("selectedItems") !== undefined) {
                    global.send(JSON.stringify({
                        event: "getItemsFromBox",
                        box_id: Number(boxID),
                        slots: draggable.data("selectedItems").slotsNumbers
                    }));
                    $(draggable).removeData("selectedItems");
                } else {
                    global.send(JSON.stringify({
                        event: "getItemFromBox",
                        box_id: Number(boxID),
                        slot: Number(draggable.data("slotData").number)
                    }))
                }
            } else if (draggable.data("slotData").parent === "storage") {
                if (draggable.data("selectedItems") !== undefined) {
                    inventorySocket.send(JSON.stringify({
                        event: "itemsToInventory",
                        storage_slots: draggable.data("selectedItems").slotsNumbers
                    }));
                } else {
                    inventorySocket.send(JSON.stringify({
                        event: "itemToInventory",
                        storage_slot: Number(draggable.data("slotData").number)
                    }));
                }
            } else if (draggable.data("slotData").parent === "Constructor") {
                inventorySocket.send(JSON.stringify({
                    event: draggable.data("slotData").event,
                    equip_slot: Number(draggable.data("slotData").equipSlot),
                    equip_slot_type: Number(draggable.data("slotData").equipType),
                    unit_slot: Number(draggable.data("slotData").unitSlot),
                    destination: "squadInventory",
                }));
            }
        }
    });

    let inventoryStorage = $('#inventoryStorageInventory');
    let parent = "squadInventory";
    inventoryStorage.find(".nameSection").remove();

    if (!inventoryItems) {
        return
    }

    for (let i in inventoryItems.slots) {
        if (inventoryItems.slots.hasOwnProperty(i) && inventoryItems.slots[i].item !== null) {

            let cell = document.getElementById(parent + i);
            if (!cell) {

                cell = document.createElement("div");
                cell.className = "InventoryCell active";

                CreateInventoryCell(cell, inventoryItems.slots[i], i, parent);
            } else {
                UpdateCell(cell, inventoryItems.slots[i]);
            }

            cell.onclick = SelectInventoryItem;
            cell.source = parent;
            cell.style.height = cellSize + "px";
            cell.style.width = cellSize + "px";

            if (categories) {
                let section = CheckRecycleSection(inventoryItems.slots[i], document.getElementById('inventoryStorageInventory'));
                section.appendChild(cell);
            } else {
                inventoryStorage.append(cell);
            }
        }
    }

    DeleteNotUpdateSlots(parent)
}

function InventoryOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = stylePositionParams.top + "px";
        inventoryTip.style.left = stylePositionParams.left + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), true, true);
    }
}