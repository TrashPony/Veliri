function UpdateStorage(inventory) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("storage " + i + 6);

        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item !== null) {

            cell.slotData = JSON.stringify(inventory.slots[i]);
            cell.number = i;
            cell.inventoryType = 'storage';

            if (JSON.parse(cell.slotData).type === "resource" || JSON.parse(cell.slotData).type === "recycle") {
                cell.style.backgroundImage = "url(/assets/resource/" + JSON.parse(cell.slotData).item.name + ".png)";
            } else {
                cell.style.backgroundImage = "url(/assets/units/" + JSON.parse(cell.slotData).type + "/" + JSON.parse(cell.slotData).item.name + ".png)";
            }

            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            CreateHealBar(cell, "inventory", true);

            cell.onclick = SelectInventoryItem;

            cell.onmousemove = StorageOverTip;
            cell.onmouseout = OffTip;
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


function StorageOverTip(e) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true, false);
    }
}