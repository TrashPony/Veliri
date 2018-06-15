function TipEquipOn(equip) {
    var tip = document.createElement("div");
    tip.id = "TipEquip";

    var table = document.createElement("table");
    var headTR = document.createElement("tr");
    var headTH = document.createElement("th");

    headTH.innerHTML = "<span class='Value'> " + equip.type + " </span>";
    headTR.appendChild(headTH);

    var iconTD = document.createElement("td");
    iconTD.style.backgroundImage = "url(/assets/" + equip.type + ".png)";
    iconTD.style.width = "20px";
    iconTD.style.height = "20px";
    iconTD.style.borderRadius = "5px";

    headTR.appendChild(iconTD);
    table.appendChild(headTR);

    var specificationTR = document.createElement("tr");
    var specificationTD = document.createElement("td");
    specificationTD.innerHTML = equip.specification;
    specificationTD.style.backgroundColor = "#4c4c4c";
    specificationTD.style.borderRadius = "5px";
    specificationTD.colSpan = 2;

    specificationTR.appendChild(specificationTD);
    table.appendChild(specificationTR);

    for (var i = 0; i < equip.effects.length; i++) {
        if (equip.effects[i].type !== "unit_always_animate" && equip.effects[i].type !== "animate" &&
            equip.effects[i].type !== "zone_always_animate" && equip.effects[i].type !== "anchor") {

            var effectsTR = ParseEffect(equip.effects[i], equip);
            table.appendChild(effectsTR);
            
        }
    }

    tip.appendChild(table);
    document.body.appendChild(tip);
}

function updatePositionTipEquip() {
    if (document.getElementById("TipEquip")) {
        document.getElementById("TipEquip").style.top = stylePositionParams.top + 'px';
        document.getElementById("TipEquip").style.left = stylePositionParams.left - 5 + 'px';
    }
}

function TipEquipOff() {
    if (document.getElementById("TipEquip")) {
        document.getElementById("TipEquip").remove();
    }
}