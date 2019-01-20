function MarkZoneEffect(coordinate, x, y) {
    let effectsLabel = game.effectsLayer.create(x + 20, y - 40, "labelEffects");

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