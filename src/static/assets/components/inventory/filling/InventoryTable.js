function InventoryTable(inventoryItems) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.hasOwnProperty(i) && inventoryItems[i].item !== null) {

            cell.slotData = JSON.stringify(inventoryItems[i]);
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            CreateHealBar(cell, "inventory");

            cell.onclick = SelectInventoryItem;

            cell.addEventListener("mousemove", InventoryOverTip);
            cell.addEventListener("mouseout", function () {
                let inventoryTip = document.getElementById("InventoryTipOver");
                if (inventoryTip) {
                    inventoryTip.remove()
                }
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
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(JSON.parse(this.slotData), e.clientX, e.clientY, true);
    }
}