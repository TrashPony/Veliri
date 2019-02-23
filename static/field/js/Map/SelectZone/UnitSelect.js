function MarkUnitSelect(unit, frame, onclickFunc) {
    unit.sprite.frame = frame;

    if (onclickFunc) {
        unit.sprite.events.onInputDown.add(onclickFunc);
        unit.sprite.input.priorityID = 1;
    }
}

function RemoveUnitMarks() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {
                    let unit = game.units[x][y];
                    unit.sprite.frame = 0;
                    unit.sprite.events.onInputDown.removeAll();
                    unit.sprite.input.priorityID = 0;
                }
            }
        }
    }
}