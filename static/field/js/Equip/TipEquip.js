function TipEquipOn(equip) {
    let tip = document.createElement("div");
    tip.id = "TipEquip";

    let table = document.createElement("table");
    let headTR = document.createElement("tr");
    let headTH = document.createElement("th");

    headTH.innerHTML = "<span class='Value'> " + equip.name + " </span>";
    headTR.appendChild(headTH);

    let iconTD = document.createElement("td");
    iconTD.style.backgroundImage = "url(/assets/units/equip/" + equip.name + ".png)";
    iconTD.style.width = "20px";
    iconTD.style.height = "20px";
    iconTD.style.borderRadius = "5px";

    headTR.appendChild(iconTD);
    table.appendChild(headTR);

    let specificationTR = document.createElement("tr");
    let specificationTD = document.createElement("td");
    specificationTD.innerHTML = equip.specification;
    specificationTD.style.backgroundColor = "#4c4c4c";
    specificationTD.style.borderRadius = "5px";
    specificationTD.colSpan = 2;

    specificationTR.appendChild(specificationTD);
    table.appendChild(specificationTR);

    for (let i = 0; i < equip.effects.length; i++) {
        if (equip.effects[i].type !== "unit_always_animate" && equip.effects[i].type !== "animate" &&
            equip.effects[i].type !== "zone_always_animate" && equip.effects[i].type !== "anchor") {

            let effectsTR = ParseEffect(equip.effects[i], equip);
            table.appendChild(effectsTR);
        }
    }

    tip.appendChild(table);
    document.body.appendChild(tip);
}


function TipNotAllowEquip(text) {
    let tip = document.createElement("div");
    tip.id = "TipEquip";
    tip.innerHTML = "<span>" + text + "</span>";
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