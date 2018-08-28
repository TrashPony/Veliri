function InventorySelectTip(slot, x, y) {
    let tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";
    tip.id = "InventoryTip";

    let name = document.createElement("span");
    name.className = "InventoryTipName";
    name.innerHTML = slot.item.name;
    tip.appendChild(name);

    let paramsTable = CreateParamsTable(slot);
    tip.appendChild(paramsTable);

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

function CreateParamsTable(slot) {
    let table = document.createElement("table");
    table.id = "paramsItem";

    if (slot.type === "weapon" || slot.type === "equip") {
        table.appendChild(createRow("Use power:", slot.item.power));
        table.appendChild(createRow("HP:", slot.hp));
    } else if (slot.type === "body") {
        table.appendChild(createRow("Power:", slot.item.max_power));
        table.appendChild(createRow("HP:", slot.hp));
    } else if (slot.type === "ammo"){
        table.appendChild(createRow("Type:", slot.item.type));
        table.appendChild(createRow("Damage:", slot.item.damage));
    }

    function createRow(params, count) {
        let tr = document.createElement("tr");

        let td1 = document.createElement("td");
        td1.innerHTML = "<span class='Value'>" + params + "</span>";
        td1.style.border = "0px";
        let td2 = document.createElement("td");
        td2.innerHTML = "<span class='Success'>" + count + "</span>";
        td2.style.border = "0px";

        tr.appendChild(td1);
        tr.appendChild(td2);

        return tr
    }

    return table;
}