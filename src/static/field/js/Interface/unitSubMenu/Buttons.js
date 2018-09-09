function ActionButton(equipPanel, unit, event, text) {
    let table = document.createElement("table");
    table.id = "targetButtonsTable";
    table.style.width = "150px";

    let trSkip = document.createElement("tr");
    let tdSkip = document.createElement("td");
    tdSkip.style.alignContent = "center";
    let skipButton = document.createElement("input");
    skipButton.type = "button";
    skipButton.value = text;
    console.log(unit);
    if (unit.action_points >= unit.body.speed) {
        skipButton.className = "button unitSubMenu";

        skipButton.onclick = function () {
            field.send(JSON.stringify({
                event: event,
                q: Number(unit.q),
                r: Number(unit.r)
            }));
        };
    } else {
        skipButton.className = "button unitSubMenu noActive";
    }

    tdSkip.appendChild(skipButton);
    trSkip.appendChild(tdSkip);

    table.appendChild(trSkip);

    equipPanel.appendChild(table);
}