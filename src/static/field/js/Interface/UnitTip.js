function unitTip(unit) {

    var unitTip = document.getElementById("unitTip").style;

    if (unit) {
        document.getElementById("tipUnit").innerHTML = "<spen class='Value'>" + unit.info.type + "</spen>";
        document.getElementById("tipOwned").innerHTML = "<spen class='Value'>" + unit.info.owner + "</spen>";
        document.getElementById("tipHP").innerHTML = "<spen class='Value'>" + unit.info.hp + "</spen>";
        document.getElementById("tipAction").innerHTML = "<spen class='Value'>" + unit.info.action + "</spen>";
        document.getElementById("tipTarget").innerHTML = "<spen class='Value'>" + unit.info.target + "</spen>";
        document.getElementById("tipDamage").innerHTML = "<spen class='Value'>" + unit.info.damage + "</spen>";
        document.getElementById("tipMove").innerHTML = "<spen class='Value'>" + unit.info.move_speed + "</spen>";
        document.getElementById("tipInit").innerHTML = "<spen class='Value'>" + unit.info.initiative + "</spen>";
        document.getElementById("tipRangeAttack").innerHTML = "<spen class='Value'>" + unit.info.range_attack + "</spen>";
        document.getElementById("tipRangeView").innerHTML = "<spen class='Value'>" + unit.info.watch_zone + "</spen>";
        document.getElementById("tipArea").innerHTML = "<spen class='Value'>" + unit.info.area_attack + "</spen>";
        document.getElementById("tipTypeAttack").innerHTML = "<spen class='Value'>" + unit.info.type_attack + "</spen>";
        unitTip.display = "block"; // Показываем слой
    }
}