function AddListSquad(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squads = JSON.parse(jsonMessage).squads;

    for (var i = 0; i < squads.length; i++) {
        var squadOption = document.createElement("option");

        squadOption.value = squads[i].name;
        squadOption.text = squads[i].name;
        squadOption.id = squads[i].id + ":squad";
        squadOption.matherShip = squads[i].mather_ship;
        squadOption.units = squads[i].units;
        squadOption.equip = squads[i].equip;

        selectSquad.appendChild(squadOption);
        selectSquad.value = "";
    }
}

function AddNewSquadInList(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squad = JSON.parse(jsonMessage).squad;

    var squadOption = document.createElement("option");

    squadOption.value = squad.name;
    squadOption.text = squad.name;
    squadOption.id = squad.id + ":squad";
    squadOption.matherShip = squad.mather_ship;
    squadOption.units = squad.units;
    squadOption.equip = squad.equip;

    selectSquad.appendChild(squadOption);
    selectSquad.value = squad.name;
}