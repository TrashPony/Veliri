function CreateInventoryMenu() {
    var mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    var inventoryBox = document.createElement("div");
    inventoryBox.id = "inventoryBox";

    var motherShipParams = document.createElement("div");
    motherShipParams.id = "MotherShipParams";
    inventoryBox.appendChild(motherShipParams);

    var constructorBackGround = document.createElement("div");
    constructorBackGround.id = "ConstructorBackGround";
    inventoryBox.appendChild(constructorBackGround);

    var inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventoryBox.appendChild(inventory);

    var squad = document.createElement("div");
    squad.id = "Squad";
    inventoryBox.appendChild(squad);

    document.body.appendChild(inventoryBox);

    CreateMotherShipParamsMenu();
    CreateConstructorMenu();
    CreateInventory();
    CreateSquadMenu();
}

function CreateMotherShipParamsMenu() {
    var menu = document.getElementById("MotherShipParams");

    var spanAttack = document.createElement("span");
    spanAttack.className = "Value params";
    spanAttack.innerHTML = " ⇓ Атака ";

    var spanDef = document.createElement("span");
    spanDef.className = "Value params";
    spanDef.innerHTML = " ⇓ Защита ";

    var spanNav = document.createElement("span");
    spanNav.className = "Value params";
    spanNav.innerHTML = " ⇓ Навигация ";

    menu.appendChild(spanAttack);
    menu.appendChild(spanDef);
    menu.appendChild(spanNav);
}

function CreateConstructorMenu() {
    var constructorBackGround = document.getElementById("ConstructorBackGround");

    var powerPanel = document.createElement("div");
    powerPanel.id = "powerPanel";

    var spanPower = document.createElement("span");
    spanPower.className = "Value";
    spanPower.innerHTML = "Энергия: Max/Use";
    powerPanel.appendChild(spanPower);
    constructorBackGround.appendChild(powerPanel);

    var constructor = document.createElement("div");
    constructor.id = "Constructor";
    constructorBackGround.appendChild(constructor);

    /* 3 type slots */
    var equippingPanelIII = document.createElement("div");
    CreateCells(3, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelIII);
    constructor.appendChild(equippingPanelIII);

    /* 5 type slots */
    var equippingPanelV = document.createElement("div");
    equippingPanelV.className = "verticalEquipPanel";
    CreateCells(5, 2, "inventoryEquipping noActive", "inventoryEquip", equippingPanelV, true);
    constructor.appendChild(equippingPanelV);

    /* shipIcon */
    var unitIcon = document.createElement("div");
    unitIcon.id = "UnitIcon";
    constructor.appendChild(unitIcon);

    /* 2 type slots */
    var equippingPanelII = document.createElement("div");
    equippingPanelII.className = "verticalEquipPanel";
    CreateCells(2, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelII, true);
    constructor.appendChild(equippingPanelII);

    /* 1 type slots */
    var equippingPanelI = document.createElement("div");
    CreateCells(1, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelI);
    constructor.appendChild(equippingPanelI)

}

function CreateInventory() {
    var inventory = document.getElementById("Inventory");

    var spanInventory = document.createElement("span");
    spanInventory.className = "Value";
    spanInventory.innerHTML = "Инвентарь: ";
    inventory.appendChild(spanInventory);

    var inventoryStorage = document.createElement("div");
    CreateCells(6, 40, "InventoryCell", "inventory ", inventoryStorage);

    inventory.appendChild(inventoryStorage);
}

function CreateSquadMenu() {
    var squad = document.getElementById("Squad");

    var squadStorage = document.createElement("div");
    CreateCells(4, 6, "inventoryUnit noActive", "squad ", squadStorage);

    squad.appendChild(squadStorage);
}


function CreateCells(typeSlot, count, className, idPrefix, parent, vertical) {
    for (var i = 0; i < count; i++) {
        var cell = document.createElement("div");
        cell.className = className;
        cell.id = idPrefix + Number(i + 1) + typeSlot;

        cell.type = typeSlot;
        cell.Number = Number(i + 1);

        cell.onclick = function () {
            console.log("NumberSlot " + this.Number + " typeSlot " + this.type)
        };

        parent.appendChild(cell);

        if (vertical) {
            var br = document.createElement("br");
            parent.appendChild(br);
        }
    }
}