function CreateUnit(unitStat, inVisible) {
    let q = unitStat.q;
    let r = unitStat.r;

    let cell = game.map.OneLayerMap[q][r].sprite;
    let x = cell.x + cell.width / 2;
    let y = cell.y + cell.height / 2;

    let unit;

    if (game.user.name === unitStat.owner) {
        unit = game.floorObjectLayer.create(x + game.shadowXOffset, y + game.shadowYOffset, 'MySelectUnit', 0) ;
    } else {
        unit = game.floorObjectLayer.create(x + game.shadowXOffset, y + game.shadowYOffset, 'HostileSelectUnit', 0);
    }

    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);
    unit.inputEnabled = true;             // включаем ивенты на спрайт


    let shadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, unitStat.body.name, unitStat.rotate);
    unit.addChild(shadow);
    game.physics.arcade.enable(shadow);
    shadow.anchor.set(0.5);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    let body = game.make.sprite(0, 0, unitStat.body.name, unitStat.rotate);
    unit.addChild(body);
    game.physics.arcade.enable(body);

    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    body.events.onInputDown.add(SelectUnit, unitStat);    // обрабатываем наведение мышки
    body.events.onInputOver.add(UnitMouseOver, unitStat); // обрабатываем наведение мышки
    body.events.onInputOut.add(UnitMouseOut, unitStat);   // обрабатываем убирание мышки

    let healBar = game.make.sprite(0, 45, 'healBar');
    unit.addChild(healBar);
    healBar.anchor.setTo(0.5);
    healBar.alpha = 0;

    let heal = game.make.sprite(-50, 0, 'heal');
    healBar.addChild(heal);
    heal.anchor.setTo(0, 0.5);
    heal.alpha = 1;

    unitStat.sprite = unit;
    unitStat.sprite.unitBody = body;
    unitStat.sprite.unitShadow = shadow;
    unitStat.sprite.healBar = healBar;
    unitStat.sprite.heal = heal;

    CalculateHealBar(unitStat);

    if (unitStat.effects !== null && unitStat.effects.length > 0) {
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
        unitStat.sprite.healBar.alpha = 0;
        unitStat.sprite.heal = heal;
    }

    addToGameUnit(unitStat);
}

function CreateAnimateEffects(unit) {
    for (let i in unit.effects) {
        if (unit.effects.hasOwnProperty(i) && unit.effects[i].type === "unit_always_animate"){
            if (unit.effects[i].name === "animate_energy_shield") {
                energyShieldAnimate(unit);
            }
        }
    }
}