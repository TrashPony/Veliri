function DetailUnitParse(jsonMessage) {
    var weapons = JSON.parse(jsonMessage).weapons;
    var chassis = JSON.parse(jsonMessage).chassis;

    var unitConstructor = document.getElementById("unitConstructor");
    unitConstructor.weapons = weapons;
    unitConstructor.chassis = chassis;

    ViewDetailUnit();
}

function ViewDetailUnit() {
    var unitConstructor = document.getElementById("unitConstructor");

    var chassisMenu = document.getElementById("chassisMenu");
    var weaponMenu = document.getElementById("weaponMenu");

    for (var i = 0; i < unitConstructor.weapons.length; i++) {
        var boxWeapon = document.createElement("div");
        boxWeapon.className = "Detail weapon";
        boxWeapon.style.backgroundImage = "url(/lobby/img/" + unitConstructor.weapons[i].type + ".png)";
        boxWeapon.weapon = unitConstructor.weapons[i];
        boxWeapon.onmouseover = function () {
            var tipWeapon = document.getElementById("tipWeapon").style;
            tipWeapon.style.display = "block";
        };
        boxWeapon.onmouseout = function () {
            var tipWeapon = document.getElementById("tipWeapon").style;
            tipWeapon.style.display = "none";
        };
        weaponMenu.appendChild(boxWeapon);
    }

    for (var j = 0; j < unitConstructor.chassis.length; j++) {
        var boxChassis = document.createElement("div");
        boxChassis.className = "Detail weapon";
        boxChassis.style.backgroundImage = "url(/lobby/img/" + unitConstructor.chassis[j].type + ".png)";
        boxChassis.chassis = unitConstructor.chassis[j];
        boxChassis.onmouseover = function () {
            var tipChassis = document.getElementById("tipChassis").style;
            tipChassis.style.display = "block";
        };
        boxChassis.onmouseout = function () {
            var tipChassis = document.getElementById("tipChassis").style;
            tipChassis.style.display = "none";
        };
        chassisMenu.appendChild(boxChassis);
    }
}