function InventoryCellsReset() {
    for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
        if ($(cell).data("slotData") !== undefined && $(cell).data("slotData").data.item) {
            cell.className = "InventoryCell active";
        } else {
            cell.className = "InventoryCell";
        }
    }

    for (let i = 0; i < document.getElementById('inventoryStorage').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorage').childNodes[i];
        if ($(cell).data("slotData") !== undefined && $(cell).data("slotData").data.item) {
            cell.className = "InventoryCell active";
        } else {
            cell.className = "InventoryCell";
        }
    }
}