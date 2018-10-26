function RecycleItems() {

    let recycleItems = [];

    let acceptFunc = function () {
        //todo
        console.log(recycleItems, "переработка");
        cancelRecycle();
    };

    SelectInventoryMod(recycleItems, "recycle", "Переработать", acceptFunc, cancelRecycle);

    this.className = "utilButtonActive";
    this.onclick = cancelRecycle;
}

function cancelRecycle() {
    document.getElementById("ConfirmInventoryMenu").remove();
    document.getElementsByClassName("utilButtonActive")[0].className = "utilButton";
    document.getElementsByClassName("utilButton")[0].onclick = RecycleItems;

    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);
        cell.onclick = SelectInventoryItem;
        cell.onmousemove = InventoryOverTip;
        cell.className = "InventoryCell";
    }

    ActionConstructorMenu();
}