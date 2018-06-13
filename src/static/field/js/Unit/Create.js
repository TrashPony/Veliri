function CreateUnit(unitStat, inVisible) {
    var x = unitStat.x;
    var y = unitStat.y;

    var cell = game.map.OneLayerMap[x][y].sprite;
    var unit;

    if (game.user.name === unitStat.owner) {
        unit = game.floorObjectLayer.create((cell.x + game.tileSize / 2) + game.shadowXOffset, (cell.y + game.tileSize / 2) + game.shadowYOffset, 'MySelectUnit', 0) ;
    } else {
        unit = game.floorObjectLayer.create((cell.x + game.tileSize / 2) + game.shadowXOffset, (cell.y + game.tileSize / 2) + game.shadowYOffset, 'HostileSelectUnit', 0);
    }
    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);
    unit.inputEnabled = true;             // включаем ивенты на спрайт
    unit.info = unitStat;

    var shadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, 'tank360', unitStat.rotate);
    unit.addChild(shadow);
    //game.physics.arcade.enable(shadow);
    shadow.anchor.set(0.5);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    var body = game.make.sprite(0, 0, 'tank360', unitStat.rotate);
    unit.addChild(body);
    game.physics.arcade.enable(body);
    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    body.body.collideWorldBounds = true;  // границы страницы
    body.input.pixelPerfectOver = true;   // уберает ивенты овера на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    body.events.onInputDown.add(SelectUnit, unit); // обрабатываем наведение мышки
    body.events.onInputOver.add(UnitMouseOver, unit); // обрабатываем наведение мышки
    body.events.onInputOut.add(UnitMouseOut, unit);   // обрабатываем убирание мышки

    unitStat.sprite = unit;
    unitStat.sprite.unitBody = body;
    unitStat.sprite.unitShadow = shadow;

    if (unitStat.effect !== null && unitStat.effect.length > 0) {
        CreateAnimateEffects(unitStat)
    }

    unitStat.RotateUnit = function(angle) {
        RotateUnit(this.sprite, angle);
    };

    if (unitStat.action && game.user.name === unitStat.owner) {
        DeactivationUnit(unitStat);
    }

    if (inVisible) {
        unitStat.sprite.alpha = 0;
        unitStat.sprite.unitBody.alpha = 0;
        unitStat.sprite.unitShadow.alpha = 0;
    }

    addToGameUnit(unitStat);
}

function CreateAnimateEffects(unit) {
    for (var i = 0; i < unit.effect.length; i++) {
        if (unit.effect[i].type === "unit_always_animate"){
            if (unit.effect[i].name === "animate_energy_shield") {
                energyShieldAnimate(unit);
            }
        }
    }
}