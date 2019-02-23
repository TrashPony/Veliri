function addToGameUnit(unitStat) {
    if (game.units !== null && game.units !== undefined) {
        if (game.units.hasOwnProperty(unitStat.q)) {
            game.units[unitStat.q][unitStat.r] = unitStat;
        } else {
            game.units[unitStat.q] = {};
            game.units[unitStat.q][unitStat.r] = unitStat;
        }
    } else {
        game.units = {};
        game.units[unitStat.q] = {};
        game.units[unitStat.q][unitStat.r] = unitStat;
    }
}