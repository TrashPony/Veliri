function SelectInventoryMod(Items, typeAction, acceptText, acceptFunc, cancelFunc) {

    checkConfirmMenu();
    DestroyInventoryTip();
    DestroyInventoryClickEvent();
    OffTip();
    RemoveActionConstructorMenu();

    for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
        if (cell.slotData) {
            cell.onclick = function () {
                SelectItems(cell, Items, i);
            };
            cell.onmousemove = null;
        }
    }

    let ConfirmMenu = document.createElement("div");
    ConfirmMenu.className = "ConfirmInventoryMenu";
    ConfirmMenu.id = "ConfirmInventoryMenu";
    ConfirmMenu.typeAction = typeAction;

    let equipButton = document.createElement("div");
    equipButton.innerHTML = acceptText;
    equipButton.onclick = acceptFunc;
    ConfirmMenu.appendChild(equipButton);

    let allButton = document.createElement("div");
    allButton.innerHTML = "Отмена";
    allButton.onclick = cancelFunc;
    ConfirmMenu.appendChild(allButton);

    document.getElementById("Inventory").appendChild(ConfirmMenu);
}

function SelectItems(cell, items, slot) {
    let cellData = JSON.parse(cell.slotData);

    cell.className = "InventoryCell Select Remove";
    items[slot] = cellData;
    console.log(cell.id);

    cell.onclick = function () {
        items[slot] = null;
        cell.className = "InventoryCell";
        this.onclick = function () {
            SelectItems(cell, items, slot);
        }
    };
}