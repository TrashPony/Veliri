function UpdateSquad(jsonMessage) {
    var squad = JSON.parse(jsonMessage).squad;

    var squadSelect = document.getElementById(squad.id + ":squad");

    squadSelect.value = squad.name;
    squadSelect.text = squad.name;
    squadSelect.id = squad.id + ":squad";
    squadSelect.matherShip = squad.mather_ship;
    squadSelect.units = squad.units;

    squadSelect.onclick = function () {
        SelectSquad(this)
    };


}