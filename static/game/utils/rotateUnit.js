function rotateUnitSprites(spriteRotate, needRotate, unit) {

    spriteRotate = Math.round(spriteRotate);
    needRotate = Math.round(needRotate);

    if (spriteRotate < 0) {
        spriteRotate += 360;
    }

    if (needRotate < 0) {
        needRotate += 360;
    }

    if (needRotate > 360) {
        needRotate -= 360;
    }

    if (spriteRotate !== needRotate) {
        if (directionRotate(spriteRotate, needRotate)) {
            SetAngle(unit, spriteRotate + 1)
        } else {
            SetAngle(unit, spriteRotate - 1)
        }
    }
}

// метод вычесления в какую сторону меньше поворачивать обьект
function directionRotate(spriteAngle, rotate) {
    // true ++
    // false --
    let count = 0;
    let direction;

    if (spriteAngle < rotate) {
        for (; spriteAngle < rotate; spriteAngle++) {
            count++;
            direction = true;
        }
    } else {
        for (; spriteAngle > rotate; rotate++) {
            count++;
            direction = false;
        }
    }

    if (direction) {
        return count <= 180
    } else {
        return !(count <= 180)
    }
}

function SetAngle(unit, angle) {
    unit.sprite.angle = angle;
    SetShadowAngle(unit, angle)
}

function SetShadowAngle(unit, angle) {
    let shadowAngle = 45 - angle;
    let connectPoints = PositionAttachSprite(shadowAngle, game.shadowXOffset);

    unit.sprite.bodyShadow.x = connectPoints.x;
    unit.sprite.bodyShadow.y = connectPoints.y;
    unit.sprite.bodyBottomShadow.x = connectPoints.x;
    unit.sprite.bodyBottomShadow.y = connectPoints.y;

    if (unit.sprite.weapon) {
        let connectWeapons = PositionAttachSprite(shadowAngle, game.shadowXOffset / 2);
        unit.sprite.weaponShadow.x = connectWeapons.x + unit.sprite.weapon.xAttach;
        unit.sprite.weaponShadow.y = connectWeapons.y + unit.sprite.weapon.yAttach;
    }

    for (let i = 0; i < unit.sprite.equipSprites.length; i++) {
        let slot = unit.sprite.equipSprites[i];

        let connectWeapons = PositionAttachSprite(shadowAngle, game.shadowXOffset / 4);
        slot.shadow.x = connectWeapons.x + slot.xAttach;
        slot.shadow.y = connectWeapons.y + slot.yAttach;
        slot.shadow.angle = slot.sprite.angle;
    }
}