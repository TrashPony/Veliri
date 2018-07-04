function InventoryTip(item, x, y) {
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

function DestroyInventoryTip() {
    if (document.getElementById("InventoryTip")) {
        document.getElementById("InventoryTip").remove();
    }
}

function DestroyInventoryClickEvent() {
    let shipIcon = document.getElementById("UnitIcon");
    shipIcon.className = "";
    shipIcon.onclick = null;

    cellEquipDestoySelect(1, 5, "inventoryEquip");
    cellEquipDestoySelect(2, 5, "inventoryEquip");
    cellEquipDestoySelect(3, 5, "inventoryEquip");
    cellEquipDestoySelect(5, 2, "inventoryEquip");
}

function cellEquipDestoySelect(typeSlot, count, idPrefix) {
    for (let i = 1; i <= count; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + typeSlot);
        if (equipSlot.className === "inventoryEquipping active select") {
            if (equipSlot.slot.hasOwnProperty("weapon")) {
                equipSlot.className = "inventoryEquipping active weapon";
                equipSlot.onclick = null;
            } else {
                equipSlot.className = "inventoryEquipping active";
                equipSlot.onclick = null;
            }
        }
    }
}