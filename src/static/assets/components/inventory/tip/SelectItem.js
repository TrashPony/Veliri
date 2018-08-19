function InventorySelectTip(item, x, y) {
    let tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";
    tip.id = "InventoryTip";

    let name = document.createElement("span");
    name.className = "InventoryTipName";
    name.innerHTML = item.name;
    tip.appendChild(name);

    let cancelButton = document.createElement("input");
    cancelButton.type = "button";
    cancelButton.className = "lobbyButton inventoryTip";
    cancelButton.value = "Отменить";
    cancelButton.style.pointerEvents = "auto";

    cancelButton.onclick = function () {
        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    tip.appendChild(cancelButton);

    document.body.appendChild(tip);
}