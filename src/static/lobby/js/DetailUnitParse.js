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
            WeaponMouseOver(this.weapon);
        };
        boxWeapon.onmouseout = function () {
            var tipWeapon = document.getElementById("tipWeapon");
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
            ChassisMouseOver(this.chassis);
        };
        boxChassis.onmouseout = function () {
            var tipChassis = document.getElementById("tipChassis");
            tipChassis.style.display = "none";
        };
        chassisMenu.appendChild(boxChassis);
    }
}

function ChassisMouseOver(chassis) {
    var tipChassis = document.getElementById("tipChassis");
    var tdTypeChassis = document.getElementById("typeChassis");
    var tdHP = document.getElementById("hp");
    var tdMoveSpeed = document.getElementById("moveSpeed");
    var tdInitiative = document.getElementById("initiative");
    var tdMaxWeaponSize = document.getElementById("maxWeaponSize");
    var tdSizeChassis = document.getElementById("sizeChassis");

    tipChassis.style.display = "block";

    tdTypeChassis.innerHTML = chassis.type;
    tdHP.innerHTML = chassis.hp;
    tdMoveSpeed.innerHTML = chassis.move_speed;
    tdInitiative.innerHTML = chassis.initiative;
    tdMaxWeaponSize.innerHTML = chassis.max_weapon_size;
    tdSizeChassis.innerHTML = chassis.size;
}

function WeaponMouseOver(weapon) {
    var tipWeapon = document.getElementById("tipWeapon");
    var tdTypeWeapon = document.getElementById("typeWeapon");
    var tdDamage = document.getElementById("damage");
    var tdRangeAttack = document.getElementById("rangeAttack");
    var tdRangeView = document.getElementById("rangeView");
    var tdAreaAttack = document.getElementById("areaAttack");
    var tdTypeAttack = document.getElementById("typeAttack");
    var tdSizeWeapon = document.getElementById("sizeWeapon");

    tipWeapon.style.display = "block";

    tdTypeWeapon.innerHTML = weapon.type;
    tdDamage.innerHTML = weapon.damage;
    tdRangeAttack.innerHTML = weapon.range_attack;
    tdRangeView.innerHTML = weapon.range_view;
    tdAreaAttack.innerHTML = weapon.area_attack;
    tdTypeAttack.innerHTML = weapon.type_attack;
    tdSizeWeapon.innerHTML = weapon.size;
}