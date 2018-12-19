let game;

function LoadGame(jsonData) {
    game = CreateGame(jsonData.map);
    game.typeService = "global";

    setTimeout(function () { // todo костыль связаной с прогрузкой карты )
        CreateSquad(jsonData.squad);
        game.input.onDown.add(initMove, game);
        console.log(game.squad)
    }, 1500);
}

function initMove(e) {
    game.squad.toMove = {x: e.worldX, y: e.worldY};
    if (game.squad.tweenTo) {
        game.squad.tweenTo.stop();
        game.squad.tweenTo = null;
    }
}

function CreateSquad(squad) {
    let q = squad.q;
    let r = squad.r;

    let x = 0;
    let y = 0;

    game.mapPoints.forEach(function (point) {
        if (point.q === q && point.r === r) {
            x = point.x;
            y = point.y;
        }
    });

    let unit;
    unit = game.unitLayer.create(x, y, 'MySelectUnit', 0);
    game.camera.focusOn(unit);

    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);
    unit.inputEnabled = true;             // включаем ивенты на спрайт

    let bodyShadow;
    bodyShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, squad.mather_ship.body.name);
    bodyShadow.scale.setTo(0.5, 0.5);
    bodyShadow.anchor.set(0.5);
    bodyShadow.tint = 0x000000;
    bodyShadow.alpha = 0.4;
    game.physics.arcade.enable(bodyShadow);
    unit.addChild(bodyShadow);

    let body = game.make.sprite(0, 0, squad.mather_ship.body.name);
    unit.addChild(body);
    game.physics.arcade.enable(body);
    body.scale.setTo(0.5, 0.5);
    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    let weapon;
    let weaponShadow;
    for (let i in squad.mather_ship.body.weapons) {
        if (squad.mather_ship.body.weapons.hasOwnProperty(i) && squad.mather_ship.body.weapons[i].weapon) {
            weapon = game.make.sprite(0, 0, squad.mather_ship.body.weapons[i].weapon.name);
            weaponShadow = game.make.sprite(game.shadowXOffset / 2, game.shadowYOffset / 2, squad.mather_ship.body.weapons[i].weapon.name);
        }
    }

    squad.sprite = unit;
    squad.sprite.unitBody = body;
    squad.sprite.bodyShadow = bodyShadow;

    if (weapon) {
        weapon.anchor.setTo(0.5, 0.61);
        weapon.scale.setTo(0.5, 0.5);
        weaponShadow.anchor.setTo(0.5, 0.61);
        weaponShadow.scale.setTo(0.5, 0.5);
        weaponShadow.tint = 0x000000;
        weaponShadow.alpha = 0.4;
        unit.addChild(weaponShadow);
        unit.addChild(weapon);

        squad.sprite.weapon = weapon;
        squad.sprite.weaponShadow = weaponShadow;
    }

    game.squad = squad;
    //body.events.onInputDown.add(SelectUnit, unitStat);    // обрабатываем наведение мышки
    //body.events.onInputOver.add(UnitMouseOver, unitStat); // обрабатываем наведение мышки
    //body.events.onInputOut.add(UnitMouseOut, unitStat);   // обрабатываем убирание мышки
}