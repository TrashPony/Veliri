function MoveButton(equipPanel, unit) {
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

    equipPanel.appendChild(table);
}

function TargetingButton(equipPanel, unit) {
    let table = document.createElement("table");
    table.id = "targetButtonsTable";
    table.style.width = "150px";

    let trDefend = document.createElement("tr");
    let tdDefend = document.createElement("td");
    tdDefend.style.alignContent = "center";
    let defendButton = document.createElement("input");
    defendButton.type = "button";
    defendButton.value = "Защита";
    defendButton.className = "button unitSubMenu";

    defendButton.onclick = function () {
        field.send(JSON.stringify({
            event: "Defend",
            x: Number(unit.x),
            y: Number(unit.y)
        }));
    };

    tdDefend.appendChild(defendButton);
    trDefend.appendChild(tdDefend);

    table.appendChild(trDefend);

    equipPanel.appendChild(table);
}