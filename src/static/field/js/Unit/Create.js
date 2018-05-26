function CreateUnit(unitStat) {
    var x = unitStat.x;
    var y = unitStat.y;

    var cell = game.map.OneLayerMap[x][y].sprite;

    var shadow = game.floorObjectLayer.create((cell.x + game.tileSize / 2) + game.shadowXOffset, (cell.y + game.tileSize / 2) + game.shadowYOffset, 'tank360', unitStat.rotate + 224); //todo смещение для тестового спрайта
    game.physics.arcade.enable(shadow);
    shadow.anchor.set(0.5);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    var unit = game.floorObjectLayer.create(cell.x + game.tileSize / 2, cell.y + game.tileSize / 2, 'tank360', unitStat.rotate + 224); //todo смещение для тестового спрайта
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

    unitStat.shadow = shadow;
    unitStat.sprite = unit;

    addToGameUnit(unitStat);
}