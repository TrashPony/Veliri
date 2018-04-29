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
    DetailBoxCreate(unitConstructor.chassis, chassisMenu, "Detail chassis", "tipChassis", ChassisMouseOver, SelectChassis);
    var weaponMenu = document.getElementById("weaponMenu");
    DetailBoxCreate(unitConstructor.weapons, weaponMenu, "Detail weapon", "tipWeapon", WeaponMouseOver, SelectWeapon);
    var towerMenu = document.getElementById("towerMenu");
    DetailBoxCreate(unitConstructor.towers, towerMenu, "Detail towers", "tipTower", TowerMouseOver, SelectTower);
    var bodyMenu = document.getElementById("bodyMenu");
    DetailBoxCreate(unitConstructor.bodies, bodyMenu, "Detail bodies", "tipBody", BodyMouseOver, SelectBody);
    var radarMenu = document.getElementById("radarMenu");
    DetailBoxCreate(unitConstructor.radars, radarMenu, "Detail radars", "tipRadar", RadarMouseOver, SelectRadar);
}

function DetailBoxCreate(details, menu, className, tip, onMouse, Select) {
    console.log(details);
    for (var j = 0; j < details.length; j++) {

        var box = document.createElement("div");
        box.className = className;
        box.style.backgroundImage = "url(/lobby/img/" + details[j].name + ".png)";
        box.chassis = details[j];

        box.onmouseover = function () {
            onMouse(this.chassis);
        };
        box.onmouseout = function () {
            var tipChassis = document.getElementById(tip);
            tipChassis.style.display = "none";
        };

        box.onclick = function () {
            Select(this);
        };

        menu.appendChild(box);
    }
}