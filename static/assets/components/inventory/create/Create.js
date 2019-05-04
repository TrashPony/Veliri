function CreateInventoryMenu(closeFunc, option) {

    if (document.getElementById("storage")) {
        document.getElementById("storage").remove();
    }

    if (document.getElementById("mask")) {
        document.getElementById("mask").remove();
    }

    if (document.getElementById("inventoryBox")) {
        document.getElementById("inventoryBox").remove();
    }

    if (document.getElementById("Inventory")) {
        document.getElementById("Inventory").remove();
    }

    if (document.getElementById('wrapperInventoryAndStorage'))
        document.getElementById('wrapperInventoryAndStorage').remove();

    if (option === 'inventory') {
        let inventory = document.createElement("div");
        inventory.id = "Inventory";
        inventory.style.position = "absolute";
        inventory.style.bottom = "70px";
        inventory.style.right = "15px";
        document.body.appendChild(inventory);
        CreateInventory();
        return
    }

    if (option === 'storage') {
        OnlyStorage();
        return
    }

    let mask = document.createElement("div");
    mask.id = "mask";
    mask.style.display = "block";
    document.body.appendChild(mask);

    let inventoryBox = document.createElement("div");
    inventoryBox.id = "inventoryBox";

    let colorPicker = document.createElement('div');
    colorPicker.className = 'colorpicker';
    colorPicker.style.display = 'none';
    colorPicker.innerHTML = `
        <input id="brightnessColorPicker" type="range" value="100" max="100" min="1">
        <canvas id="colorUnitPicker" width="140" height="140"></canvas>
`;
    inventoryBox.appendChild(colorPicker);

    let userStatus = document.createElement("div");
    userStatus.id = "SquadHead";
    inventoryBox.appendChild(userStatus);

    let motherShipParams = document.createElement("div");
    motherShipParams.id = "MotherShipParams";
    inventoryBox.appendChild(motherShipParams);

    let headSquadList = document.createElement("span");
    headSquadList.className = "InventoryHead";
    headSquadList.innerText = "АНГАР";
    motherShipParams.appendChild(headSquadList);

    let squadsList = document.createElement("div");
    squadsList.id = "SquadsList";
    motherShipParams.appendChild(squadsList);

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
    CreateSquadHead();

    let closeButton = document.createElement("input");
    closeButton.id = "inventoryCloseButton";
    closeButton.type = "button";
    closeButton.value = "Закрыть";
    closeButton.onclick = () => {
        InventoryClose();
        if (closeFunc) {
            closeFunc()
        }
    };
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