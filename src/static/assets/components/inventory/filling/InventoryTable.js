function InventoryTable(inventoryItems) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.hasOwnProperty(i) && inventoryItems[i].item !== null) {

            cell.slotData = JSON.stringify(inventoryItems[i]);
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).quantity + "</span>";

            cell.onclick = SelectInventoryItem

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