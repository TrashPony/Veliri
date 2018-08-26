function GetGameUnitID(id) {
    let unit;

    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    if (id === game.units[x][y].id) {
                        unit = game.units[x][y];
                        return unit
                    }
                }
            }
        }
    }
}

function GetGameUnitXY(xUnit,yUnit) {
    let unit;

    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    if (game.units[x][y].x === xUnit && game.units[x][y].y === yUnit) {
                        unit = game.units[x][y];
                        return unit
                    }
                }
            }
        }
    }
}