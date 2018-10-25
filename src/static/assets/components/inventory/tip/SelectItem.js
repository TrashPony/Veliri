function InventorySelectTip(slot, x, y, first) {

    if (!slot || !slot.item) {
        return
    }

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

        let detailedButton = document.createElement("input");
        detailedButton.type = "button";
        detailedButton.className = "lobbyButton inventoryTip";
        detailedButton.value = "Подробнее";
        detailedButton.style.pointerEvents = "auto";
        // TODO detailedButton.onclick = функция вывода подробной информации
        tip.appendChild(detailedButton);
    }

    document.body.appendChild(tip);
}

function CreateParamsTable(slot, tip) {
    let table = document.createElement("table");
    table.id = "paramsItemTable";

    let description = document.createElement("div");
    description.id = "description";


    if (slot.item.type_slot === 1) {
        table.appendChild(createRow("Slot type:", "I"));
    } else if (slot.item.type_slot === 2) {
        table.appendChild(createRow("Slot type:", "II"));
    } else if (slot.item.type_slot === 3) {
        table.appendChild(createRow("Slot type:", "III"));
    } else if (slot.item.type_slot === 4) {
        table.appendChild(createRow("Slot type:", "IV"));
    } else if (slot.item.type_slot === 5) {
        table.appendChild(createRow("Slot type:", "V"));
    }

    if (slot.type === "weapon") {
        description.innerHTML = "<span class='Value'>Оружие универсального предназначения</span>";
        table.appendChild(createRow("Accuracy:", slot.item.accuracy));
        table.appendChild(createRow("Ammo capacity:", slot.item.ammo_capacity));
        table.appendChild(createRow("Artillery:", slot.item.artillery));
        table.appendChild(createRow("Range:", slot.item.range));
        table.appendChild(createRow("Min attack range:", slot.item.min_attack_range));
        table.appendChild(createRow("HP:", slot.hp));
        table.appendChild(createRow("Use power:", slot.item.power));
    } else if (slot.type === "equip") {
        description.innerHTML = "<span class='Value'>" + slot.item.specification + "</span>";
        table.appendChild(createRow("Active:", slot.item.active));
        table.appendChild(createRow("Applicable:", slot.item.applicable));
        table.appendChild(createRow("Radius:", slot.item.radius));
        table.appendChild(createRow("Region:", slot.item.region));
        table.appendChild(createRow("Reload:", slot.item.reload));
        table.appendChild(createRow("Active power:", slot.item.use_power));
        table.appendChild(createRow("HP:", slot.hp));
        table.appendChild(createRow("Use power:", slot.item.power));
    } else if (slot.type === "body") {

        if (slot.item.mother_ship) {
            description.innerHTML = "<span class='Value'>Корпус для материнской машины</span>";
        } else {
            description.innerHTML = "<span class='Value'>Корпус для десантного дрона</span>";
        }
        table.appendChild(createRow("initiative:", slot.item.initiative));
        table.appendChild(createRow("Max hp:", slot.item.max_hp));
        table.appendChild(createRow("Max power:", slot.item.max_power));
        table.appendChild(createRow("Range View:", slot.item.range_view));
        table.appendChild(createRow("Recovery power:", slot.item.range_view));
        table.appendChild(createRow("Speed:", slot.item.speed));
        table.appendChild(createRow("Wall hack:", slot.item.wall_hack));
        table.appendChild(createRow("HP:", slot.hp));
        table.appendChild(createRow("Power:", slot.item.max_power));

    } else if (slot.type === "ammo"){
        description.innerHTML = "<span class='Value'>Боеприпасы для оружия</span>";
        table.appendChild(createRow("Area covers:", slot.item.area_covers));
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