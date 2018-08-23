function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        let squad = JSON.parse(jsonData).squad;
        InventoryTable(squad.inventory);
        SquadTable(squad);
        if (squad.mather_ship.body != null) {
            ConstructorTable(squad.mather_ship.body);
        } else {
            NoActiveCell();
        }
    }
}