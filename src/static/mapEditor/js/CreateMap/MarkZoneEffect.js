function MarkZoneEffect(coordinate, x, y) {
    let effectsLabel = game.effectsLayer.create(x - 23, y + 23, "labelEffects");

    coordinate.effectsLabel = effectsLabel;

    effectsLabel.inputEnabled = true;
    effectsLabel.events.onInputOut.add(TipZoneOut, coordinate);
    effectsLabel.events.onInputOver.add(TipZoneEffects, coordinate);
    effectsLabel.events.onInputDown.add(DetailZoneEffects, coordinate);

    effectsLabel.input.priorityID = 1;
    effectsLabel.alpha = 1;
    effectsLabel.alpha = 0.8;
}

function DetailZoneEffects() {
    if (game.input.activePointer.leftButton.isDown) {
        this.effectsLabel.scale.setTo(1.05);
        this.effectsLabel.alpha = 1;

        if (document.getElementById("effectZonePanel")) {
            document.getElementById("effectZonePanel").remove()
        }

        if (document.getElementById("effectDetailZonePanel")) {
            document.getElementById("effectDetailZonePanel").remove();
        }

        createTipEffects(this.effects, true);
    }
}

function TipZoneEffects() {
    this.effectsLabel.scale.setTo(1.05);
    this.effectsLabel.alpha = 1;
    if (!document.getElementById("effectZonePanel") && !document.getElementById("effectDetailZonePanel")) {
        createTipEffects(this.effects, false);
    }
}

function TipZoneOut() {
    this.effectsLabel.scale.setTo(1);
    this.effectsLabel.alpha = 0.8;
    if (document.getElementById("effectZonePanel")) {
        document.getElementById("effectZonePanel").remove()
    }
}

function createTipEffects(effects, detail) {
    let effectPanel = document.createElement("div");

    if (detail) {
        effectPanel.id = "effectDetailZonePanel";
        let cancel = document.createElement("div");
        cancel.id = "cancelTipEffectsButton";
        cancel.onclick = function () {
            effectPanel.remove();
        };
        effectPanel.appendChild(cancel);
    } else {
        effectPanel.id = "effectZonePanel";
    }

    effectPanel.style.top = stylePositionParams.top + "px";
    effectPanel.style.left = stylePositionParams.left + "px";

    let panel = document.createElement("table");
    panel.className = "panel Effect";

    let rowEffect;
    let count = 0;

    for (let j = 0; j < effects.length; j++) {
        if (effects[j].type !== "unit_always_animate") {
            if (count % 4 === 0) {
                rowEffect = document.createElement("tr");
                rowEffect.className = "row Effect";
            }

            let cellEffect = document.createElement("td");
            cellEffect.className = "cell Effect";
            cellEffect.style.backgroundImage = "url(/assets/effects/" + effects[j].name + "_" + effects[j].level + ".png)";
            cellEffect.effects = effects[j];

            cellEffect.onmouseover = function () {
                TipEffectOn(this.effects);
            };

            cellEffect.onmouseout = function () {
                TipEffectOff();
            };

            rowEffect.appendChild(cellEffect);

            if (count % 4 === 0) {
                panel.appendChild(rowEffect);
            }

            count++;
        }
    }

    effectPanel.appendChild(panel);
    document.body.appendChild(effectPanel);
}

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


function ParseEffect(effect, equip) {
    let effectsTR = document.createElement("tr");
    let effectsTD = document.createElement("td");

    effectsTD.colSpan = 2;
    effectsTD.style.fontSize = "8pt";
    effectsTD.style.backgroundColor = "#4c4c4c";
    effectsTD.style.borderRadius = "5px";

    let type = "";
    let quantity = "";
    let time = "";
    let region = "";

    if (equip !== undefined && equip.region > 0) {
        if (equip.region === 1) {
            region = "<br> в радиусе <span class='Value'>" + equip.region + " клетки</span>"
        } else {
            region = "<br> в радиусе <span class='Value'>" + equip.region + " клеток</span>"
        }
    }

    if (equip !== undefined) {
        if (effect.forever) {
            if (equip.steps_time === 1) {
                time = ""
            } else {
                time = "<br> в течение <span class='Value'>" + equip.steps_time + " ходов</span>";
            }
        } else {
            if (equip.steps_time > 4) {
                time = "<br> на <span class='Value'>" + equip.steps_time + " ходов</span>";
            } else {
                time = "<br> на <span class='Value'>" + equip.steps_time + " хода</span>";
            }
        }
    } else {
        if (effect.steps_time === 1) {
            time = "<br> остался <span class='Value'>" + effect.steps_time + " ход</span>";
        }
        if (effect.steps_time > 1 && 5 > effect.steps_time) {
            time = "<br> осталось <span class='Value'>" + effect.steps_time + " хода</span>";
        }
        if (5 <= effect.steps_time) {
            time = "<br> осталось <span class='Value'>" + effect.steps_time + " ходов</span>";
        }
    }

    if (effect.percentages) {
        quantity = effect.quantity + "%";
    } else {
        quantity = effect.quantity;
    }

    if (effect.type === "enhances") {
        type = "+"
    }

    if (effect.type === "takes_away") {
        type = "-"
    }

    if (effect.type === "replenishes") {
        type = "++"
    }

    effectsTD.innerHTML = "<span class='Value'>" + type + quantity + " " + effect.parameter + "</span>" + time + region;

    effectsTR.appendChild(effectsTD);
    return effectsTR
}