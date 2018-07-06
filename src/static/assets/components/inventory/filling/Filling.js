function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        console.log(jsonData);
        let squad = JSON.parse(jsonData).squad;
        FillingInventoryTable(squad.inventory);

        if (squad.mather_ship.body != null) {
            FillingConstructorTable(squad.mather_ship.body)
        }
    }
}



function FillingSquadTable() {

}

