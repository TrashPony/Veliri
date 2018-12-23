function ThrowItems() {

    let throwItems = [];

    let acceptFunc = function () {
        if (global){
            global.send(JSON.stringify({
                event: "ThrowItems",
                throw_items: throwItems
            }))
        }
        console.log(throwItems, "Выбросить");
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

    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);
        cell.onclick = SelectInventoryItem;
        cell.onmousemove = InventoryOverTip;
        cell.className = "InventoryCell";
    }

    ActionConstructorMenu();
}