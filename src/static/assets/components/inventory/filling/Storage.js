function UpdateStorage(inventory) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("storage " + i + 6);

        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item !== null) {

            cell.slotData = JSON.stringify(inventory.slots[i]);
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/units/" + JSON.parse(cell.slotData).type + "/" + JSON.parse(cell.slotData).item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            CreateHealBar(cell, "inventory", true);

            //cell.onclick = SelectInventoryItem;

            cell.onmousemove = InventoryOverTip;
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