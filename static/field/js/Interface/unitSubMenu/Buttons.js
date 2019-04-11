function ActionButton(equipPanel, unit, event, text) {

    let skipButton = document.createElement("input");
    skipButton.type = "button";
    skipButton.value = text;


    skipButton.onclick = function () {
        field.send(JSON.stringify({
            event: event,
            q: Number(unit.q),
            r: Number(unit.r),
            unit_id: Number(unit.id)
        }));
    };

    equipPanel.appendChild(skipButton);
}