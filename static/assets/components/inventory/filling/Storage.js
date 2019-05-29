let categories = false;
let cellSize = 34;

function UpdateStorage(inventory) {
    if (!inventory) return;

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
            cell.onmousemove = InventoryOverTip;
            cell.onmouseout = OffTip;
            cell.source = 'storage';

            cell.style.height = cellSize + "px";
            cell.style.width = cellSize + "px";

            if (categories) {
                let section = CheckRecycleSection(inventory.slots[i], document.getElementById('inventoryStorage'));
                section.appendChild(cell);
            } else {
                inventoryStorage.append(cell);
            }
        }
    }
}