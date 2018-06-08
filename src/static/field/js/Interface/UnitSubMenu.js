function CreateUnitSubMenu(unit) {

    var unitSubMenu = document.getElementById("UnitSubMenu");

    if (unitSubMenu) {
        unitSubMenu.remove();
    }

    unitSubMenu = document.createElement("table");
    unitSubMenu.id = "UnitSubMenu";
    unitSubMenu.style.left = stylePositionParams.left + 'px';
    unitSubMenu.style.top = stylePositionParams.top + 'px';
    unitSubMenu.style.display = "block";


    if (game.Phase === "move") {

        unitSubMenu.style.width = "100px";
        unitSubMenu.style.height = "40px";

        MoveSubMenu(unitSubMenu, unit);
    }

    if (game.Phase === "targeting") {

        unitSubMenu.style.width = "100px";
        unitSubMenu.style.height = "70px";

        TargetingSubMenu(unitSubMenu, unit);
    }


    document.body.appendChild(unitSubMenu);
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
            x: Number(unit.info.x),
            y: Number(unit.info.y)
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
            x: Number(unit.info.x),
            y: Number(unit.info.y)
        }));
    };

    tdDefend.appendChild(defendButton);
    trDefend.appendChild(tdDefend);

    var trEquip = document.createElement("tr");
    var tdEquip = document.createElement("td");
    tdEquip.style.alignContent = "center";
    var equipButton = document.createElement("input");
    equipButton.type = "button";
    equipButton.value = "Инвентарь";
    equipButton.className = "button subMenu";

    equipButton.onclick = ChoiceEquip;

    tdEquip.appendChild(equipButton);
    trEquip.appendChild(tdEquip);

    table.appendChild(tr);
    table.appendChild(trDefend);
    table.appendChild(trEquip);

    unitSubMenu.appendChild(table);
}