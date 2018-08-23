function CreateUnitSubMenu(unit) {
    if (game.Phase === "move" || game.Phase === "targeting") {

        let unitSubMenu = document.getElementById("UnitSubMenu");

        if (unitSubMenu) {
            unitSubMenu.remove();
        }

        unitSubMenu = document.createElement("div");
        unitSubMenu.id = "UnitSubMenu";
        unitSubMenu.style.left = stylePositionParams.left + 'px';
        unitSubMenu.style.top = stylePositionParams.top + 'px';
        unitSubMenu.style.display = "block";

        if (!unit.action && game.user.name === unit.owner) {
            if (game.Phase === "move") {

                unitSubMenu.style.width = "100px";
                unitSubMenu.style.height = "45px";

                MoveSubMenu(unitSubMenu, unit);
            }

            if (game.Phase === "targeting") {

                unitSubMenu.style.width = "100px";
                unitSubMenu.style.height = "45px";

                TargetingSubMenu(unitSubMenu, unit);
            }
        }

        document.body.appendChild(unitSubMenu);

        if (unit.effects !== null && unit.effects.length > 0) {

            if (!unit.action && game.user.name === unit.owner) {
                unitSubMenu.style.height = "65px";
            } else {
                unitSubMenu.style.height = "20px";
            }
            EffectsPanel(unitSubMenu, unit);
        }
    }
}

function MoveSubMenu(unitSubMenu, unit) {
    let table = document.createElement("table");
    table.style.width = "95px";

    let tr = document.createElement("tr");
    let th = document.createElement("th");
    th.style.alignContent = "center";
    th.innerHTML = "Действия:";
    th.className = "h";

    tr.appendChild(th);

    let trSkip = document.createElement("tr");
    let tdSkip = document.createElement("td");
    tdSkip.style.alignContent = "center";
    let skipButton = document.createElement("input");
    skipButton.type = "button";
    skipButton.value = "Пропустить ход";
    skipButton.className = "button subMenu";

    skipButton.onclick = function () {
        field.send(JSON.stringify({
            event: "SkipMoveUnit",
            x: Number(unit.x),
            y: Number(unit.y)
        }));
    };

    tdSkip.appendChild(skipButton);
    trSkip.appendChild(tdSkip);

    table.appendChild(tr);
    table.appendChild(trSkip);

    unitSubMenu.appendChild(table);
}

function TargetingSubMenu(unitSubMenu, unit) {
    let table = document.createElement("table");
    table.style.width = "95px";

    let tr = document.createElement("tr");
    let th = document.createElement("th");
    th.style.alignContent = "center";
    th.innerHTML = "Действия:";
    th.className = "h";

    tr.appendChild(th);

    let trDefend = document.createElement("tr");
    let tdDefend = document.createElement("td");
    tdDefend.style.alignContent = "center";
    let defendButton = document.createElement("input");
    defendButton.type = "button";
    defendButton.value = "Защита";
    defendButton.className = "button subMenu";

    defendButton.onclick = function () {
        field.send(JSON.stringify({
            event: "Defend",
            x: Number(unit.x),
            y: Number(unit.y)
        }));
    };

    tdDefend.appendChild(defendButton);
    trDefend.appendChild(tdDefend);

    table.appendChild(tr);
    table.appendChild(trDefend);

    unitSubMenu.appendChild(table);
}

function EffectsPanel(unitSubMenu, unit) {
    let table = document.createElement("table");
    table.className = "panel Effect";

    let tr = document.createElement("tr");
    let th = document.createElement("th");
    th.style.alignContent = "center";
    th.colSpan = 4;
    th.innerHTML = "Эфеекты:";
    th.className = "h";

    tr.appendChild(th);
    table.appendChild(tr);
    unitSubMenu.appendChild(table);

    let panel = document.createElement("table");
    panel.className = "panel Effect";

    let rowInventory;
    let count = 0;

    for (let j = 0; j < unit.effects.length; j++) {
        if (unit.effects[j].type !== "unit_always_animate") {
            if (count % 4 === 0) {
                rowInventory = document.createElement("tr");
                rowInventory.className = "row Effect";
            }

            let cellInventory = document.createElement("td");
            cellInventory.className = "cell Effect";
            cellInventory.style.backgroundImage = "url(/assets/effects/" + unit.effects[j].name + "_" + unit.effects[j].level + ".png)";
            cellInventory.effects = unit.effects[j];

            cellInventory.onmouseover = function () {
                TipEffectOn(this.effects);
            };

            cellInventory.onmouseout = function () {
                TipEffectOff();
            };

            rowInventory.appendChild(cellInventory);

            if (count % 4 === 0) {
                panel.appendChild(rowInventory);
                let height = unitSubMenu.offsetHeight + 22;
                unitSubMenu.style.height = height + "px";
            }

            count++;
        }
    }

    unitSubMenu.appendChild(panel);
}