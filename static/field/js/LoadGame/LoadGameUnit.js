function CreateMyGameUnits() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    CreateLocalUnit(game.units[x][y]);
                    if (game.units[x][y].target && game.Phase === "targeting") {
                        CreateTargetLine(game.units[x][y]);
                    }
                }
            }
        }
    }
}

function CreateHostileGameUnits() {
    for (let x in game.hostileUnits) {
        if (game.hostileUnits.hasOwnProperty(x)) {
            for (let y in game.hostileUnits[x]) {
                if (game.hostileUnits[x].hasOwnProperty(y)) {
                    CreateLocalUnit(game.hostileUnits[x][y]);
                }
            }
        }
    }
}