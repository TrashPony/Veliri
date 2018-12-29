function CreateSquad(squad, x, y, bodyName, weaponName, rotate, focus) {
    let unit;

    if (!game.unitLayer) return;

    unit = game.unitLayer.create(x, y, 'MySelectUnit', 0);
    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);

    if (focus) {
        game.camera.focusOn(unit);
    }

    let bodyShadow;
    bodyShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, bodyName);
    bodyShadow.scale.setTo(0.5, 0.5);
    bodyShadow.anchor.set(0.5);
    bodyShadow.tint = 0x000000;
    bodyShadow.alpha = 0.4;
    game.physics.arcade.enable(bodyShadow);

    let body = game.make.sprite(0, 0, bodyName);
    game.physics.arcade.enable(body);
    body.scale.setTo(0.5, 0.5);
    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта



    let weapon;
    let weaponShadow;
    if (weaponName && weaponName !== '') {
        weapon = game.make.sprite(0, 0, weaponName);
        weaponShadow = game.make.sprite(game.shadowXOffset / 2, game.shadowYOffset / 2, weaponName);
        weapon.anchor.setTo(0.5, 0.61);
        weapon.scale.setTo(0.5, 0.5);
        weaponShadow.anchor.setTo(0.5, 0.61);
        weaponShadow.scale.setTo(0.5, 0.5);
        weaponShadow.tint = 0x000000;
        weaponShadow.alpha = 0.4;
    }

    squad.sprite = unit;
    squad.sprite.unitBody = body;
    squad.sprite.bodyShadow = bodyShadow;

    unit.addChild(bodyShadow);
    unit.addChild(body);

    if (weapon) {
        squad.sprite.weapon = weapon;
        squad.sprite.weaponShadow = weaponShadow;

        unit.addChild(weaponShadow);
        unit.addChild(weapon);
    }

    SetAngle(squad, rotate);
}