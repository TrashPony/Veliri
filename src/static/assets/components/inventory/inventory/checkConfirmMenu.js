function checkConfirmMenu() {
    let ConfirmMenu = document.getElementById("ConfirmInventoryMenu");

    if (ConfirmMenu) {
        if (ConfirmMenu.typeAction === "recycle") {
            cancelRecycle();
        } else if (ConfirmMenu.typeAction === "throw") {
            cancelThrow();
        }
        ConfirmMenu.remove()
    }
}