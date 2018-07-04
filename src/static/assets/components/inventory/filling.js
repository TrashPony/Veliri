function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory") {
        let squad = JSON.parse(jsonData).squad;
        FillingInventoryTable(squad.inventory);
    }

    if (event === "UpdateSquad") {
        console.log(jsonData)
    }
}

function FillingInventoryTable(inventoryItems) {
    for (let slot in inventoryItems) {
        if (inventoryItems.hasOwnProperty(slot)) {
            let cell = document.getElementById("inventory " + slot + 6);

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