function CreateSortPanel() {
    let sortPanel = document.createElement("div");
    sortPanel.className = "sortPanel";

    let sortButton1 = document.createElement("div");
    sortButton1.innerHTML = 'C';
    sortButton1.id = 'sortButton1';
    sortButton1.onclick = function () {
        Categories();
    };

    let sortButton2 = document.createElement("div");
    sortButton2.id = 'sortButton2';
    sortButton2.innerHTML = '+';
    $(sortButton2).click(function () {
        changeSize('#Inventory .InventoryCell', 1);
        changeSize('#storage .InventoryCell', 1);
    });

    let sortButton3 = document.createElement("div");
    sortButton3.id = 'sortButton3';
    sortButton3.innerHTML = '-';

    $(sortButton3).click(function () {
        changeSize('#Inventory .InventoryCell', -1);
        changeSize('#storage .InventoryCell', -1);
    });

    sortPanel.appendChild(sortButton1);
    sortPanel.appendChild(sortButton2);
    sortPanel.appendChild(sortButton3);

    return sortPanel
}

function changeSize(selector, change) {
    if ($(selector).width() + change > 16 && $(selector).width() + change < 50) {
        $(selector).css('width', $(selector).width() + change);
        $(selector).css('height', $(selector).height() + change);
        cellSize = $(selector).height();
    }
}

function Categories() {
    categories = !categories;
    inventorySocket.send(JSON.stringify({
        event: "openInventory"
    }));
}