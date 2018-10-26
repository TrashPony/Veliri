function InventoryRepair() {
    inventorySocket.send(JSON.stringify({
        event: "InventoryRepair"
    }));
}

function overInventoryButton() {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);
        if (cell.slotData) {
            let percentHP = CreateHealBar(cell, "inventory", false);
            if (percentHP < 100) {
                cell.className = "InventoryCell Select";
            }
        }
    }
}

function outInventoryButton() {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);
        cell.className = "InventoryCell";
    }
}