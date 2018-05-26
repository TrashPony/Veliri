function addToGameUnit(unitStat) {
    if (game.units !== null && game.units !== undefined) {
        if (game.units.hasOwnProperty(unitStat.x)) {
            game.units[unitStat.x][unitStat.y] = unitStat;
        } else {
            game.units[unitStat.x] = {};
            game.units[unitStat.x][unitStat.y] = unitStat;
        }
    } else {
        game.units = {};
        game.units[unitStat.x] = {};
        game.units[unitStat.x][unitStat.y] = unitStat;
    }
}