function InventoryCellsReset() {
    let resetFunc = function (id) {
        for (let i = 0; i < document.getElementById(id).childNodes.length; i++) {
            let cell = document.getElementById(id).childNodes[i];

            if (!cell.slotData) continue;

            if ($(cell).data("slotData") !== undefined && $(cell).data("slotData").data.item) {
                cell.className = "InventoryCell active";
            } else {
                cell.className = "InventoryCell";
            }
        }
    };

    resetFunc('inventoryStorageInventory');
    resetFunc('inventoryStorage');
}