function InitCreateUnit(box) {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var unitConstructor = document.getElementById("unitConstructor");

    if (!unitConstructor) {
        unitConstructor = CreateUnitConstructor();
    } else {
        unitConstructor.style.display = "block";
    }

    var slotParse = box.id.split(':'); // "slot:unitSlot"

    unitConstructor.unitSlot = slotParse[0];
    unitConstructor.unit = box.unit;

    UnitConfig(unitConstructor.unit)
}

function BackToLobby() {

    var mask = document.getElementById("mask");
    mask.style.display = "none";

    var unitConstructor = document.getElementById("unitConstructor");
    unitConstructor.remove();
}
