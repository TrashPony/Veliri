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

    // todo жажда рефакторинга

    for (var i = 0; i < equip.effects.length; i++) {
        var effectsTR = document.createElement("tr");

        var effectsTD = document.createElement("td");
        effectsTD.colSpan = 2;
        effectsTD.style.fontSize = "8pt";
        effectsTD.style.backgroundColor = "#4c4c4c";
        effectsTD.style.borderRadius = "5px";

        var type;
        var quantity;
        var time;

        if (equip.effects[i].steps_time > 4) {
            time = equip.effects[i].steps_time + " ходов"
        } else {
            time = equip.effects[i].steps_time + " хода"
        }

        if (equip.effects[i].percentages) {
            quantity = equip.effects[i].quantity + "%";
        } else {
            quantity = equip.effects[i].quantity;
        }

        if (equip.effects[i].type === "enhances") {
            type = "+"
        }

        if (equip.effects[i].type === "replenishes") {
            type = "++"
        }

        effectsTD.innerHTML = "<span class='Value'>" + type + quantity + " " + equip.effects[i].parameter + "</span>" +
            "<br> на <span class='Value'>" + time + "</span>";

        effectsTR.appendChild(effectsTD);
        table.appendChild(effectsTR);
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
    document.getElementById("TipEquip").remove();
}