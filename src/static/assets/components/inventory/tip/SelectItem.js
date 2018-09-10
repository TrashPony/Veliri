function InventorySelectTip(slot, x, y, first) {
    let tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";
    if (!first) {
        tip.id = "InventoryTip";
    } else {
        let inventoryTip = document.getElementById("InventoryTip");
        if (!inventoryTip) {
            tip.id = "InventoryTipOver";
        } else {
            tip.remove();
            return
        }
    }

    let name = document.createElement("span");
    name.className = "InventoryTipName";
    name.innerHTML = slot.item.name;
    tip.appendChild(name);

    CreateParamsTable(slot, tip);

    if (!first) {
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
    }

    document.body.appendChild(tip);
}

function CreateParamsTable(slot, tip) {
    let table = document.createElement("table");
    table.id = "paramsItem";

    let description = document.createElement("div");
    description.id = "description";

    if (slot.type === "weapon") {
        description.innerHTML = "<span class='Value'>Оружие универсального предназначения</span>";
        table.appendChild(createRow("Use power:", slot.item.power));
        table.appendChild(createRow("HP:", slot.hp));
    } else if (slot.type === "equip") {
        description.innerHTML = "<span class='Value'>" + slot.item.specification + "</span>";
        table.appendChild(createRow("Use power:", slot.item.power));
        table.appendChild(createRow("HP:", slot.hp));
    } else if (slot.type === "body") {
        if (slot.item.mother_ship) {
            description.innerHTML = "<span class='Value'>Корпус для материнской машины</span>";
        } else {
            description.innerHTML = "<span class='Value'>Корпус для десантного дрона</span>";
        }
        table.appendChild(createRow("Power:", slot.item.max_power));
        table.appendChild(createRow("HP:", slot.hp));
    } else if (slot.type === "ammo"){
        description.innerHTML = "<span class='Value'>Боеприпасы для оружия</span>";
        table.appendChild(createRow("Type:", slot.item.type));
        table.appendChild(createRow("Damage:", slot.item.min_damage + "-" + slot.item.max_damage));
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

    tip.appendChild(description);
    tip.appendChild(table);
    return table;
}