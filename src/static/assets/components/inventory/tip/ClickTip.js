function ClickTip(event, removeFunction) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    DestroyInventoryClickEvent();
    DestroyInventoryTip();

    let tip = document.createElement("div");
    tip.style.top = event.clientY + "px";
    tip.style.left = event.clientX + "px";
    tip.id = "InventoryTipSelect";

    let detailedButton = document.createElement("input");
    detailedButton.type = "button";
    detailedButton.className = "lobbyButton inventoryTip";
    detailedButton.value = "Подробнее";
    detailedButton.style.pointerEvents = "auto";
    // TODO detailedButton.onclick = функция вывода подробной информации

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

    tip.appendChild(detailedButton);

    if (removeFunction) {
        let removeButton = document.createElement("input");
        removeButton.type = "button";
        removeButton.className = "lobbyButton inventoryTip";
        removeButton.value = "Удалить";
        removeButton.style.pointerEvents = "auto";
        removeButton.onclick = removeFunction;
        tip.appendChild(removeButton);
    }

    tip.appendChild(cancelButton);
    document.body.appendChild(tip);
}

function OffTip() {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.remove()
    }

    let itemSize = document.getElementById("itemSize");
    if (itemSize) {
        itemSize.remove()
    }
}