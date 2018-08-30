function unitTip(unit) {

    let unitTip = document.getElementById("unitTip").style;
    let weapon;
    let ammo;

    for (let i in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon){
            weapon = unit.body.weapons[i].weapon;
            if (unit.body.weapons[i].ammo)  {
                ammo = unit.body.weapons[i].ammo
            }
        }
    }

    if (unit) {
        // todo сделать так что бы не выводились значения который нет, и авто мастабировать блок только под то что есть
        if (unit.body.mother_ship){
            document.getElementById("tipUnit").innerHTML = "<spen class='Value'> MS </spen>";
        } else {
            document.getElementById("tipUnit").innerHTML = "<spen class='Value'> Юнит </spen>";
        }

        if (weapon) {
            document.getElementById("tipRangeAttack").innerHTML = "<spen class='Value'>" + weapon.range + "</spen>";
        }

        if (ammo) {
            document.getElementById("tipDamage").innerHTML = "<spen class='Value'>" + ammo.damage + "</spen>";
            document.getElementById("tipMove").innerHTML = "<spen class='Value'>" + unit.body.speed + "</spen>";
            document.getElementById("tipArea").innerHTML = "<spen class='Value'>" + ammo.area_covers + "</spen>";
            document.getElementById("tipTypeAttack").innerHTML = "<spen class='Value'>" + ammo.type_attack + "</spen>";
        }

        document.getElementById("tipOwned").innerHTML = "<spen class='Value'>" + unit.owner + "</spen>";
        document.getElementById("tipHP").innerHTML = "<spen class='Value'>" + unit.hp + "/" + unit.body.max_hp+ "</spen>";
        document.getElementById("tipInit").innerHTML = "<spen class='Value'>" + unit.body.initiative + "</spen>";
        document.getElementById("tipRangeView").innerHTML = "<spen class='Value'>" + unit.body.range_view + "</spen>";
        unitTip.display = "block"; // Показываем слой
    }
}