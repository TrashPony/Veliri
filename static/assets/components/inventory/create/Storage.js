function CreateStorage() {
    let storage = document.getElementById("storage");

    let spanInventory = document.createElement("span");
    spanInventory.className = "InventoryHead";
    spanInventory.innerHTML = "СКЛАД";
    spanInventory.style.margin = "-3px 0px 3px 45px";
    storage.appendChild(spanInventory);

    let inventoryStorage = document.createElement("div");
    inventoryStorage.className = "inventoryStorage";
    inventoryStorage.id = "inventoryStorage";
    inventoryStorage.style.height = "232px";
    inventoryStorage.style.margin = "0";
    $(inventoryStorage).mousedown(function (event) {
        // это костыль что бы работали полосы прокрутки, https://bugs.jqueryui.com/ticket/4441#no1
        if (event.offsetX >= event.target.clientWidth || event.offsetY >= event.target.clientHeight) {
            event.stopImmediatePropagation();
        }
    });
    $(inventoryStorage).selectable({
        filter: '.InventoryCell.active',
        start: function () {
            $('.ui-selected').removeClass('ui-selected')
        }
    });
    storage.appendChild(inventoryStorage);
    storage.appendChild(CreateSortPanel());
}