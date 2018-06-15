function CreateUnitSubMenu(unit) {

    var unitSubMenu = document.getElementById("UnitSubMenu");

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

    if (unit.effect !== null && unit.effect.length > 0) {

        if (!unit.action && game.user.name === unit.owner) {
            unitSubMenu.style.height = "65px";
        } else {
            unitSubMenu.style.height = "20px";
        }
        EffectsPanel(unitSubMenu, unit);
    }
}

function MoveSubMenu(unitSubMenu, unit) {
    var table = document.createElement("table");
    table.style.width = "95px";

    var tr = document.createElement("tr");
    var th = document.createElement("th");
    th.style.alignContent = "center";
    th.innerHTML = "Действия:";
    th.className = "h";

    tr.appendChild(th);

    var trSkip = document.createElement("tr");
    var tdSkip = document.createElement("td");
    tdSkip.style.alignContent = "center";
    var skipButton = document.createElement("input");
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
    var table = document.createElement("table");
    table.style.width = "95px";

    var tr = document.createElement("tr");
    var th = document.createElement("th");
    th.style.alignContent = "center";
    th.innerHTML = "Действия:";
    th.className = "h";

    tr.appendChild(th);

    var trDefend = document.createElement("tr");
    var tdDefend = document.createElement("td");
    tdDefend.style.alignContent = "center";
    var defendButton = document.createElement("input");
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
    var table = document.createElement("table");
    table.className = "panel Effect";

    var tr = document.createElement("tr");
    var th = document.createElement("th");
    th.style.alignContent = "center";
    th.colSpan = 4;
    th.innerHTML = "Эфеекты:";
    th.className = "h";

    tr.appendChild(th);
    table.appendChild(tr);
    unitSubMenu.appendChild(table);

    var panel = document.createElement("table");
    panel.className = "panel Effect";

    var rowInventory;
    var count = 0;

    for (var j = 0; j < unit.effect.length; j++) {
        if (unit.effect[j].type !== "unit_always_animate") {
            if (count % 4 === 0) {
                rowInventory = document.createElement("tr");
                rowInventory.className = "row Effect";
            }

            var cellInventory = document.createElement("td");
            cellInventory.className = "cell Effect";
            cellInventory.style.backgroundImage = "url(/assets/effects/" + unit.effect[j].name + "_" + unit.effect[j].level + ".png)";
            cellInventory.effect = unit.effect[j];

            cellInventory.onmouseover = function () {
                TipEffectOn(this.effect);
            };

            cellInventory.onmouseout = function () {
                TipEffectOff();
            };

            rowInventory.appendChild(cellInventory);

            if (count % 4 === 0) {
                panel.appendChild(rowInventory);
                var height = unitSubMenu.offsetHeight + 22;
                unitSubMenu.style.height = height + "px";
            }

            count++;
        }
    }

    unitSubMenu.appendChild(panel);
}