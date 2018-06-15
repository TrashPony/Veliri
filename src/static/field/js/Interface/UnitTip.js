function unitTip(unit) {

    var unitTip = document.getElementById("unitTip").style;

    if (unit) {
        document.getElementById("tipUnit").innerHTML = "<spen class='Value'> Юнит </spen>";
        document.getElementById("tipOwned").innerHTML = "<spen class='Value'>" + unit.owner + "</spen>";
        document.getElementById("tipHP").innerHTML = "<spen class='Value'>" + unit.hp + "/" + unit.max_hp+ "</spen>";
        document.getElementById("tipDamage").innerHTML = "<spen class='Value'>" + unit.damage + "</spen>";
        document.getElementById("tipMove").innerHTML = "<spen class='Value'>" + unit.move_speed + "</spen>";
        document.getElementById("tipInit").innerHTML = "<spen class='Value'>" + unit.initiative + "</spen>";
        document.getElementById("tipRangeAttack").innerHTML = "<spen class='Value'>" + unit.range_attack + "</spen>";
        document.getElementById("tipRangeView").innerHTML = "<spen class='Value'>" + unit.range_view + "</spen>";
        document.getElementById("tipArea").innerHTML = "<spen class='Value'>" + unit.area_attack + "</spen>";
        document.getElementById("tipTypeAttack").innerHTML = "<spen class='Value'>" + unit.type_attack + "</spen>";
        unitTip.display = "block"; // Показываем слой
    }
}