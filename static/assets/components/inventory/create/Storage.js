function CreateStorage(){
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
        start: function() {$('.ui-selected').removeClass('ui-selected')}
    });
    storage.appendChild(inventoryStorage);

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
    storage.appendChild(sortPanel);

    ConnectStorage();
}