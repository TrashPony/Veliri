function CreateSquadMenu() {
    let squad = document.getElementById("Squad");

    let squadStorage = document.createElement("div");
    CreateCells(4, 6, "inventoryUnit noActive", "squad ", squadStorage);

    squad.appendChild(squadStorage);
}
