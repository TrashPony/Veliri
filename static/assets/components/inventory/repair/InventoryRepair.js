function InventoryRepair() {
    inventorySocket.send(JSON.stringify({
        event: "InventoryRepair"
    }));
}

function overInventoryButton() {
    for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
        if (cell.slotData) {
            let percentHP = CreateHealBar(cell, "inventory", false);
            if (percentHP < 100) {
                cell.className = "InventoryCell Select";
            }
        }
    }
}

function outInventoryButton() {
    for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
        cell.className = "InventoryCell";
    }
}