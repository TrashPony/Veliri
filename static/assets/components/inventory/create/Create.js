function CreateInventoryMenu(closeFunc, option) {

    if (document.getElementById("inventoryBox") && option === 'constructor') {
        let jBox = $('#inventoryBox');
        setState('inventoryBox', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    if (document.getElementById("Inventory") && option === 'inventory') {
        let jBox = $('#Inventory');
        setState('Inventory', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    if (document.getElementById('wrapperInventoryAndStorage') && option === 'storage') {
        let jBox = $('#wrapperInventoryAndStorage');
        setState('wrapperInventoryAndStorage', jBox.position().left, jBox.position().top, jBox.height(), jBox.width(), false);
        return
    }

    if (option === 'inventory') {
        let inventory = document.createElement("div");
        inventory.id = "Inventory";
        inventory.style.position = "absolute";
        inventory.style.bottom = "70px";
        inventory.style.right = "15px";
        document.body.appendChild(inventory);
        CreateInventory();
        return
    }

    if (option === 'storage') {
        OnlyStorage();
        return
    }

    if (option === 'constructor') {
        OnlyConstructor();
    }
}

function InventoryClose() {
    document.getElementById("inventoryBox").remove();
    inventorySocket.close();
}

function CreateCells(typeSlot, count, className, idPrefix, parent, vertical) {
    for (let i = 0; i < count; i++) {
        let cell = document.createElement("div");
        cell.className = className;
        cell.id = idPrefix + Number(i + 1) + typeSlot;

        cell.type = typeSlot;
        cell.Number = Number(i + 1);

        parent.appendChild(cell);

        if (vertical) {
            let br = document.createElement("br");
            parent.appendChild(br);
        }
    }
}