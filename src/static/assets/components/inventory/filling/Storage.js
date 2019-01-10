function UpdateStorage(inventory) {

    $('#storage').droppable({
        drop: function (event, ui) {
            let draggable = ui.draggable;
            if (draggable.data("slotData").parent === "squadInventory") {
                // todo
                console.log("dfdfd")
            } else if (draggable.data("slotData").parent === "storage"){
            }
        }
    });

    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("storage " + i + 6);
        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item !== null) {
            CreateInventoryCell(cell, inventory.slots[i], i, "storage");
            cell.onclick = SelectInventoryItem;
            cell.onmousemove = StorageOverTip;
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


function StorageOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true, false);
    }
}