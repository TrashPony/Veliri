function FillingInventory(jsonData) {
    var squad = JSON.parse(jsonData).squad;
    FillingInventoryTable(squad.inventory)
}

function FillingInventoryTable(inventoryItems) {
    for (var slot in inventoryItems) {
        if (inventoryItems.hasOwnProperty(slot)) {
            var cell = document.getElementById("inventory " + slot + 6);

            cell.item = inventoryItems[slot].item;
            cell.type = inventoryItems[slot].type;
            cell.quantity = inventoryItems[slot].quantity;
            cell.slot = slot;

            cell.style.backgroundImage = "url(/assets/" + cell.item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + cell.quantity + "</span>";

            cell.onclick = SelectInventoryItem
        }
    }
}

function FillingSquadTable() {

}

function FillingConstructorTable() {

}