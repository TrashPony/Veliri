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
    DetailBoxCreate(unitConstructor.chassis, chassisMenu, "Detail chassis", ChassisMouseOver, "chassisElement", "picChassis", "picDetail chassis");
    var weaponMenu = document.getElementById("weaponMenu");
    DetailBoxCreate(unitConstructor.weapons, weaponMenu, "Detail weapon", WeaponMouseOver, "weaponElement", "picWeapon", "picDetail weapon");
    var towerMenu = document.getElementById("towerMenu");
    DetailBoxCreate(unitConstructor.towers, towerMenu, "Detail towers", TowerMouseOver, "towerElement", "picTower", "picDetail tower");
    var bodyMenu = document.getElementById("bodyMenu");
    DetailBoxCreate(unitConstructor.bodies, bodyMenu, "Detail bodies", BodyMouseOver, "bodyElement", "picBody", "picDetail body");
    var radarMenu = document.getElementById("radarMenu");
    DetailBoxCreate(unitConstructor.radars, radarMenu, "Detail radars", RadarMouseOver, "radarElement", "picRadar", "picDetail radar");
}

function DetailBoxCreate(details, menu, className, onMouse, unitElement, pic, picDetail) {

    for (var j = 0; j < details.length; j++) {

        var box = document.createElement("div");
        box.className = className;
        box.style.backgroundImage = "url(/lobby/img/" + details[j].name + ".png)";
        box.detail = details[j];

        box.onmouseover = function () {
            onMouse(this.detail);
        };

        box.onmouseout = function () {
            TipOff();
        };

        box.onclick = function () {
            SelectDetail(this.detail, unitElement, pic, picDetail, onMouse);
        };

        menu.appendChild(box);
    }
}