function CreateInventory() {
    let inventory = document.getElementById("Inventory");
    $(inventory).resizable({
        alsoResize: "#inventoryStorageInventory",
        alsoResizeReverse: "#storage, #inventoryStorage",
        minHeight: 105,
        maxHeight: 307,
        maxWidth: 163,
        minWidth: 163,
        handles: "s",
        resize() {
            $(this).resizable("option", "maxHeight", ($(this).height() + $('#storage').height()) - 50);
        }
    });

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
    inventoryStorage.id = "inventoryStorageInventory";
    $(inventoryStorage).mousedown(function (event) {
        // это костыль что бы работали полосы прокрутки, https://bugs.jqueryui.com/ticket/4441#no1
        if (event.offsetX >= event.target.clientWidth || event.offsetY >= event.target.clientHeight) {
            event.stopImmediatePropagation();
        }
    });
    $(inventoryStorage).selectable({
        filter: '.InventoryCell.active',
        start: function () {
            $('.ui-selected').removeClass('ui-selected');
        }
    });

    inventory.appendChild(inventoryStorage);
}