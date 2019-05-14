function CreateRepairMenu() {
    let repairMenu = document.createElement("div");
    repairMenu.className = "RepairMenu";
    repairMenu.id = "repairMenu";

    let equipButton = document.createElement("div");
    equipButton.innerHTML = "Price";
    equipButton.onclick = EquipsRepair;
    equipButton.onmouseover = overEquipButton;
    equipButton.onmouseout = outEquipButton;
    repairMenu.appendChild(equipButton);

    document.getElementById("ConstructorBackGround").appendChild(repairMenu);

    this.className = "repairButtonActive";

    this.onclick = function () {
        this.className = "repairButton";
        repairMenu.remove();
        this.onclick = CreateRepairMenu;
    }
}

function EquipsRepair() {
    inventorySocket.send(JSON.stringify({
        event: "EquipsRepair"
    }));
}