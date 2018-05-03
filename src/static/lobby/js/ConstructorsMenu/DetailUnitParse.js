function DetailUnitParse(jsonMessage) {
    var weapons = JSON.parse(jsonMessage).weapons;
    var chassis = JSON.parse(jsonMessage).chassis;
    var towers = JSON.parse(jsonMessage).towers;
    var bodies = JSON.parse(jsonMessage).bodies;
    var radars = JSON.parse(jsonMessage).radars;

    var unitConstructor = document.getElementById("unitConstructor");

    unitConstructor.weapons = weapons;
    unitConstructor.chassis = chassis;
    unitConstructor.towers = towers;
    unitConstructor.bodies = bodies;
    unitConstructor.radars = radars;

    ViewDetailUnit();
}

function ViewDetailUnit() {
    var unitConstructor = document.getElementById("unitConstructor");

    var chassisMenu = document.getElementById("chassisMenu");
    DetailBoxCreate(unitConstructor.chassis, chassisMenu, "Detail chassis", ChassisMouseOver, "chassisElement");
    var weaponMenu = document.getElementById("weaponMenu");
    DetailBoxCreate(unitConstructor.weapons, weaponMenu, "Detail weapon", WeaponMouseOver, "weaponElement");
    var towerMenu = document.getElementById("towerMenu");
    DetailBoxCreate(unitConstructor.towers, towerMenu, "Detail towers", TowerMouseOver, "towerElement");
    var bodyMenu = document.getElementById("bodyMenu");
    DetailBoxCreate(unitConstructor.bodies, bodyMenu, "Detail bodies", BodyMouseOver, "bodyElement");
    var radarMenu = document.getElementById("radarMenu");
    DetailBoxCreate(unitConstructor.radars, radarMenu, "Detail radars", RadarMouseOver, "radarElement");
}

function DetailBoxCreate(details, menu, className, onMouse, unitElement) {

    for (var j = 0; j < details.length; j++) {

        var box = document.createElement("div");
        box.className = className;
        box.style.backgroundImage = "url(/assets/" + details[j].name + ".png)";
        box.detail = details[j];

        box.onmouseover = function () {
            onMouse(this.detail);
        };

        box.onmouseout = function () {
            TipOff();
        };

        box.onclick = function () {
            SelectDetail(this.detail, unitElement, onMouse);
            SendEventAddOrDelDetail()
        };

        menu.appendChild(box);
    }
}