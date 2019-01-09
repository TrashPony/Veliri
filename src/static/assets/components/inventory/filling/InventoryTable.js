function InventoryTable(inventoryItems) {

    $('#Inventory').droppable({
        drop: function (event, ui) {
            let draggable = ui.draggable;
            if (draggable.data("slotData").parent === "squadInventory") {
                // todo перемещение слотов внутри инвентаря
            } else if(draggable.data("slotData").parent.split(':')[0] === "box") {
                let boxID = draggable.data("slotData").parent.split(':')[1];
                global.send(JSON.stringify({
                    event: "getItemFromBox",
                    box_id: Number(boxID),
                    slot: Number(draggable.data("slotData").number)
                }))
            }
        }
    });

    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.slots.hasOwnProperty(i) && inventoryItems.slots[i].item !== null) {

            cell.slotData = JSON.stringify(inventoryItems.slots[i]);
            cell.number = i;

            if (JSON.parse(cell.slotData).type === "resource" || JSON.parse(cell.slotData).type === "recycle") {
                cell.style.backgroundImage = "url(/assets/resource/" + JSON.parse(cell.slotData).item.name + ".png)";
            } else if(JSON.parse(cell.slotData).type === "boxes"){
                cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).type + "/" + JSON.parse(cell.slotData).item.name + ".png)";
            } else {
                cell.style.backgroundImage = "url(/assets/units/" + JSON.parse(cell.slotData).type + "/" + JSON.parse(cell.slotData).item.name + ".png)";
            }

            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            CreateHealBar(cell, "inventory", true);

            cell.onclick = SelectInventoryItem;
            cell.onmousemove = InventoryOverTip;
            cell.onmouseout = OffTip;

            $(cell).data( "slotData", {parent: "squadInventory", data: inventoryItems.slots[i], number: i});
            $(cell).draggable({
                revert: "invalid",
                zIndex: 999,
                helper: 'clone'
            });
        } else {

            cell.slotData = null;
            cell.style.backgroundImage = null;
            cell.innerHTML = "";

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