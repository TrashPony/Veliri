function ThrowItems() {

    let throwItems = [];

    let acceptFunc = function () {
        if (typeof (global) !== 'undefined') {
            global.send(JSON.stringify({
                event: "ThrowItems",
                throw_items: throwItems
            }))
        } else {
            inventorySocket.send(JSON.stringify({
                event: "destroyItems",
                throw_items: throwItems,
            }));
        }
        cancelThrow();
    };

    SelectInventoryMod(throwItems, "throw", "Выбросить", acceptFunc, cancelThrow);

    this.className = "throwButtonActive";
    this.onclick = cancelThrow;
}

function cancelThrow() {
    document.getElementById("ConfirmInventoryMenu").remove();
    document.getElementsByClassName("throwButtonActive")[0].className = "destroyButton";
    document.getElementsByClassName("destroyButton")[0].onclick = ThrowItems;

    for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
        let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
        cell.onclick = SelectInventoryItem;
        cell.onmousemove = InventoryOverTip;
        cell.className = "InventoryCell";
    }

    ActionConstructorMenu();
}