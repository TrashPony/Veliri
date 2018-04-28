function AddListSquad(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squads = JSON.parse(jsonMessage).squads;

    for (var i = 0; i < squads.length; i++) {
        var squadSelect = document.createElement("option");

        squadSelect.value = squads[i].name;
        squadSelect.text = squads[i].name;
        squadSelect.id = squads[i].id + ":squad";
        squadSelect.matherShip = squads[i].mather_ship;
        squadSelect.units = squads[i].units;

        squadSelect.onclick = function () {
            SelectSquad(this)
        };

        selectSquad.appendChild(squadSelect);
        selectSquad.value = "";
    }
}

function AddNewSquadInList(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squad = JSON.parse(jsonMessage).squad;

    var squadSelect = document.createElement("option");

    squadSelect.value = squad.name;
    squadSelect.text = squad.name;
    squadSelect.id = squad.id + ":squad";
    squadSelect.matherShip = squad.mather_ship;
    squadSelect.units = squad.units;

    squadSelect.onclick = function () {
        SelectSquad(this)
    };

    selectSquad.appendChild(squadSelect);
    selectSquad.value = squad.name;
}