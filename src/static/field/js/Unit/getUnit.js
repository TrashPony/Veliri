function GetGameUnitID(id) {
    var unit;

    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
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
    var unit;

    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    console.log(game.units[x][y]);
                    if (game.units[x][y].x === xUnit && game.units[x][y].y === yUnit) {
                        unit = game.units[x][y];
                        return unit
                    }
                }
            }
        }
    }
}