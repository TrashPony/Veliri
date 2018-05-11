function CreateMyGameUnits() {
    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    CreateUnit(game.units[x][y]);
                }
            }
        }
    }
}

function CreateHostileGameUnits() {
    for (var x in game.hostileUnits) {
        if (game.hostileUnits.hasOwnProperty(x)) {
            for (var y in game.hostileUnits[x]) {
                if (game.hostileUnits[x].hasOwnProperty(y)) {
                    CreateUnit(game.hostileUnits[x][y]);
                }
            }
        }
    }
}