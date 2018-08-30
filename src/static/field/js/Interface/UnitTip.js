function unitTip(unit) {
    console.log(unit.body.weapons[1]);
    let unitTip = document.getElementById("unitTip").style;

    if (unit) {
        document.getElementById("tipUnit").innerHTML = "<spen class='Value'> Юнит </spen>";
        document.getElementById("tipOwned").innerHTML = "<spen class='Value'>" + unit.owner + "</spen>";
        document.getElementById("tipHP").innerHTML = "<spen class='Value'>" + unit.hp + "/" + unit.body.max_hp+ "</spen>";
        document.getElementById("tipDamage").innerHTML = "<spen class='Value'>" + unit.body.weapons[1].ammo.damage + "</spen>";
        document.getElementById("tipMove").innerHTML = "<spen class='Value'>" + unit.body.speed + "</spen>";
        document.getElementById("tipInit").innerHTML = "<spen class='Value'>" + unit.body.initiative + "</spen>";
        document.getElementById("tipRangeAttack").innerHTML = "<spen class='Value'>" + unit.body.weapons[1].weapon.range + "</spen>";
        document.getElementById("tipRangeView").innerHTML = "<spen class='Value'>" + unit.body.range_view + "</spen>";
        document.getElementById("tipArea").innerHTML = "<spen class='Value'>" + unit.body.weapons[1].ammo.area_covers + "</spen>";
        document.getElementById("tipTypeAttack").innerHTML = "<spen class='Value'>" + unit.body.weapons[1].ammo.type_attack + "</spen>";
        unitTip.display = "block"; // Показываем слой
    }
}