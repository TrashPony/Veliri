function EffectsPanel(unitSubMenu, unit) {
    let effectPanel = document.createElement("div");
    effectPanel.id = "effectPanel";

    let head = document.createElement("span");
    head.className = "Value";
    head.innerHTML = "Эффекты:";
    effectPanel.appendChild(head);

    let panel = document.createElement("table");
    panel.className = "panel Effect";

    let rowEffect;
    let count = 0;

    for (let j = 0; j < unit.effects.length; j++) {
        if (unit.effects[j] && unit.effects[j].type !== "unit_always_animate") {
            if (count % 4 === 0) {
                rowEffect = document.createElement("tr");
                rowEffect.className = "row Effect";
            }

            let cellEffect = document.createElement("td");
            cellEffect.className = "cell Effect";
            cellEffect.style.backgroundImage = "url(/assets/effects/" + unit.effects[j].name + "_" + unit.effects[j].level + ".png)";
            cellEffect.effects = unit.effects[j];

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
    unitSubMenu.appendChild(effectPanel);
}