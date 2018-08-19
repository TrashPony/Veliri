function RemoveTip(event, removeFunction) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    DestroyInventoryClickEvent();
    DestroyInventoryTip();

    let tip = document.createElement("div");
    tip.style.top = event.clientY + "px";
    tip.style.left = event.clientX + "px";
    tip.id = "InventoryTip";

    let removeButton = document.createElement("input");
    removeButton.type = "button";
    removeButton.className = "lobbyButton inventoryTip";
    removeButton.value = "Удалить";
    removeButton.style.pointerEvents = "auto";

    removeButton.onclick = removeFunction;

    let cancelButton = document.createElement("input");
    cancelButton.type = "button";
    cancelButton.className = "lobbyButton inventoryTip";
    cancelButton.value = "Отменить";
    cancelButton.style.pointerEvents = "auto";

    cancelButton.onclick = function (event) {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    tip.appendChild(removeButton);
    tip.appendChild(cancelButton);
    document.body.appendChild(tip);
}