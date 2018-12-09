function CreateInventoryMenu(closeFunc) {

    if (document.getElementById("mask")) {
        document.getElementById("mask").remove();
    }

    if (document.getElementById("inventoryBox")) {
        document.getElementById("inventoryBox").remove();
    }

    let mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    let inventoryBox = document.createElement("div");
    inventoryBox.id = "inventoryBox";

    let userStatus = document.createElement("div");
    userStatus.id = "userStatus";
    inventoryBox.appendChild(userStatus);

    let motherShipParams = document.createElement("div");
    motherShipParams.id = "MotherShipParams";
    inventoryBox.appendChild(motherShipParams);

    let constructorBackGround = document.createElement("div");
    constructorBackGround.id = "ConstructorBackGround";
    inventoryBox.appendChild(constructorBackGround);

    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventoryBox.appendChild(inventory);

    let storage = document.createElement("div");
    storage.id = "storage";
    inventoryBox.appendChild(storage);

    let squad = document.createElement("div");
    squad.id = "Squad";
    inventoryBox.appendChild(squad);

    document.body.appendChild(inventoryBox);

    CreateMotherShipParamsMenu();
    CreateConstructorMenu();
    CreateInventory();
    CreateSquadMenu();
    CreateUserStatus();

    let closeButton = document.createElement("div");
    closeButton.id = "inventoryCloseButton";
    closeButton.className = "button";
    closeButton.innerHTML = "Закрыть";
    closeButton.onclick = () => {InventoryClose(); closeFunc()};
    motherShipParams.appendChild(closeButton);
}

function InventoryClose() {
    document.getElementById("mask").remove();
    document.getElementById("inventoryBox").remove();
    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) constructorUnit.remove();
    inventorySocket.close();
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