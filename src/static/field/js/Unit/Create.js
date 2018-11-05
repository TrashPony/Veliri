function CreateUnit(unitStat, inVisible) {
    let q = unitStat.q;
    let r = unitStat.r;

    let cell = game.map.OneLayerMap[q][r].sprite;
    let x = cell.x + cell.width / 2;
    let y = cell.y + cell.height / 2;

    let unit;

    if (game.user.name === unitStat.owner) {
        unit = game.unitLayer.create(x, y, 'MySelectUnit', 0);
    } else {
        unit = game.unitLayer.create(x, y, 'HostileSelectUnit', 0);
    }

    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);
    unit.inputEnabled = true;             // включаем ивенты на спрайт

    let bodyShadow;
    if (unitStat.body.mother_ship) {
        bodyShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, unitStat.body.name);
    } else {
        bodyShadow = game.make.sprite(game.shadowXOffset / 2, game.shadowYOffset / 2, unitStat.body.name);
    }
    unit.addChild(bodyShadow);

    if (unitStat.body.mother_ship) {
        bodyShadow.scale.setTo(0.5, 0.5);
    } else {
        bodyShadow.scale.setTo(0.3, 0.3);
    }

    bodyShadow.anchor.set(0.5);
    bodyShadow.tint = 0x000000;
    bodyShadow.alpha = 0.4;
    game.physics.arcade.enable(bodyShadow);

    let weapon;
    let weaponShadow;
    for (let i in unitStat.body.weapons) {
        if (unitStat.body.weapons.hasOwnProperty(i) && unitStat.body.weapons[i].weapon) {
            weapon = game.make.sprite(0, 0, unitStat.body.weapons[i].weapon.name);
            weaponShadow = game.make.sprite(game.shadowXOffset / 2, game.shadowYOffset / 2, unitStat.body.weapons[i].weapon.name);
        }
    }

    if (weapon) {
        weapon.anchor.setTo(0.5, 0.61);

        if (unitStat.body.mother_ship) {
            weapon.scale.setTo(0.5, 0.5);
        } else {
            weapon.scale.setTo(0.3, 0.3);
        }

        weaponShadow.anchor.setTo(0.5, 0.61);

        if (unitStat.body.mother_ship) {
            weaponShadow.scale.setTo(0.5, 0.5);
        } else {
            weaponShadow.scale.setTo(0.3, 0.3);
        }

        weaponShadow.tint = 0x000000;
        weaponShadow.alpha = 0.4;
    }

    let body = game.make.sprite(0, 0, unitStat.body.name);
    unit.addChild(body);
    game.physics.arcade.enable(body);

    if (unitStat.body.mother_ship) {
        body.scale.setTo(0.5, 0.5);
    } else {
        body.scale.setTo(0.3, 0.3);
    }

    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    body.events.onInputDown.add(SelectUnit, unitStat);    // обрабатываем наведение мышки
    body.events.onInputOver.add(UnitMouseOver, unitStat); // обрабатываем наведение мышки
    body.events.onInputOut.add(UnitMouseOut, unitStat);   // обрабатываем убирание мышки

    if (weapon) {
        unit.addChild(weaponShadow);
        unit.addChild(weapon);
    }

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
    unitStat.sprite.bodyShadow = bodyShadow;
    unitStat.sprite.healBar = healBar;
    unitStat.sprite.heal = heal;
    unitStat.sprite.weapon = weapon;
    unitStat.sprite.weaponShadow = weaponShadow;


    CalculateHealBar(unitStat);

    if (unitStat.effects !== null && unitStat.effects.length > 0) {
        CreateAnimateEffects(unitStat)
    }

    if (unitStat.action && game.user.name === unitStat.owner) {
        DeactivationUnit(unitStat);
    }

    if (inVisible) {
        unitStat.sprite.alpha = 0;
        unitStat.sprite.unitBody.alpha = 0;
        unitStat.sprite.bodyShadow.alpha = 0;
        unitStat.sprite.healBar.alpha = 0;
        unitStat.sprite.heal.alpha = 0;
        unitStat.sprite.weapon.alpha = 0;
        unitStat.sprite.weaponShadow.alpha = 0;
    }

    addToGameUnit(unitStat);
    SetAngle(unitStat, unitStat.rotate + 90);

    return unitStat
}

function CreateAnimateEffects(unit) {
    for (let i in unit.effects) {
        if (unit.effects.hasOwnProperty(i) && unit.effects[i] != null && unit.effects[i].type === "unit_always_animate") {
            if (unit.effects[i].name === "animate_energy_shield") {
                energyShieldAnimate(unit);
            }
        }
    }
}