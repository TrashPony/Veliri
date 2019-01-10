function InventoryTable(inventoryItems) {

    $('#Inventory').droppable({
        drop: function (event, ui) {
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
            } else if (draggable.data("slotData").parent === "storage"){
                // todo
                console.log("dfdfd")
            }
        }
    });

    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.slots.hasOwnProperty(i) && inventoryItems.slots[i].item !== null) {

            CreateInventoryCell(cell, inventoryItems.slots[i], i, "squadInventory");
            cell.onclick = SelectInventoryItem;
            cell.onmousemove = InventoryOverTip;
            cell.onmouseout = OffTip;

        } else {

            cell.slotData = null;
            cell.style.backgroundImage = null;
            cell.innerHTML = "";
            cell.className = "InventoryCell";

            $(cell).draggable({
                disabled: true
            });

            cell.onclick = function () {
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            };
        }
    }
}

function InventoryOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = stylePositionParams.top + "px";
        inventoryTip.style.left = stylePositionParams.left + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true, true);
    }
}