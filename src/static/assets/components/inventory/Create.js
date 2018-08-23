function CreateInventoryMenu(closeFunc) {
    let mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    let inventoryBox = document.createElement("div");
    inventoryBox.id = "inventoryBox";

    let motherShipParams = document.createElement("div");
    motherShipParams.id = "MotherShipParams";
    inventoryBox.appendChild(motherShipParams);

    let constructorBackGround = document.createElement("div");
    constructorBackGround.id = "ConstructorBackGround";
    inventoryBox.appendChild(constructorBackGround);

    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventoryBox.appendChild(inventory);

    let squad = document.createElement("div");
    squad.id = "Squad";
    inventoryBox.appendChild(squad);

    document.body.appendChild(inventoryBox);

    CreateMotherShipParamsMenu();
    CreateConstructorMenu();
    CreateInventory();
    CreateSquadMenu();

    let closeButton = document.createElement("div");
    closeButton.id = "inventoryCloseButton";
    closeButton.className = "button";
    closeButton.innerHTML = "Закрыть";
    closeButton.onclick = () => {InventoryClose(); closeFunc()};
    inventory.appendChild(closeButton);
}

function InventoryClose() {
    document.getElementById("mask").remove();
    document.getElementById("inventoryBox").remove();
    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) constructorUnit.remove();
    inventorySocket.close();
}

function CreateMotherShipParamsMenu() {
    let menu = document.getElementById("MotherShipParams");

    let spanAttack = document.createElement("span");
    spanAttack.className = "Value params";
    spanAttack.innerHTML = " ⇓ Атака ";

    let spanDef = document.createElement("span");
    spanDef.className = "Value params";
    spanDef.innerHTML = " ⇓ Защита ";

    let spanNav = document.createElement("span");
    spanNav.className = "Value params";
    spanNav.innerHTML = " ⇓ Навигация ";

    menu.appendChild(spanAttack);
    menu.appendChild(spanDef);
    menu.appendChild(spanNav);
}

function CreateConstructorMenu() {
    let constructorBackGround = document.getElementById("ConstructorBackGround");

    let powerPanel = document.createElement("div");
    powerPanel.id = "powerPanel";

    let spanPower = document.createElement("span");
    spanPower.className = "Value";
    spanPower.innerHTML = "Энергия: Max/Use";
    powerPanel.appendChild(spanPower);
    constructorBackGround.appendChild(powerPanel);

    let constructorMS = document.createElement("div");
    constructorMS.id = "ConstructorMS";
    constructorBackGround.appendChild(constructorMS);

    /* 3 type slots */
    let equippingPanelIII = document.createElement("div");
    CreateCells(3, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelIII);
    constructorMS.appendChild(equippingPanelIII);

    /* 5 type slots */
    let equippingPanelV = document.createElement("div");
    equippingPanelV.className = "verticalEquipPanel";
    CreateCells(5, 2, "inventoryEquipping noActive", "inventoryEquip", equippingPanelV, true);
    constructorMS.appendChild(equippingPanelV);

    /* shipIcon */
    let unitIcon = document.createElement("div");
    unitIcon.id = "MSIcon";
    constructorMS.appendChild(unitIcon);

    /* 2 type slots */
    let equippingPanelII = document.createElement("div");
    equippingPanelII.className = "verticalEquipPanel";
    CreateCells(2, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelII, true);
    constructorMS.appendChild(equippingPanelII);

    /* 1 type slots */
    let equippingPanelI = document.createElement("div");
    CreateCells(1, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelI);
    constructorMS.appendChild(equippingPanelI)
}

function CreateInventory() {
    let inventory = document.getElementById("Inventory");

    let spanInventory = document.createElement("span");
    spanInventory.className = "Value";
    spanInventory.innerHTML = "Инвентарь: ";
    inventory.appendChild(spanInventory);

    let inventoryStorage = document.createElement("div");
    CreateCells(6, 40, "InventoryCell", "inventory ", inventoryStorage);

    inventory.appendChild(inventoryStorage);
}

function CreateSquadMenu() {
    let squad = document.getElementById("Squad");

    let squadStorage = document.createElement("div");
    CreateCells(4, 6, "inventoryUnit noActive", "squad ", squadStorage);

    squad.appendChild(squadStorage);
}


function CreateCells(typeSlot, count, className, idPrefix, parent, vertical) {
    for (let i = 0; i < count; i++) {
        let cell = document.createElement("div");
        cell.className = className;
        cell.id = idPrefix + Number(i + 1) + typeSlot;

        cell.type = typeSlot;
        cell.Number = Number(i + 1);

        parent.appendChild(cell);

        if (vertical) {
            let br = document.createElement("br");
            parent.appendChild(br);
        }
    }
}