function TipEffectOn(effect) {
    let tip = document.createElement("div");
    tip.id = "TipEffect";

    let table = document.createElement("table");
    let headTR = document.createElement("tr");
    let headTH = document.createElement("th");

    headTH.innerHTML = "<span class='Value'> " + effect.name + " " + effect.level + " </span>";
    headTR.appendChild(headTH);

    let iconTD = document.createElement("td");
    iconTD.style.backgroundImage = "url(/assets/effects/" + effect.name + "_" + effect.level + ".png)";
    iconTD.style.width = "20px";
    iconTD.style.height = "20px";
    iconTD.style.borderRadius = "5px";

    headTR.appendChild(iconTD);
    table.appendChild(headTR);

    let effectsTR = ParseEffect(effect);
    table.appendChild(effectsTR);

    tip.appendChild(table);
    document.body.appendChild(tip);
}

function updatePositionTipEffect() {
    if (document.getElementById("TipEffect")) {
        document.getElementById("TipEffect").style.top = stylePositionParams.top + 'px';
        document.getElementById("TipEffect").style.left = stylePositionParams.left - 5 + 'px';
    }
}

function TipEffectOff() {
    if (document.getElementById("TipEffect")) {
        document.getElementById("TipEffect").remove();
    }
}