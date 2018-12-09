function InventorySelectTip(slot, x, y, first, size, numberSlot, storage) {

    if (!slot || !slot.item) {
        return
    }

    let tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";

    if (size) {
        notificationInventorySize(slot.size);
    }

    if (!first) {
        tip.id = "InventoryTip";

        let detailedButton = createInput("Подробнее", tip);
        detailedButton.onclick = function () {
            tip.remove();

            let tipDetail = document.createElement("div");
            tipDetail.id = "InventoryTip";
            tipDetail.style.top = y + "px";
            tipDetail.style.left = x + "px";

            CreateParamsTable(slot, tipDetail);
            document.body.appendChild(tipDetail);
        };

        if (storage) {
            let sellButton = createInput("Продать", tip);
            sellButton.onclick = function () {
                tip.remove();
                CreateSellDialog(x, y, numberSlot, slot)
            };

            let to = createInput("В инвентарь", tip);
            to.onclick = function () {
                inventorySocket.send(JSON.stringify({
                    event: "itemToInventory",
                    storage_slot: Number(numberSlot)
                }));
            };
        } else {
            let to = createInput("На склад", tip);
            to.onclick = function () {
                inventorySocket.send(JSON.stringify({
                    event: "itemToStorage",
                    inventory_slot: Number(numberSlot)
                }));
            };
        }

        let cancelButton = createInput("Отменить", tip);
        cancelButton.onclick = function () {
            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };
    } else {
        let inventoryTip = document.getElementById("InventoryTip");
        if (!inventoryTip) {
            tip.id = "InventoryTipOver";
        } else {
            tip.remove();
            return
        }

        let name = document.createElement("span");
        name.className = "InventoryTipName";
        name.innerHTML = slot.item.name;
        tip.appendChild(name);
    }

    document.body.appendChild(tip);
}

function CreateSellDialog(x, y, numberSlot, slot) {
    let sellBlock = document.createElement("div");
    sellBlock.id = "sellDialog";
    sellBlock.style.top = y + "px";
    sellBlock.style.left = x + "px";

    sellBlock.innerHTML = "<h2>Быстрая продажа</h2>" +
        "<div><span>Кол-во</span><input id='sellQuantity' type='number' min='0' value='" + slot.quantity + "' max='" + slot.quantity + "'></div><br>" +
        "<div><span>Цена за шт.</span><input id='sellPrice' min='0' type='number'></div><br>";

    let closeButton = createInput("Отменить", sellBlock);
    closeButton.onclick = function () {
        sellBlock.remove();
    };

    let sellButton = createInput("Продать", sellBlock);
    sellButton.onclick = function () {
        let quantity = Number(document.getElementById("sellQuantity").value);
        let price = Number(document.getElementById("sellPrice").value);

        if (quantity > 0 && price > 0) {
            marketSocket.send(JSON.stringify({
                event: 'placeNewSellOrder',
                storage_slot: Number(numberSlot),
                quantity: quantity,
                price: price,
            }));
            sellBlock.remove();
        } else {
            alert("не заполнены поля")
        }
    };

    document.body.appendChild(sellBlock)
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

        createTipWeaponType(table, slot, "weapon");

        table.appendChild(createRow("Accuracy:", slot.item.accuracy));
        table.appendChild(createRow("Ammo capacity:", slot.item.ammo_capacity));
        table.appendChild(createRow("Artillery:", slot.item.artillery));
        table.appendChild(createRow("Range:", slot.item.min_attack_range + "-" + slot.item.range));
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

        createTipWeaponType(table, slot, "body");

        table.appendChild(createRow("initiative:", slot.item.initiative));
        table.appendChild(createRow("Max hp:", slot.item.max_hp));
        table.appendChild(createRow("Max power:", slot.item.max_power));
        table.appendChild(createRow("Range View:", slot.item.range_view));
        table.appendChild(createRow("Recovery power:", slot.item.range_view));
        table.appendChild(createRow("Speed:", slot.item.speed));
        table.appendChild(createRow("Wall hack:", slot.item.wall_hack));
        table.appendChild(createRow("HP:", slot.hp));
        table.appendChild(createRow("Power:", slot.item.max_power));

    } else if (slot.type === "ammo") {
        description.innerHTML = "<span class='Value'>Боеприпасы для оружия</span>";

        createTipWeaponType(table, slot, "weapon");

        table.appendChild(createRow("Area covers:", slot.item.area_covers));
        table.appendChild(createRow("Type:", slot.item.type));
        table.appendChild(createRow("Damage:", slot.item.min_damage + "-" + slot.item.max_damage));
    }

    table.appendChild(createRow("Size:", slot.size));

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

function createTipWeaponType(table, slot, type) {
    if (type === "body") {
        let tr = document.createElement("tr");
        let td1 = document.createElement("td");
        td1.innerHTML = "<span class='Value'>Типы оружия: </span>";
        td1.style.border = "0px";
        let td2 = document.createElement("td");
        let weaponType = document.createElement("div");
        weaponType.className = "weaponTipTypePanel";
        weaponTypePanel(slot.item, weaponType, "weaponUnitTypeAllow", "weaponUnitTypeNotAllow");

        td2.appendChild(weaponType);
        tr.appendChild(td1);
        tr.appendChild(td2);
        table.appendChild(tr);
    } else if (type === "weapon") {
        let tr = document.createElement("tr");
        let td1 = document.createElement("td");
        td1.innerHTML = "<span class='Value'>Тип оружия: </span>";
        td1.style.border = "0px";
        let td2 = document.createElement("td");
        let weaponType = document.createElement("div");
        weaponType.className = "weaponTipTypePanel";

        let fakeBody = {};
        if (slot.item.standard_size === 1) {
            fakeBody.standard_size_small = true;
            fakeBody.standard_size_medium = false;
            fakeBody.standard_size_big = false;
        }

        if (slot.item.standard_size === 2) {
            fakeBody.standard_size_small = false;
            fakeBody.standard_size_medium = true;
            fakeBody.standard_size_big = false;
        }

        if (slot.item.standard_size === 3) {
            fakeBody.standard_size_small = false;
            fakeBody.standard_size_medium = false;
            fakeBody.standard_size_big = true;
        }

        weaponTypePanel(fakeBody, weaponType, "weaponUnitTypeAllow", "weaponUnitTypeNotAllow");

        td2.appendChild(weaponType);
        tr.appendChild(td1);
        tr.appendChild(td2);
        table.appendChild(tr);
    }
}

function notificationInventorySize(size) {
    let realSize = document.getElementById("sizeInventoryInfo");

    let unitIcon = document.getElementById("MSIcon");
    if (unitIcon.slotData) {

        if (document.getElementById("itemSize")) {
            document.getElementById("itemSize").remove();
        }

        let slotData = JSON.parse(unitIcon.slotData);
        let percentFill = 100 / (slotData.body.capacity_size / size);

        let itemSize = document.createElement("div");
        itemSize.id = "itemSize";
        itemSize.style.width = percentFill + "%";

        realSize.appendChild(itemSize);
    }
}