function CreateUnitSubMenu(unit) {

    var unitSubMenu = document.getElementById("UnitSubMenu");

    if (unitSubMenu) {
        unitSubMenu.remove();
    }

    unitSubMenu = document.createElement("table");
    unitSubMenu.id = "UnitSubMenu";
    unitSubMenu.style.left = stylePositionParams.left;
    unitSubMenu.style.top = stylePositionParams.top;
    unitSubMenu.style.display = "block";


    if (game.Phase === "move") {

        unitSubMenu.style.width = "100px";
        unitSubMenu.style.height = "40px";

        MoveSubMenu(unitSubMenu, unit);
    }

    if (game.Phase === "targeting") {
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

    var trSkip = document.createElement("tr");
    var tdSkip = document.createElement("td");
    tdSkip.style.alignContent = "center";
    var skipButton = document.createElement("input");
    skipButton.type = "button";
    skipButton.value = "Отменить цель";
    skipButton.className = "button subMenu";

    skipButton.onclick = function () {
        field.send(JSON.stringify({
            event: "DeleteTarget",
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