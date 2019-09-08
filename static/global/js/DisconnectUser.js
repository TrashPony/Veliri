function RemoveUnit(jsonData) {

    if (!game || !game.units) return;

    let unit = game.units[jsonData.short_unit.id];
    if (unit) {

        while (0 < unit.oldPoint.length) {
            let label = unit.oldPoint.shift();
            if (label) label.destroy();
        }

        if (unit.sprite) {
            unit.sprite.destroy();
        }
        if (unit.colision) {
            unit.colision.destroy();
        }

        delete game.units[jsonData.short_unit.id]
    }
}