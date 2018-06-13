function MarkUnitSelect(unit) {
    unit.frame = 1;
}

function RemoveUnitMarks() {
    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {
                    var unit = game.units[x][y];
                    unit.sprite.frame = 0;
                    unit.sprite.events.onInputDown.removeAll();
                    unit.sprite.input.priorityID = 0;
                }
            }
        }
    }
}