function InitCreateUnit() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var unitConstructor = document.getElementById("unitConstructor");

    if (!unitConstructor) {
        CreateUnitConstructor();
    } else {
        unitConstructor.style.display = "block";
    }
}

function BackToLobby() {
    var mask = document.getElementById("mask");
    mask.style.display = "none";

    var unitConstructor = document.getElementById("unitConstructor");
    unitConstructor.style.display = "none";
}
