function CreateSquadMenu() {
    let squad = document.getElementById("Squad");

    let spanInventory = document.createElement("span");
    spanInventory.className = "SquadHead";
    spanInventory.innerHTML = "ОТСЕКИ ДЛЯ ЮНИТОВ";
    spanInventory.style.width = "200px";
    squad.appendChild(spanInventory);

    let squadStorage = document.createElement("div");
    CreateCells(4, 6, "inventoryUnit noActive", "squad ", squadStorage);

    squad.appendChild(squadStorage);
}
