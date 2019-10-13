function CreateUnit(unit, x, y, rotate, bColor, b2Color, wColor, w2Color, userID, boxType, healBar) {
    if (!game.unitLayer) return;

    let weaponSlot;
    for (let i in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon) {
            weaponSlot = unit.body.weapons[i]
        }
    }

    let unitBox = game.unitLayer.create(x, y, boxType, 0);
    game.physics.enable(unitBox, Phaser.Physics.ARCADE);
    unitBox.anchor.setTo(0.5, 0.5);
    if (!unit.body.mother_ship) {
        unitBox.scale.setTo(0.75);
    }

    let bodyBottomShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, unit.body.name + "_bottom_animate", 11);
    bodyBottomShadow.animations.add('move');
    bodyBottomShadow.play('move', 25, true).paused = true;
    bodyBottomShadow.scale.setTo(0.25);
    bodyBottomShadow.anchor.set(0.5);
    bodyBottomShadow.tint = 0x000000;
    bodyBottomShadow.alpha = 0.2;

    let bodyBottom = game.make.sprite(0, 0, unit.body.name + "_bottom_animate", 11);
    bodyBottom.animations.add('move');
    bodyBottom.play('move', 25, true).paused = true;
    bodyBottom.scale.setTo(0.25);
    bodyBottom.inputEnabled = true;             // включаем ивенты на спрайт
    bodyBottom.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта

    let bodyShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, unit.body.name);
    bodyShadow.scale.setTo(0.25);
    bodyShadow.anchor.set(0.5);
    bodyShadow.tint = 0x000000;
    bodyShadow.alpha = 0.2;

    let body = game.make.sprite(0, 0, unit.body.name);
    let bodyMask = game.make.sprite(0, 0, unit.body.name + '_mask');
    let bodyMask2 = game.make.sprite(0, 0, unit.body.name + '_mask2');

    body.scale.setTo(0.25);
    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5);               // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта
    body.input.priorityID = 1;

    mouseBodyOver(body, unit, unitBox);
    body.events.onInputDown.add(function () {
        SelectOneUnit(unit, unitBox, true);
    }, this);

    bodyMask.anchor.setTo(0.5);          // устанавливаем центр спрайта
    bodyMask.tint = bColor;

    bodyMask2.anchor.setTo(0.5);          // устанавливаем центр спрайта
    bodyMask2.tint = b2Color;
    bodyMask2.alpha = 0.3;

    // накладываем цветовые маски на спрайт
    body.addChild(bodyMask2);
    body.addChild(bodyMask);

    let weapon;
    let weaponShadow;
    let weaponColorMask;
    let weaponColorMask2;

    // размер картинки спрайта для правильного расположение точек слотов
    let weaponScale;
    //костыльная переменная
    let spriteOffset;

    if (unit.body.mother_ship) {
        weaponScale = 0.25;
        spriteOffset = 50;
    } else {
        weaponScale = 0.20;
        spriteOffset = 20;
    }

    if (weaponSlot && weaponSlot.weapon) {

        let xAttach = ((weaponSlot.x_attach) / (1 / weaponScale)) - spriteOffset;
        let yAttach = ((weaponSlot.y_attach) / (1 / weaponScale)) - spriteOffset;

        weapon = game.make.sprite(xAttach, yAttach, weaponSlot.weapon.name);
        weaponColorMask = game.make.sprite(0, 0, weaponSlot.weapon.name + '_mask');
        weaponColorMask2 = game.make.sprite(0, 0, weaponSlot.weapon.name + '_mask2');

        weaponShadow = game.make.sprite(xAttach + game.shadowXOffset / 2, yAttach + game.shadowYOffset / 2, weaponSlot.weapon.name);

        weapon.xAttach = xAttach;
        weapon.yAttach = yAttach;
        weapon.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weapon.scale.setTo(weaponScale);

        weaponColorMask.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponColorMask.tint = wColor;

        weaponColorMask2.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponColorMask2.tint = w2Color;
        weaponColorMask2.alpha = 0.3;

        weaponShadow.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponShadow.scale.setTo(weaponScale);
        weaponShadow.tint = 0x000000;
        weaponShadow.alpha = 0.5;

        // накладываем цветовые маски на спрайт
        weapon.addChild(weaponColorMask2);
        weapon.addChild(weaponColorMask);
    }

    unit.sprite = unitBox;
    unit.sprite.unitBody = body;
    unit.sprite.bodyShadow = bodyShadow;
    unit.sprite.bodyBottom = bodyBottom;
    unit.sprite.bodyBottomShadow = bodyBottomShadow;

    unitBox.addChild(bodyBottomShadow);
    unitBox.addChild(bodyBottom);
    unitBox.addChild(bodyShadow);
    unitBox.addChild(body);

    CreateEquip(unit);

    if (weapon) {
        unit.sprite.weapon = weapon;
        unit.sprite.weaponShadow = weaponShadow;

        unitBox.addChild(weaponShadow);
        unitBox.addChild(weapon);
    }

    unit.sprite.angle = rotate;
    SetShadowAngle(unit, rotate);
    // принимаем угол башни
    RotateGun(unit, unit.gun_rotate, 10);

    return unit
}

function CreateEquip(unit) {
    unit.sprite.equipSprites = [];

    let createSprite = function (slot) {
        if (slot.equip && (slot.equip.x_attach > 0 && slot.equip.y_attach > 0)) {
            let xAttach = ((slot.x_attach) / 4) - 50;
            let yAttach = ((slot.y_attach) / 4) - 50;

            let slotSprite = game.make.sprite(xAttach, yAttach, slot.equip.name);
            let attachPoint = game.make.sprite(xAttach, yAttach, slot.equip.name);
            let slotShadow = game.make.sprite(xAttach + game.shadowXOffset / 4, yAttach + game.shadowYOffset / 4, slot.equip.name);

            slotSprite.anchor.setTo(slot.equip.x_attach / 256, slot.equip.y_attach / 256);
            slotSprite.scale.setTo(0.25);

            attachPoint.anchor.setTo(slot.equip.x_attach / 256, slot.equip.y_attach / 256);
            attachPoint.scale.setTo(0.25);

            slotShadow.anchor.setTo(slot.equip.x_attach / 256, slot.equip.y_attach / 256);
            slotShadow.scale.setTo(0.25);
            slotShadow.tint = 0x000000;
            slotShadow.alpha = 0.3;

            unit.sprite.addChild(slotShadow);
            unit.sprite.addChild(attachPoint);
            unit.sprite.addChild(slotSprite);

            unit.sprite.equipSprites.push({
                sprite: slotSprite,
                shadow: slotShadow,
                xAttach: xAttach,
                yAttach: yAttach,
                slot: slot,
                attachPoint: attachPoint
            });
        }
    };

    let createSlots = function (slots) {
        for (let slot in slots) {
            createSprite(slots[slot])
        }
    };

    createSlots(unit.body.equippingI);
    createSlots(unit.body.equippingII);
    createSlots(unit.body.equippingIII);
    createSlots(unit.body.equippingIV);
    createSlots(unit.body.equippingV);
}