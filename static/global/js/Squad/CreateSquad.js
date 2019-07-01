function CreateSquad(squad, x, y, squadBody, weaponSlot, rotate, bColor, b2Color, wColor, w2Color) {
    let unit;

    if (!game.unitLayer) return;

    unit = game.unitLayer.create(x, y, 'MySelectUnit', 0);
    game.physics.enable(unit, Phaser.Physics.ARCADE);
    unit.anchor.setTo(0.5, 0.5);

    let bodyBottomShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, squadBody.name + "_bottom_animate", 11);
    bodyBottomShadow.animations.add('move');
    bodyBottomShadow.play('move', 25, true).paused = true;
    bodyBottomShadow.scale.setTo(0.25);
    bodyBottomShadow.anchor.set(0.5);
    bodyBottomShadow.tint = 0x000000;
    bodyBottomShadow.alpha = 0.2;

    let bodyBottom = game.make.sprite(0, 0, squadBody.name + "_bottom_animate", 11);
    bodyBottom.animations.add('move');
    bodyBottom.play('move', 25, true).paused = true;
    bodyBottom.scale.setTo(0.25);
    bodyBottom.inputEnabled = true;             // включаем ивенты на спрайт
    bodyBottom.anchor.setTo(0.5, 0.5);          // устанавливаем центр спрайта

    let bodyShadow = game.make.sprite(game.shadowXOffset, game.shadowYOffset, squadBody.name);
    bodyShadow.scale.setTo(0.25);
    bodyShadow.anchor.set(0.5);
    bodyShadow.tint = 0x000000;
    bodyShadow.alpha = 0.2;

    let body = game.make.sprite(0, 0, squadBody.name);
    let bodyMask = game.make.sprite(0, 0, squadBody.name + '_mask');
    let bodyMask2 = game.make.sprite(0, 0, squadBody.name + '_mask2');

    body.scale.setTo(0.25);
    body.inputEnabled = true;             // включаем ивенты на спрайт
    body.anchor.setTo(0.5);               // устанавливаем центр спрайта
    body.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    body.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    mouseBodyOver(body, squad, unit);

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

    if (weaponSlot && weaponSlot.weapon) {

        let xAttach = ((weaponSlot.x_attach) / 4) - 50;
        let yAttach = ((weaponSlot.y_attach) / 4) - 50;

        weapon = game.make.sprite(xAttach, yAttach, weaponSlot.weapon.name);
        weaponColorMask = game.make.sprite(0, 0, weaponSlot.weapon.name + '_mask');
        weaponColorMask2 = game.make.sprite(0, 0, weaponSlot.weapon.name + '_mask2');

        weaponShadow = game.make.sprite(xAttach + game.shadowXOffset / 2, yAttach + game.shadowYOffset / 2, weaponSlot.weapon.name);

        weapon.xAttach = xAttach;
        weapon.yAttach = yAttach;
        weapon.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weapon.scale.setTo(0.25);

        weaponColorMask.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponColorMask.tint = wColor;

        weaponColorMask2.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponColorMask2.tint = w2Color;
        weaponColorMask2.alpha = 0.3;

        weaponShadow.anchor.setTo(weaponSlot.weapon.x_attach / 200, weaponSlot.weapon.y_attach / 200);
        weaponShadow.scale.setTo(0.25);
        weaponShadow.tint = 0x000000;
        weaponShadow.alpha = 0.5;

        // накладываем цветовые маски на спрайт
        weapon.addChild(weaponColorMask2);
        weapon.addChild(weaponColorMask);
    }

    squad.sprite = unit;
    squad.sprite.unitBody = body;
    squad.sprite.bodyShadow = bodyShadow;
    squad.sprite.bodyBottom = bodyBottom;
    squad.sprite.bodyBottomShadow = bodyBottomShadow;

    unit.addChild(bodyBottomShadow);
    unit.addChild(bodyBottom);
    unit.addChild(bodyShadow);
    unit.addChild(body);

    CreateEquip(squadBody, squad);

    if (weapon) {
        squad.sprite.weapon = weapon;
        squad.sprite.weaponShadow = weaponShadow;

        unit.addChild(weaponShadow);
        unit.addChild(weapon);
    }

    SetAngle(squad, rotate);
}

function CreateEquip(squadBody, squad) {
    squad.sprite.equipSprites = [];

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

            squad.sprite.addChild(slotShadow);
            squad.sprite.addChild(attachPoint);
            squad.sprite.addChild(slotSprite);

            squad.sprite.equipSprites.push({
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

    createSlots(squadBody.equippingI);
    createSlots(squadBody.equippingII);
    createSlots(squadBody.equippingIII);
    createSlots(squadBody.equippingIV);
    createSlots(squadBody.equippingV);
}

function mouseBodyOver(body, squad, unit) {
    let positionInterval = null;
    let checkTimeOut = null;

    body.events.onInputOver.add(function () {

        clearTimeout(checkTimeOut);

        if (document.getElementById("UserLabel" + squad.user_name)) {
            return;
        }

        //todo загрузка аватарки

        let userLabel = document.createElement('div');
        userLabel.id = "UserLabel" + squad.user_name;
        userLabel.className = "UserLabel";
        document.body.appendChild(userLabel);
        userLabel.innerHTML = `
            <div>
                <div>
                    <div class="logo"></div>
                    <h4>${squad.user_name}</h4>
                    <div class="detailUser" onmousedown="informationFunc('${squad.user_name}', '${squad.user_id}')">i</div>
                </div>
            </div>
        `;

        positionInterval = setInterval(function () {
            userLabel.style.left = unit.worldPosition.x - 50 + "px";
            userLabel.style.top = unit.worldPosition.y - 70 + "px";
            userLabel.style.display = "block";
        }, 10);
    }, this);

    body.events.onInputOut.add(function () {
        checkTimeOut = setTimeout(function () {
            if (document.getElementById("UserLabel" + squad.user_name)) document.getElementById("UserLabel" + squad.user_name).remove();
            clearInterval(positionInterval);
            clearTimeout(checkTimeOut);
        }, 2000);
    }, this);
}

function informationFunc(userName, userId) {
    if (userName === Data.user.login) {
        UsersStatus()
    } else {
        OtherUserStatus(userName, userId)
    }
}