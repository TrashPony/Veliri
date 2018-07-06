function FillingInventoryTable(inventoryItems) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.hasOwnProperty(i) && inventoryItems[i].item !== null) {

            cell.slot = inventoryItems[i];
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/" + cell.slot.item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + cell.slot.quantity + "</span>";

            cell.onclick = SelectInventoryItem
        } else {
            cell.slot = null;

            cell.style.backgroundImage = null;
            cell.innerHTML = "";

            cell.onclick = function () {
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            };
        }
    }
}