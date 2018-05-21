function CreateUnit(unitStat) {
    var x = unitStat.x;
    var y = unitStat.y;

    var cell = game.map.OneLayerMap[x][y].sprite;

    var shadow = game.floorObjectLayer.create((cell.x + game.tileSize / 2) + game.shadowXOffset, (cell.y + game.tileSize / 2) + game.shadowYOffset, 'tank360', unitStat.rotate);
    shadow.anchor.set(0.5);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    var unit = game.floorObjectLayer.create(cell.x + game.tileSize / 2, cell.y + game.tileSize / 2, 'tank360', unitStat.rotate);
    game.physics.arcade.enable(unit);
    unit.inputEnabled = true;             // включаем ивенты на спрайт
    unit.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    unit.body.collideWorldBounds = true;  // границы страницы
    unit.info = unitStat;
    unit.input.pixelPerfectOver = true;   // уберает ивенты овера на пустую зону спрайта
    unit.input.pixelPerfectClick = true;   // уберает ивенты кликов на пустую зону спрайта

    unit.events.onInputDown.add(SelectUnit, unit); // обрабатываем наведение мышки
    unit.events.onInputOver.add(unitTip, unit); // обрабатываем наведение мышки
    unit.events.onInputOut.add(TipOff);   // обрабатываем убирание мышки

    if (game.units !== null && game.units !== undefined) {
        if (game.units.hasOwnProperty(unitStat.x)) {
            if (game.units[unitStat.x].hasOwnProperty(unitStat.y)) {
                game.units[unitStat.x][unitStat.y].shadow = shadow;
                game.units[unitStat.x][unitStat.y].sprite = unit;
            } else {
                game.units[unitStat.x][unitStat.y] = {};
                game.units[unitStat.x][unitStat.y].shadow = shadow;
                game.units[unitStat.x][unitStat.y].sprite = unit;
            }
        } else {
            game.units[unitStat.x] = {};
            game.units[unitStat.x][unitStat.y] = {};
            game.units[unitStat.x][unitStat.y].shadow = shadow;
            game.units[unitStat.x][unitStat.y].sprite = unit;
        }
    } else {
        game.units = {};
        game.units[unitStat.x] = {};
        game.units[unitStat.x][unitStat.y] = {};
        game.units[unitStat.x][unitStat.y].shadow = shadow;
        game.units[unitStat.x][unitStat.y].sprite = unit;
    }
}