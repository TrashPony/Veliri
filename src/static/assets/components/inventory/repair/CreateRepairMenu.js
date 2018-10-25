function CreateRepairMenu() {
    let repairMenu = document.createElement("div");
    repairMenu.className = "RepairMenu";

    let equipButton = document.createElement("div");
    equipButton.innerHTML = "MS";
    repairMenu.appendChild(equipButton);

    let allButton = document.createElement("div");
    allButton.innerHTML = "ВСЕ";
    repairMenu.appendChild(allButton);

    let inventoryButton = document.createElement("div");
    inventoryButton.innerHTML = "ТРЮМ";
    repairMenu.appendChild(inventoryButton);


    document.getElementById("ConstructorBackGround").appendChild(repairMenu);

    this.onclick = function () {
        repairMenu.remove();
        this.onclick = CreateRepairMenu;
    }
}