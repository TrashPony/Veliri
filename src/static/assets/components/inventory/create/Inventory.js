function CreateInventory() {
    let inventory = document.getElementById("Inventory");

    let spanInventory = document.createElement("span");
    spanInventory.className = "InventoryHead";
    spanInventory.id = "InventoryHead";
    spanInventory.innerHTML = "ТРЮМ";
    spanInventory.style.margin = "2px 0 3px 12px";
    inventory.appendChild(spanInventory);

    let recycleButton = document.createElement("div");
    recycleButton.className = "utilButton";
    recycleButton.id = "utilButton";
    recycleButton.innerHTML = "<div></div>";
    recycleButton.onclick = RecycleItems;
    inventory.appendChild(recycleButton);

    let throwButton = document.createElement("div");
    throwButton.className = "destroyButton";
    throwButton.id = "destroyButton";
    throwButton.innerHTML = "<div></div>";
    throwButton.onclick = ThrowItems;
    inventory.appendChild(throwButton);

    let sizeInventoryInfo = document.createElement("div");
    sizeInventoryInfo.id = "sizeInventoryInfo";
    inventory.appendChild(sizeInventoryInfo);

    let inventoryStorage = document.createElement("div");
    inventoryStorage.className = "inventoryStorage";
    $(inventoryStorage).selectable({
        filter: '.InventoryCell.active',
        start: function() {$('.ui-selected').removeClass('ui-selected')}
    });

    CreateCells(6, 40, "InventoryCell", "inventory ", inventoryStorage);
    inventory.appendChild(inventoryStorage);

    let sortPanel = document.createElement("div");
    sortPanel.className = "sortPanel";

    let sortButton1 = document.createElement("div");
    sortButton1.sort = "1";
    sortButton1.onclick = SortingItems;
    sortPanel.appendChild(sortButton1);
    let sortButton2 = document.createElement("div");
    sortButton2.sort = "2";
    sortButton2.onclick = SortingItems;
    sortPanel.appendChild(sortButton2);
    let sortButton3 = document.createElement("div");
    sortButton3.sort = "3";
    sortButton3.onclick = SortingItems;
    sortPanel.appendChild(sortButton3);
    inventory.appendChild(sortPanel);
}