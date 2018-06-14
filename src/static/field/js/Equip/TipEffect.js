function TipEffectOn(effect) {
    var tip = document.createElement("div");
    tip.id = "TipEffect";

    var table = document.createElement("table");
    var headTR = document.createElement("tr");
    var headTH = document.createElement("th");

    headTH.innerHTML = "<span class='Value'> " + effect.name + " </span>";
    headTR.appendChild(headTH);

    var iconTD = document.createElement("td");
    iconTD.style.backgroundImage = "url(/assets/effects/" + effect.name + ".png)";
    iconTD.style.width = "20px";
    iconTD.style.height = "20px";
    iconTD.style.borderRadius = "5px";

    headTR.appendChild(iconTD);
    table.appendChild(headTR);

    var effectsTR = ParseEffect(effect);
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