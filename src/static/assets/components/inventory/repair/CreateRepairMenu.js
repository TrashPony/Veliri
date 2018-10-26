function CreateRepairMenu() {
    let repairMenu = document.createElement("div");
    repairMenu.className = "RepairMenu";
    repairMenu.id = "repairMenu";

    let equipButton = document.createElement("div");
    equipButton.innerHTML = "MS";
    equipButton.onclick = EquipsRepair;
    equipButton.onmouseover = overEquipButton;
    equipButton.onmouseout = outEquipButton;
    repairMenu.appendChild(equipButton);

    let allButton = document.createElement("div");
    allButton.innerHTML = "ВСЕ";
    allButton.onclick = AllRepair;
    allButton.onmouseover = function () {
        overEquipButton();
        overInventoryButton();
    };
    allButton.onmouseout = function () {
        outEquipButton();
        outInventoryButton();
    };
    repairMenu.appendChild(allButton);

    let inventoryButton = document.createElement("div");
    inventoryButton.innerHTML = "ТРЮМ";
    inventoryButton.onclick = InventoryRepair;
    inventoryButton.onmouseover = overInventoryButton;
    inventoryButton.onmouseout = outInventoryButton;
    repairMenu.appendChild(inventoryButton);

    document.getElementById("ConstructorBackGround").appendChild(repairMenu);

    this.className = "repairButtonActive";

    this.onclick = function () {
        this.className = "repairButton";
        repairMenu.remove();
        this.onclick = CreateRepairMenu;
    }
}

function AllRepair() {
    inventorySocket.send(JSON.stringify({
        event: "AllRepair"
    }));
}