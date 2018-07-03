function InventoryTip(item, x, y) {
    var tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";
    tip.id = "InventoryTip";

    var name = document.createElement("span");
    name.className = "InventoryTipName";
    name.innerHTML = item.name;
    tip.appendChild(name);

    var cancelButton = document.createElement("input");
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

function DestroyInventoryTip() {
    if (document.getElementById("InventoryTip")) {
        document.getElementById("InventoryTip").remove();
    }
}

function DestroyInventoryClickEvent() {
    var shipIcon = document.getElementById("UnitIcon");
    shipIcon.className = "";
    shipIcon.onclick = null;
}