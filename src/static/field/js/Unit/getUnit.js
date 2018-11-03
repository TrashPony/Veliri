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

function GetStorageUnit(id) {
    let unit;

    for (let i in game.unitStorage) {
        if (game.unitStorage.hasOwnProperty(i)) {
            if (id === game.unitStorage[i].id) {
                unit = game.unitStorage[i];
                return unit
            }
        }
    }
}

function GetGameUnitXY(qUnit,rUnit) {
    let unit;

    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r)) {
                    if (game.units[q][r].q === qUnit && game.units[q][r].r === rUnit) {
                        unit = game.units[q][r];
                        return unit
                    }
                }
            }
        }
    }
}