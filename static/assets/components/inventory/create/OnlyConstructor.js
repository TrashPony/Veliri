function OnlyConstructor(closeFunc) {
    let inventoryBox = document.createElement("div");
    inventoryBox.id = "inventoryBox";
    document.body.appendChild(inventoryBox);

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

    let squad = document.createElement("div");
    squad.id = "Squad";
    inventoryBox.appendChild(squad);

    CreateMotherShipParamsMenu();
    CreateConstructorMenu();
    CreateSquadMenu();
    CreateSquadHead();

    let closeButton = document.createElement("input");
    closeButton.id = "inventoryCloseButton";
    closeButton.type = "button";
    closeButton.value = "Закрыть";
    closeButton.onclick = () => {
        InventoryClose();
    };
    motherShipParams.appendChild(closeButton);

    let buttons = CreateControlButtons("1px", "32px", "-3px", "29px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'inventoryBox');
    });

    $(buttons.close).mousedown(function () {
        inventoryBox.remove();
    });
    inventoryBox.appendChild(buttons.move);
    inventoryBox.appendChild(buttons.close);
}