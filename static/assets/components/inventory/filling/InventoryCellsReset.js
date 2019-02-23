function InventoryCellsReset() {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);
        if (cell) {
            if ($(cell).data("slotData") !== undefined && $(cell).data("slotData").data.item) {
                cell.className = "InventoryCell active";
            } else {
                cell.className = "InventoryCell";
            }
        }
    }
}