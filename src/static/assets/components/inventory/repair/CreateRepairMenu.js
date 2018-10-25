function CreateRepairMenu() {
    let repairMenu = document.createElement("div");
    repairMenu.className = "RepairMenu";

    let equipButton = document.createElement("div");
    equipButton.innerHTML = "MS";
    equipButton.onclick = EquipsRepair;
    repairMenu.appendChild(equipButton);

    let allButton = document.createElement("div");
    allButton.innerHTML = "ВСЕ";
    allButton.onclick = AllRepair;
    repairMenu.appendChild(allButton);

    let inventoryButton = document.createElement("div");
    inventoryButton.innerHTML = "ТРЮМ";
    inventoryButton.onclick = InventoryRepair;
    repairMenu.appendChild(inventoryButton);

    document.getElementById("ConstructorBackGround").appendChild(repairMenu);

    this.className = "repairButtonActive";

    this.onclick = function () {
        this.className = "repairButton";
        repairMenu.remove();
        this.onclick = CreateRepairMenu;
    }
}

function EquipsRepair() {

    // todo подсвечивать вещи которые будут починены

    inventorySocket.send(JSON.stringify({
        event: "EquipsRepair"
    }));
}

function AllRepair() {

    // todo подсвечивать вещи которые будут починены

    inventorySocket.send(JSON.stringify({
        event: "AllRepair"
    }));
}

function InventoryRepair() {

    // todo подсвечивать вещи которые будут починены

    inventorySocket.send(JSON.stringify({
        event: "InventoryRepair"
    }));
}