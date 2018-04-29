function UpdateSquad(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");

    var squad = JSON.parse(jsonMessage).squad;
    var squadOption = document.getElementById(squad.id + ":squad");

    squadOption.value = squad.name;
    squadOption.text = squad.name;
    squadOption.id = squad.id + ":squad";
    squadOption.matherShip = squad.mather_ship;
    squadOption.units = squad.units;
    squadOption.equip = squad.equip;

    SelectSquad(selectSquad)
}